package worker_reloader

import (
	"context"
	"time"

	"github.com/vladimirok5959/golang-ip2location/internal/client"
	"github.com/vladimirok5959/golang-utils/utils/http/logger"
	"github.com/vladimirok5959/golang-worker/worker"
)

var Delay = 60 * time.Minute

func New(cl *client.Client) *worker.Worker {
	time.Sleep(1000 * time.Millisecond)
	return worker.New(func(ctx context.Context, w *worker.Worker, o *[]worker.Iface) {
		if cl, ok := (*o)[0].(*client.Client); ok {
			Run(ctx, cl)
		}
		select {
		case <-ctx.Done():
		case <-time.After(Delay):
			return
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
		logger.LogInfo("worker reloader: done\n")
	} else {
		logger.LogInfo("worker reloader: done with error\n")
	}
}
