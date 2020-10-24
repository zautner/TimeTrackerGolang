package controller

import (
	"strconv"

	"test-time-tracker/models"
	"test-time-tracker/services"
)

func Track(uname string) {
	validateCall(services.Track(uname))
}

func ThisMonth(uname string) (data interface{}) {
	var err error
	data, err = services.ThisMonth(uname)
	validateCall(err, data)
	return
}

func GetDaysBack(uname string, back string) (data interface{}) {
	days, err := strconv.Atoi(back)
	validateCall(days, err)
	data, err = services.GetDaysBack(uname, days)
	validateCall(data, err)
	return
}

func GetMonthsBack(uname string, back string) (data interface{}) {
	mo, err := strconv.Atoi(back)
	validateCall(mo, err)
	data, err = services.GetMonthsBack(uname, mo)
	validateCall(data, err)

	return
}

func ResetForUser(uname string) {
	validateCall(services.SetImplicitState(uname, models.Absent))
}
