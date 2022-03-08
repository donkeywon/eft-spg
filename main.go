package main

import (
	"context"
	"eft-spg/cmd"
	"fmt"
	"github.com/donkeywon/gtil/logger"
	"github.com/donkeywon/gtil/logger/core"
	"github.com/donkeywon/gtil/service"
	"github.com/traefik/yaegi/interp"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	l, _ := logger.FromConfig(logger.DefaultConsoleConfig(), zap.WrapCore(core.NewStackExtractCore))
	zap.ReplaceGlobals(l)

	config := cmd.NewConfig()
	f, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		l.Error("Start fail", zap.Error(err))
		return
	}

	err = yaml.Unmarshal(f, config)
	if err != nil {
		l.Error("Start fail", zap.Error(err))
		return
	}

	c := cmd.New(config)
	err = service.DoOpen(c, context.Background(), l)
	if err != nil {
		l.Error("Start fail", zap.Error(err))
		return
	}

	l.Info("Start success")

	src, _ := ioutil.ReadFile("./mod/mod.go")
	itp := interp.New(interp.Options{})
	_, err = itp.Eval(string(src))
	if err != nil {
		panic(err)
	}

	gv, err := itp.Eval("mod.GetValue")
	if err != nil {
		panic(err)
	}

	fmt.Println(gv.Interface().(func() string)())

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
