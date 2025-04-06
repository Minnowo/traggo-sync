package sync

import (
	"github.com/minnowo/traggo-sync/graph"
	"github.com/rs/zerolog/log"
)

func (s *Syncer) SyncDashboards() error {

	s.Info().
		Msg("Syncing dashboards")
	s.Start()

	res, err := graph.GetDashboards(s.ctx, s.F.Ql())

	if err != nil {
		return err
	}

	for _, dash := range res.Dashboards {

		log.Info().Str("name", dash.Name).Msg("creating dashboard")
		cRes, err := graph.CreateDashboard(s.ctx, s.T.Ql(), dash.Name)

		if err != nil {
			return err
		}

		rangeMap := make(map[int]int)

		for _, rang := range dash.Ranges {

			log.Info().Str("name", rang.Name).Msg("adding range")
			rangeRes, err := graph.AddDashboardRange(s.ctx, s.T.Ql(), cRes.CreateDashboard.Id, rang.ToInputNamedDateRange())

			if err != nil {
				return err
			}

			rangeMap[rang.Id] = rangeRes.AddDashboardRange.Id
		}

		for _, entry := range dash.Items {

			log.Info().Str("name", entry.Title).Msg("adding entry")
			sel := graph.InputStatsSelection(entry.StatsSelection)
			sel.RangeId = rangeMap[sel.RangeId]

			_, err := graph.AddDashboardEntry(s.ctx, s.T.Ql(),
				cRes.CreateDashboard.Id,
				entry.EntryType,
				entry.Title,
				entry.Total,
				sel,
				entry.Pos.ToInputResponsiveDashboardEntryPos(),
			)
			if err != nil {
				return err
			}
		}

	}

	s.Stop()
	s.Info().
		Str("took", s.duration.String()).
		Msg("Finished syncing dashboards")
	return nil
}
