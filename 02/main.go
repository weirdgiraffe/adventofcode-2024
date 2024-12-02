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
	reports := readReports("input")
	sum := 0
	for i := range reports {
		sum += int(safeReport(reports[i]))
	}
	fmt.Println(sum)
}

func safeReport(l []int) int {
	prev := l[0]
	asc := l[1] > l[0]
	for i := 1; i < len(l); i++ {
		if l[i] == prev {
			return 0
		}
		if asc {
			if l[i] < prev {
				return 0
			}
			if l[i]-prev > 3 {
				return 0
			}
		} else {
			if l[i] > prev {
				return 0
			}
			if prev-l[i] > 3 {
				return 0
			}
		}
		prev = l[i]
	}
	return 1
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
