package dateformat

import (
	"strconv"
)

var turkishLongMonthNames = [...]string{
	"Ocak",
	"Şubat",
	"Mart",
	"Nisan",
	"Mayıs",
	"Haziran",
	"Temmuz",
	"Ağustos",
	"Eylül",
	"Ekim",
	"Kasım",
	"Aralık",
}

var turkishShortMonthNames = [...]string{
	"Oca",
	"Şub",
	"Mar",
	"Nis",
	"May",
	"Haz",
	"Tem",
	"Ağu",
	"Eyl",
	"Eki",
	"Kas",
	"Ara",
}

var turkishLongDayNames = [...]string{
	"Pazar",
	"Pazartesi",
	"Salı",
	"Çarşamba",
	"Perşembe",
	"Cuma",
	"Cumartesi",
}

var turkishShortDayNames = [...]string{
	"Paz",
	"Pts",
	"Sal",
	"Çar",
	"Per",
	"Cum",
	"Cts",
}

type turkish struct{}

func (e turkish) MonthName(index int) string {
	return turkishLongMonthNames[index]
}

func (e turkish) ShortMonthName(index int) string {
	return turkishShortMonthNames[index]
}

func (e turkish) DayName(index int) string {
	return turkishLongDayNames[index]
}

func (e turkish) ShortDayName(index int) string {
	return turkishShortDayNames[index]
}

func (e turkish) Ordinal(num int) string {
	var suffix = func() string {
		a := num % 10
		b := num % 100

		switch a {
		case 1:
			return "inci"
		case 2:
			return "nci"
		case 3:
			return "üncü"
		case 4:
			return "üncü"
		case 5:
			return "inci"
		case 6:
			return "ncı"
		case 7:
			return "nci"
		case 8:
			return "inci"
		case 9:
			return "uncu"
		}

		switch b {
		case 10:
			return "uncu"
		case 20:
			return "nci"
		case 30:
			return "uncu"
		case 40:
			return "ıncı"
		case 50:
			return "nci"
		case 60:
			return "ıncı"
		case 70:
			return "inci"
		case 80:
			return "inci"
		case 90:
			return "ıncı"
		}

		return "üncü"
	}()

	return strconv.Itoa(num) + "'" + suffix
}

// Turkish language locale that can be used as "FormatLocale" parameter
var Turkish turkish
