package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompute(t *testing.T) {
	tt := []struct {
		Input    string
		Expected int
	}{
		{
			Input:    `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
			Expected: 48,
		},
	}
	for i, tc := range tt {
		name := fmt.Sprintf("testcase%02d", i)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.Expected, compute2(tc.Input))
		})
	}

}
