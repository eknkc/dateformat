# dateformat
--
    import "github.com/eknkc/dateformat"

DateFormat provides convenient date formatting with localization support. Built in support for French, German, Spanish, Dutch, Turkish and English

## Usage

```go
var Dutch dutch
```
Dutch language locale that can be used as "FormatLocale" parameter

```go
var English english
```
English language locale that can be used as "FormatLocale" parameter

```go
var French french
```
French language locale that can be used as "FormatLocale" parameter

```go
var German german
```
German language locale that can be used as "FormatLocale" parameter

```go
var Spanish spanish
```
Spanish language locale that can be used as "FormatLocale" parameter

```go
var Turkish turkish
```
Turkish language locale that can be used as "FormatLocale" parameter

#### func  Format

```go
func Format(t time.Time, format string) string
```
Formats given date in "English" locale. See FormatLocale for details.

#### func  FormatLocale

```go
func FormatLocale(t time.Time, format string, locale DateLocale) string
```
Formats a given date according to the format string and date locale. Format
strings are specified like MMMM Do YYYY, h:mm:ss a, which would yield October
24th 2012, 12:42:03 am in English locale. Supported format string tokens:

    Month
    	M: 1 2 ... 12
    	MM: 01 01 ... 12
    	Mo: 1st 2nd ... 12th
    	MMM: Jan Feb ... Dec
    	MMMM: January February ... December
    Day of Month
    	D: 1 2 ... 31
    	DD: 01 02 ... 31
    	Do: 1st 2nd ... 31st
    Day of Week
    	ddd: Sun Mon ... Sat
    	dddd: Sunday Monday ... Saturday
    Year
    	YY: 70 71 ... 12
    	YYYY: 1970 1971 ... 2012
    Hour
    	H: 0 1 2 ... 23
    	HH: 00 01 02 .. 23
    	h: 1 2 ... 12
    	hh: 01 02 ... 12
    Minute
    	m: 0 1 2 ... 59
    	mm: 00 01 02 ... 59
    Second
    	s: 0 1 2 ... 59
    	ss: 00 01 02 ... 59
    AM / PM
    	A: AM PM
    	a: am pm
    Timezone
    	Z: -07:00 -06:00 ... +07:00
    	ZZ: -0700 -0600 ... +0700
    Escaping
    	It is possible to place arbitrary text within [] to display it as it is.
    Sample:
    	MMMM Do YYYY, h:mm:ss a [Foo] -> October 24th 2012, 1:52:27 am Foo

#### func  FormatMagic

```go
func FormatMagic(t time.Time, magic string, locale DateLocale) string
```
Uses built in time package semantics for format string, that is a magic date.
Instead of "MM", "H" style tokens that Format function accepts, you can supply a
sample date as "magic" parameter And it will be figured out as the format. Given
magic date layout must be

    Mon Jan 2 15:04:05 MST 2006

That is

    01/02 03:04:05PM '06 -0700

in a more memorable layout.

    Sample:
    magic = Mon, 02 Jan 2006 15:04:05 -0700
    result = Wed, 24 Oct 2012  2:10:11 +0300

#### type DateLocale

```go
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
```

DateLocale specifies how to format dates in different languages. FormatLocale
function accepts a DateLocale instance and formats the date according to given
locale.

Developed By
============

Ekin Koc - <ekin@eknkc.com>

License
=======

    Copyright 2012 Ekin Koc

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
