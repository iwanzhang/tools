package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	EverySecond = "* * * * * *"
	EveryDay    = "0 1 * * *"
	EveryWeek   = "0 1 * * 1"
	EveryMonth  = "0 1 1 * *"
)

func run() {
	i := 0
	i++
	log.Println("cron running:", i)
}

func compute(c *cron.Cron) {
	c.AddJob(EverySecond, cron.FuncJob(func() {
		fmt.Println("eee")
	}))
}

func GetTodayBegin(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func GetMonthBegin(t time.Time) time.Time {
	begin := GetTodayBegin(t)
	return begin.AddDate(0, 0, -begin.Day()+1)
}

//GetLastSixMonthsStr offset
func GetLastSixMonthsStr(t time.Time, offset int) []string {

	var strs []string

	for i := 0; i < 6; i++ {
		//month := t.AddDate(0, -(offset + i), 0).Format("200601")
		month := t.AddDate(0, 0, -(offset + i)).Format("20060102")
		strs = append(strs, month)
	}
	return strs

}

func RunInfo() string {
	pc, file, line, ok := runtime.Caller(1) //上级调用入口
	if ok {
		return fmt.Sprintf("file: %s, func: %s, line: %d", file, runtime.FuncForPC(pc).Name(), line)
	}
	return ""
}

func test(ch chan []map[string]map[string]int, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	ch <- []map[string]map[string]int{
		{
			strconv.Itoa(i): {"hh": i},
		},
	}
}

//多协程并发执行

func multiTask(wg *sync.WaitGroup, ch chan interface{}, f func()) {

}

func main() {

	mm := make(map[string]map[string]int)

	if _, ok := mm["1"]; !ok {
		mm["1"] = make(map[string]int)
	}
	mm["1"]["2"] = 2
	mm["1"]["3"] = 3

	//m := make(map[string]int, 0)
	//m["1"] = 1
	//m["2"] = 2
	//m["3"] = 3

	fmt.Println(mm)
	os.Exit(0)

	type A struct {
		Biz  string
		Line []struct {
			DataStr string
			Num     int
		}
	}

	ch := make(chan []map[string]map[string]int, 3)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {

		wg.Add(1)
		go test(ch, &wg, i)
	}

	wg.Wait()

	close(ch)

	aa := make([]A, 0)

	for m := range ch {
		fmt.Println(m)
		for _, biz := range m {
			for bizName, opData := range biz {
				tmp := new(A)
				tmp.Biz = bizName
				for orgName, count := range opData {
					tmp.Line = append(tmp.Line, struct {
						DataStr string
						Num     int
					}{DataStr: orgName, Num: count})
					aa = append(aa, *tmp)
				}

			}
		}

	}

	//fmt.Printf("%+v", aa)

	//slice := append([]byte("hello"), "world"...)
	//fmt.Println(RunInfo(), slice)

	//fmt.Println(GetLastSixMonthsStr(time.Now(), 3))

	//yesterdayBegin := GetTodayBegin(time.Now()).AddDate(0, 0, -1).Unix()
	//yesterdayEnd := GetTodayBegin(time.Now()).Unix()
	//
	//lastMonthBegin := GetMonthBegin(time.Now()).AddDate(0, -1, 0).Unix()
	//lastMonthEnd := GetMonthBegin(time.Now()).Unix()
	//
	//fmt.Println(yesterdayEnd)
	//fmt.Println(yesterdayBegin)
	//fmt.Println(lastMonthBegin)
	//fmt.Println(lastMonthEnd)
	//

	//ctx, cancel := context.WithCancel(context.Background())
	//c := cron.New(cron.WithSeconds())
	//c.Start()
	//defer func() {
	//	c.Stop()
	//	cancel()
	//}()
	//go func() {
	//	for {
	//		compute(c)
	//		//监听主协程状态
	//		select {
	//		case <-ctx.Done():
	//			return
	//		default:
	//			time.Sleep(10 * time.Second)
	//		}
	//	}
	//}()
	//
	//time.Sleep(5 * time.Second)
}
