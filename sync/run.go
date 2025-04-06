package sync

import (
	"context"
	"time"

	"github.com/minnowo/traggo-sync/client"
	"github.com/rs/zerolog/log"
)

type TagSyncMode int

const (
	TAG_SYNC_ALL     TagSyncMode = iota
	TAG_SYNC_MISSING TagSyncMode = iota
)

var (
	MIN_TIME time.Time = time.Unix(0, 0)
	MAX_TIME time.Time = time.Date(9999, 12, 31, 23, 59, 59, int(time.Nanosecond*999999999), time.UTC)
)

func SyncFrom(run RunConfig) error {

	if err := run.Validate(); err != nil {
		return err
	}

	var src *client.TraggoHttpClient
	var dst []*client.TraggoHttpClient

	src, err := client.NewTraggoClient(run.From.Host)

	if err != nil {
		return err
	}
	if err := src.LogVersion(); err != nil {
		return err
	}
	if err := src.Login(run.From.User, run.From.Pass); err != nil {
		return err
	}

	dst = make([]*client.TraggoHttpClient, len(run.To))
	for i := 0; i < len(dst); i++ {

		t := run.To[i]
		c, err := client.NewTraggoClient(t.Host)
		if err != nil {
			return err
		}
		if err := c.LogVersion(); err != nil {
			return err
		}
		if err := c.Login(t.User, t.Pass); err != nil {
			return err
		}

		dst[i] = c
	}

	for i := 0; i < len(dst); i++ {

		syncer := Syncer{
			ctx:    context.Background(),
			F:      src,
			T:      dst[i],
			DryRun: run.DryRun,
		}

		if run.AllTags || run.MissingTags {

			var tagSync TagSyncMode

			if run.AllTags {
				tagSync = TAG_SYNC_ALL
			}

			if run.MissingTags {
				tagSync = TAG_SYNC_MISSING
			}

			err = syncer.SyncTags(tagSync)

			if err != nil {
				log.Error().Err(err).Msg("error while syncing tags")
			}
		}

		if run.MissingTimespans {

			start := run.StartTime
			end := run.EndTime

			err = syncer.SyncTimespansInRange(start, end)
			if err != nil {
				log.Error().Err(err).Msg("error while syncing timespans")
			}
		}
	}

	return nil
}
