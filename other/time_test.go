package other

import (
	"testing"
	"time"
)

func Test_Time(t *testing.T) {
	now := time.Now()
	t.Log(now)

	local := now.Local()
	t.Log(local)
	t.Logf("year = %d", local.Year())
	t.Logf("month = %d", local.Month())
	t.Logf("day = %d", local.Day())
	t.Logf("hour = %d", local.Hour())
	t.Logf("minute = %d", local.Minute())
	t.Logf("second = %d", local.Second())
	t.Logf("weekday = %v", local.Weekday())
	t.Logf("sub = %v", now.Sub(local))
}

func Test_Time_Format(t *testing.T) {
	now := time.Now()
	t.Logf("%v", now.Format(time.RFC3339))
	// 2023-03-05T15:39:40+08:00

	t.Logf("%v", now.Format("2006-01-02 15:04:05"))
	// 2023-03-05 15:39:40
	//t.Logf("%v", now.Format(time.DateTime))
}
