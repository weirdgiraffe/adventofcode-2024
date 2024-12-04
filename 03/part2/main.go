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

func compute2(line string) int {
	acc := 0
	for len(line) > 0 {
		i := strings.Index(line, "don't()")
		if i < 0 {
			acc += compute(line)
			break
		}

		fmt.Println("enabled:", line[:i])
		acc += compute(line[:i])

		line = line[i+7:]

		j := strings.Index(line, "do()")
		if j < 0 {
			break
		}
		line = line[j+4:]
	}
	return acc
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
		acc += (a * b)
		fmt.Println("op:", line[i:j+1], "acc=", acc)
		line = line[j+1:]
	}

	return acc
}

func handleInput(path string) int {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("failed to read file: %v", err))
	}
	return compute2(strings.Join(lines, ""))
}
