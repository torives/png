package model

import (
	"testing"
)

type testCase struct {
	name           string
	expectedResult error
}

var teamTestCases = []testCase{
	{"AAA", nil}, {"AGM", nil}, {"ZZZ", nil}, {"AZA", nil}, {"aAA", ErrInvalidTeamName},
	{"AaA", ErrInvalidTeamName}, {"AAa", ErrInvalidTeamName}, {"AZ1", ErrInvalidTeamName},
	{"", ErrInvalidTeamName}, {"AZ ", ErrInvalidTeamName}, {" AZ", ErrInvalidTeamName},
	{"A-Z", ErrInvalidTeamName}, {"A*Z", ErrInvalidTeamName}, {"AZ(", ErrInvalidTeamName},
	{"AZHG", ErrInvalidTeamName}, {"AA A", ErrInvalidTeamName}, {"103", ErrInvalidTeamName},
	{"treta", ErrInvalidTeamName}, {"jjj", ErrInvalidTeamName},
}

var workTypeTestCases = []testCase{
	{"AA", nil}, {"AG", nil}, {"ZZ", nil}, {"AZ", nil}, {"JJ", nil},
	{"aA", ErrInvalidWorkTypeName}, {"Aa", ErrInvalidWorkTypeName}, {"A1", ErrInvalidWorkTypeName},
	{"", ErrInvalidWorkTypeName}, {"A ", ErrInvalidWorkTypeName}, {" Z", ErrInvalidWorkTypeName},
	{"A-", ErrInvalidWorkTypeName}, {"*Z", ErrInvalidWorkTypeName}, {"Z(", ErrInvalidWorkTypeName},
	{"AZHG", ErrInvalidWorkTypeName}, {"  ", ErrInvalidWorkTypeName}, {"10", ErrInvalidWorkTypeName},
	{"treta", ErrInvalidWorkTypeName}, {"jj", ErrInvalidWorkTypeName},
}

func TestTeam(t *testing.T) {
	t.Run("isValidName", func(t *testing.T) {
		for _, tc := range teamTestCases {
			result := ValidateTeamName(tc.name)
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
			result := ValidateWorkTypeName(tc.name)
			if result != tc.expectedResult {
				t.Logf(`expected name "%s" to be %v, but got %v`, tc.name, tc.expectedResult, result)
				t.Fail()
			}
		}
	})
}
