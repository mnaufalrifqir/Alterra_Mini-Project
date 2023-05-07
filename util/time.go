package util

import "time"

func DeadlineRange(deadline int) (time.Time, time.Time) {
	deadlineTime := time.Now().AddDate(0, 0, deadline)
	return time.Now(), deadlineTime
}
