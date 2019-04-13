/*
time 包为我们提供了一个数据类型 time.Time（作为值使用）以及显示和测量时间和日期的功能函数。

当前时间可以使用 time.Now() 获取，或者使用 t.Day()、t.Minute() 等等来获取时间的一部分；
你甚至可以自定义时间格式化字符串，
例如： fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) 将会输出 21.07.2011。

Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。Location 类型映射某个时区的时间，UTC 表示通用协调世界时间。

包中的一个预定义函数 func (t Time) Format(layout string) string 可以根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串，
你可以使用一些预定义的格式，如：time.ANSIC 或 time.RFC822
*/

package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	// 2019-04-13 17:23:36.494106 +0800 CST m=+0.000186325
	fmt.Printf("%4d.%02d.%02d\n", t.Year(), t.Month(), t.Day())
	// 2019.04.13
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// 2019.04.13 17:23:36
	//t = time.Now().UTC()
	//fmt.Println(t)
	//fmt.Println(time.Now())
	week = 60 * 60 * 24 * 7 * 1e9 //  从纳秒开始计算
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", week_from_now.Year(), week_from_now.Month(), week_from_now.Day(), week_from_now.Hour(), week_from_now.Minute(), week_from_now.Second())
	fmt.Println(t.Format(time.RFC822))
	//13 Apr 19 17:36 CST
	fmt.Println(t.Format(time.ANSIC))
	//Sat Apr 13 17:36:40 2019
	fmt.Println(time.Now().Unix())
	//1555148397
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 必须吐槽下，必须写这个时间
	// 2019-04-13 17:41:02

	// 时间戳转str格式时间
	str_time := time.Unix(1555148645, 0).Format("2006-01-02 15:04:05")
	fmt.Println(str_time)

	// str 格式转时间戳
	the_time := time.Date(2019, 4, 13, 17, 44, 5, 0, time.Local)
	unix_time := the_time.Unix()
	fmt.Println(unix_time)
	//1555148645
	// 格式化当前时间
	lasttime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(lasttime)
	/*如果你需要在应用程序在经过一定时间或周期执行某项任务（事件处理的特例），则可以使用 time.After 或者 time.Ticker
	time.Sleep（Duration d） 可以实现对某个进程（实质上是 goroutine）时长为 d 的暂停*/
}
