package sync

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

type RunConfig struct {
	DryRun           bool
	MissingTags      bool
	AllTags          bool
	MissingTimespans bool
	ReplaceTimespans bool
	DeleteTimespans  bool
	Dashboards       bool
	StartTime        time.Time
	EndTime          time.Time
	From             TraggoServer
	To               []TraggoServer
}

func (rc *RunConfig) Validate() error {

	check := make(map[string]bool)

	log.Info().
		Str("host", rc.From.Host).
		Str("user", rc.From.User).
		Str("pass", rc.From.Pass).
		Msg("source")
	if err := rc.From.CheckValid(); err != nil {
		return err
	}

	check[rc.From.Host] = true

	for i := 0; i < len(rc.To); i++ {
		log.Info().
			Str("user", rc.To[i].User).
			Str("pass", rc.To[i].Pass).
			Str("host", rc.To[i].Host).
			Msg("target")
		if err := rc.To[i].CheckValid(); err != nil {
			return err
		}

		_, ok := check[rc.To[i].Host]
		if !ok {
			check[rc.From.Host] = true
		} else {
			return fmt.Errorf("duplicate Traggo given: %s", rc.To[i].Host)
		}

	}

	if rc.StartTime.After(rc.EndTime) {
		return fmt.Errorf("Start time range is after the end time range")
	}

	return nil
}
