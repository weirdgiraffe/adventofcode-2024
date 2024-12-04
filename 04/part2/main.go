package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(solve("../input"))
}

func Match(square [3][3]byte) bool {
	/*
		M.S
		.A.
		M.S
	*/
	if square[0][0] == 'M' && square[0][2] == 'S' {
		if square[2][0] == 'M' && square[2][2] == 'S' {
			return square[1][1] == 'A'
		}
	}
	/*
		S.M
		.A.
		S.M
	*/
	if square[0][0] == 'S' && square[0][2] == 'M' {
		if square[2][0] == 'S' && square[2][2] == 'M' {
			return square[1][1] == 'A'
		}
	}
	/*
		S.S
		.A.
		M.M
	*/
	if square[0][0] == 'S' && square[0][2] == 'S' {
		if square[2][0] == 'M' && square[2][2] == 'M' {
			return square[1][1] == 'A'
		}
	}
	/*
		M.M
		.A.
		S.S
	*/
	if square[0][0] == 'M' && square[0][2] == 'M' {
		if square[2][0] == 'S' && square[2][2] == 'S' {
			return square[1][1] == 'A'
		}
	}
	return false
}

func printSquare(square [3][3]byte) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c := square[i][j]
			if c == 0 {
				c = '.'
			}
			fmt.Printf("[%c]", c)
		}
		fmt.Println()
	}
}

func search(text string) int {
	arr := asArray(text)

	h := len(arr)
	w := len(arr[0])
	acc := 0

	for i := 0; i < h-2; i++ {
		for j := 0; j < w-2; j++ {
			var square [3][3]byte
			square[0][0] = arr[i][j]
			square[0][2] = arr[i][j+2]
			square[1][1] = arr[i+1][j+1]
			square[2][0] = arr[i+2][j]
			square[2][2] = arr[i+2][j+2]
			if Match(square) {
				acc++
			}
		}
	}

	return acc
}

func asArray(text string) []string {
	l := make([]string, 0)
	for _, s := range strings.Split(text, "\n") {
		s = strings.TrimSpace(s)
		if len(s) > 0 {
			l = append(l, s)
		}
	}
	return l
}

func solve(path string) int {
	text, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	return search(string(text))
}
