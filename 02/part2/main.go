package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reports := readReports("../input")
	sum := 0
	for i := range reports {
		if IsSafeReport(reports[i]) {
			fmt.Println("safe:", reports[i])
			sum++
		} else {
			fmt.Println("unsafe:", reports[i])
		}
	}
	fmt.Println(sum)
}

func IsSafeReport(l []int) bool {
	// check first element to figure out if we need to
	// change the direction at the beginning
	if IsSafe(l[0], l[1], true) {
		return safeReport(l[1:], true, 0)
	}
	if IsSafe(l[0], l[1], false) {
		return safeReport(l[1:], false, 0)
	}
	if IsSafe(l[1], l[2], true) {
		return safeReport(l[2:], true, 1)
	}
	if IsSafe(l[1], l[2], false) {
		return safeReport(l[2:], false, 1)
	}
	return false
}

func safeReport(l []int, ascending bool, errs int) bool {
	if errs > 1 {
		return false
	}
	if len(l) < 2 {
		return true
	}
	if !IsSafe(l[0], l[1], ascending) {
		errs++
		if len(l) > 2 {
			if IsSafe(l[1], l[2], ascending) {
				return safeReport(l[2:], ascending, errs)
			}
			if IsSafe(l[0], l[2], ascending) {
				copy(l[1:], l[2:])
				l = l[:len(l)-1]
				return safeReport(l[2:], ascending, errs)
			}
			return false
		}
	}
	return safeReport(l[1:], ascending, errs)
}

func IsSafe(a, b int, ascending bool) bool {
	if a < b && !ascending {
		return false
	}
	if a > b && ascending {
		return false
	}
	if abs(a-b) > 3 {
		return false
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func readReports(path string) (l [][]int) {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	ln := 0
	for sc.Scan() {
		ln++
		report := slices.Collect(func(yield func(int) bool) {
			for _, s := range strings.Split(sc.Text(), " ") {
				i, err := strconv.Atoi(s)
				if err != nil {
					panic(fmt.Sprintf("failed to parse number %q at line %d: %w", s, ln, err))
				}
				if !yield(i) {
					return
				}
			}
		})
		l = append(l, report)
	}
	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("failed to read file: %v", err))
	}
	return l
}
