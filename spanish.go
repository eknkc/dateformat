package dateformat

import (
	"strconv"
)

var spanishLongMonthNames = []string{
	"Enero",
	"Febrero",
	"Marzo",
	"Abril",
	"Mayo",
	"Junio",
	"Julio",
	"Agosto",
	"Septiembre",
	"Octubre",
	"Noviembre",
	"Diciembre",
}

var spanishShortMonthNames = []string{
	"Ene.",
	"Feb.",
	"Mar.",
	"Abr.",
	"May.",
	"Jun.",
	"Jul.",
	"Ago.",
	"Sep.",
	"Oct.",
	"Nov.",
	"Dic.",
}

var spanishLongDayNames = []string{
	"Domingo",
	"Lunes",
	"Martes",
	"Miércoles",
	"Jueves",
	"Viernes",
	"Sábado",
}

var spanishShortDayNames = []string{
	"Dom.",
	"Lun.",
	"Mar.",
	"Mié.",
	"Jue.",
	"Vie.",
	"Sáb.",
}

type spanish struct{}

func (e spanish) MonthName(index int) string {
	return spanishLongMonthNames[index]
}

func (e spanish) ShortMonthName(index int) string {
	return spanishShortMonthNames[index]
}

func (e spanish) DayName(index int) string {
	return spanishLongDayNames[index]
}

func (e spanish) ShortDayName(index int) string {
	return spanishShortDayNames[index]
}

func (e spanish) Ordinal(num int) string {
	return strconv.Itoa(num) + "º"
}

// Spanish language locale that can be used as "FormatLocale" parameter
var Spanish spanish
