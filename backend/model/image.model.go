package model

import (
	"time"
)

type Image struct {
	Url      string
	Time     time.Time
	Unsigned string
	Lon      float64
	Lat      float64
}
