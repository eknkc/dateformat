package dateformat

import (
	"strconv"
)

var germanLongMonthNames = []string{
	"Januar",
	"Februar",
	"März",
	"April",
	"Mai",
	"Juni",
	"Juli",
	"August",
	"September",
	"Oktober",
	"November",
	"Dezember",
}

var germanShortMonthNames = []string{
	"Jan.",
	"Febr.",
	"Mrz.",
	"Apr.",
	"Mai",
	"Jun.",
	"Jul.",
	"Aug.",
	"Sept.",
	"Okt.",
	"Nov.",
	"Dez.",
}

var germanLongDayNames = []string{
	"Sonntag",
	"Montag",
	"Dienstag",
	"Mittwoch",
	"Donnerstag",
	"Freitag",
	"Samstag",
}

var germanShortDayNames = []string{
	"So.",
	"Mo.",
	"Di.",
	"Mi.",
	"Do.",
	"Fr.",
	"Sa.",
}

type german struct{}

func (e german) MonthName(index int) string {
	return germanLongMonthNames[index]
}

func (e german) ShortMonthName(index int) string {
	return germanShortMonthNames[index]
}

func (e german) DayName(index int) string {
	return germanLongDayNames[index]
}

func (e german) ShortDayName(index int) string {
	return germanShortDayNames[index]
}

func (e german) Ordinal(num int) string {
	return strconv.Itoa(num) + "."
}

// German language locale that can be used as "FormatLocale" parameter
var German german
