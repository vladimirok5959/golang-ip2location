package main

import (
	"context"

	"github.com/vladimirok5959/golang-ctrlc/ctrlc"
	"github.com/vladimirok5959/golang-ip2location/internal/client"
	"github.com/vladimirok5959/golang-ip2location/internal/consts"
	"github.com/vladimirok5959/golang-ip2location/internal/server"
	"github.com/vladimirok5959/golang-ip2location/internal/workers/worker_reloader"
	"github.com/vladimirok5959/golang-utils/utils/http/logger"
	"github.com/vladimirok5959/golang-utils/utils/penv"
)

func init() {
	if err := penv.ProcessConfig(&consts.Config); err != nil {
		panic(err)
	}

	var err error
	if consts.Config.DataDir == "" {
		consts.Config.DataDir, err = consts.DataPath()
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	logger.AccessLogFile = consts.Config.AccessLogFile
	logger.ErrorLogFile = consts.Config.ErrorLogFile

	ctrlc.App(func(ctx context.Context, shutdown context.CancelFunc) *[]ctrlc.Iface {
		cl, err := client.New(ctx, shutdown)
		if err != nil {
			return ctrlc.MakeError(shutdown, ctrlc.AppError(err))
		}

		sv, err := server.New(ctx, shutdown, cl)
		if err != nil {
			return ctrlc.MakeError(shutdown, ctrlc.AppError(err), cl)
		}

		workerReloader := worker_reloader.New(cl)

		return &[]ctrlc.Iface{
			workerReloader,
			sv,
			cl,
		}
	})
}
