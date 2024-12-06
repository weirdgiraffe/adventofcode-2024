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

type PositionWithDirection struct {
	Position
	Direction int
}

func patrolPath(m *Map, pos Position, trace bool) (path []Position, loop bool) {
	visited := make(map[PositionWithDirection]bool)
	dir := DirectionUp
	for {
		pd := PositionWithDirection{Position: pos, Direction: dir}
		if visited[pd] {
			return path, true
		}
		path = append(path, pos)
		visited[pd] = true
		if trace {
			m.cells[pos.Y][pos.X] = DirectionSym(dir)
		}
		nextPos := m.Step(pos, dir)
		if !m.Contains(nextPos) {
			return path, false
		}

		if m.At(nextPos) == '#' {
			dir = TurnRight(dir)
			path = path[:len(path)-1]
			continue
		}

		pos = nextPos
	}

}

func patrol(text string) int {
	m := asMap(text)
	startingPos := m.StartingPosition()
	acc := 0
	path, _ := patrolPath(&m, startingPos, false)

	uniq := make(map[Position]bool)
	for i := range path {
		uniq[path[i]] = true
	}
	for pos := range uniq {
		if pos == startingPos {
			continue
		}
		mm := m.WithSymbolAt(pos, '#')
		if _, loop := patrolPath(&mm, startingPos, true); loop {
			// 			fmt.Println("loop detected")
			// 			mm.Print(startingPos)
			acc++
		}
	}
	return acc
}

const (
	DirectionUp    = 0
	DirectionRight = 1
	DirectionDown  = 2
	DirectionLeft  = 3
)

func DirectionSym(d int) byte {
	switch d {
	case DirectionUp:
		return '^'
	case DirectionRight:
		return '>'
	case DirectionDown:
		return 'v'
	case DirectionLeft:
		return '<'
	default:
		panic("invalid direction")
	}
}

func asMap(text string) Map {
	l := make([][]byte, 0)
	for _, s := range strings.Split(text, "\n") {
		s = strings.TrimSpace(s)
		if len(s) > 0 {
			l = append(l, []byte(s))
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
	cells  [][]byte
	height int
	width  int
}

func (m Map) Print(startingPos Position) {
	for i := range m.cells {
		for j := range m.cells[i] {
			if j == startingPos.X && i == startingPos.Y {
				fmt.Print(" * ")
			} else {
				fmt.Printf(" %c ", m.cells[i][j])
			}
		}
		fmt.Println()
	}
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

func (m Map) WithSymbolAt(p Position, c byte) Map {
	cells := make([][]byte, len(m.cells))
	for i := range cells {
		cells[i] = make([]byte, len(m.cells[i]))
		copy(cells[i], m.cells[i])
	}
	cells[p.Y][p.X] = c
	return Map{
		cells:  cells,
		height: m.height,
		width:  m.width,
	}
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
