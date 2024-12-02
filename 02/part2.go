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
	// 	reports := [][]int{
	// 		{7, 6, 4, 2, 1},
	// 		{1, 2, 7, 8, 9},
	// 		{9, 7, 6, 2, 1},
	// 		{1, 3, 2, 4, 5},
	// 		{8, 6, 4, 4, 1},
	// 		{1, 3, 6, 7, 9},
	// 	}
	sum := 0
	for i := range reports {
		n := safeReport(reports[i])
		if n == 0 {
			fmt.Println("unsafe:", reports[i])
		} else {
			fmt.Println("safe:", reports[i])
		}
		sum += n
	}
	fmt.Println(sum)
}

func safeReport(l []int) int {
	asc := l[0] < l[1]
	demped := false
	i := 0
	for i < len(l)-1 {
		if !isSafeLevel(l[i], l[i+1], asc) {
			if !demped {
				demped = true
				if i-1 >= 0 && isSafeLevel(l[i-1], l[i+1], asc) {
					fmt.Printf("demped: skip %d\n", l[i])
					i++
					continue
				}
				if i+2 == len(l) || isSafeLevel(l[i], l[i+2], asc) {
					fmt.Printf("demped: skip %d\n", l[i+1])
					i += 2
					continue
				}
			}
			return 0
		}
		i++
	}
	return 1
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func isSafeLevel(a, b int, asc bool) bool {
	fmt.Printf("%d %d: ", a, b)
	if a == b {
		fmt.Println("cond 1")
		return false
	}
	if asc && a > b {
		fmt.Println("cond 2")
		return false
	}
	if !asc && a < b {
		fmt.Println("cond 3")
		return false
	}
	if abs(a-b) > 3 {
		fmt.Println("cond 4")
		return false
	}
	fmt.Println("safe")
	return true
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
