package main

import (
	"context"
	"eft-spg/cmd"
	"github.com/donkeywon/gtil/logger"
	"github.com/donkeywon/gtil/service"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	config := cmd.NewConfig()

	f, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, config)
	if err != nil {
		panic(err)
	}

	lc := logger.DefaultConsoleConfig()
	lc.Development = false
	lc.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	l, _ := logger.FromConfig(lc)
	l = l.Named("main")

	c := cmd.New(config)
	err = service.DoOpen(c, context.Background(), l)
	if err != nil {
		panic(err)
	}

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	<-signalCh
	l.Info("Receive signal, exit")

	go func() {
		err := service.DoClose(c)
		if err != nil {
			l.Error("Close fail", zap.Error(err))
		}
	}()

	select {
	case <-c.Closed():
		l.Info("Closed")
	case <-signalCh:
		l.Info("Receive signal twice, exit")
	case <-time.After(time.Second * 10):
		l.Info("Close timeout, exit")
	}
}
