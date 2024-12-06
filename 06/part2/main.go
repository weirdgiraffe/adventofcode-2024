package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(solveFile("../input"))
}

func solveFile(path string) int {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer f.Close()
	return solve(f)
}

func solve(r io.Reader) int {
	text, err := io.ReadAll(r)
	if err != nil {
		panic("failed to read input")
	}
	return patrol(string(text))
}

func patrolPath(m *Map) (path []Position) {
	dir := DirectionUp
	pos := m.StartingPosition()
	for {
		if len(path) == 0 || path[len(path)-1] != pos {
			path = append(path, pos)
		}
		nextPos := m.Step(pos, dir)
		if !m.Contains(nextPos) {
			return path
		}
		if m.At(nextPos) == '#' {
			dir = TurnRight(dir)
			continue
		}
		pos = nextPos
	}
}

func patrol(text string) int {
	m := asMap(text)
	path := patrolPath(&m)
	fmt.Println(path)
	return 0
}

const (
	DirectionUp    = 0
	DirectionRight = 1
	DirectionDown  = 2
	DirectionLeft  = 3
)

func DirectionSym(d int) string {
	switch d {
	case DirectionUp:
		return "^"
	case DirectionRight:
		return ">"
	case DirectionDown:
		return "v"
	case DirectionLeft:
		return "<"
	default:
		panic("invalid direction")
	}
}

func asMap(text string) Map {
	l := make([]string, 0)
	for _, s := range strings.Split(text, "\n") {
		s = strings.TrimSpace(s)
		if len(s) > 0 {
			l = append(l, s)
		}
	}
	return Map{
		cells:  l,
		height: len(l),
		width:  len(l[0]),
	}
}

type Position struct {
	X int
	Y int
}

type Map struct {
	cells  []string
	height int
	width  int
}

func (m Map) StartingPosition() Position {
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if m.cells[y][x] == '^' {
				return Position{X: x, Y: y}
			}
		}
	}
	panic("there should be a starting position")
}

func (m Map) Contains(p Position) bool {
	return (p.X >= 0 && p.X < m.width) && (p.Y >= 0 && p.Y < m.height)
}

func (m Map) At(p Position) byte {
	return m.cells[p.Y][p.X]
}

func (m Map) Step(p Position, direction int) Position {
	switch direction {
	case DirectionUp:
		p.Y--
	case DirectionRight:
		p.X++
	case DirectionDown:
		p.Y++
	case DirectionLeft:
		p.X--
	}
	return p
}

func TurnRight(direction int) int {
	// 0 + 1 = 1 % 4 = 1
	// 1 + 1 = 2 % 4 = 2
	// 2 + 1 = 3 % 4 = 3
	// 3 + 1 = 4 % 4 = 0
	return (direction + 1) % 4
}
