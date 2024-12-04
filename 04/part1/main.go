package main

import (
	"container/ring"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(solve("../input"))
}

func NewRingBuffer() *ring.Ring {
	r := ring.New(len("XMAS"))
	for i := 0; i < r.Len(); i++ {
		r.Value = uint8('.')
		r = r.Next()
	}
	return r
}

func GetWord(r *ring.Ring) string {
	b := make([]byte, 0, r.Len())
	r.Do(func(c any) {
		b = append(b, c.(byte))
	})
	return string(b)
}

func MatchWord(r *ring.Ring) bool {
	word := GetWord(r)
	return word == "XMAS" || word == "SAMX"
}

func searchHorizontal(l []string) int {
	acc := 0
	w := len(l[0])
	h := len(l)
	for i := 0; i < h; i++ {
		r := NewRingBuffer()
		for j := 0; j < w; j++ {
			r.Value = l[i][j]
			r = r.Next()
			if MatchWord(r) {
				acc++
			}
		}
	}
	return acc
}

func searchVertical(l []string) int {
	acc := 0
	w := len(l[0])
	h := len(l)
	for j := 0; j < w; j++ {
		r := NewRingBuffer()
		for i := 0; i < h; i++ {
			r.Value = l[i][j]
			r = r.Next()
			if MatchWord(r) {
				acc++
			}
		}
	}
	return acc
}

func searchDiagDown(l []string) int {
	acc := 0
	w := len(l[0])
	h := len(l)

	for iofft := 0; iofft < h; iofft++ {
		r := NewRingBuffer()
		i := iofft
		j := 0
		for i >= 0 && j < w {
			r.Value = l[i][j]
			r = r.Next()
			if MatchWord(r) {
				acc++
			}
			i--
			j++
		}
	}

	for jofft := 1; jofft < w; jofft++ {
		r := NewRingBuffer()
		i := h - 1
		j := jofft
		for i >= 0 && j < w {
			r.Value = l[i][j]
			r = r.Next()
			if MatchWord(r) {
				acc++
			}
			i--
			j++
		}
	}
	return acc
}

func searchDiagUp(l []string) int {
	acc := 0
	w := len(l[0])
	h := len(l)

	for iofft := h - 1; iofft >= 0; iofft-- {
		r := NewRingBuffer()
		i := iofft
		j := 0
		for i < h && j < w {
			r.Value = l[i][j]
			r = r.Next()
			if MatchWord(r) {
				acc++
			}
			i++
			j++
		}
	}

	for jofft := 1; jofft < w; jofft++ {
		r := NewRingBuffer()
		i := 0
		j := jofft
		for i < h && j < w {
			r.Value = l[i][j]
			r = r.Next()
			if MatchWord(r) {
				acc++
			}
			i++
			j++
		}
	}
	return acc
}

func searchDiag(l []string) int {
	return searchDiagDown(l) + searchDiagUp(l)
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

func search(text string) int {
	arr := asArray(text)
	return searchDiag(arr) + searchHorizontal(arr) + searchVertical(arr)
}

func solve(path string) int {
	text, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	return search(string(text))
}
