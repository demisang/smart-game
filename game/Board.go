package game

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gookit/color.v1"
	"log"
)

const empty = 0
const RedCell = 1
const darkRed = 2
const CELL_ORANGE = 3
const CELL_PURPLE = 4
const CELL_BLUE = 5
const CELL_LIGHT_BLUE = 6
const CELL_DARK_BLUE = 7
const PinkCell = 8
const YellowCell = 9
const CELL_GREEN = 10
const CELL_LIGHT_GREEN = 11
const CELL_DARK_GREEN = 12

type Board struct {
	Figures    []Figure             `json:"figures"`
	Cells      [11][5]int8          `json:"cells"`
	CellColors map[int8]FigureColor `json:"cellColors"`
}

func (b *Board) AddFigure(f Figure, rotation int8, flip bool, startX, startY int8) {
	f.Init()
	f.Rotate(rotation, false)
	if flip {
		f.FlipHorizontal()
	}
	if b.CellColors == nil {
		b.CellColors = make(map[int8]FigureColor)
		b.CellColors[0] = FigureColor{HEX: "white"}
	}
	if b.Figures == nil {
		b.Figures = make([]Figure, 11)
	}

	b.CellColors[f.TypeId] = f.Color
	//fmt.Println(b.Cells)
	//fmt.Println(f.Cells)
	//fmt.Println(f.Cells[1])
	for x := startX; x < startX+f.SquareSize; x++ {
		for y := startY; y < startY+f.SquareSize; y++ {
			//fmt.Printf("x: %d; y: %d\n", x, y)
			// NOTICE: inverted for humans!
			cellType := f.Cells[x-startX][y-startY]
			if cellType > 0 {
				if b.Cells[x][y] > 0 {
					panic(fmt.Sprintf("Board cell(%d:%d) is not empty(%d)", x, y, b.Cells[x][y]))
				}

				b.Cells[x][y] = f.TypeId
			}
		}
	}

	b.Figures[f.TypeId] = f
}

func (b *Board) PrintToScreen() {
	for y := 0; y < 5; y++ {
		for x := 0; x < 11; x++ {
			//print(b.Cells[x][y])
			color.Print("● ")
			//b.CellColors[b.Cells[x][y]].Print("● ")
			// fmt.Printf(b.CellColors[b.Cells[x][y]], "●")
		}
		print("\n")
	}
}

func (b Board) GetFigureById(id int8) Figure {
	return b.Figures[id]
}

func (b *Board) Move(figure Figure, direction string) {

}

type BoardInfo struct {
	Cells  [5][11]int8          `json:"cells"`
	Colors map[int8]FigureColor `json:"colors"`
}

func (b *Board) ToJson() []byte {
	boardInfo := BoardInfo{}
	// reverse desk for humans!
	for y := 0; y < 5; y++ {
		for x := 0; x < 11; x++ {
			boardInfo.Cells[y][x] = b.Cells[x][y]
		}
	}
	// colors
	boardInfo.Colors = b.CellColors

	jsonString, err := json.Marshal(boardInfo)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	return jsonString
}
