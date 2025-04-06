package sync

import (
	"context"
	"time"

	"github.com/minnowo/traggo-sync/client"
	"github.com/rs/zerolog/log"
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

			var mode TagSyncMode

			if run.AllTags {
				mode = TAG_SYNC_ALL
			}

			if run.MissingTags {
				mode = TAG_SYNC_MISSING
			}

			err = syncer.SyncTags(mode)

			if err != nil {
				log.Error().Err(err).Msg("error while syncing tags")
			}
		}

		if run.MissingTimespans || run.ReplaceTimespans {

			var mode TimeSpanSyncMode

			if run.ReplaceTimespans {
				mode = TIMESPAN_SYNC_REPLACE
			}
			if run.MissingTimespans {
				mode = TIMESPAN_SYNC_MISSING
			}

			start := run.StartTime
			end := run.EndTime

			err = syncer.SyncTimespansInRange(mode, start, end)

			if err != nil {
				log.Error().Err(err).Msg("error while syncing timespans")
			}

		}
		if run.Dashboards {

			err = syncer.SyncDashboards()

			if err != nil {
				log.Error().Err(err).Msg("error while syncing dashboards")
			}
		}
	}

	return nil
}
