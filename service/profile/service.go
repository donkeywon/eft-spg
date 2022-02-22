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
	return &Svc{
		BaseService: service.NewBase(),
		Config:      config,
		profiles:    make(map[string]*ast.Node),
	}
}

func (s *Svc) Name() string {
	return Name
}

func (s *Svc) Open() error {
	svc = s

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
			return errors.Wrapf(err, util.ErrReadProfile, filename)
		}

		v, err := sonic.Get(bs)
		if err != nil {
			return errors.Wrapf(err, util.ErrReadProfile, filename)
		}

		s.SetProfile(sessID, &v)
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
		return errors.Wrapf(err, util.ErrSaveProfile, s.FullProfileFileName(sessID))
	}

	bs, err := p.MarshalJSON()
	if err != nil {
		return errors.Wrapf(err, util.ErrSaveProfile, s.FullProfileFileName(sessID))
	}

	_, err = f.Write(bs)
	if err != nil {
		return errors.Wrapf(err, util.ErrSaveProfile, s.FullProfileFileName(sessID))
	}

	return nil
}

func (s *Svc) GetSessProfileByUsername(username string) (string, *ast.Node) {
	for sessID, profile := range s.profiles {
		un, err := profile.GetByPath("info", "username").String()
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
	p := s.GetProfile(sessID)
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
	p := s.GetProfile(sessID)
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
	return filepath.Join(s.Config.Path, s.FullProfileFileName(sessID))
}

func (s *Svc) FullProfileFileName(sessID string) string {
	return sessID + "." + ProfileFileExt
}
