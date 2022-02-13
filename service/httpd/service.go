package httpd

import "go.uber.org/zap"

const (
	Name = "httpd"
)

type Service struct {
	logger *zap.Logger
}

func (s *Service) Name() string {
	return Name
}

func (s *Service) Open() error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Close() error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Shutdown() error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) WithLogger(logger *zap.Logger) {
	s.logger = logger.Named(Name)
}

func (s *Service) Statistics() map[string]float64 {
	//TODO implement me
	panic("implement me")
}
