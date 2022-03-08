package profile

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/service/database"
	"github.com/donkeywon/eft-spg/service/profile/hook"
	"github.com/donkeywon/eft-spg/util"
	"github.com/donkeywon/gtil/service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	Name            = "profile"
	ProfileFileExt  = util.JsonFileExt
	ProfileFilePerm = 0644
)

var (
	svc *Svc
)

func GetSvc() *Svc {
	return svc
}

type Svc struct {
	*service.BaseService
	Config *Config

	profiles map[string]*ast.Node
}

func New(config *Config) *Svc {
	svc = &Svc{
		BaseService: service.NewBase(),
		Config:      config,
		profiles:    make(map[string]*ast.Node),
	}

	return svc
}

func (s *Svc) Name() string {
	return Name
}

func (s *Svc) Open() error {

	if !util.DirExist(s.Config.Path) {
		err := os.MkdirAll(s.Config.Path, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, util.ErrMkdirFail)
		}
	}

	err := filepath.Walk(s.Config.Path, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fn, fe := util.FileNameAndExt(info.Name())
		if fe != ProfileFileExt {
			return nil
		}

		err = LoadProfile(fn)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	// TODO
	// RagfairServer.addPlayerOffers
	// RagfairServer.update

	return nil
}

func (s *Svc) Close() error {
	return nil
}

func (s *Svc) Shutdown() error {
	return nil
}

func LoadProfile(sessID string) error {
	var err error
	filename := FullPath(sessID)
	if util.FileExist(filename) {
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			return errors.Wrap(err, util.ErrReadProfile)
		}

		v, err := sonic.Get(bs)
		if err != nil {
			return errors.Wrap(err, util.ErrReadProfile)
		}

		if !v.Get("info").Exists() {
			GetSvc().Info("Profile file crash", zap.String("file", sessID+"."+ProfileFileExt))
			return nil
		}

		SetProfile(sessID, &v)
	}

	hook.PostLoadHook(GetProfile(sessID))

	return err
}

func InitialProfile() *ast.Node {
	p := util.GetEmptyJsonNode()
	p.Set("pmc", util.GetEmptyJsonNode())
	p.Set("scav", util.GetEmptyJsonNode())
	return &p
}

func GetProfile(sessID string) *ast.Node {
	return GetSvc().profiles[sessID]
}

func SetProfile(sessID string, v *ast.Node) {
	GetSvc().profiles[sessID] = v
}

// TODO go update
func SaveProfile(sessID string) error {
	filePath := FullPath(sessID)

	p := GetProfile(sessID)
	if p == nil {
		p = InitialProfile()
		SetProfile(sessID, p)
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, ProfileFilePerm)
	if err != nil {
		return errors.Wrap(err, util.ErrSaveProfile)
	}

	bs, err := p.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, util.ErrSaveProfile)
	}

	_, err = f.Write(bs)
	if err != nil {
		return errors.Wrap(err, util.ErrSaveProfile)
	}

	return nil
}

func GetSessProfileByUsername(username string) (string, *ast.Node) {
	for sessID, profile := range GetSvc().profiles {
		unn := profile.GetByPath("info", "username")
		if !unn.Exists() {
			continue
		}
		un, err := unn.String()
		if err != nil {
			continue
		}

		if username == un {
			return sessID, profile
		}
	}

	return "", nil
}

func GetProfileItemByPath(sesssID string, paths ...interface{}) *ast.Node {
	p := GetProfile(sesssID)
	if p == nil {
		return nil
	}

	n := p.GetByPath(paths...)
	if !n.Exists() {
		return nil
	}

	return n
}

func GetCharacterProfile(sessID string, typ string) *ast.Node {
	return GetProfileItemByPath(sessID, "characters", typ)
}

func GetPMCProfile(sessID string) *ast.Node {
	return GetCharacterProfile(sessID, "pmc")
}

func GetScavProfile(sessID string) *ast.Node {
	return GetCharacterProfile(sessID, "scav")
}

func SetScavProfile(sessID string, sp *ast.Node) error {
	p := GetProfile(sessID)
	if p == nil {
		return nil
	}

	if !p.Get("characters").Exists() {
		_, err := p.Set("characters", util.GetEmptyJsonNode())
		if err != nil {
			return errors.Wrap(err, util.ErrSetScavProfile)
		}
	}

	_, err := p.Get("characters").Set("scav", *sp)
	if err != nil {
		return errors.Wrap(err, util.ErrSetScavProfile)
	}

	return nil
}

func GetCompleteProfile(sessID string) *ast.Node {
	v := util.GetEmptyJsonArray()

	if !IsWipe(sessID) {
		return &v
	}

	i := 0
	pmcProfile := GetPMCProfile(sessID)
	if pmcProfile != nil {
		v.SetByIndex(i, *pmcProfile)
		i++
	}

	scavProfile := GetScavProfile(sessID)
	if scavProfile != nil {
		v.SetByIndex(i, *scavProfile)
	}

	return &v
}

func GetProfileInfo(sessID string) *ast.Node {
	return GetProfileItemByPath(sessID, "info")
}

func CreateProfile(sessID string, info *ast.Node) error {
	if !info.Get("side").Exists() ||
		!info.Get("nickname").Exists() ||
		!info.Get("voiceId").Exists() ||
		!info.Get("headId").Exists() ||
		sessID == "" {
		return errors.New(util.ErrIllegalArg)
	}

	side, _ := info.Get("side").String()
	nickname, _ := info.Get("nickname").String()
	voiceId, _ := info.Get("voiceId").String()
	headId, _ := info.Get("headId").String()

	if !database.GetDatabase().GetByPath("templates", "customization", voiceId, "_name").Exists() {
		return errors.New(util.ErrIllegalArg)
	}

	pInfo := GetProfileInfo(sessID)
	edition, _ := pInfo.Get("edition").String()
	tp := database.GetDatabase().GetByPath("profiles", edition, strings.ToLower(side))
	pmcD := tp.Get("character")

	if GetProfile(sessID) != nil {
		RemoveProfile(sessID)
	}

	pmcD.Set("_id", ast.NewString("pmc"+sessID))
	pmcD.Set("aid", ast.NewString(sessID))
	pmcD.Set("savage", ast.NewString("scav"+sessID))
	pmcD.Get("Info").Set("Nickname", ast.NewString(nickname))
	pmcD.Get("Info").Set("LowerNickname", ast.NewString(strings.ToLower(nickname)))
	pmcD.Get("Info").Set("RegistrationDate", ast.NewNumber(strconv.Itoa(int(time.Now().Unix()))))
	pmcD.Get("Info").Set("Voice", *database.GetDatabase().GetByPath("templates", "customization", voiceId, "_name"))
	pmcD.Get("Stats").Set("SessionCounters", ast.NewObject([]ast.Pair{{Key: "Items", Value: ast.NewArray(nil)}}))
	pmcD.Get("Customization").Set("Head", ast.NewString(headId))
	pmcD.Get("Health").Set("UpdateTime", ast.NewNumber(strconv.Itoa(int(time.Now().Unix()))))
	pmcD.Set("Quests", ast.NewArray(nil))
	pmcD.Set("RepeatableQuests", ast.NewArray(nil))
	pmcD.Set("CarExtractCounts", ast.NewArray(nil))

	return nil
}

func GenerateScavProfile(sessID string) {
	// TODO
}

func GetDefaultCounters() *ast.Node {
	n, _ := sonic.Get([]byte(`
{
    "SessionCounters": {
        "Items": []
    },
    "OverallCounters": {
        "Items": []
    }
}`))
	return &n
}

func IsWipe(sessID string) bool {
	p := GetProfile(sessID)
	if p == nil {
		return false
	}

	n := p.Get("info")
	if !n.Exists() {
		return false
	}

	w, err := n.Get("wipe").Bool()
	if err != nil {
		return false
	}

	return w
}

func FullPath(sessID string) string {
	return filepath.Join(GetSvc().Config.Path, FullProfileFileName(sessID))
}

func FullProfileFileName(sessID string) string {
	return sessID + "." + ProfileFileExt
}

func GetMiniProfile(sessID string) (*ast.Node, error) {
	p := GetProfile(sessID)
	if p == nil {
		return nil, errors.New(util.ErrGetMiniProfile)
	}

	info := GetProfileInfo(sessID)
	username, _ := info.Get("username").String()

	maxLevel, err := database.GetMaxLevel()
	if err != nil {
		return nil, errors.Wrap(err, util.ErrGetMiniProfile)
	}

	pmc := p.GetByPath("characters", "pmc")
	if !pmc.Exists() {
		return nil, errors.New(util.ErrProfileCrash)
	}

	pbs := `{
    "username": "%s",
    "nickname": "%s",
    "side": "%s",
    "currlvl": %d,
    "currexp": %d,
    "prevexp": %d,
    "nextlvl": %d,
    "maxlvl": %d,
    "akiData": %s
}`

	if !pmc.Get("Info").Exists() || !pmc.GetByPath("Info", "Level").Exists() {
		n, err := sonic.GetFromString(
			fmt.Sprintf(pbs, username, "unknown", "unknown", 0, 0, 0, 0, maxLevel, GetDefaultAkiData()))
		return &n, errors.Wrap(err, util.ErrGetMiniProfile)
	}

	pmcInfo := pmc.Get("Info")
	nickname, _ := pmcInfo.Get("Nickname").String()
	side, _ := pmcInfo.Get("Side").String()
	currLvl, _ := pmcInfo.GetByPath("Level").Int64()
	currExp, _ := pmcInfo.GetByPath("Experience").Int64()
	var prevExp int64
	if currLvl > 0 {
		prevExp, _ = GetExperience(int(currLvl))
	}
	nextLvl, _ := GetExperience(int(currLvl) + 1)
	n, err := sonic.GetFromString(
		fmt.Sprintf(pbs, username, nickname, side, currLvl, currExp, prevExp, nextLvl, maxLevel, GetDefaultAkiData()))
	return &n, errors.Wrap(err, util.ErrGetMiniProfile)
}

func GetAllMiniProfiles() (*ast.Node, error) {
	ps := ast.NewArray([]ast.Node{})

	for sessID, _ := range GetSvc().profiles {
		mp, err := GetMiniProfile(sessID)
		if err != nil {
			return nil, err
		}
		ps.Add(*mp)
	}

	return &ps, nil
}

func GetDefaultAkiData() string {
	return fmt.Sprintf(`{"version":"%s"}`, util.ServerVersion)
}

func GetExperience(lvl int) (int64, error) {
	expTable, err := database.GetDatabase().GetByPath("globals", "config", "exp", "level", "exp_table").ArrayUseNode()
	if err != nil {
		return 0, errors.Wrap(err, "Get experience fail")
	}

	var exp int64

	for i := 0; i < lvl && i < len(expTable)-1; i++ {
		e, _ := expTable[i].Get("exp").Int64()
		exp += e
	}

	return exp, nil
}

func RemoveProfile(sessID string) error {
	if GetProfile(sessID) == nil {
		return errors.New(util.ErrUserNotExist)
	}

	delete(GetSvc().profiles, sessID)
	err := os.Remove(FullPath(sessID))
	if err != nil {
		return errors.Wrap(err, util.ErrRemoveProfile)
	}

	return nil
}
