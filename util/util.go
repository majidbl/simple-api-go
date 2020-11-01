package util

import (
	"errors"
	"fmt"
	"time"
)

var classTime = map[string]string{
	"8-10":  "1",
	"10-12": "2",
	"12-14": "3",
	"14-16": "4",
	"16-18": "5",
}

var classDay = map[string]string{
	"Saturday":  "1",
	"Sunday":    "2",
	"Monday":    "3",
	"Tuesday":   "4",
	"Wednesday": "5",
	"Thursday":  "6",
	"Friday":    "7",
}

var lessonCode = map[string]string{
	"Biology":      "1000",
	"mathematical": "2000",
	"Chemistry":    "3000",
	"Cultural":     "4000",
	"History":      "5000",
	"Geography":    "6000",
	"IT":           "7000",
}

// GetPersonalId generate uniqe Id for teachers and student
func GetPersonalId(role string) (string, error) {
	t := time.Now()
	pid := fmt.Sprintf("%d%d%d%d%d%d", t.Year()/100, int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
	if role == "teacher" {
		pid = pid + "01"

	} else if role == "student" {
		pid = pid + "02"
	} else {
		err1 := errors.New("Role Error: role unknown")
		return "", err1
	}
	return pid, nil
}

//GetClassId generate uniqe id for each class
func GetClassId(lc, cd, ct string) (string, error) {
	cid := lessonCode[lc] + classDay[cd] + classTime[ct]
	return cid, nil
}
