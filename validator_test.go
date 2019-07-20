package numbertowords

import (
	"fmt"
	"testing"
)

func TestValidator(t *testing.T) {
	var testCases = []struct {
		given  string
		expect bool
	}{
		{"th", true},
		{"en", true},
		{"cn", false},
		{"jp", false},
		{"kr", false},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("tc %d", idx+1), func(t *testing.T) {
			expect := validateLang(tc.given)

			if tc.expect != expect {
				t.Errorf("Expect `%v` got `%v`", tc.expect, expect)
			}
		})
	}
}
