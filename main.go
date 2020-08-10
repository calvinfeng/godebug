package main

import (
	"context"
	"godebug/spammer"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/gops/agent"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour))
	defer cancel()

	s := spammer.New()
	go s.Spam(ctx)

	// Setup gops
	go func() {
		if err := agent.Listen(agent.Options{}); err != nil {
			log.Fatal(err)
		}
	}()

	logrus.Infof("process ID %d", os.Getpid())
	logrus.Info("now serving and listening on port 1234")
	http.ListenAndServe(":1234", nil)
}
