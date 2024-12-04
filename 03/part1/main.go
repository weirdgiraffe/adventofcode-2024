package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(handleInput("../input"))
}

func compute(line string) int {
	acc := 0
	for {
		i := strings.Index(line, "mul(")
		if i < 0 {
			break
		}
		// we have op start
		j := strings.IndexByte(line[i+4:], ')')
		if j < 0 {
			line = line[i+4:]
			continue
		}
		j = j + (i + 4)

		if j-i > 12 {
			// skip line because it's way too long
			line = line[i+4:]
			continue
		}

		args := strings.Split(line[i+4:j], ",")
		if len(args) != 2 {
			// skip line because it doesn't have , char
			line = line[i+4:]
			continue
		}

		a, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("arg[0] %q: %v\n", args[0], err)
			line = line[i+4:]
			continue
		}

		b, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("arg[1] %q: %v\n", args[1], err)
			line = line[i+4:]
			continue
		}

		// we have op end
		fmt.Println("op:", line[i:j+1])
		acc += (a * b)
		line = line[j+1:]
	}

	return acc
}

func handleInput(path string) int {
	acc := 0

	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		acc += compute(sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("failed to read file: %v", err))
	}
	return acc
}
