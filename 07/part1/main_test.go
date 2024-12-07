package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompute(t *testing.T) {
	tt := []struct {
		Input    string
		Expected int
	}{
		{
			Input: `
				190: 10 19
				3267: 81 40 27
				83: 17 5
				156: 15 6
				7290: 6 8 6 15
				161011: 16 10 13
				192: 17 8 14
				21037: 9 7 18 13
				292: 11 6 16 20
			`,
			Expected: 3749,
		},
	}
	for i, tc := range tt {
		name := fmt.Sprintf("testcase%02d", i)
		t.Run(name, func(t *testing.T) {
			input := strings.NewReader(strings.TrimSpace(tc.Input))
			require.Equal(t, tc.Expected, MustSolveInput(input))
		})
	}

}
