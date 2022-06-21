package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2021-12-31T19:50:00")

	if convertedTime.Year() != 2021 {
		t.Errorf("Espera que o ano seja 2021")
	}

	if convertedTime.Month() != 12 {
		t.Errorf("Espera que o mes seja 12")
	}

	if convertedTime.Hour() != 19 {
		t.Errorf("Espera que a hora seja 19")
	}
}
