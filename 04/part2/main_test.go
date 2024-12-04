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
				MMMSXXMASM
				MSAMXMSMSA
				AMXSXMAAMM
				MSAMASMSMX
				XMASAMXAMM
				XXAMMXXAMA
				SMSMSASXSS
				SAXAMASAAA
				MAMMMXMMMM
				MXMXAXMASX
			`,
			Expected: 9,
		},
	}
	for i, tc := range tt {
		name := fmt.Sprintf("testcase%02d", i)
		t.Run(name, func(t *testing.T) {
			l := make([]string, 0)
			for _, s := range strings.Split(tc.Input, "\n") {
				s = strings.TrimSpace(s)
				if len(s) > 0 {
					l = append(l, s)
				}
			}
			input := strings.Join(l, "\n")
			require.Equal(t, tc.Expected, search(input))
		})
	}

}
