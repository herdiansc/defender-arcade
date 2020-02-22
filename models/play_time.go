package models

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// PlayTime holds range of playing time.
type PlayTime struct {
	StartTime  int
	FinishTime int
}

func (pt PlayTime) FromStringRange(str string) (PlayTime, error) {
	listRange := strings.Split(str, " ")
	if len(listRange) != 2 {
		return pt, errors.New("invalid string")
	}
	startTime, err := strconv.Atoi(listRange[0])
	if err != nil {
		return pt, errors.New("range consist of invalid time")
	}
	finishTime, err := strconv.Atoi(listRange[1])
	if err != nil {
		return pt, errors.New("range consist of invalid time")
	}

	pt.StartTime = startTime
	pt.FinishTime = finishTime
	return pt, nil
}

// PlayTimeList holds list of PlayTime.
type PlayTimeList []PlayTime

// GetTotalOverlap returns total of overlap times.
func (playTimes PlayTimeList) GetTotalOverlap() int {
	sort.Slice(playTimes, func(i, j int) bool {
		return playTimes[i].StartTime < playTimes[j].StartTime
	})
	var arcade int
	max := 1
	for i := range playTimes {
		arcade = 1
		for j := i + 1; j < len(playTimes); j++ {
			if playTimes[i].FinishTime <= playTimes[j].StartTime {
				continue
			}
			if playTimes[i].FinishTime > playTimes[j].StartTime {
				arcade++
			}
		}
		if arcade > max {
			max = arcade
		}
	}
	return max
}
