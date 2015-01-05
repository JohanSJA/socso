package socso

/**
 * This is a library to help in the checking against the Rate of Contribution
 * table as published at
 * http://www.perkeso.gov.my/en/social-security-protection/employer-employee-eligibilty/rate-of-contributions.html .
 */

import (
	"errors"
	"math"
)

/**
 * A RateEntry represents a row in the Rate of Contribution table.
 */
type RateEntry struct {
	WagesFrom         float64
	WagesTo           float64
	Category1Employer float64
	Category1Employee float64
	Category2Employer float64
}

/**
 * Simple function adding up the employer and employee contribution for
 * category 1.
 */
func (r *RateEntry) Category1Total() float64 {
	return r.Category1Employer + r.Category1Employee
}

var RATES []RateEntry

func init() {
	RATES = []RateEntry{
		{0.0, 30.0, 0.4, 0.1, 0.3},
		{30.0, 50.0, 0.7, 0.2, 0.5},
		{50.0, 70.0, 1.1, 0.3, 0.8},
		{70.0, 100.0, 1.5, 0.4, 1.1},
		{100.0, 140.0, 2.1, 0.6, 1.5},
		{140.0, 200.0, 2.95, 0.85, 2.1},
		{200.0, 300.0, 4.35, 1.25, 3.1},
		{300.0, 400.0, 6.15, 1.75, 4.4},
		{400.0, 500.0, 7.85, 2.25, 5.6},
		{500.0, 600.0, 9.65, 2.75, 6.9},
		{600.0, 700.0, 11.35, 3.25, 8.1},
		{700.0, 800.0, 13.15, 3.75, 9.4},
		{800.0, 900.0, 14.85, 4.25, 10.6},
		{900.0, 1000.0, 16.65, 4.75, 11.9},
		{1000.0, 1100.0, 18.35, 5.25, 13.1},
		{1100.0, 1200.0, 20.15, 5.75, 14.4},
		{1200.0, 1300.0, 21.85, 6.25, 15.6},
		{1300.0, 1400.0, 23.65, 6.75, 16.9},
		{1400.0, 1500.0, 25.35, 7.25, 18.1},
		{1500.0, 1600.0, 27.15, 7.75, 19.4},
		{1600.0, 1700.0, 28.85, 8.25, 20.6},
		{1700.0, 1800.0, 30.65, 8.75, 21.9},
		{1800.0, 1900.0, 32.35, 9.25, 23.1},
		{1900.0, 2000.0, 34.15, 9.75, 24.4},
		{2000.0, 2100.0, 35.85, 10.25, 25.6},
		{2100.0, 2200.0, 37.65, 10.75, 26.9},
		{2200.0, 2300.0, 39.35, 11.25, 28.1},
		{2300.0, 2400.0, 41.15, 11.75, 29.4},
		{2400.0, 2500.0, 42.85, 12.25, 30.6},
		{2500.0, 2600.0, 44.65, 12.75, 31.9},
		{2600.0, 2700.0, 46.35, 13.25, 33.1},
		{2700.0, 2800.0, 48.15, 13.75, 34.4},
		{2800.0, 2900.0, 49.85, 14.25, 35.6},
		{2900.0, math.MaxFloat64, 51.65, 14.75, 36.9},
	}
}

/**
 * Return corresponding RateEntry instance if possible.
 * Wages is not supposed to be less than 0.
 */
func Rate(wages float64) (error, RateEntry) {
	if wages <= 0 {
		return errors.New("Wages cannot be less than 0."), RateEntry{}
	}
	for _, entry := range RATES {
		if wages > entry.WagesFrom && wages <= entry.WagesTo {
			return nil, entry
		}
	}
	return errors.New("Unknown error."), RateEntry{}
}
