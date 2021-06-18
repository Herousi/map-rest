package echarts

import (
	"fmt"
	"time"
)

// 获取n天时间
func GetNumDayArr(num int, format string) (arr []string) {
	arr = make([]string, 0)
	now := time.Now()
	now = now.AddDate(0, 0, -num)
	for i := 0; i < num; i++ {
		arr = append(arr, now.Format(format))
		now = now.AddDate(0, 0, 1)
	}
	return
}

type WeekInfo struct {
	Day string `json:"day"`
	Num string `json:"num"`
}

// 获取n周时间
func GetNumWeekArr(num int, format string) (arr []WeekInfo) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekStart = weekStart.AddDate(0, 0, -7*(num-1))
	for i := 0; i < num; i++ {
		arr = append(arr, WeekInfo{Day: weekStart.Format(format), Num: WeekByDate(weekStart)})
		weekStart = weekStart.AddDate(0, 0, 7)
	}
	return
}

// 获取n月事件
func GetNumMonthArr(num int, format string) (arr []string) {
	arr = make([]string, 0)
	now := time.Now().AddDate(0, -num, 0)
	for i := 0; i < num; i++ {
		arr = append(arr, now.Format(format))
		now = now.AddDate(0, 1, 0)
	}
	return
}

//判断时间是当年的第几周
func WeekByDate(t time.Time) string {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())
	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return fmt.Sprintf("第%d周", week)

}
