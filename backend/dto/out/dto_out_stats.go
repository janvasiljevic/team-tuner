package out

import (
	"time"
)

// This really shouldnt exist here
// Perhaps we can have a Store layer that handles all the database stuff
type RowReturn struct {
	Day   time.Time `json:"day" validate:"required"`
	Count int       `json:"count" validate:"required"`
}

type DailyStat struct {
	Day   string `json:"day" validate:"required"`
	Count int    `json:"count" validate:"required"`
}

type DailyActivity struct {
	Activity  []DailyStat `json:"activity" validate:"required"`
	StartTime time.Time   `json:"startTime" validate:"required"`
	EndTime   time.Time   `json:"endTime" validate:"required"`
}

type QuestionerCompletionStats struct {
	TotalCount int `json:"totalCount" validate:"required"`
	Completed  int `json:"completed" validate:"required"`
	Percent    int `json:"percent" validate:"required"`
}

func NewDailyActivityFromRows(rows []RowReturn, startTime time.Time, endTime time.Time) DailyActivity {
	// Create a map for quick lookup of count by date
	rowMap := make(map[string]int)
	for _, row := range rows {
		rowMap[row.Day.Format("2006-01-02")] = row.Count
	}

	// Create a slice of DailyStat for each day between start and end times
	var activity []DailyStat

	for d := startTime; d.Before(endTime) || d.Equal(endTime); d = d.AddDate(0, 0, 1) {
		day := d.Format("2006-01-02")
		count := rowMap[day]
		activity = append(activity, DailyStat{Day: day, Count: count})
	}

	return DailyActivity{
		Activity:  activity,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func NewQuestionerCompletionStatsFromRows(total, completed int) QuestionerCompletionStats {
	return QuestionerCompletionStats{
		TotalCount: total,
		Completed:  completed,
		Percent:    int(float64(completed) / float64(total) * 100),
	}
}
