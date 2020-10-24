package data

import (
	"fmt"
	"time"

	"test-time-tracker/models"
)

func GetUserState(uname string) (state models.UserState, err error) {
	anyChecker()
	tm.RLock()
	defer tm.RUnlock()
	if records := trackerData.trackerRecords[uname]; records == nil {
		err = fmt.Errorf("`%s` user not found, %w", uname, ErrorDataLayer)
	} else {
		if l := len(*records); l < 1 {
			err = fmt.Errorf("%s empty storage, %w", uname, ErrorDataLayer)
		} else {
			state = ((*records)[len(*records)-1]).State
		}
	}
	return
}

func SetUserState(uname string, state models.UserState) (err error) {
	anyChecker()
	tm.Lock()
	defer tm.Unlock()
	if records := trackerData.trackerRecords[uname]; records == nil {
		err = fmt.Errorf("`%s` user not found, %w", uname, ErrorDataLayer)
	} else {
		if l := len(*records); l < 1 {
			err = fmt.Errorf("%s empty storage, %w", uname, ErrorDataLayer)
		} else {
			r := append(*records, models.NewTrackerRecord(state))
			trackerData.trackerRecords[uname] = &r
		}
	}
	return
}
func CreateUser(uname string) (err error) {
	anyChecker()
	tm.Lock()
	defer tm.Unlock()
	if _, ok := trackerData.trackerRecords[uname]; ok {
		err = fmt.Errorf("user exists %s, %w", uname, ErrorDataLayer)
	} else {
		trackerData.trackerRecords[uname] = &[]models.TrackerRecord{models.NewTrackerRecord(models.Absent)}
	}
	return
}

func ListUsers() (list map[string]string, err error) {
	anyChecker()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v:%w", r, ErrorDataLayer)
		}
	}()
	tm.RLock()
	defer tm.RUnlock()
	list = make(map[string]string, len(trackerData.trackerRecords))
	for key, value := range trackerData.trackerRecords {
		v := *value
		list[key] = v[len(v)-1].State.String()
	}

	return
}

func DeleteUser(uname string) (err error) {
	anyChecker()
	tm.Lock()
	defer tm.Unlock()
	if _, ok := trackerData.trackerRecords[uname]; !ok {
		err = fmt.Errorf("`%s` user not found, %w", uname, ErrorDataLayer)
	} else {
		delete(trackerData.trackerRecords, uname)
	}
	return
}

func ThisMonth(uname string) (model []models.TrackerRecord, err error) {
	anyChecker()
	tm.RLock()
	defer tm.RUnlock()
	n := time.Now()
	model, err = recordsTillThreshold(uname, time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, n.Location()))
	return
}
func GetDaysBack(uname string, days int) (model []models.TrackerRecord, err error) {
	anyChecker()
	tm.RLock()
	defer tm.RUnlock()
	model, err = recordsTillThreshold(uname, time.Now().Add(time.Duration(-days)))
	return
}

func GetMonthsBack(uname string, months int) (model []models.TrackerRecord, err error) {
	anyChecker()
	tm.RLock()
	defer tm.RUnlock()
	for i := months - 1; i > -1; i-- {
		f := fileForPreviousMonth(i)
		var e error
		if model, e = trackMonthsBack(f, uname); e != nil {
			if len(model) < 1 {
				err = fmt.Errorf("GetMonthsBack: %w", e)
				return
			}
			continue
		}
	}
	return
}

func Shutdown() {
	anyChecker()
	saveToFile(trackerData.name)
}
