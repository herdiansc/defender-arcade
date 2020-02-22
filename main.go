package main

import (
	"defender-arcade/models"
	"defender-arcade/utils"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <input_filename> \n", os.Args[0])
		os.Exit(0)
	}
	fileUtil := utils.NewFileUtil(ioutil.ReadFile)
	list, err := fileUtil.ContentToList(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	playTime := models.PlayTime{}
	playTimeList := models.PlayTimeList{}
	for i := range list {
		pt, err := playTime.FromStringRange(list[i])
		if err != nil {
			continue
		}
		playTimeList = append(playTimeList, pt)
	}
	totalOverlap := playTimeList.GetTotalOverlap()
	fmt.Printf("Minimum arcade needed: %d\n", totalOverlap)
}
