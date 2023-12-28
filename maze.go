package mazego

import "math/rand"

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
}

type Maze struct {
	Columns    uint8
	Rows       uint8
	Sparseness float32
	Cells      [][]Cell
	start      Cell
	end        Cell
}

func NewMaze(columns, rows uint8, sparseness float32) *Maze {
	m := Maze{Columns: columns, Rows: rows, Sparseness: sparseness}
	m.generateCells()
	m.start = Cell{0, 0, false, Mouse}
	m.end = Cell{rows - 1, columns - 1, false, Cheese}
	m.Cells[m.start.Row][m.start.Column] = m.start
	m.Cells[m.end.Row][m.end.Column] = m.end
	return &m
}

func (m *Maze) generateCells() {
	m.Cells = make([][]Cell, m.Rows)
	for row := uint8(0); row < m.Rows; row++ {
		m.Cells[row] = make([]Cell, m.Columns)
		for column := uint8(0); column < m.Columns; column++ {
			if rand.Float32() < m.Sparseness {
				m.Cells[row][column] = Cell{row, column, false, Wall}
			} else {
				m.Cells[row][column] = Cell{row, column, false, Empty}
			}
		}
	}
}

func (m *Maze) DFS() {

}

func (m *Maze) BFS() {

}
