package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	left, right, err := readLists("input")
	if err != nil {
		log.Fatalf("failed to read lists: %v", err)
	}
	lm := toCounts(left)
	rm := toCounts(right)
	sum := 0
	for k := range lm {
		v := rm[k]
		sum += (k * v)
	}
	fmt.Println(sum)
}

func toCounts(l []int) map[int]int {
	m := make(map[int]int)
	for _, n := range l {
		m[n] = m[n] + 1
	}
	return m
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
