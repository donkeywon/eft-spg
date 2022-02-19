package save

import (
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/donkeywon/eft-spg/util"
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

type svc struct {
	*service.BaseService
	config *Config

	profiles map[string]*jsonvalue.V
}

func New(config *Config) *svc {
	return &svc{
		BaseService: service.NewBase(),
		config:      config,
		profiles:    make(map[string]*jsonvalue.V),
	}
}

func (s *svc) Name() string {
	return Name
}

func (s *svc) Open() error {
	if !util.FileOrPathExist(s.config.Path) {
		err := os.MkdirAll(s.config.Path, os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, util.ErrMkdirFail, s.config.Path)
		}
	}

	err := filepath.Walk(s.config.Path, func(path string, info fs.FileInfo, err error) error {
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

func (s *svc) LoadProfile(sessID string) error {
	var err error
	filename := FullProfileFileName(sessID)
	path := filepath.Join(s.config.Path, filename)
	if util.FileOrPathExist(path) {
		bs, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, util.ErrReadProfile, filename)
		}

		v, err := jsonvalue.Unmarshal(bs)
		if err != nil {
			return errors.Wrapf(err, util.ErrReadProfile, filename)
		}

		s.SetProfile(sessID, v)
	}

	return err
}

func (s *svc) Close() error {
	return nil
}

func (s *svc) Shutdown() error {
	return nil
}

func (s *svc) InitialProfile() *jsonvalue.V {
	p := util.GetEmptyJsonNode()
	p.Set(util.GetEmptyJsonNode()).At("pmc")
	p.Set(util.GetEmptyJsonNode()).At("scav")
	return p
}

func (s *svc) GetProfile(sessID string) *jsonvalue.V {
	return s.profiles[sessID]
}

func (s *svc) SetProfile(sessID string, v *jsonvalue.V) {
	s.profiles[sessID] = v
}

// TODO go update
func (s *svc) SaveProfile(sessID string) error {
	filePath := s.FullPath(sessID)

	p := s.GetProfile(sessID)
	if p == nil {
		p = s.InitialProfile()
		s.SetProfile(sessID, p)
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, ProfileFilePerm)
	if err != nil {
		return errors.Wrapf(err, util.ErrSaveProfile, FullProfileFileName(sessID))
	}

	_, err = f.Write(p.Bytes())
	if err != nil {
		return errors.Wrapf(err, util.ErrSaveProfile, FullProfileFileName(sessID))
	}

	return nil
}

func (s *svc) GetPMCProfile(sessID string) *jsonvalue.V {
	p := s.GetProfile(sessID)
	if p == nil {
		return nil
	}

	return p.MustGet("characters", "pmc")
}

func (s *svc) GetScavProfile(sessID string) *jsonvalue.V {
	p := s.GetProfile(sessID)
	if p == nil {
		return nil
	}

	return p.MustGet("characters", "scav")
}

func (s *svc) SetScavProfile(sessID string, sp *jsonvalue.V) error {
	p := s.GetProfile(sessID)
	if p == nil {
		return nil
	}

	_, err := p.Set(sp).At("characters", "scav")
	if err != nil {
		return errors.Wrapf(err, util.ErrSetScavProfile, sessID)
	}

	return nil
}

func (s *svc) FullPath(sessID string) string {
	return filepath.Join(s.config.Path, FullProfileFileName(sessID))
}

func FullProfileFileName(sessID string) string {
	return sessID + "." + ProfileFileExt
}
