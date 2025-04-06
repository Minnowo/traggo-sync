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

type RelativeOrStaticRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type InputRelativeOrStaticRange RelativeOrStaticRange

type NamedDateRange struct {
	Id       int                   `json:"id"`
	Name     string                `json:"name"`
	Editable bool                  `json:"editable"`
	Range    RelativeOrStaticRange `json:"range"`
}

func (in *NamedDateRange) ToInputNamedDateRange() InputNamedDateRange {
	return InputNamedDateRange{
		Name:     in.Name,
		Editable: in.Editable,
		Range:    in.Range,
	}
}

type InputNamedDateRange struct {
	Name     string                `json:"name"`
	Editable bool                  `json:"editable"`
	Range    RelativeOrStaticRange `json:"range"`
}

type Range struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type StatsInterval string

const (
	StatsIntervalHourly  StatsInterval = "Hourly"
	StatsIntervalDaily   StatsInterval = "Daily"
	StatsIntervalWeekly  StatsInterval = "Weekly"
	StatsIntervalMonthly StatsInterval = "Monthly"
	StatsIntervalYearly  StatsInterval = "Yearly"
	StatsIntervalSingle  StatsInterval = "Single"
)

type StatsSelection struct {
	Interval    StatsInterval              `json:"interval"`
	Tags        []string                   `json:"tags"`
	ExcludeTags []TimeSpanTag              `json:"excludeTags"`
	IncludeTags []TimeSpanTag              `json:"includeTags"`
	RangeId     int                        `json:"rangeId"`
	Range       InputRelativeOrStaticRange `json:"range"`
}

type InputStatsSelection StatsSelection

type DashboardEntryPos struct {
	W    int `json:"w"`
	H    int `json:"h"`
	X    int `json:"x"`
	Y    int `json:"y"`
	MinW int `json:"minW"`
	MinH int `json:"minH"`
}

func (in *DashboardEntryPos) ToDashboardEntryPos() InputDashboardEntryPos {
	return InputDashboardEntryPos{
		W: in.W,
		H: in.H,
		X: in.X,
		Y: in.Y,
	}
}

type InputDashboardEntryPos struct {
	W int `json:"w"`
	H int `json:"h"`
	X int `json:"x"`
	Y int `json:"y"`
}

type ResponsiveDashboardEntryPos struct {
	Desktop DashboardEntryPos `json:"desktop"`
	Mobile  DashboardEntryPos `json:"mobile"`
}

func (in *ResponsiveDashboardEntryPos) ToInputResponsiveDashboardEntryPos() InputResponsiveDashboardEntryPos {
	return InputResponsiveDashboardEntryPos{
		Desktop: InputDashboardEntryPos{
			W: in.Desktop.W,
			H: in.Desktop.H,
			X: in.Desktop.X,
			Y: in.Desktop.Y,
		},
		Mobile: InputDashboardEntryPos{
			W: in.Mobile.W,
			H: in.Mobile.H,
			X: in.Mobile.X,
			Y: in.Mobile.Y,
		},
	}
}

type InputResponsiveDashboardEntryPos struct {
	Desktop InputDashboardEntryPos `json:"desktop"`
	Mobile  InputDashboardEntryPos `json:"mobile"`
}
