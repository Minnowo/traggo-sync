package sync

import (
	"context"
	"time"

	"github.com/minnowo/traggo-sync/client"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Syncer struct {
	ctx      context.Context
	F        *client.TraggoHttpClient
	T        *client.TraggoHttpClient
	start    time.Time
	duration time.Duration
}

func (s *Syncer) Start() {
	s.start = time.Now()
}
func (s *Syncer) Stop() {
	s.duration += time.Now().Sub(s.start)
}
func (s *Syncer) Duration() time.Duration {
	return s.duration
}
func (s *Syncer) Info() *zerolog.Event {
	return s.Log(zerolog.InfoLevel)
}

func (s *Syncer) Log(l zerolog.Level) *zerolog.Event {
	return log.WithLevel(l).
		Str("from", s.F.Url()).
		Str("to", s.T.Url())
}
