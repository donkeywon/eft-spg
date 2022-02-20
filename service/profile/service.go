package profile

import (
	"eft-spg/util"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
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
	profiles = make(map[string]*jsonvalue.V)
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

		v, err := jsonvalue.Unmarshal(bs)
		if err != nil {
			return errors.Wrapf(err, util.ErrReadProfile, filename)
		}

		SetProfile(sessID, v)
	}

	return err
}

func (s *Svc) Close() error {
	return nil
}

func (s *Svc) Shutdown() error {
	return nil
}

func (s *Svc) InitialProfile() *jsonvalue.V {
	p := util.GetEmptyJsonNode()
	p.Set(util.GetEmptyJsonNode()).At("pmc")
	p.Set(util.GetEmptyJsonNode()).At("scav")
	return p
}

func GetProfile(sessID string) *jsonvalue.V {
	return profiles[sessID]
}

func SetProfile(sessID string, v *jsonvalue.V) {
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

	_, err = f.Write(p.Bytes())
	if err != nil {
		return errors.Wrapf(err, util.ErrSaveProfile, FullProfileFileName(sessID))
	}

	return nil
}

func (s *Svc) GetProfileItemByPath(sesssID string, path interface{}, paths ...interface{}) *jsonvalue.V {
	p := GetProfile(sesssID)
	if p == nil {
		return nil
	}

	i, err := p.Get(path, paths...)
	if err != nil {
		return nil
	}

	return i
}

func (s *Svc) GetCharacterProfile(sessID string, typ string) *jsonvalue.V {
	return s.GetProfileItemByPath(sessID, "characters", typ)
}

func (s *Svc) GetPMCProfile(sessID string) *jsonvalue.V {
	return s.GetCharacterProfile(sessID, "pmc")
}

func (s *Svc) GetScavProfile(sessID string) *jsonvalue.V {
	return s.GetCharacterProfile(sessID, "scav")
}

func (s *Svc) SetScavProfile(sessID string, sp *jsonvalue.V) error {
	p := GetProfile(sessID)
	if p == nil {
		return nil
	}

	_, err := p.Set(sp).At("characters", "scav")
	if err != nil {
		return errors.Wrapf(err, util.ErrSetScavProfile, sessID)
	}

	return nil
}

func (s *Svc) GetCompleteProfile(sessID string) *jsonvalue.V {
	v := util.GetEmptyJsonArray()

	if !s.IsWipe(sessID) {
		return v
	}

	pmcProfile := s.GetPMCProfile(sessID)
	if pmcProfile != nil {
		v.Append(pmcProfile)
	}

	scavProfile := s.GetScavProfile(sessID)
	if scavProfile != nil {
		v.Append(scavProfile)
	}

	return v
}

func (s *Svc) GetProfileInfo(sessID string) *jsonvalue.V {
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

	return p.MustGet("info", "wipe").Bool()
}

func (s *Svc) FullPath(sessID string) string {
	return filepath.Join(s.Config.Path, FullProfileFileName(sessID))
}

func FullProfileFileName(sessID string) string {
	return sessID + "." + ProfileFileExt
}
