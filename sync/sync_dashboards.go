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

		var err error
		var cRes *graph.CreateDashboardResponse

		if s.DryRun {
			err = nil
			cRes = &graph.CreateDashboardResponse{CreateDashboard: graph.CreateDashboardCreateDashboard{Id: dash.Id}}
		} else {
			cRes, err = graph.CreateDashboard(s.ctx, s.T.Ql(), dash.Name)
		}

		if err != nil {
			log.Warn().Err(err).Msg("could not create dashboard")
			continue
		}

		rangeMap := make(map[int]int)

		for _, rang := range dash.Ranges {

			log.Info().Str("name", rang.Name).Msg("adding range")

			var err error
			var rangeRes *graph.AddDashboardRangeResponse

			if s.DryRun {
				err = nil
				rangeRes = &graph.AddDashboardRangeResponse{AddDashboardRange: graph.NamedDateRange{Id: rang.Id}}
			} else {
				rangeRes, err = graph.AddDashboardRange(s.ctx, s.T.Ql(), cRes.CreateDashboard.Id, rang.ToInputNamedDateRange())
			}

			if err != nil {
                log.Warn().Err(err).Msg("error creating range")
                continue
			}

			rangeMap[rang.Id] = rangeRes.AddDashboardRange.Id
		}

		for _, entry := range dash.Items {

			log.Info().Str("name", entry.Title).Msg("adding entry")

			var err error

			if s.DryRun {
				err = nil
			} else {
                sel := graph.InputStatsSelection(entry.StatsSelection)
                id, ok := rangeMap[sel.RangeId]

                if !ok {
                    log.Warn().Msg("Could not find range id for this dashboard entry, skippping...")
                    continue
                }
                sel.RangeId = id

				_, err = graph.AddDashboardEntry(s.ctx, s.T.Ql(),
					cRes.CreateDashboard.Id,
					entry.EntryType,
					entry.Title,
					entry.Total,
					sel,
					entry.Pos.ToInputResponsiveDashboardEntryPos(),
				)
			}

			if err != nil {
                log.Warn().Err(err).Msg("error creating dashboard entry")
			}
		}

	}

	s.Stop()
	s.Info().
		Str("took", s.duration.String()).
		Msg("Finished syncing dashboards")
	return nil
}
