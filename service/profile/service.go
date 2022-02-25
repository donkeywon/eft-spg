package profile

import (
	"eft-spg/service/database"
	"eft-spg/service/profile/hook"
	"eft-spg/util"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
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

		err = s.LoadProfile(fn)
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

func (s *Svc) LoadProfile(sessID string) error {
	var err error
	filename := s.FullPath(sessID)
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
			s.Info("Profile file crash", zap.String("file", sessID+"."+ProfileFileExt))
			return nil
		}

		s.SetProfile(sessID, &v)
	}

	hook.PostLoadHook(s.GetProfile(sessID))

	return err
}

func InitialProfile() *ast.Node {
	p := util.GetEmptyJsonNode()
	p.Set("pmc", util.GetEmptyJsonNode())
	p.Set("scav", util.GetEmptyJsonNode())
	return &p
}

func (s *Svc) GetProfile(sessID string) *ast.Node {
	return s.profiles[sessID]
}

func (s *Svc) SetProfile(sessID string, v *ast.Node) {
	s.profiles[sessID] = v
}

// TODO go update
func (s *Svc) SaveProfile(sessID string) error {
	filePath := s.FullPath(sessID)

	p := s.GetProfile(sessID)
	if p == nil {
		p = InitialProfile()
		s.SetProfile(sessID, p)
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

func (s *Svc) GetSessProfileByUsername(username string) (string, *ast.Node) {
	for sessID, profile := range s.profiles {
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

func (s *Svc) GetProfileItemByPath(sesssID string, paths ...interface{}) *ast.Node {
	p := s.GetProfile(sesssID)
	if p == nil {
		return nil
	}

	n := p.GetByPath(paths...)
	if !n.Exists() {
		return nil
	}

	return n
}

func (s *Svc) GetCharacterProfile(sessID string, typ string) *ast.Node {
	return s.GetProfileItemByPath(sessID, "characters", typ)
}

func (s *Svc) GetPMCProfile(sessID string) *ast.Node {
	return s.GetCharacterProfile(sessID, "pmc")
}

func (s *Svc) GetScavProfile(sessID string) *ast.Node {
	return s.GetCharacterProfile(sessID, "scav")
}

func (s *Svc) SetScavProfile(sessID string, sp *ast.Node) error {
	p := s.GetProfile(sessID)
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

func (s *Svc) GetCompleteProfile(sessID string) *ast.Node {
	v := util.GetEmptyJsonArray()

	if !s.IsWipe(sessID) {
		return &v
	}

	i := 0
	pmcProfile := s.GetPMCProfile(sessID)
	if pmcProfile != nil {
		v.SetByIndex(i, *pmcProfile)
		i++
	}

	scavProfile := s.GetScavProfile(sessID)
	if scavProfile != nil {
		v.SetByIndex(i, *scavProfile)
	}

	return &v
}

func (s *Svc) GetProfileInfo(sessID string) *ast.Node {
	return s.GetProfileItemByPath(sessID, "info")
}

func (s *Svc) CreateProfile() {
	// TODO
}

func (s *Svc) IsWipe(sessID string) bool {
	p := s.GetProfile(sessID)
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

func (s *Svc) FullPath(sessID string) string {
	return filepath.Join(s.Config.Path, s.FullProfileFileName(sessID))
}

func (s *Svc) FullProfileFileName(sessID string) string {
	return sessID + "." + ProfileFileExt
}

func (s *Svc) GetMiniProfile(sessID string) (*ast.Node, error) {
	p := s.GetProfile(sessID)
	if p == nil {
		return nil, errors.New(util.ErrGetMiniProfile)
	}

	info := s.GetProfileInfo(sessID)
	username, _ := info.Get("username").String()

	maxLevel, err := database.GetSvc().GetMaxLevel()
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
			fmt.Sprintf(pbs, username, "unknown", "unknown", 0, 0, 0, 0, maxLevel, s.GetDefaultAkiData()))
		return &n, errors.Wrap(err, util.ErrGetMiniProfile)
	}

	pmcInfo := pmc.Get("Info")
	nickname, _ := pmcInfo.Get("Nickname").String()
	side, _ := pmcInfo.Get("Side").String()
	currLvl, _ := pmcInfo.GetByPath("Level").Int64()
	currExp, _ := pmcInfo.GetByPath("Experience").Int64()
	var prevExp int64
	if currLvl > 0 {
		prevExp, _ = s.GetExperience(int(currLvl))
	}
	nextLvl, _ := s.GetExperience(int(currLvl) + 1)
	n, err := sonic.GetFromString(
		fmt.Sprintf(pbs, username, nickname, side, currLvl, currExp, prevExp, nextLvl, maxLevel, s.GetDefaultAkiData()))
	return &n, errors.Wrap(err, util.ErrGetMiniProfile)
}

func (s *Svc) GetAllMiniProfiles() (*ast.Node, error) {
	ps := ast.NewArray([]ast.Node{})

	for sessID, _ := range s.profiles {
		mp, err := s.GetMiniProfile(sessID)
		if err != nil {
			return nil, err
		}
		ps.Add(*mp)
	}

	return &ps, nil
}

func (s *Svc) GetDefaultAkiData() string {
	return fmt.Sprintf(`{"version":"%s"}`, util.ServerVersion)
}

func (s *Svc) GetExperience(lvl int) (int64, error) {
	expTable, err := database.GetSvc().GetDatabase().GetByPath("globals", "config", "exp", "level", "exp_table").ArrayUseNode()
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

func (s *Svc) RemoveProfile(sessID string) error {
	if s.GetProfile(sessID) == nil {
		return errors.New(util.ErrUserNotExist)
	}

	delete(s.profiles, sessID)
	err := os.Remove(s.FullPath(sessID))
	if err != nil {
		return errors.Wrap(err, util.ErrRemoveProfile)
	}

	return nil
}
