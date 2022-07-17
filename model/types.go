package model

import "time"

type Activity struct {
	ShortName   string
	Description string
	Duration    time.Duration
}
