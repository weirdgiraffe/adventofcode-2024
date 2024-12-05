package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	tt := []struct {
		Input    string
		Expected int
	}{{
		Input: `
			47|53
			97|13
			97|61
			97|47
			75|29
			61|13
			75|53
			29|13
			97|29
			53|29
			61|53
			97|53
			61|29
			47|13
			75|47
			97|75
			47|61
			75|61
			47|29
			75|13
			53|13

			75,47,61,53,29
			97,61,53,29,13
			75,29,13
			75,97,47,61,53
			61,13,29
			97,13,75,29,47
		`,
		Expected: 123,
	}}

	for i, tc := range tt {
		name := fmt.Sprintf("testcase%02d", i)
		t.Run(name, func(t *testing.T) {
			input := strings.TrimSpace(tc.Input)
			require.Equal(t, tc.Expected, solve(strings.NewReader(input)))
		})
	}
}
