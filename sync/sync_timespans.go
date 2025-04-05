package sync

import (
	"time"

	"github.com/minnowo/traggo-sync/graph"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ComparableTimeSpan struct {
	Start time.Time
	End   time.Time
	Note  string
	Tags  map[string]string
}

func ComparableTimeSpanFromGQL(ts graph.TimeSpan) ComparableTimeSpan {
	tags := make(map[string]string)
	for _, tag := range ts.Tags {
		tags[tag.Key] = tag.Value
	}
	return ComparableTimeSpan{
		Start: ts.Start,
		End:   ts.End,
		Note:  ts.Note,
		Tags:  tags,
	}
}
func (t *ComparableTimeSpan) Eq(o *ComparableTimeSpan) bool {
	if !(t.Start == o.Start && t.End == o.End && t.Note == o.Note) {
		return false
	}
	if len(t.Tags) != len(o.Tags) {
		return false
	}
	for k, v := range t.Tags {
		ov, ok := o.Tags[k]

		if !ok || v != ov {
			return false
		}
	}
	return true
}

func (s *Syncer) SyncTimespansInRange(from, to time.Time) error {

	s.Info().
		Str("start", from.String()).
		Str("stop", to.String()).
		Msg("Syncing timespans")
	s.Start()

	var fTimes *graph.GetTimeSpansResponse

	fTimes, err := graph.GetTimeSpans(s.ctx, s.F.Ql(), from, to)

	if err != nil {
		return err
	}

	tTimes, err := graph.GetTimeSpans(s.ctx, s.T.Ql(), from, to)

	if err != nil {
		return err
	}

	oCompObjs := make([]ComparableTimeSpan, len(tTimes.TimeSpans.TimeSpans))
	for j, t := range fTimes.TimeSpans.TimeSpans {

		tComp := ComparableTimeSpanFromGQL(t)
		shouldSync := true

		for i, o := range tTimes.TimeSpans.TimeSpans {

			if j == 0 {
				oCompObjs[i] = ComparableTimeSpanFromGQL(o)
			}

			if oCompObjs[i].Eq(&tComp) {
				shouldSync = false
				break
			}
		}

		if !shouldSync {
			continue
		}

		arr := zerolog.Arr()
		for _, tTags := range t.Tags {
			arr.Str(tTags.Key + tTags.Value)
		}
		log.Info().
			Str("start", t.Start.String()).
			Str("end", t.End.String()).
			Array("tags", arr).
			Msg("")

		_, err := graph.CreateTimeSpan(s.ctx, s.T.Ql(), t.Start, t.End, t.Tags, t.Note)
		if err != nil {
			log.Warn().Err(err).Msg("error creating timespan")
		}
	}

	s.Stop()
	s.Info().
		Str("took", s.duration.String()).
		Msg("Syncing timespans finished")

	return nil
}
