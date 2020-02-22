package models

import (
	"fmt"
	"testing"
)

func TestPlayTime_FromStringRange(t *testing.T) {
	playTime := PlayTime{}
	cases := []struct {
		testName    string
		stringRange string
		expected    PlayTime
	}{
		{
			testName:    "1. Positive Case",
			stringRange: "900 1000",
			expected:    PlayTime{StartTime: 900, FinishTime: 1000},
		},
		{
			testName:    "2. Negative Case: Empty string range",
			stringRange: "",
			expected:    PlayTime{StartTime: 0, FinishTime: 0},
		},
		{
			testName:    "3. Negative Case: Invalid string range",
			stringRange: "1 2 3",
			expected:    PlayTime{StartTime: 0, FinishTime: 0},
		},
		{
			testName:    "4. Negative Case: Invalid time",
			stringRange: "A 1000",
			expected:    PlayTime{StartTime: 0, FinishTime: 0},
		},
		{
			testName:    "5. Negative Case: Invalid time",
			stringRange: "900 B",
			expected:    PlayTime{StartTime: 0, FinishTime: 0},
		},
	}

	for _, c := range cases {
		t.Logf("Currently testing: %s\n", c.testName)
		result, _ := playTime.FromStringRange(c.stringRange)
		if c.expected.StartTime != result.StartTime {
			t.Errorf("Expected: %v, Got: %v\n", c.expected.StartTime, result.StartTime)
		}
		if c.expected.FinishTime != result.FinishTime {
			t.Errorf("Expected: %v, Got: %v\n", c.expected.FinishTime, result.FinishTime)
		}
	}
}

func TestPlayTimeList_GetTotalOverlap(t *testing.T) {
	cases := []struct {
		testName     string
		playTimeList PlayTimeList
		expected     int
	}{
		{
			testName: "1. Positive Case",
			playTimeList: PlayTimeList{
				PlayTime{StartTime: 900, FinishTime: 1000},
				PlayTime{StartTime: 930, FinishTime: 1030},
				PlayTime{StartTime: 1000, FinishTime: 1100},
			},
			expected: 2,
		},
		{
			testName: "2. Positive Case: No overlap, needed 1 arcade",
			playTimeList: PlayTimeList{
				PlayTime{StartTime: 900, FinishTime: 1000},
				PlayTime{StartTime: 1000, FinishTime: 1100},
				PlayTime{StartTime: 1100, FinishTime: 1200},
			},
			expected: 1,
		},
		{
			testName: "3. Positive Case: All overlap, needed most(3) arcade",
			playTimeList: PlayTimeList{
				PlayTime{StartTime: 900, FinishTime: 1000},
				PlayTime{StartTime: 900, FinishTime: 1000},
				PlayTime{StartTime: 900, FinishTime: 1000},
			},
			expected: 3,
		},
	}

	for _, c := range cases {
		t.Logf("Currently testing: %s\n", c.testName)
		result := c.playTimeList.GetTotalOverlap()
		fmt.Printf("result: %v\n", result)
		if c.expected != result {
			t.Errorf("Expected: %v, Got: %v\n", c.expected, result)
		}
	}
}
