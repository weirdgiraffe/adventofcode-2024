package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(MustSolveFile("../input"))
}

func MustSolveFile(path string) int {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer f.Close()
	return MustSolveInput(f)
}

func MustSolveInput(r io.Reader) int {
	text, err := io.ReadAll(r)
	if err != nil {
		panic("failed to read input")
	}
	var l []Testcase
	for _, line := range strings.Split(strings.TrimSpace(string(text)), "\n") {
		if line == "" {
			continue
		}
		l = append(l, MustParseTestCase(line))
	}
	return Solve(l)
}

func MustParseTestCase(line string) (res Testcase) {
	var err error
	l := strings.Split(line, ":")
	res.Value, err = strconv.Atoi(strings.TrimSpace(l[0]))
	if err != nil {
		fmt.Println("l=", l)
		fmt.Println("l[0]=", l[0])
		panic("failed to parse value")
	}
	l = strings.Split(l[1], " ")
	for _, s := range l {
		if s == "" {
			continue
		}
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("l=", l)
			fmt.Println("s=", s)
			panic("failed to parse number")
		}
		res.Numbers = append(res.Numbers, num)
	}
	return res
}

func Solve(tt []Testcase) int {
	acc := 0
	for _, tc := range tt {
		if tc.CanBeCombined() {
			acc += tc.Value
		}
	}
	return acc
}

type Testcase struct {
	Value   int
	Numbers []int
}

func (t Testcase) CanBeCombined() bool {
	fmt.Println("t.Value", t.Value)
	fmt.Println("t.Numbers", t.Numbers)
	return canBeCombined(t.Value, t.Numbers, []string{})
}

func canBeCombined(value int, number []int, ops []string) bool {
	fmt.Println(
		"value=", value,
		"number=", number,
		"ops=", ops,
	)

	if len(number) == 0 {
		res := value == 0
		if len(ops) > 0 && ops[len(ops)-1] == "*" {
			res = value == 1
		}
		fmt.Println(
			"value=", value,
			"number=", number,
			"ops=", ops,
			"result=", res,
		)
		return res
	}

	last := number[len(number)-1]
	divisible := value%last == 0
	if divisible {
		if canBeCombined(value/last, number[:len(number)-1], append(ops, "*")) {
			fmt.Println(
				"value=", value,
				"number=", number,
				"ops=", ops,
				"result=", true,
			)
			return true
		}
	}

	return canBeCombined(value-last, number[:len(number)-1], append(ops, "+"))
}
