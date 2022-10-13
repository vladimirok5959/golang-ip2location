package worker_reloader

import (
	"context"
	"time"

	"github.com/vladimirok5959/golang-ip2location/internal/client"
	"github.com/vladimirok5959/golang-ip2location/internal/consts"
	"github.com/vladimirok5959/golang-utils/utils/http/logger"
	"github.com/vladimirok5959/golang-worker/worker"
)

func New(cl *client.Client) *worker.Worker {
	return worker.New(func(ctx context.Context, w *worker.Worker, o *[]worker.Iface) {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(consts.Config.DbUpdateTime) * time.Minute):
		}
		if cl, ok := (*o)[0].(*client.Client); ok {
			Run(ctx, cl)
		}
	}, &[]worker.Iface{
		cl,
	})
}

func Run(ctx context.Context, cl *client.Client) {
	logger.LogInfo("worker reloader: trying to reload database\n")

	var err error
	if err = cl.ReloadDatabase(ctx); err != nil {
		logger.LogInternalError(err)
	}

	if err == nil {
		logger.LogInfo("worker reloader: reloading done\n")
	} else {
		logger.LogInfo("worker reloader: reloading done with errors\n")
	}
}
