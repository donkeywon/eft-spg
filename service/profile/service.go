package profile

import (
	"eft-spg/util"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/service"
	"github.com/pkg/errors"
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
	profiles = make(map[string]*ast.Node)
	path     = "user/profiles"
)

type Svc struct {
	*service.BaseService
	Config *Config
}

func New(config *Config) *Svc {
	return &Svc{
		BaseService: service.NewBase(),
		Config:      config,
	}
}

func (s *Svc) Name() string {
	return Name
}

func (s *Svc) Open() error {
	if !util.DirExist(s.Config.Path) {
		err := os.MkdirAll(s.Config.Path, os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, util.ErrMkdirFail, s.Config.Path)
		}
	}

	path = s.Config.Path

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
			return errors.Wrapf(err, util.ErrReadProfile, filename)
		}

		v, err := sonic.Get(bs)
		if err != nil {
			return errors.Wrapf(err, util.ErrReadProfile, filename)
		}

		SetProfile(sessID, &v)
	}

	// TODO ?

	return err
}

func InitialProfile() *ast.Node {
	p := util.GetEmptyJsonNode()
	p.Set("pmc", util.GetEmptyJsonNode())
	p.Set("scav", util.GetEmptyJsonNode())
	return &p
}

func GetProfile(sessID string) *ast.Node {
	return profiles[sessID]
}

func SetProfile(sessID string, v *ast.Node) {
	profiles[sessID] = v
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
		return errors.Wrapf(err, util.ErrSaveProfile, FullProfileFileName(sessID))
	}

	bs, err := p.MarshalJSON()
	if err != nil {
		return errors.Wrapf(err, util.ErrSaveProfile, FullProfileFileName(sessID))
	}

	_, err = f.Write(bs)
	if err != nil {
		return errors.Wrapf(err, util.ErrSaveProfile, FullProfileFileName(sessID))
	}

	return nil
}

func GetProfileItemByPath(sesssID string, paths ...interface{}) *ast.Node {
	p := GetProfile(sesssID)
	if p == nil {
		return nil
	}

	n := p.GetByPath(paths)
	if n.Check() != nil {
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

	if p.Get("characters").Check() != nil {
		_, err := p.Set("characters", util.GetEmptyJsonNode())
		if err != nil {
			return errors.Wrapf(err, util.ErrSetScavProfile, sessID)
		}
	}

	_, err := p.Get("characters").Set("scav", *sp)
	if err != nil {
		return errors.Wrapf(err, util.ErrSetScavProfile, sessID)
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

func CreateProfile() {
	// TODO
}

func IsWipe(sessID string) bool {
	p := GetProfile(sessID)
	if p == nil {
		return false
	}

	n := p.Get("info")
	if n.Check() != nil {
		return false
	}

	w, err := n.Get("wipe").Bool()
	if err != nil {
		return false
	}

	return w
}

func FullPath(sessID string) string {
	return filepath.Join(path, FullProfileFileName(sessID))
}

func FullProfileFileName(sessID string) string {
	return sessID + "." + ProfileFileExt
}
