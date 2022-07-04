/*
 * @Autor: 郭彬
 * @Description: 时间相关工具函数
 * @Date: 2022-04-28 11:33:50
 * @LastEditTime: 2022-06-17 10:03:49
 * @FilePath: \Test\util\TimeUtil.go
 */
package util

import (
	"fmt"
	"time"
)

const (
	COMM_TIMER_DEF     = "2006-01-02 15:04:05"
	COMM_TIMER_DEF_YMD = "2006-01-02 00:00:00"
	COMM_TIMER_Y_M_D   = "2006-01-02"
	COMM_TIMER_Y_M_D2  = "20060102"
	COMM_TIMER_Y_M2    = "200601"
	COMM_TIMER_Y       = "2006"
)

//golang time.Time 时间格式 eg: 2022-04-28 11:37:20.1815267 +0800 CST m=+0.002939001

/*******************************************************************
//  获取当前时间
//
*******************************************************************/

// 获取当前 golang 的 time 时间
func GetTime() time.Time {
	return time.Now()
}

// 获取当前 linux 时间戳 秒数
func GetTimeStamp() int64 {
	return time.Now().Unix()
}

// 获取当前 YY-MM-DD HH:ii:ss 时间字符串
func GetDateTimeStr() string {
	return time.Now().Format(COMM_TIMER_DEF)
}

/*******************************************************************
//  time.Time 转 其他， time.Time比较
//
*******************************************************************/

// golang time.Time 时间转 unix 时间戳
func Time2Unix(t time.Time) int64 {
	return t.Unix()
}

// golang time.Time 时间转  yy-mm-dd HH:ii:ss 日期时间 格式
func Time2DateTime(t time.Time) string {
	return t.Format(COMM_TIMER_DEF)
}

// golang time.Time 时间转 HH:ii:ss 时间
func Time2DayTime(t time.Time) string {
	return t.Format("15:04:05")
}

// golang time.Time 时间转 YYMMDDHHiiss 格式的时间字符串
func Time2StrTimeFormat(t time.Time) string {
	var s string = t.Format("20060102150405")
	return s
}

//golang time.Time 获取此月开始时间字符串 YY-mm-dd
func Time2MonthBeginDt(t time.Time) string {
	Y := t.Year()
	M := t.Month()
	return fmt.Sprintf("%d-%02d-01", Y, M)
}

//golang time.Time 判断时间是当年的第几周
func Time2YearWeek(t time.Time) int {
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
	return week
}

// golang time.Time 是否是同一周
func CheckTimeSameWeek(t1 time.Time, t2 time.Time) bool {
	return Time2YearWeek(t1) == Time2YearWeek(t2)
}

// golang time.Time 是否是同一天
func CheckTimeSameDay(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// golang time.Time 两个天数相差，0表同一天，正数表示 t1 > t2，负数表示 t1 < t2
func GetDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}

/***********************************************************************
//   （YY-mm-dd HH:ii:ss）字符串格式化时间 转换其他格式
//
************************************************************************/

// 时间字符串 yy-mm-dd HH:ii:ss 转换为 time.Time格式时间
func DateTime2Time(str string) time.Time {
	todayZero, _ := time.ParseInLocation(COMM_TIMER_DEF, str, time.Local)
	return todayZero
}

// 时间字符串 yy-mm-dd HH:ii:ss 转换为 unix 时间戳
func DateTime2Unix(formatTimeStr string) int64 {
	formatTime, err := time.ParseInLocation(COMM_TIMER_DEF, formatTimeStr, time.Local)
	if err == nil {
		return formatTime.Unix()
	} else {
		return 0
	}
}

/*********************************************
//   unix 时间戳 转换其他格式
//
**********************************************/

// unix 时间戳 转 golang time.Time格式
func Unix2Time(ts int64) time.Time {
	return time.Unix(ts, 0)
}

// unix 时间戳 转 yyy-mm-dd HH:ii:ss 日期时间 格式
func Unix2DateTime(ts int64) string {
	var s string = time.Unix(ts, 0).Format(COMM_TIMER_DEF)
	return s
}

// 判断两个 unix 时间戳是否是同一天
func CheckUnixStampSameDay(ts1 int64, ts2 int64) bool {
	t1 := time.Unix(ts1, 0)
	t2 := time.Unix(ts2, 0)
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

/***********************************************************
 * 其他复杂应用
 *
 ***********************************************************/

// getYearMonthToDay 查询指定年份指定月份有多少天
// @params year int 指定年份
// @params month int 指定月份
func GetYearMonthToDaysCnt(year int, month int) uint {
	var day uint
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			day = 30
		} else {
			day = 31
		}
	} else {
		// 计算是平年还是闰年,得出2月的天数
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			day = 29
		} else {
			day = 28
		}
	}
	return day
}

// 返回本周一的日期
func GetFirstDateOfWeek() (weekMonday string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format(COMM_TIMER_Y_M_D)
	return
}

// 获取上周的周一日期
func GetLastWeekFirstDate() (weekMonday string) {
	thisWeekMonday := GetFirstDateOfWeek()
	TimeMonday, _ := time.Parse(COMM_TIMER_Y_M_D, thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -7)
	weekMonday = lastWeekMonday.Format(COMM_TIMER_Y_M_D)
	return
}

/**
 * @Description 获得指定日的初始和结束日期（时间字符串）
 * @Param  date 2021-08-02 00:00:00
 * @return 天开始时间字符串，天结束时间字符串
 **/
func GetDayBeginAndDayEnd(date string, format string) (string, string) {
	now := time.Now()
	if date != "" {
		now = DateTime2Time(date)
	}
	if format == "" {
		format = COMM_TIMER_DEF
	}
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfDay := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)           //今天开始
	lastOfDay := time.Date(currentYear, currentMonth, 1, 23, 59, 59, 999999999, currentLocation) //今天结束
	return firstOfDay.Format(format), lastOfDay.Format(format)
}

/**
 * @Description 获得指定时间周的初始和结束日期
 * @Param  date 2021-08-02 00:00:00
 * @return
 **/
func GetWeekDayBeginAndEnd(date string, format string) (string, string) {
	now := time.Now()
	if date != "" {
		now = DateTime2Time(date)
	}
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}
	lastOffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastOffset == 6 {
		lastOffset = -1
	}
	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastOffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	if format == "" {
		return time.Unix(f, 0).Format(COMM_TIMER_Y_M_D) + " 00:00:00", time.Unix(l, 0).Format(COMM_TIMER_Y_M_D) + " 23:59:59"
	}
	return time.Unix(f, 0).Format(format), time.Unix(l, 0).Format(format)
}

/**
 * @Description 获得指定月的初始和结束日期
 * @Param  date 2021-08-02 00:00:00
 * @return
 **/
func GetMonthDayBeginAndEnd(date string, format string) (string, string) {
	now := time.Now()
	if date != "" {
		now = DateTime2Time(date)
	}
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	if format == "" {
		return time.Unix(f, 0).Format(COMM_TIMER_Y_M_D) + " 00:00:00", time.Unix(l, 0).Format(COMM_TIMER_Y_M_D) + " 23:59:59"
	}
	return time.Unix(f, 0).Format(format), time.Unix(l, 0).Format(format)
}

/**
 * @Description //获得指定时间季度的初始和结束日期
 * @Param  date 2021-08-02 00:00:00
 * @return
 **/
func GetQuarterDayBeginAndEnd(date string) (string, string) {
	now := time.Now()
	if date != "" {
		now = DateTime2Time(date)
	}
	year := now.Format("2006")
	month := int(now.Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
}

/**
 * @Description //根据开始日期和结束日期计算出时间段内所有日期
 * @Param  sdate 开始 edate 结束 (参数为日期格式，如：2020-01-01)
 * @return
 **/
func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}
