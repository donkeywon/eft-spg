package database

import (
	"eft-spg/util"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/service"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	Name = "database"
)

var (
	svc *Svc
)

func GetSvc() *Svc {
	return svc
}

type Svc struct {
	*service.BaseService
	Config   *Config
	database *ast.Node
}

func New(config *Config) *Svc {
	svc = &Svc{
		BaseService: service.NewBase(),
		Config:      config,
	}
	return svc
}

func (s *Svc) Name() string {
	return Name
}

func (s *Svc) Open() error {
	s.Info("Opening")
	defer s.Info("Opened")

	bs := make([]byte, 500000000, 500000000)
	n, err := readDirToJson(bs, s.Config.Path)
	if err != nil {
		return err
	}

	node, err := sonic.Get(bs[:n])
	if err != nil {
		return errors.Wrapf(err, util.ErrParseJson)
	}
	bs = nil

	s.database = &node

	return err
}

func (s *Svc) Close() error {
	return nil
}

func (s *Svc) Shutdown() error {
	return nil
}

func readFile(bs []byte, path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, errors.Wrapf(err, util.ErrOpenFile)
	}
	defer f.Close()

	var n int
	for {
		n1, err := f.Read(bs[n:])
		if err != nil && err != io.EOF {
			return 0, errors.Wrapf(err, util.ErrReadFile)
		}
		if n1 == 0 {
			break
		}
		n += n1
	}

	return n, nil
}

func readDirToJson(bs []byte, path string) (int, error) {
	if util.FileExist(path) {
		n, err := readFile(bs, path)
		if err != nil {
			return 0, errors.Wrapf(err, util.ErrReadDirToJson)
		}
		return n, nil
	}

	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var offset int

	bs[0] = '{'
	offset += 1
	for _, info := range infos {
		fn, fe := util.FileNameAndExt(info.Name())
		if fe != util.JsonFileExt && util.FileExist(filepath.Join(path, info.Name())) {
			continue
		}

		part := []byte(`"` + fn + `":`)

		var i int
		for ; i < len(part); i++ {
			bs[offset+i] = part[i]
		}
		offset += i

		n1, err := readDirToJson(bs[offset:], filepath.Join(path, info.Name()))
		if err != nil {
			return offset, err
		}
		offset += n1

		bs[offset] = ','
		offset += 1
	}
	if bs[offset-1] == ',' {
		offset -= 1
	}
	bs[offset] = '}'
	offset += 1

	return offset, nil
}

func GetDatabase() *ast.Node {
	return svc.database
}

func GetProfileEditionsTemplate() *ast.Node {
	return GetDatabase().GetByPath("templates", "profiles")
}

func GetProfileEditions() ([]string, error) {
	ps := GetProfileEditionsTemplate()

	editions := make([]string, 0, 4)
	err := ps.ForEach(func(path ast.Sequence, node *ast.Node) bool {
		editions = append(editions, *path.Key)
		return true
	})

	return editions, err
}

func GetMaxLevel() (int, error) {
	expTable, err := GetDatabase().GetByPath("globals", "config", "exp", "level", "exp_table").Array()
	if err != nil {
		return 0, errors.Wrap(err, "Get max level fail")
	}

	return len(expTable) - 1, nil
}
