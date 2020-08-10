package spammer

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type Spammer interface {
	Spam(context.Context)
	KillOne()
}

func New() Spammer {
	return &spammer{
		signal: make(chan struct{}),
	}
}

type spammer struct {
	signal chan struct{}
}

func (s *spammer) Spam(ctx context.Context) {
	for range time.NewTicker(1000 * time.Millisecond).C {
		logrus.Warn("generate new spam goroutine")
		go s.subroutine(ctx)
	}
}

func (s *spammer) KillOne() {
	s.signal <- struct{}{}
}

func (s *spammer) subroutine(ctx context.Context) {
	memFootprint := make([]int64, 10000000)
	i := 0
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-s.signal:
			return
		case <-ticker.C:
			memFootprint[i] = 1
			i++
		}
	}
}
