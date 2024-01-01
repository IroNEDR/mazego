package mazego

import (
	"math/rand"
)

const (
	DefaultColumns    uint8   = 10
	DefaultRows       uint8   = 10
	DefaultSparseness float32 = 0.2
)

const (
	Mouse uint8 = iota
	Cheese
	Empty
	Path
	Wall
)

type Cell struct {
	Column  uint8
	Row     uint8
	Visited bool
	Content uint8
	Parent  *Cell
}

type Maze struct {
	Rows       uint8
	Columns    uint8
	Sparseness float32
	Cells      [][]Cell
	Start      Cell
	Goal       Cell
}

func NewMaze(columns, rows uint8, sparseness float32) *Maze {
	m := Maze{Columns: columns, Rows: rows, Sparseness: sparseness}
	m.generateRandomMaze()
	m.Start = Cell{0, 0, true, Mouse, nil}
	m.Goal = Cell{rows - 1, columns - 1, false, Cheese, nil}
	m.Cells[m.Start.Row][m.Start.Column] = m.Start
	m.Cells[m.Goal.Row][m.Goal.Column] = m.Goal
	return &m
}

func (m *Maze) generateRandomMaze() {
	m.Cells = make([][]Cell, m.Rows)
	for row := uint8(0); row < m.Rows; row++ {
		m.Cells[row] = make([]Cell, m.Columns)
		for column := uint8(0); column < m.Columns; column++ {
			if rand.Float32() < m.Sparseness {
				m.Cells[row][column] = Cell{Row: row, Column: column, Visited: false, Content: Wall, Parent: nil}
			} else {
				m.Cells[row][column] = Cell{Row: row, Column: column, Visited: false, Content: Empty, Parent: nil}
			}
		}
	}
}

func (m *Maze) GetNeighbours(cell *Cell) []*Cell {
	var neighbours []*Cell
	if cell.Column+1 < m.Columns && (m.Cells[cell.Row][cell.Column+1]).Content != Wall {
		neighbours = append(neighbours, &m.Cells[cell.Row][cell.Column+1])
	}
	if cell.Column > 0 && (m.Cells[cell.Row][cell.Column-1]).Content != Wall {
		neighbours = append(neighbours, &m.Cells[cell.Row][cell.Column-1])
	}
	if cell.Row+1 < m.Rows && (m.Cells[cell.Row+1][cell.Column]).Content != Wall {
		neighbours = append(neighbours, &m.Cells[cell.Row+1][cell.Column])
	}
	if cell.Row > 0 && (m.Cells[cell.Row-1][cell.Column]).Content != Wall {
		neighbours = append(neighbours, &m.Cells[cell.Row-1][cell.Column])
	}
	return neighbours
}

func (m *Maze) DFS() (*Cell, error) {
	cellStack := NewStack[Cell]()
	cellStack.Push(&m.Start)
	for cellStack.Len() > 0 {
		current, err := cellStack.Pop()
		if err != nil {
			return nil, err
		}
		if current.Content == m.Goal.Content {
			return current, nil
		}
		neighbours := m.GetNeighbours(current)
		for _, neighbour := range neighbours {
			if neighbour.Visited {
				continue
			}
			neighbour.Visited = true
			neighbour.Parent = current
			cellStack.Push(neighbour)
		}
	}
	return nil, nil
}

func (m *Maze) BFS() {

}
