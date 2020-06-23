package game

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gookit/color.v1"
	"log"
)

const empty = 0
const PinkCell = 1
const BlueCell = 2
const LightBlueCell = 3
const YellowCell = 4
const GreenBlueCell = 5 // салатовый
const OrangeCell = 6
const RedCell = 7
const PurpleCell = 8
const DarkBlueCell = 9
const GreenCell = 10
const DarkRedCell = 11
const ColdGreenCell = 12

const EmptyCell = int8(0)

const BoardX = int8(11)
const BoardY = int8(5)

const DirectionUp = int8(1)
const DirectionLeft = int8(2)
const DirectionRight = int8(3)
const DirectionDown = int8(4)

type BoardCells [BoardX][BoardY]int8

type Board struct {
	Figures     [13]Figure           `json:"figures"`
	FreeFigures []Figure           `json:"freeFigures"`
	Cells       BoardCells           `json:"cells"`
	CellColors  map[int8]FigureColor `json:"cellColors"`
}

func (b *Board) AddFigure(f Figure, rotation int8, flip bool, startX, startY int8) error {
	f.Init()
	f.Rotate(rotation, false)
	if flip {
		f.FlipHorizontal()
	}
	if b.CellColors == nil {
		b.CellColors = make(map[int8]FigureColor)
		b.CellColors[0] = FigureColor{HEX: "white"}
	}
	//if b.Figures == nil {
	//	b.Figures = make([]Figure, BoardX)
	//}

	// Временное хранилище новых ячеек
	newCells := b.Cells
	b.CellColors[f.TypeId] = f.Color
	//fmt.Println(b.Cells)
	//fmt.Println(f.Cells)
	//fmt.Println(f.Cells[1])
	for x := startX; x < startX+f.SquareSize; x++ {
		for y := startY; y < startY+f.SquareSize; y++ {
			//fmt.Printf("x: %d; y: %d\n", x, y)
			cellType := f.Cells[x-startX][y-startY]
			if cellType > 0 {
				if newCells[x][y] > 0 {
					return fmt.Errorf("board cell(%d:%d) is not empty(%d)", x, y, newCells[x][y])
				}

				newCells[x][y] = f.TypeId
			}
		}
	}

	if b.checkHasWhiteHoles(newCells) {
		return fmt.Errorf("board has white holes")
	}

	// Если всё прошло без ошибок - комитим новые ячейки в доску
	b.Cells = newCells

	// Записываем у фигуры где она начинаетя
	f.BoardX = startX
	f.BoardY = startY
	b.Figures[f.TypeId] = f

	return nil
}

func (b *Board) AddFreeFigure(f Figure) {
	b.FreeFigures = append(b.FreeFigures, f)
}

// Проверяет, есть ли пустые ячейки, где точно не влезет фигура
func (b Board) checkHasWhiteHoles(cells BoardCells) bool {
	for x := int8(0); x < BoardX; x++ {
		for y := int8(0); y < BoardY; y++ {
			if cells[x][y] == EmptyCell {
				filledCount := int8(0)
				// Если ячейка слева выходит за границы или не пустая
				if x-1 < 0 || cells[x-1][y] != EmptyCell {
					filledCount++ // увеличиваем на единицу
				}

				// Ячейка справа выходит за границы или не пустая
				if x+1 > BoardX-1 || cells[x+1][y] != EmptyCell {
					filledCount++ // увеличиваем на единицу
				}

				// Ячейка сверху выходит за границы или не пустая
				if y-1 < 0 || cells[x][y-1] != EmptyCell {
					filledCount++ // увеличиваем на единицу
				}

				// Ячейка снизу выходит за границы или не пустая
				if y+1 > BoardY-1 || cells[x][y+1] != EmptyCell {
					filledCount++ // увеличиваем на единицу
				}

				if filledCount == 4 {
					return true
				}
			}
		}
	}

	return false
}

func (b *Board) PrintToScreen() {
	for y := int8(0); y < BoardY; y++ {
		for x := int8(0); x < BoardX; x++ {
			//print(b.Cells[x][y])
			color.Print("● ")
			//b.CellColors[b.Cells[x][y]].Print("● ")
			// fmt.Printf(b.CellColors[b.Cells[x][y]], "●")
		}
		print("\n")
	}
	print("\n")
}

func (b Board) GetFigureById(id int8) *Figure {
	return &b.Figures[id]
}

func (b *Board) Move(figure *Figure, direction int8) bool {
	if direction == 0 {
		return false
	}
	currentCells := b.Cells
	newCells := BoardCells{}
	unit := int8(1)
	lowX := int8(0)
	lowY := int8(0)
	for x := int8(0); x < BoardX; x++ {
		for y := int8(0); y < BoardY; y++ {
			value := currentCells[x][y]
			if value != figure.TypeId {
				if value > EmptyCell {
					newCells[x][y] = value
				}
				continue
			}
			var newX = x
			var newY = y
			switch direction {
			case DirectionUp:
				newY -= unit
			case DirectionDown:
				newY += unit
			case DirectionLeft:
				newX -= unit
			case DirectionRight:
				newX += unit
			}

			fmt.Printf("Direction: %d; CurrentX: %d; CurrentY: %d --- NewX: %d; NewY: %d\n", direction, x, y, newX, newY)

			// out of range
			if newX < lowX || newY < lowY || (newX > BoardX-1) || (newY > BoardY-1) {
				fmt.Println("Out of range")
				return false
			}

			// target cell is not empty
			if currentCells[newX][newY] != EmptyCell && currentCells[newX][newY] != value {
				fmt.Printf("new cell is not empty: %d\n", currentCells[newX][newY])
				return false // new cell is not empty
			}
			newCells[newX][newY] = value   // move to new position
			currentCells[x][y] = EmptyCell // remove old value
		}
	}

	// update figure start coordinates
	switch direction {
	case DirectionUp:
		figure.BoardY -= unit
	case DirectionDown:
		figure.BoardY += unit
	case DirectionLeft:
		figure.BoardX -= unit
	case DirectionRight:
		figure.BoardX += unit
	}

	b.Cells = newCells
	return true
}

func (b *Board) Rotate(f *Figure, rotation int8, flip bool) {
	fmt.Printf("Start rotation figure StartX:%d; StartY: %d\n", f.BoardX, f.BoardY)
	fmt.Println(f.Cells)
	//oldFigureCells := f.Cells
	boardCells := b.Cells
	isSuccess := true
	defer func() {
		// restore figure cells if something went wrong
		if !isSuccess {
			fmt.Println("Something went wrong, rollback figure cell changes")
			//f.Cells = oldFigureCells
		}
	}()

	// Rotate and/or flip
	f.Rotate(rotation, false)
	if flip {
		f.FlipHorizontal()
	}

	fmt.Println(f.Cells)

	// replace old figure points to new positions
	for x := f.BoardX; x < f.BoardX+f.SquareSize; x++ {
		for y := f.BoardY; y < f.BoardY+f.SquareSize; y++ {
			value := boardCells[x][y]
			if value == f.TypeId {
				boardCells[x][y] = EmptyCell // remove old point for this figure
			} else if value > EmptyCell && f.Cells[x-f.BoardX][y-f.BoardY] > EmptyCell {
				isSuccess = false // new target cell already used by other figure
				fmt.Printf("Target cell %d:%d already used(%d)\n", x, y, value)
				return
			}

			if f.Cells[x-f.BoardX][y-f.BoardY] > 0 {
				boardCells[x][y] = f.TypeId
			}
		}
	}

	b.Cells = boardCells
}

type BoardInfo struct {
	Cells  [BoardY][BoardX]int8 `json:"cells"`
	Colors map[int8]FigureColor `json:"colors"`
}

func (b *Board) ToJson() []byte {
	boardInfo := BoardInfo{}
	// reverse desk for humans!
	for y := int8(0); y < BoardY; y++ {
		for x := int8(0); x < BoardX; x++ {
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
