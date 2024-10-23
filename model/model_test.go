package model

import (
	"testing"
)

type testCase struct {
	name           string
	expectedResult bool
}

var teamTestCases = []testCase{
	{"AAA", true}, {"AGM", true}, {"ZZZ", true}, {"AZA", true}, {"aAA", false},
	{"AaA", false}, {"AAa", false}, {"AZ1", false}, {"", false}, {"AZ ", false},
	{" AZ", false}, {"A-Z", false}, {"A*Z", false}, {"AZ(", false},
	{"AZHG", false}, {"AA A", false}, {"103", false}, {"treta", false},
	{"jjj", false},
}

var workTypeTestCases = []testCase{
	{"AA", true}, {"AG", true}, {"ZZ", true}, {"AZ", true}, {"JJ", true},
	{"aA", false}, {"Aa", false}, {"A1", false}, {"", false}, {"A ", false},
	{" Z", false}, {"A-", false}, {"*Z", false}, {"Z(", false}, {"AZHG", false},
	{"  ", false}, {"10", false}, {"treta", false}, {"jj", false},
}

func TestTeam(t *testing.T) {
	t.Run("isValidName", func(t *testing.T) {
		for _, tc := range teamTestCases {
			result := isValidTeamName(tc.name)
			if result != tc.expectedResult {
				t.Logf(`expected name "%s" to be %v, but got %v`, tc.name, tc.expectedResult, result)
				t.Fail()
			}
		}
	})
}

func TestWorkType(t *testing.T) {
	t.Run("isValidName", func(t *testing.T) {
		for _, tc := range workTypeTestCases {
			result := isValidWorkTypeName(tc.name)
			if result != tc.expectedResult {
				t.Logf(`expected name "%s" to be %v, but got %v`, tc.name, tc.expectedResult, result)
				t.Fail()
			}
		}
	})
}
