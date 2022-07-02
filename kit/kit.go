package main

import (
	"fmt"
	"time"
)

var week2Num = map[time.Weekday]int{
	time.Sunday:    0,
	time.Monday:    1,
	time.Tuesday:   2,
	time.Wednesday: 3,
	time.Thursday:  4,
	time.Friday:    5,
	time.Saturday:  6,
}

func main() {

	stringTime := "20170830"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("20060102", stringTime, loc)
	//if err == nil {
	//	unixTime := theTime.Unix() //1504082441
	//	fmt.Println(unixTime)
	//}

	fmt.Println(week2Num[theTime.Weekday()])
}
