package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	left, right, err := readLists("input1")
	if err != nil {
		log.Fatalf("failed to read lists: %v", err)
	}
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := range left {
		sum += abs(left[i] - right[i])
	}
	fmt.Println(sum)
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func readLists(path string) (left, right []int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	ln := 0
	for sc.Scan() {
		ln++
		s := sc.Text()
		i := 0
		for ; i < len(s) && (s[i] >= '0' && s[i] <= '9'); i++ {
		}
		a, err := strconv.Atoi(s[:i])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse left number %q at line %d: %w", s[:i], ln, err)
		}
		s = s[i:]
		i = 0
		for ; i < len(s) && (s[i] < '0' || s[i] > '9'); i++ {
		}
		b, err := strconv.Atoi(s[i:])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse left number %q at line %d: %w", s[i:], ln, err)
		}
		left = append(left, a)
		right = append(right, b)
	}
	if err := sc.Err(); err != nil {
		return nil, nil, err
	}
	return left, right, nil
}
