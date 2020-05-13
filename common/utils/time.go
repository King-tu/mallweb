package utils

import (
	"errors"
	"fmt"
	"github.com/king-tu/mallweb/common"
	"go.uber.org/zap"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GetNoExpireTime() time.Time {
	return time.Date(2116, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func GetZeroTime() time.Time {
	return time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
}
func GetZeroTime2() time.Time {
	return time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
}
func GetUnixZeroTime() time.Time {
	return time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func TimeDuration(startTime, endTime string) (int, error) {
	slices := strings.Split(startTime, ":")
	startHour, err := strconv.ParseInt(slices[1], 10, 0)
	if err != nil {
		return 0, err
	}
	startSecond, err := strconv.ParseInt(slices[0], 10, 0)
	if err != nil {
		return 0, err
	}
	slices = strings.Split(endTime, ":")
	endHour, err := strconv.ParseInt(slices[1], 10, 0)
	if err != nil {
		return 0, err
	}
	endSecond, err := strconv.ParseInt(slices[0], 10, 0)
	if err != nil {
		return 0, err
	}
	return ((int)(60*(endHour-startHour) + (endSecond - startSecond))), nil
}

func GetTime(time string) (int, int, error) {
	if time == "" {
		return 0, 0, errors.New("time is non")
	}
	slices := strings.Split(time, ":")
	if len(slices) < 2 {

		return 0, 0, errors.New("time format is error")
	}
	minute, err := strconv.ParseInt(slices[1], 10, 0)
	if err != nil {
		return 0, 0, err
	}
	hour, err := strconv.ParseInt(slices[0], 10, 0)
	if err != nil {
		return 0, 0, err
	}
	if hour > 23 || minute > 59 {
		err := errors.New("wrong value for the time")
		return 0, 0, err
	}
	return (int)(hour), (int)(minute), nil
}

func HandleDate(timestamp time.Time) (dateTime time.Time) {
	year, month, day := timestamp.Date()
	dateTime = time.Date(year, month, day, 0, 0, 0, 0, timestamp.Location())
	return dateTime
}

func GetTimeStamp(date time.Time, strTime string) (time.Time, error) {
	hour, minute, err := GetTime(strTime)
	if err != nil {
		return date, err
	}
	timeStamp := date.Add(time.Duration(hour)*time.Hour + time.Duration(minute)*time.Minute)
	return timeStamp, nil
}

func GetDate(strYear string, strDay string, defaultDate time.Time) (time.Time, error) {
	slices := strings.Split(strDay, ":")
	if len(slices) < 2 {

		return defaultDate, errors.New("day format is error")
	}
	day, err := strconv.ParseInt(slices[1], 10, 0)
	if err != nil {
		return defaultDate, err
	}
	month, err := strconv.ParseInt(slices[0], 10, 0)
	if err != nil {
		return defaultDate, err
	}
	year, err := strconv.ParseInt(strYear, 10, 0)
	if err != nil {
		return defaultDate, err
	}

	if month > 12 || day > 31 {
		err := errors.New("wrong value for the day")
		return defaultDate, err
	}
	date := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, defaultDate.Location())
	return date, nil
}

func IsLastDayOfMonth(year, month, day int) bool {
	if day == 31 {
		return true
	}
	if day == 29 && time.Month(month) == time.February {
		return true
	}
	if day == 30 {
		switch time.Month(month) {
		case time.April:
		case time.June:
		case time.September:
		case time.November:
			return true
		default:
			break
		}
	}
	if day == 28 && time.Month(month) == time.February && year%4 > 0 {
		return true
	}
	return false
}

func IsSameDate(time1 time.Time, time2 time.Time) bool {
	year1, month1, day1 := time1.Date()
	year2, month2, day2 := time2.Date()

	return year1 == year2 && month1 == month2 && day1 == day2
}

func IsSameMonth(time1 time.Time, time2 time.Time) bool {
	year1, month1, _ := time1.Date()
	year2, month2, _ := time2.Date()

	return year1 == year2 && month1 == month2
}

func GetDaysBetweenTwoDates(time1 time.Time, time2 time.Time) int {
	diff := time1.Sub(time2)

	return int(diff.Hours() / 24)
}

func GetMonthsBetweenTwoDates(time1 time.Time, time2 time.Time) int {
	if time1.Before(time2) {
		return 0
	}
	year1, month1, _ := time1.Date()
	year2, month2, _ := time2.Date()

	return int(month1) - int(month2) + 12*(int(year1)-int(year2)) + 1
}

func GetDuration(unit string) time.Duration {
	switch unit {
	case "m":
		return time.Minute
	case "d":
		return time.Hour * 24
	case "h":
		return time.Hour
	case "s":
		return time.Second
	}

	return time.Minute
}

func GetMinMaxTimeInOneDay(t time.Time) (time.Time, time.Time) {
	year, month, day := t.Date()
	minTime := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	maxTime := time.Date(year, month, day, 23, 59, 59, 0, t.Location())
	return minTime, maxTime
}

func IsDayAfter(time1, time2 time.Time) bool {
	ret := false
	year1 := time1.Year()
	year2 := time2.Year()

	if year1 > year2 {
		ret = true
	} else if year1 == year2 {
		ret = time1.YearDay() > time2.YearDay()
	}

	return ret
}

var monthDays = map[int]int{
	1:  31,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func LastDayInMonth(year, month int) int {
	if month == 2 {
		isLeapYear := year%4 == 0 &&
			(year%100 != 0 || year%400 == 0)
		if isLeapYear {
			return 29
		}
		return 28
	}
	return monthDays[month]
}

func GetTimeStampIn24h(t time.Time) string {
	format := "20060102150405"
	return t.Format(format)
}

func StrToTime(timeStr, timeTem string) time.Time {
	//toBeCharge := timeStr                                           //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	//timeLayout := timeTem                                           //转化所需模板
	loc, _ := time.LoadLocation("Local")                      //重要：获取时区
	theTime, _ := time.ParseInLocation(timeTem, timeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime
}

func GetDateTimeByHourMinute(timeStr string) time.Time {
	atArr := strings.Split(timeStr, ":")
	hour, _ := strconv.Atoi(atArr[0])
	minute, _ := strconv.Atoi(atArr[1])
	t := time.Now()
	year, month, day := t.Date()
	return time.Date(year, month, day, hour, minute, 0, 0, t.Location())
}

/**
根据输入的带月份的时间，导出标准的时间
返回值：int64, error
*/
func GetFormatMonthTimeByImport(inDate string) (int64, error) {
	time, err := GetFormatMonthByImport(inDate)
	if err == nil {
		return time.Unix(), nil
	}
	return 0, err
}
func GetFormatMonthTimeByImportV2(inDate interface{}) (int64, error) {
	var month int64
	var ret error
	_kind := reflect.TypeOf(inDate).Kind()
	zap.L().Sugar().Debugf("_kind:%v", _kind)
	if _kind == reflect.String {
		opMonth, err := GetFormatMonthByImport(inDate.(string))
		if err != nil {
			zap.L().Sugar().Debugf("err:%v", err)
			ret = err
		}
		month = GetYearMonthByTime(opMonth).Unix()
	} else if _kind == reflect.Float64 {
		str := fmt.Sprintf("%v", inDate.(float64))
		opMonth, err := GetFormatMonthByImport(str)
		if err != nil {
			zap.L().Sugar().Debugf("err:%v", err)
			ret = err
		}
		month = GetYearMonthByTime(opMonth).Unix()
	} else {
		zap.L().Debug("err")
	}

	return month, ret
}

/**
根据输入的带月份的时间，导出标准的时间
返回值：Time, error
*/
func GetFormatMonthByImport(inDate string) (time.Time, error) {
	endTime, err := time.Parse(common.DATE_TIME_FORMAT, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.DATE_FORMAT2, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.MONTH_FORMAT2, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.DATE_FORMAT3, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.CHINA_MONTH_FORMAT, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.CHINA_DATE_FORMAT, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.MONTH_FORMAT3, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.DATE_FORMAT, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.MONTH_FORMAT, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.MONTH_FORMAT1, inDate)
	if err == nil {
		return endTime, nil
	}
	endTime, err = time.Parse(common.DATE_FORMAT1, inDate)
	if err == nil {
		return endTime, nil
	}

	endTime, err = time.Parse(common.CHINA_MONTH_FORMAT1, inDate)
	if err == nil {
		endTime = endTime.AddDate(time.Now().Year(), 0, 0)
		return endTime, nil
	}

	endTime, err = time.Parse(common.CHINA_MONTH_FORMAT2, inDate)
	if err == nil {
		return endTime, nil
	}

	zap.L().Sugar().Errorf("Failed to parse month:%v with err:%v", inDate, err)
	return endTime, err
}

func FormatTimeToString(date time.Time, layout string) string {
	return date.Format(layout)
}

func FormatTimeToMonthLastday(date time.Time, layout string) string {
	year, month, _ := date.Date()
	lastday, _ := monthDays[int(month)]
	newdate := time.Date(year, month, lastday, 0, 0, 0, 0, date.Location())
	return newdate.Format(layout)
}

//func FormatTime(v interface{}, timeType string) (time.Time, error) {
//	var tm time.Time
//	if v == nil {
//		return tm, errors.New("时间格式没有识别")
//	}
//	_kind := reflect.TypeOf(v).Kind()
//	if _kind == reflect.Int64 {
//		tm = time.Unix(v.(int64), 0)
//	} else if _kind == reflect.String {
//		val := StrTypeMapping(v.(string) /*"date"*/, timeType, false)
//		if t, ok := val.(time.Time); ok {
//			tm = t
//		}
//	} else if t, ok := v.(time.Time); ok {
//		tm = t
//	} else {
//		return tm, errors.New("时间格式没有识别")
//	}
//	return tm, nil
//}

func FirstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	return date
}

//格式话时间，只要年月日
func FormatUnixTimeToYearMonth(i int64) int64 {
	t := time.Unix(i, 0)
	location, _ := time.LoadLocation(common.DEFAULT_TIME_ZONE)
	new := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, location)
	return new.Unix()
}

//传入time.TIME,返回时间戳
func GetUnixTimeByTime(t time.Time) int64 {
	return t.Unix()
}

//传入时间戳，返回time.TIME
func GetTimeByUnixTime(i int64) time.Time {
	return time.Unix(i, 0)
}

//传入时间戳，返回time.Month
func GetMonthByUnixTime(i int64) time.Month {
	t := GetTimeByUnixTime(i)
	return t.Month()
}

//传入时间戳，返回day in month
func GetDayInMonth(i int64) int {
	t := GetTimeByUnixTime(i)
	return t.Day()
}

//传入TIME,返回MAX_DAY in month
func GetMaxDayInMonth(i int64) int {
	t := GetTimeByUnixTime(i)
	year, month, _ := t.Date()
	switch month {
	case time.January:
	case time.March:
	case time.May:
	case time.July:
	case time.August:
	case time.October:
	case time.December:
		return 31
	case time.February:
		if year%4 == 0 {
			return 29
		}
		return 28
	case time.April:
	case time.June:
	case time.September:
	case time.November:
		return 30
	}
	return 0
}

/*
保留时间的年月 将日设置为1 时间设置为0
*/
func GetYearMonthByTime(opMonthTime time.Time) time.Time {
	location, _ := time.LoadLocation(common.DEFAULT_TIME_ZONE)
	opMonth := time.Date(opMonthTime.Year(), opMonthTime.Month(), 1, 0, 0, 0, 0, location)
	return opMonth
}

func GetYearMonthDayByTime(opMonthTime time.Time) time.Time {
	location, _ := time.LoadLocation(common.DEFAULT_TIME_ZONE)
	opMonth := time.Date(opMonthTime.Year(), opMonthTime.Month(), opMonthTime.Day(), 0, 0, 0, 0, location)
	return opMonth
}

func ListDateFormats() []string {
	date_formats := make([]string, 0)

	years := []string{"2006", "06"}
	months := []string{"01", "1"}
	days := []string{"02", "2"}
	separators := []string{"-", ".", "/", "c"}

	for _, _year := range years {
		for _, _month := range months {
			for _, _day := range days {
				for _, separator := range separators {
					_seps := []string{separator, separator, separator}
					if separator == "c" {
						_seps = []string{"年", "月", "日"}
					}

					format := fmt.Sprintf("%s%s%s%s%s%s", _year, _seps[0], _month, _seps[1], _day, _seps[2])

					if separator != "c" {
						format = fmt.Sprintf("%s%s%s%s%s", _year, _seps[0], _month, _seps[1], _day)
					}

					date_formats = append(date_formats, format)
				}
			}

			for _, separator := range separators {
				_seps := []string{separator, separator}
				format := fmt.Sprintf("%s%s%s", _year, _seps[0], _month)
				if separator == "c" {
					_seps = []string{"年", "月"}
					format = fmt.Sprintf("%s%s%s%s", _year, _seps[0], _month, _seps[1])
				}
				date_formats = append(date_formats, format)
			}
		}
	}

	return date_formats
}

func GetStartEndTime(startTime, endTime string) (time.Time, time.Time) {
	var _startTime, _endTime time.Time
	location, _ := time.LoadLocation(common.DEFAULT_TIME_ZONE)

	timeSeconds, err := strconv.ParseInt(startTime, 10, 64)
	if err == nil {
		year, month, day := time.Unix(timeSeconds, 0).Date()
		_startTime = time.Date(year, month, day, 0, 0, 0, 0, location)
	}

	timeSeconds, err = strconv.ParseInt(endTime, 10, 64)
	if err == nil {
		year, month, day := time.Unix(timeSeconds, 0).Date()
		_endTime = time.Date(year, month, day, 23, 59, 59, 0, location)
	}

	return _startTime, _endTime
}

func IsWeekend(t time.Time) bool {
	w := t.Weekday()
	return w == 0 || w == 6
}

func IsInDayTimeRange(t time.Time, startTime, endTime string) bool {
	year, month, day := t.Date()

	starts := strings.Split(startTime, ":")
	startHour, _ := strconv.ParseInt(starts[0], 10, 0)
	startMinute, _ := strconv.ParseInt(starts[1], 10, 0)

	ends := strings.Split(endTime, ":")
	endHour, _ := strconv.ParseInt(ends[0], 10, 0)
	endMinute, _ := strconv.ParseInt(ends[1], 10, 0)

	_startTime := time.Date(year, month, day, int(startHour), int(startMinute), 0, 0, t.Location())
	_endTime := time.Date(year, month, day, int(endHour), int(endMinute), 0, 0, t.Location())

	return t.Equal(_startTime) || t.Equal(_endTime) || (t.After(_startTime) && t.Before(_endTime))
}

func FormatDate(timestamp time.Time) (dateTime time.Time) {
	year, month, day := timestamp.Date()
	location, _ := time.LoadLocation(common.DEFAULT_TIME_ZONE)
	dateTime = time.Date(year, month, day, 8, 0, 0, 0, location)
	return dateTime
}

//NextMonth ...
func NextMonth(t time.Time) time.Time {
	m, y := int(t.Month()+1), t.Year()
	if m > 12 {
		m = 1
		y++
	}

	return time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)
}

//NextDay ...
func NextDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, 1)
}

// CompareMonth 按月份比较
func CompareMonth(t1, t2 time.Time) int64 {
	t1 = time.Date(t1.Year(), t1.Month(), 0, 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), 0, 0, 0, 0, 0, time.Local)

	return t1.Unix() - t2.Unix()
}

// CompareDay 按天比较
func CompareDay(t1, t2 time.Time) int64 {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return t1.Unix() - t2.Unix()
}
