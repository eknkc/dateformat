package dateformat

import (
	"strconv"
)

var frenchLongMonthNames = []string{
	"janvier",
	"février",
	"mars",
	"avril",
	"mai",
	"juin",
	"juillet",
	"août",
	"septembre",
	"octobre",
	"novembre",
	"décembre",
}

var frenchShortMonthNames = []string{
	"janv.",
	"févr.",
	"mars",
	"avr.",
	"mai",
	"juin",
	"juil.",
	"août",
	"sept.",
	"oct.",
	"nov.",
	"déc.",
}

var frenchLongDayNames = []string{
	"dimanche",
	"lundi",
	"mardi",
	"mercredi",
	"jeudi",
	"vendredi",
	"samedi",
}

var frenchShortDayNames = []string{
	"dim.",
	"lun.",
	"mar.",
	"mer.",
	"jeu.",
	"ven.",
	"sam.",
}

type french struct{}

func (e french) MonthName(index int) string {
	return frenchLongMonthNames[index]
}

func (e french) ShortMonthName(index int) string {
	return frenchShortMonthNames[index]
}

func (e french) DayName(index int) string {
	return frenchLongDayNames[index]
}

func (e french) ShortDayName(index int) string {
	return frenchShortDayNames[index]
}

func (e french) Ordinal(num int) string {
	if num == 1 {
		return strconv.Itoa(num) + "er"
	}

	return strconv.Itoa(num) + "ème"
}

// French language locale that can be used as "FormatLocale" parameter
var French = french{}
