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
				....#.....
				.........#
				..........
				..#.......
				.......#..
				..........
				.#..^.....
				........#.
				#.........
				......#...
			`,
			Expected: 41,
		},
	}
	for i, tc := range tt {
		name := fmt.Sprintf("testcase%02d", i)
		t.Run(name, func(t *testing.T) {
			input := strings.NewReader(strings.TrimSpace(tc.Input))
			require.Equal(t, tc.Expected, solve(input))
		})
	}

}
