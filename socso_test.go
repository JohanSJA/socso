package socso

import (
	"testing"
)

func TestRateEntryTotal(t *testing.T) {
	re := RateEntry{0.0, 30.0, 0.4, 0.1, 0.3}
	if re.Category1Total() != 0.5 {
		t.Fail()
	}
}

func TestRateWagesLessThanZero(t *testing.T) {
	err, _ := Rate(-1.0)
	if err == nil {
		t.Fail()
	}
}

func TestRateWagesNormal(t *testing.T) {
	expected_re := RateEntry{1400.0, 1500.0, 25.35, 7.25, 18.1}
	err, re := Rate(1450.0)
	if err != nil {
		t.Fail()
	}
	if re != expected_re {
		t.Errorf("%v !== %v", re, expected_re)
	}
}
