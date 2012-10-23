// DateFormat provides convinient date formatting with localization support. Built in support for French, German, Spanish, Dutch, Turkish and English
package dateformat

import (
	"fmt"
	"regexp"
	"time"
)

var tokenRegexp = regexp.MustCompile(`(\[[^\[]*\])|(d{3,4}|Do|D{1,2}|Mo|M{1,4}|YYYY|YY|H{1,2}|h{1,2}|m{1,2}|s{1,2}|A|a|Z{1,2})`)
var magicRegexp = regexp.MustCompile(`(January|Jan|03|3|15|1|01|Monday|Mon|2006|06|02|2|04|4|05|5|PM|pm|\-0700|\-07:00)`)

// DateLocale specifies how to format dates in different languages.
// FormatLocale function accepts a DateLocale instance and formats the date according to given locale.
type DateLocale interface {
	// Return the long month name of given index (1 = January, 12 = December)
	MonthName(index int) string
	// Return the long weekday name of given index (0 = Sunday, 6 = Saturday)
	DayName(index int) string
	// Return the short month name of given index (1 = Jan, 12 = Dec)
	ShortMonthName(index int) string
	// Return the short weekday name of given index (0 = Sun, 6 = Sat)
	ShortDayName(index int) string
	// Return the ordinal value (1 = 1st, 2 = 2nd)
	Ordinal(num int) string
}

// Formats given date in "English" locale. See FormatLocale for details.
func Format(t time.Time, format string) string {
	return FormatLocale(t, format, English)
}

// Formats a given date according to the format string and date locale.
// Format strings are specified like MMMM Do YYYY, h:mm:ss a, which would yield October 24th 2012, 12:42:03 am in English locale.
// Supported format string tokens:
//
//	Month
//		M: 1 2 ... 12
//		MM: 01 01 ... 12
//		Mo: 1st 2nd ... 12th
//		MMM: Jan Feb ... Dec
//		MMMM: January February ... December
//
//	Day of Month
//		D: 1 2 ... 31
//		DD: 01 02 ... 31
//		Do: 1st 2nd ... 31st
//
//	Day of Week
//		ddd: Sun Mon ... Sat
//		dddd: Sunday Monday ... Saturday
//
//	Year
//		YY: 70 71 ... 12
//		YYYY: 1970 1971 ... 2012
//
//	Hour
//		H: 0 1 2 ... 23
//		HH: 00 01 02 .. 23
//		h: 1 2 ... 12
//		hh: 01 02 ... 12
//
//	Minute
//		m: 0 1 2 ... 59
//		mm: 00 01 02 ... 59
//
//	Second
//		s: 0 1 2 ... 59
//		ss: 00 01 02 ... 59
//
//	AM / PM
//		A: AM PM
//		a: am pm
//
//	Timezone
//		Z: -07:00 -06:00 ... +07:00
//		ZZ: -0700 -0600 ... +0700
//
//	Escaping
//		It is possible to place arbitrary text within [] to display it as it is.
//
//	Sample:
//		MMMM Do YYYY, h:mm:ss a [Foo] -> October 24th 2012, 1:52:27 am Foo
func FormatLocale(t time.Time, format string, locale DateLocale) string {
	formatted := tokenRegexp.ReplaceAllStringFunc(format, func(token string) string {
		switch token {
		case "M":
			return t.Format("1")
		case "MM":
			return t.Format("01")
		case "Mo":
			return locale.Ordinal(int(t.Month()))
		case "MMM":
			return locale.ShortMonthName(int(t.Month()) - 1)
		case "MMMM":
			return locale.MonthName(int(t.Month()) - 1)
		case "D":
			return t.Format("2")
		case "DD":
			return t.Format("02")
		case "Do":
			return locale.Ordinal(t.Day())
		case "ddd":
			return locale.ShortDayName(int(t.Weekday()))
		case "dddd":
			return locale.DayName(int(t.Weekday()))
		case "YY":
			return t.Format("06")
		case "YYYY":
			return t.Format("2006")
		case "H":
			return fmt.Sprintf("%d", t.Hour())
		case "HH":
			return fmt.Sprintf("%02d", t.Hour())
		case "h":
			return t.Format("3")
		case "hh":
			return t.Format("03")
		case "m":
			return t.Format("4")
		case "mm":
			return t.Format("04")
		case "s":
			return t.Format("5")
		case "ss":
			return t.Format("05")
		case "A":
			return t.Format("PM")
		case "a":
			return t.Format("pm")
		case "Z":
			return t.Format("-07:00")
		case "ZZ":
			return t.Format("-0700")
		}

		if token[0] == '[' && token[len(token)-1] == ']' {
			return token[1 : len(token)-1]
		}

		return token
	})

	return formatted
}

// Uses built in time package semantics for format string, that is a magic date.
// Instead of "MM", "H" style tokens that Format function accepts, you can supply a sample date as "magic" parameter
// And it will be figured out as the format.
// Given magic date layout must be
//	Mon Jan 2 15:04:05 MST 2006
// That is
//	01/02 03:04:05PM '06 -0700
// in a more memorable layout.
//
//	Sample:
//	magic = Mon, 02 Jan 2006 15:04:05 -0700
//	result = Wed, 24 Oct 2012  2:10:11 +0300
func FormatMagic(t time.Time, magic string, locale DateLocale) string {
	format := magicRegexp.ReplaceAllStringFunc(magic, func(token string) string {
		switch token {
		case "1":
			return "M"
		case "01":
			return "MM"
		case "Jan":
			return "MMM"
		case "January":
			return "MMMM"
		case "2":
			return "D"
		case "02":
			return "DD"
		case "Mon":
			return "ddd"
		case "Monday":
			return "dddd"
		case "06":
			return "YY"
		case "2006":
			return "YYYY"
		case "15":
			return "HH"
		case "3":
			return "h"
		case "03":
			return "hh"
		case "4":
			return "m"
		case "04":
			return "mm"
		case "5":
			return "s"
		case "05":
			return "ss"
		case "PM":
			return "A"
		case "pm":
			return "a"
		case "-07:00":
			return "Z"
		case "-0700":
			return "ZZ"
		}

		return "[" + token + "]"
	})

	return FormatLocale(t, format, locale)
}
