package graph

import "time"

type TimeSpan struct {
	Id       int           `json:"id"`
	Start    time.Time     `json:"start"`
	End      time.Time     `json:"end"`
	OldStart time.Time     `json:"oldStart"`
	Note     string        `json:"note"`
	Tags     []TimeSpanTag `json:"tags"`
}

type TimeSpanTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
