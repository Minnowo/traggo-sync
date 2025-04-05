package sync

import (
	"github.com/minnowo/traggo-sync/graph"
	"github.com/rs/zerolog/log"
)

func TagsToMap(tags []graph.GetAllTagsTagsTagDefinition) map[string]int {

	m := make(map[string]int, len(tags))

	for i, t := range tags {
		m[t.Key] = i
	}
	return m
}

func (s *Syncer) SyncTags() error {

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

	tagsCreated := 0
	tTags, err := graph.GetAllTags(s.ctx, s.T.Ql())

	if err != nil {
		return err
	}

	tTagMap := TagsToMap(tTags.Tags)

	for _, tag := range fTags.Tags {

		if _, ok := tTagMap[tag.Key]; ok {
			continue
		}

		log.Info().Str("key", tag.Key).Str("color", tag.Color).Msg("Creating tag")

		_, err := graph.CreateTag(s.ctx, s.T.Ql(), tag.Key, tag.Color)

		if err != nil {
			log.Warn().Err(err).Msg("Could not create tag")
		} else {
			tagsCreated++
		}
	}

	s.Stop()
	s.Info().
		Int("tagsCreated", tagsCreated).
		Str("took", s.duration.String()).
		Msg("Syncing tags completed")

	return nil
}
