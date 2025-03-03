package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"ship-manage-local/pkg/config"
	"ship-manage-local/pkg/database"
	"ship-manage-local/pkg/global"
	"ship-manage-local/pkg/web"
	"syscall"

	"github.com/go-kratos/kratos/v2"
	"github.com/sirupsen/logrus"
)

func init() {
	c, err := config.New()
	if err != nil {
		logrus.Fatalf("generate config error: %v", err)
	}

	c.Show()
}

func main() {
	global.ShowProgramInfo()

	srv := web.NewService()
	app := newApp(srv)

	go func() {
		if err := app.Run(); err != nil {
			logrus.Fatalf("ship_management program start error: %v", err)
		}
	}()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)
	<-stopSignal
	logrus.Info("ship_management program is stopping ......")

	if err := app.Stop(); err != nil {
		logrus.Errorf("ship_management stop error: %v", err)
	}
}

func newApp(gs *web.Service) *kratos.App {
	return kratos.New(
		kratos.Name("ship_management"),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs),
		kratos.AfterStart(AfterStart),
		kratos.AfterStop(AfterStop),
	)
}

func AfterStart(ctx context.Context) error {
	if _, err := os.Stat(global.ResourceDir); err != nil {
		if err := os.Mkdir(global.ResourceDir, 0755); err != nil {
			logrus.Errorf("create directory %s failed: %v", global.ResourceDir, err)
		}
	}

	if _, err := os.Stat(global.DBFilesDir); err != nil {
		if err := os.Mkdir(global.DBFilesDir, 0755); err != nil {
			logrus.Errorf("create directory %s failed: %v", global.DBFilesDir, err)
		}
	}

	if err := database.InitDatabase(global.DBFilesDir); err != nil {
		return fmt.Errorf("InitDatabase: %v", err)
	}

	return nil
}

func AfterStop(ctx context.Context) error {
	c := config.Get()
	c.Log.Close()

	return nil
}
