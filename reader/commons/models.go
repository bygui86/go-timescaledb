package commons

import (
	"fmt"
	"time"
)

type QueryParameters struct {
	Location  string
	StartTime time.Time // unix-nano
	EndTime   time.Time // unix-nano
	Limit     int64
}

func (p *QueryParameters) String() string {
	return fmt.Sprintf("Location[%s], Start[%v], End[%v], Limit[%d]",
		p.Location, p.StartTime, p.EndTime, p.Limit)
}

type Condition struct {
	Time        time.Time `json:"time"`
	Location    string    `json:"location"`
	Temperature float64   `json:"temperature"`
}

func (c *Condition) String() string {
	return fmt.Sprintf("Time[%v], Location[%s], Temperature[%f]",
		c.Time, c.Location, c.Temperature)
}
