package main

import (
	"fmt"

	"github.com/IroNEDR/mazego"
)

func printMaze(m *mazego.Maze) {
	for row := uint8(0); row < m.Rows; row++ {
		fmt.Print("█")
		for column := uint8(0); column < m.Columns; column++ {
			switch m.Cells[row][column].Content {
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
		fmt.Println("█")
	}
}

func main() {
	m := mazego.NewMaze(mazego.DefaultColumns, mazego.DefaultRows, mazego.DefaultSparseness)
	printMaze(m)
}
