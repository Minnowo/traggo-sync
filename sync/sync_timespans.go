package sync

import (
	"time"

	"github.com/minnowo/traggo-sync/graph"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type TimeSpanSyncMode int

const (
	TIMESPAN_SYNC_MISSING TimeSpanSyncMode = iota
	TIMESPAN_SYNC_REPLACE TimeSpanSyncMode = iota
	TIMESPAN_SYNC_DELETE  TimeSpanSyncMode = iota
)

func TsCmp(t, o *graph.TimeSpan, tagsO map[string]string) bool {
	if t.Start != o.Start || t.End != o.End || t.Note != o.Note {
		return false
	}
	if len(t.Tags) != len(o.Tags) {
		return false
	}
	for _, v := range t.Tags {
		ov, ok := tagsO[v.Key]
		if !ok || v.Value != ov {
			return false
		}
	}
	return true
}

func MakeTagMap(ts []graph.TimeSpanTag) map[string]string {
	tags := make(map[string]string)
	for _, tag := range ts {
		tags[tag.Key] = tag.Value
	}
	return tags
}

func logTs(ts *graph.TimeSpan) *zerolog.Event {

	arr := zerolog.Arr()
	for _, tTags := range ts.Tags {
		arr.Str(tTags.Key + ":" + tTags.Value)
	}
	return log.Info().
		Str("start", ts.Start.String()).
		Str("end", ts.End.String()).
		Array("tags", arr)
}

func (s *Syncer) DeleteTimespansInRange(from, to time.Time) error {

	s.Info().
		Str("start", from.String()).
		Str("stop", to.String()).
		Msg("Deleting from target")
	s.Start()

	toTimeSpans, err := s.T.GetBulkTimeSpans(s.ctx, from, to)

	if err != nil {
		return err
	}

	deleted := 0
	failed := 0

	for _, o := range toTimeSpans {

		logTs(&o).Msg("deleting timespan")

		if s.DryRun {
			err = nil
		} else {
			_, err = graph.RemoveTimeSpan(s.ctx, s.T.Ql(), o.Id)
		}

		if err != nil {
			failed++
			log.Warn().Err(err).Msg("error deleting timespan")
		} else {
			deleted++
		}
	}

	s.Stop()
	s.Info().
		Str("took", s.duration.String()).
		Int("deleted", deleted).
		Int("failed", failed).
		Msg("Deleting timespans finished")

	return nil
}

func (s *Syncer) SyncTimespansInRange(mode TimeSpanSyncMode, from, to time.Time) error {

	var toTimeSpans []graph.TimeSpan
	var err error

	switch mode {
	case TIMESPAN_SYNC_MISSING:
		toTimeSpans, err = s.T.GetBulkTimeSpans(s.ctx, from, to)
		break
	case TIMESPAN_SYNC_REPLACE:
		err = s.DeleteTimespansInRange(from, to)
		toTimeSpans = make([]graph.TimeSpan, 0)
		break
	case TIMESPAN_SYNC_DELETE:
		return s.DeleteTimespansInRange(from, to)
	default:
		panic("unreachable")
	}

	if err != nil {
		return err
	}

	s.Info().
		Str("start", from.String()).
		Str("stop", to.String()).
		Msg("Syncing timespans")
	s.Start()

	var fromTimeSpans []graph.TimeSpan
	fromTimeSpans, err = s.F.GetBulkTimeSpans(s.ctx, from, to)

	if err != nil {
		return err
	}

	log.Info().Int("count", len(fromTimeSpans)).Msg("Found source timespans")
	log.Info().Int("count", len(toTimeSpans)).Msg("Found target timespans")

	failed := 0
	created := 0
	skipped := 0

	toTagMap := make([]map[string]string, len(toTimeSpans))

	for i, o := range toTimeSpans {
		toTagMap[i] = MakeTagMap(o.Tags)
	}

	for _, fromTs := range fromTimeSpans {

		shouldSync := true

		for i, toTs := range toTimeSpans {

			if TsCmp(&fromTs, &toTs, toTagMap[i]) {
				shouldSync = false
				break
			}
		}

		if !shouldSync {
			skipped++
			continue
		}

		logTs(&fromTs).Msg("creating timespan")

		var err error

		if s.DryRun {
			err = nil
		} else {
			_, err = graph.CreateTimeSpan(s.ctx, s.T.Ql(), fromTs.Start, fromTs.End, fromTs.Tags, fromTs.Note)
		}

		if err != nil {
			failed++
			log.Warn().Err(err).Msg("error creating timespan")
		} else {
			created++
		}
	}

	s.Stop()
	s.Info().
		Str("took", s.duration.String()).
		Int("created", created).
		Int("skipped", skipped).
		Int("failed", failed).
		Msg("Syncing timespans finished")

	return nil
}
