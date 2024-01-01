package main

import (
	"fmt"

	"github.com/IroNEDR/mazego"
)

func printCell(content uint8) {
	switch content {
	case mazego.Mouse:
		fmt.Print("🐁")
	case mazego.Cheese:
		fmt.Print("🧀")
	case mazego.Empty:
		fmt.Print("  ")
	case mazego.Path:
		fmt.Print("🔸")
	case mazego.Wall:
		fmt.Print("██")
	}
}

func printMaze(m *mazego.Maze) {
	for row := uint8(0); row < m.Rows; row++ {
		fmt.Print("█")
		for column := uint8(0); column < m.Columns; column++ {
			printCell(m.Cells[row][column].Content)
		}
		fmt.Println("█")
	}
}

func printMazePath(m *mazego.Maze, cell *mazego.Cell) {
	if cell == nil {
		fmt.Println("No path found")
		return
	}
	for cell.Parent != nil {
		m.Cells[cell.Row][cell.Column].Content = mazego.Path
		cell = cell.Parent
	}
	m.Cells[m.Goal.Row][m.Goal.Column].Content = mazego.Cheese
	printMaze(m)
}

func main() {
	m := mazego.NewMaze(mazego.DefaultColumns, mazego.DefaultRows, mazego.DefaultSparseness)
	printMaze(m)
	cell, err := m.DFS()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("DFS Solution:")
	printMazePath(m, cell)

}
