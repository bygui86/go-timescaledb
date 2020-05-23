package commons

import (
	"fmt"
	"time"
)

type Condition struct {
	Time        time.Time
	Location    string
	Temperature float64
}

func (c *Condition) String() string {
	return fmt.Sprintf("Time[%v], Location[%s], Temperature[%f]",
		c.Time, c.Location, c.Temperature)
}
