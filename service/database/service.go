package database

import (
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	Name = "database"
)

type Service struct {
	config *Config
	n      ast.Node
	logger *zap.Logger
}

func (s *Service) Name() string {
	return Name
}

func (s *Service) Open() error {
	s.logger.Info("Open")
	var err error
	s.n, err = readDir(s.config.AssetPath)
	return err
}

func readDir(dir string) (ast.Node, error) {
	var m ast.Node

	fs, _ := ioutil.ReadDir(dir)
	for _, f := range fs {
		path := filepath.Join(dir, f.Name())
		fn, fe := fileNameAndExt(f.Name())
		if f.IsDir() {
			n, err := readDir(path)
			if err != nil {
				return m, err
			}

			_, err = m.Set(fn, n)
			if err != nil {
				return m, err
			}

		} else {
			if fe != "json" {
				continue
			}

			fbs, err := ioutil.ReadFile(path)
			if err != nil {
				return m, err
			}

			n, err := sonic.Get(fbs)
			if err != nil {
				return m, err
			}

			_, err = m.Set(fn, n)
			if err != nil {
				return m, err
			}
		}
	}

	return m, nil
}

func fileNameAndExt(fileName string) (string, string) {
	splited := strings.Split(fileName, ".")
	if len(splited) < 2 {
		return fileName, ""
	}
	return strings.Join(splited[0:len(splited)-1], "."), splited[len(splited)-1]
}

func (s *Service) Close() error {
	s.logger.Info("Close")
	return nil
}

func (s *Service) Shutdown() error {
	s.logger.Info("Shutdown")
	return nil
}

func (s *Service) WithLogger(logger *zap.Logger) {
	s.logger = logger.Named(s.Name())
}

func (s *Service) Statistics() map[string]float64 {
	return nil
}

func (s *Service) LastError() error {
	return nil
}
