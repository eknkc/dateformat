package dateformat

import (
	"strconv"
)

var englishLongMonthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

var englishShortMonthNames = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var englishLongDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var englishShortDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

type english struct{}

func (e english) MonthName(index int) string {
	return englishLongMonthNames[index]
}

func (e english) ShortMonthName(index int) string {
	return englishShortMonthNames[index]
}

func (e english) DayName(index int) string {
	return englishLongDayNames[index]
}

func (e english) ShortDayName(index int) string {
	return englishShortDayNames[index]
}

func (e english) Ordinal(num int) string {
	n := num % 10

	switch {
	case ((num % 100) / 10) == 1:
		return strconv.Itoa(num) + "th"
	case n == 1:
		return strconv.Itoa(num) + "st"
	case n == 2:
		return strconv.Itoa(num) + "nd"
	case n == 3:
		return strconv.Itoa(num) + "rd"
	}

	return strconv.Itoa(num) + "th"
}

// English language locale that can be used as "FormatLocale" parameter
var English = english{}
