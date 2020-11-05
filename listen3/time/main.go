package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now()
	fmt.Println("now:", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timeStamp := now.Unix()

	fmt.Printf("timeStamp is:%d\n", timeStamp)

}

func testTimeStamp(timeStamp2 int64) {
	timeObj := time.Unix(timeStamp2, 0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("current timestamp:%d\n", timeStamp2)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}
func processTask() {
	fmt.Printf("do tesk\n")
}

func testTicker() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}
}

func testConst() {
	fmt.Printf("纳秒 nano second:%d\n", time.Nanosecond)
	fmt.Printf("微妙 Micr second:%d\n", time.Microsecond)
	fmt.Printf("毫秒 Mill second:%d\n", time.Millisecond)
	fmt.Printf("秒 second:%d\n", time.Second)
}

//func testFormat(){
//	now := time.Now()
//	timeStr := now.Format("2006/01/02 15:04:05")
//	fmt.Println("timeStr:%s\n",timeStr)
//}
func testFormat2() {
	now := time.Now()
	timeStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("time:%s\n", timeStr)

}

func testCost() {
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()

	cost := (end-start)/1000
	fmt.Printf("cost :%d us\n", cost)
}

func main() {
	//testTime()
	//timeStamp2 := time.Now().Unix()
	//testTimeStamp(timeStamp2)
	//testTicker()
	//testConst()
	//testFormat()
	//testFormat2()
	testCost()
}
