package dateformat

import (
	"strconv"
)

var dutchLongMonthNames = []string{
	"januari",
	"februari",
	"maart",
	"april",
	"mei",
	"juni",
	"juli",
	"augustus",
	"september",
	"oktober",
	"november",
	"december",
}

var dutchShortMonthNames = []string{
	"jan.",
	"feb.",
	"mrt.",
	"apr.",
	"mei",
	"jun.",
	"jul.",
	"aug.",
	"sep.",
	"okt.",
	"nov.",
	"dec.",
}

var dutchLongDayNames = []string{
	"zondag",
	"maandag",
	"dinsdag",
	"woensdag",
	"donderdag",
	"vrijdag",
	"zaterdag",
}

var dutchShortDayNames = []string{
	"zo.",
	"ma.",
	"di.",
	"wo.",
	"do.",
	"vr.",
	"za.",
}

type dutch struct{}

func (e dutch) MonthName(index int) string {
	return dutchLongMonthNames[index]
}

func (e dutch) ShortMonthName(index int) string {
	return dutchShortMonthNames[index]
}

func (e dutch) DayName(index int) string {
	return dutchLongDayNames[index]
}

func (e dutch) ShortDayName(index int) string {
	return dutchShortDayNames[index]
}

func (e dutch) Ordinal(num int) string {
	if num == 1 || num == 8 || num >= 20 {
		return strconv.Itoa(num) + "ste"
	}

	return strconv.Itoa(num) + "de"
}

// Dutch language locale that can be used as "FormatLocale" parameter
var Dutch dutch
