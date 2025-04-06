package client

import (
	"context"
	"time"

	"github.com/minnowo/traggo-sync/graph"
	"github.com/rs/zerolog/log"
)

func (c *TraggoHttpClient) GetBulkTimeSpans(ctx context.Context, from, to time.Time) ([]graph.TimeSpan, error) {

	res, err := graph.GetTimeSpans(ctx, c.Ql(), from, to)

	if err != nil {
		return nil, err
	}

	if !res.TimeSpans.Cursor.HasMore {
		return res.TimeSpans.TimeSpans, nil
	}

	var timeSpans []graph.TimeSpan
	timeSpans = res.TimeSpans.TimeSpans

	var cursor graph.InputCursor
	cursor.Offset = res.TimeSpans.Cursor.Offset
	cursor.PageSize = res.TimeSpans.Cursor.PageSize
	cursor.StartId = res.TimeSpans.Cursor.StartId

	for {

		res2, err := graph.GetPagedTimeSpans(ctx, c.Ql(), from, to, cursor)

		if err != nil {
			log.Error().Err(err).Msg("")
			break
		}

		n := len(res2.TimeSpans.TimeSpans)

		if n > 0 {
			log.Debug().Int("n", n).Msg("got new page of timespans")
			timeSpans = append(timeSpans, res2.TimeSpans.TimeSpans...)
		}

		cursor.Offset = res2.TimeSpans.Cursor.Offset
		cursor.PageSize = res2.TimeSpans.Cursor.PageSize
		cursor.StartId = res2.TimeSpans.Cursor.StartId

		if !res2.TimeSpans.Cursor.HasMore {
			break
		}
	}

	return timeSpans, err
}
