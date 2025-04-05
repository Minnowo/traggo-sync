package sync

import (
	"context"
	"time"

	"github.com/minnowo/traggo-sync/client"
)

type SyncItem int

const (
	TAG_NOT_EXISTS SyncItem = iota
	TAG_UPDATE     SyncItem = iota
	TIMESPANS      SyncItem = iota
)

func SyncFrom(run RunConfig) error {

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
			ctx: context.Background(),
			F:   src,
			T:   dst[i],
		}
		syncer.SyncTags()

		est, err := time.LoadLocation("America/New_York")
		if err != nil {
			return err
		}
		start := time.Date(2024, time.December, 30, 16, 9, 53, 0, est)
		end := time.Date(2024, time.December, 30, 21, 19, 55, 0, est)

		syncer.SyncTimespansInRange(start, end)
	}

	return nil
}
