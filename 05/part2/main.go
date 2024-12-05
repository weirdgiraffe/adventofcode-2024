package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solveInput("../input"))
}

func solve(r io.Reader) int {
	acc := 0

	rules, updates := parseInput(r)
	fmt.Println("rules:", rules)

	for _, update := range updates {
		fmt.Println("update:", update)
		if !IsCorrectUpate(update, rules) {
			correct := Reorder(update, rules)
			if len(correct)%2 == 0 {
				panic("odd number of pages in update")
			}
			acc += update[len(correct)/2]
		}
	}
	return acc
}

func Reorder(update []int, r Rules) []int {
	for i, e := range update {
		after := r[e]
		for j := i + 1; j < len(update); j++ {
			n := update[j]
			if !after[n] {
				// swap elemenets
				fmt.Println("not in order:", update[:j+1])
				update[i], update[j] = update[j], update[i]
				fmt.Println("     swapped:", update[:j+1])
				return Reorder(update, r)
			}
		}
	}
	fmt.Println("fixed:", update)
	return update
}

func IsCorrectUpate(update []int, r Rules) bool {
	for i := len(update) - 1; i >= 0; i-- {
		page := update[i]
		if !verify(update[:i], r[page]) {
			fmt.Printf("update is not correct. Rule for %d is not fulfilled\n", page)
			return false
		}
	}
	return true
}

func verify(pages []int, forbidden Set) bool {
	for _, page := range pages {
		if forbidden[page] {
			return false
		}
	}
	return true
}

func solveInput(path string) int {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer f.Close()
	return solve(f)
}

func parseInput(r io.Reader) (rules Rules, updates [][]int) {
	sc := bufio.NewScanner(r)
	parseUpdates := false
	rules = make(Rules)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			parseUpdates = true
			continue
		}
		if parseUpdates {
			updates = append(updates, MustParseUpdate(sc.Text()))
		} else {
			MustParseRule(rules, sc.Text())
		}
	}
	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("failed to read file: %v", err))
	}
	return
}

type Set map[int]bool

type Rules map[int]Set

func MustParseRule(r Rules, line string) {
	l := strings.Split(line, "|")
	if len(l) != 2 {
		panic(fmt.Sprintf("invalid rule: %q", line))
	}
	first, err := strconv.Atoi(strings.TrimSpace(l[0]))
	if err != nil {
		panic(fmt.Sprintf("failed to parse first index: %v", err))
	}
	second, err := strconv.Atoi(strings.TrimSpace(l[1]))
	if err != nil {
		panic(fmt.Sprintf("failed to parse second index: %v", err))
	}
	if _, ok := r[first]; !ok {
		r[first] = make(Set)
	}
	r[first][second] = true
}

func MustParseUpdate(line string) []int {
	l := strings.Split(line, ",")
	res := make([]int, len(l))
	for i := range l {
		index, err := strconv.Atoi(strings.TrimSpace(l[i]))
		if err != nil {
			panic(fmt.Sprintf("failed to parse index[%d]: %v", i, err))
		}
		res[i] = index
	}
	return res
}
