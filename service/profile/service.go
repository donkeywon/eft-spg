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
	Name            = "save"
	ProfileFileExt  = util.JsonFileExt
	ProfileFilePerm = 0644
)

var (
	profiles = make(map[string]*ast.Node)
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

func (s *Svc) LoadProfile(sessID string) error {
	var err error
	filename := FullProfileFileName(sessID)
	path := filepath.Join(s.Config.Path, filename)
	if util.FileExist(path) {
		bs, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, util.ErrReadProfile, filename)
		}

		v, err := sonic.Get(bs)
		if err != nil {
			return errors.Wrapf(err, util.ErrReadProfile, filename)
		}

		SetProfile(sessID, &v)
	}

	return err
}

func (s *Svc) Close() error {
	return nil
}

func (s *Svc) Shutdown() error {
	return nil
}

func (s *Svc) InitialProfile() *ast.Node {
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
func (s *Svc) SaveProfile(sessID string) error {
	filePath := s.FullPath(sessID)

	p := GetProfile(sessID)
	if p == nil {
		p = s.InitialProfile()
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

func (s *Svc) GetProfileItemByPath(sesssID string, paths ...interface{}) *ast.Node {
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

func (s *Svc) FullPath(sessID string) string {
	return filepath.Join(s.Config.Path, FullProfileFileName(sessID))
}

func FullProfileFileName(sessID string) string {
	return sessID + "." + ProfileFileExt
}
