package data

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"test-time-tracker/models"
)

type (
	TrackerData struct {
		trackerRecords map[string]*[]models.TrackerRecord
		name           string
	}
)
const layout = "ttt-02-Jan.ttt"

var (
	trackerData TrackerData
	tm          sync.RWMutex

	ErrorDataLayer = errors.New("data error")
)

func init() {
	trackerData.name = fileForCurrentMonth()
	if initFromFile(trackerData.name) {
		fmt.Println("Staring time-tracker for", trackerData.name)
	} else {
		if initFromFile(fileForPreviousMonth(1)) {
			fmt.Println("Staring NEW time-tracker for", trackerData.name)
			resetTrackerData()
		} else {
			fmt.Println("Staring EMPTY time-tracker for", trackerData.name)
			trackerData.trackerRecords = map[string]*[]models.TrackerRecord{}
		}
	}
}

func anyChecker() {
	tm.Lock()
	defer tm.Unlock()
	if name := fileForCurrentMonth(); name != trackerData.name {
		saveToFile(trackerData.name)
		resetTrackerData()
		trackerData.name = name
	}
}

func initFromFile(name string) bool {
	if buf, err := ioutil.ReadFile(name); err == nil {
		if e := gob.NewDecoder(bytes.NewBuffer(buf)).Decode(&trackerData); e != nil {
			panic(e)
		}
		return true
	}
	return false
}

func saveToFile(name string) {
	var buffer bytes.Buffer

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&buffer)

	if err := enc.Encode(&trackerData); err != nil {
		panic(fmt.Errorf("`%s` encoding failure, %w", name, ErrorDataLayer))
	}
	if err := ioutil.WriteFile(name, buffer.Bytes(), 0666); err != nil { //nolint:gosec
		panic(fmt.Errorf("`%s` FS failure, %w", name, ErrorDataLayer))
	}
}

func resetTrackerData() {
	for k, records := range trackerData.trackerRecords {
		s := (*records)[len(*records)-1]
		trackerData.trackerRecords[k] = &[]models.TrackerRecord{models.NewTrackerRecord(s.State)}
	}
}

func fileForCurrentMonth() string {
	return time.Now().Format(layout)
}

func fileForPreviousMonth(i int) string {
	return time.Now().AddDate(0, -i, 0).Format(layout)
}

func recordsTillThreshold(uname string, threshold time.Time) (model []models.TrackerRecord, err error) {
	if records := trackerData.trackerRecords[uname]; records == nil {
		err = fmt.Errorf("`%s` user not found, %w", uname, ErrorDataLayer)
	} else {
		for _, r := range *records {
			if r.Timestamp.After(threshold) {
				model = append(model, r)
			}
		}
	}
	return model, err
}
func trackMonthsBack(f string, uname string) (model []models.TrackerRecord, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(f); err == nil {
		tData := TrackerData{}
		if e := gob.NewDecoder(bytes.NewBuffer(buf)).Decode(&tData); e != nil {
			err = fmt.Errorf("error %s in `%s`, %w", e.Error(), f, ErrorDataLayer)
			return
		}
		model = append(model, *(tData.trackerRecords[uname])...)
	}
	return
}
