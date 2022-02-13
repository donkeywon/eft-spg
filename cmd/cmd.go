package cmd

import (
	"go.uber.org/zap"
)

const (
	Name = "cmd"
)

type Command struct {
	logger *zap.Logger
}

func (c *Command) Name() string {
	return Name
}

func (c *Command) Open() error {
	//TODO implement me
	panic("implement me")
}

func (c *Command) Close() error {
	//TODO implement me
	panic("implement me")
}

func (c *Command) Shutdown() error {
	//TODO implement me
	panic("implement me")
}

func (c *Command) WithLogger(logger *zap.Logger) {
	c.logger = logger.Named(Name)
}

func (c *Command) Statistics() map[string]float64 {
	//TODO implement me
	panic("implement me")
}
