package projecteuler

import "fmt"

// structure for holding a certain
type date struct {
	day     int // day of month, starting from 0 -> January
	month   int // month of year
	year    int // year
	weekday int // weekday, starting from 0 -> Sunday
}

// string representation of `date`
func (dt date) String() string {
	return fmt.Sprintf("%d / %d / %d", dt.day, dt.month, dt.year)
}

// checks whether a given year is leap year or not
func (dt date) isLeapYear() bool {
	// needs to be, either evenly divisible by 4 or in case of century, needs to be divisible by 400
	if dt.year%100 == 0 && dt.year%400 == 0 || dt.year%4 == 0 {
		return true
	}
	return false
}

// checks how many days present in february month, when year is given
func (dt date) getDayCountInFeb() int {
	if dt.isLeapYear() {
		return 29
	}
	return 28
}

// calculates how many days present in any given month and year
func (dt date) getDayCountInMonth() int {
	days := 0
	switch dt.month {
	case 0, 2, 4, 6, 7, 9, 11:
		days = 31
	case 3, 5, 8, 10:
		days = 30
	default:
		days = dt.getDayCountInFeb()
	}
	return days
}

// updates `date` pointer properly, so that next sunday
// gets reflected into updated `date`
//
// make sure, when you start calculation, initial date is `Sunday`
func (dt *date) getNextSunday() {
	if dt.getDayCountInMonth() > dt.day+7 {
		dt.day += 7
	} else {
		dt.day = dt.day + 7 - dt.getDayCountInMonth()
		dt.month = (dt.month + 1) % 12
		if dt.month == 0 {
			dt.year++
		}
	}
}

// checks whether given `date` is first day of month or not
func (dt date) isFirstOfMonth() bool {
	if dt.day == 1 {
		return true
	}
	return false
}

// checks whether given `date` is within timespan to be considered or not
/// given timespan is : 01/01/1901 - 31/12/2000
func (dt date) isWithinTimeSpan() bool {
	if dt.year >= 1901 && dt.year < 2001 {
		return true
	}
	return false
}

// CountingSundays - Counts number of sundays
// from 01/01/1901 to 31/12/2000, which are also
// first day of a month
func CountingSundays() int {
	cnt, dt := 0, date{31, 11, 1899, 0}
	for dt.year < 2001 {
		if dt.isWithinTimeSpan() && dt.isFirstOfMonth() {
			cnt++
		}
		dt.getNextSunday()
	}
	return cnt
}
