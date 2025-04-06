package sync

import (
	"github.com/minnowo/traggo-sync/graph"
	"github.com/rs/zerolog/log"
)

type TagSyncMode int

const (
	TAG_SYNC_ALL     TagSyncMode = iota
	TAG_SYNC_MISSING TagSyncMode = iota
)

func TagsToMap(tags []graph.GetAllTagsTagsTagDefinition) map[string]int {

	m := make(map[string]int, len(tags))

	for i, t := range tags {
		m[t.Key] = i
	}
	return m
}

func (s *Syncer) SyncTags(syncMethod TagSyncMode) error {

	s.Info().Msg("Syncting tags")
	s.Start()

	var fTags *graph.GetAllTagsResponse

	fTags, err := graph.GetAllTags(s.ctx, s.F.Ql())

	if err != nil {
		return err
	}

	log.Info().
		Str("traggo", s.T.Url()).
		Msg("Syncing tags")

	created := 0
	updated := 0
	skipped := 0
	failed := 0

	tTags, err := graph.GetAllTags(s.ctx, s.T.Ql())

	if err != nil {
		return err
	}

	tTagMap := TagsToMap(tTags.Tags)

	for _, tag := range fTags.Tags {

		doUpdate := false

		if i, ok := tTagMap[tag.Key]; ok {

			if syncMethod != TAG_SYNC_ALL {
				skipped++
				continue
			}

			if o := tTags.Tags[i]; o.Color == tag.Color {
				skipped++
				continue
			}

			doUpdate = true
		}

		var err error

		if doUpdate {
			log.Info().Str("key", tag.Key).Str("color", tag.Color).Msg("Updating tag")
			if s.DryRun {
				err = nil
			} else {
				_, err = graph.UpdateTag(s.ctx, s.T.Ql(), tag.Key, tag.Color)
			}
		} else {
			log.Info().Str("key", tag.Key).Str("color", tag.Color).Msg("Creating tag")
			if s.DryRun {
				err = nil
			} else {
				_, err = graph.CreateTag(s.ctx, s.T.Ql(), tag.Key, tag.Color)
			}
		}

		if err != nil {
			log.Warn().Err(err).Msg("Could not create tag")
			failed++
		} else if doUpdate {
			updated++
		} else {
			created++
		}
	}

	s.Stop()
	s.Info().
		Int("created", created).
		Int("failed", failed).
		Int("updated", updated).
		Int("skipped", skipped).
		Str("took", s.duration.String()).
		Msg("Syncing tags completed")

	return nil
}
