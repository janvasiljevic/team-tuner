package in

import (
	"time"
)

type GetDailyActivityRequest struct {
	StartTime *time.Time `json:"startTime" query:"startTime"`
	EndTime   *time.Time `json:"endTime" query:"endTime"`
}
