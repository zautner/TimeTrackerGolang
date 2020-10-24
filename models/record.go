package models

import (
	"time"
)

type (
	TrackerRecord struct {
		Timestamp time.Time
		State     UserState
	}
)

// NewTrackerRecord
// Create new
func NewTrackerRecord(s UserState) TrackerRecord {
	return TrackerRecord{
		Timestamp: time.Now(),
		State:     s,
	}
}
