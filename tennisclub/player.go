package tennisclub

import "time"

type Player struct {
	ID                               string
	FirstName, LastName, Nationality string
	Birthday                         Birthday
}

type Birthday struct {
	Day   int
	Month time.Month
	Year  int
}
