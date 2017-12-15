package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second              = 8
	Third               = 15
	Fourth              = 22
	Last                = -6
	Teenth              = 13
)

func Day(w WeekSchedule, ww time.Weekday, m time.Month, y int) int {

	if w == Last {
		m++
	}
	t := time.Date(y, m, int(w), 12, 0, 0, 0, time.UTC)
	return t.Day() + int(ww-t.Weekday()+7)%7

}
