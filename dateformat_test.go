package dateformat

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	at := time.Date(2011, 2, 3, 4, 5, 5, 7, time.UTC)

	if Format(at, "YY YYYY MM MMM MMMM") != "11 2011 02 Feb February" {
		t.Fatal("YY YYYY MM MMM MMMM -> " + Format(at, "YY YYYY MM MMM MMMM"))
	}

	if Format(at, "M D DD ddd dddd") != "2 3 03 Thu Thursday" {
		t.Fatal("M D DD ddd dddd -> " + Format(at, "M D DD ddd dddd"))
	}

	if Format(at, "HH h hh m mm s ss A a Z ZZ") != "04 4 04 5 05 5 05 AM am +00:00 +0000" {
		t.Fatal("HH h hh m mm s ss A a Z ZZ -> " + Format(at, "HH h hh m mm s ss A a Z ZZ"))
	}
}
