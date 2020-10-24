package services

import (
	"fmt"

	"test-time-tracker/data"
	"test-time-tracker/models"
)

// ErrorTimeTrack is a base for any tracker related error
var ErrorTimeTrack = fmt.Errorf("tracker error")

// Track
// uname: User Name
// track event from the clock
func Track(uname string) (e error) {
	s, e := getUserState(uname)
	if e != nil {
		return
	}
	switch s {
	case models.Absent, models.Sick, models.OOO, models.ReserveDuty:
		e = setUserState(uname, models.Working)
	case models.Working:
		e = setUserState(uname, models.Absent)
	default:
		e = fmt.Errorf("wrong state %s cannot be changed for %s, %w", s, uname, ErrorTimeTrack)
	}
	return
}

// ThisMonth
// uname: User Name
// get track events since beginning of month
func ThisMonth(uname string) (d interface{}, err error) {
	d, err = data.ThisMonth(uname)
	return
}

// SetImplicitState
// uname: User Name
// from the HR manager
func SetImplicitState(uname string, state models.UserState) (e error) {
	e = setUserState(uname, state)
	return
}

// GetDaysBack
// uname: User Name
// days: day count
//
func GetDaysBack(uname string, days int) (d interface{}, err error) {
	d, err = data.GetDaysBack(uname, days)
	return
}

// GetMonthsBack
// uname: User Name
// days: month count
func GetMonthsBack(uname string, months int) (d interface{}, err error) {
	d, err = data.GetMonthsBack(uname, months)
	return
}

// ////////////////////////////////////////////////////////////////////////////////////////
func setUserState(uname string, state models.UserState) (e error) {
	e = data.SetUserState(uname, state)
	return
}

func getUserState(uname string) (state models.UserState, e error) {
	state, e = data.GetUserState(uname)
	return
}
