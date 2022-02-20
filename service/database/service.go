package database

import (
	"eft-spg/util"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/donkeywon/gtil/service"
)

const (
	Name = "database"
)

var (
	database *jsonvalue.V
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
	s.Info("Opening")
	defer s.Info("Opened")
	d, err := util.ReadDatabaseBox()
	if err != nil {
		return err
	}
	database = d

	return err
}

func (s *Svc) Close() error {
	return nil
}

func (s *Svc) Shutdown() error {
	return nil
}

func GetDatabase() *jsonvalue.V {
	return database
}

func GetProfileEditions() ([]string, error) {
	pe, err := database.Get("templates", "profiles")
	if err != nil {
		return nil, err
	}

	editions := make([]string, 0, 4)
	for edition, _ := range pe.ForRangeObj() {
		editions = append(editions, edition)
	}

	return editions, nil
}
