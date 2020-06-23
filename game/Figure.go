package game

type Figure struct {
	BoardX        int8
	BoardY        int8
	SquareSize    int8
	Cells         [4][4]int8
	TypeId        int8
	Color         FigureColor
	isInitialized bool
	FilledCells   [4][4]int8
}

type FigureColor struct {
	HEX string
}

func (f *Figure) Init() {
	if f.isInitialized {
		return
	}

	if f.SquareSize == 0 {
		f.SquareSize = 4
	}

	// invert BoardX and BoardY (because cells filled by human!)
	newCells := [4][4]int8{}
	for x := int8(0); x < f.SquareSize; x++ {
		for y := int8(0); y < f.SquareSize; y++ {
			newCells[x][y] = f.Cells[y][x]
		}
	}
	f.Cells = newCells

	// TODO Setup color by type

	f.isInitialized = true
}

func (f *Figure) FlipHorizontal() {
	f.Init()

	for x := int8(0); x < f.SquareSize/2; x++ {
		for y := int8(0); y < f.SquareSize; y++ {
			// opposite cell coordinates
			oppositeCellX := f.SquareSize - x - 1
			// store current first cell value
			firstCellValue := f.Cells[x][y]
			// overwrite current first cell value
			f.Cells[x][y] = f.Cells[oppositeCellX][y]
			// restore current first cell value write it to opposite cell
			f.Cells[oppositeCellX][y] = firstCellValue
		}
	}
}

func (f *Figure) FlipVertical() {
	f.Init()

	for x := int8(0); x < f.SquareSize; x++ {
		for y := int8(0); y < f.SquareSize/2; y++ {
			// opposite cell coordinates
			oppositeCellY := f.SquareSize - y - 1
			// store current first cell value
			firstCellValue := f.Cells[x][y]
			// overwrite current first cell value
			f.Cells[x][y] = f.Cells[x][oppositeCellY]
			// restore current first cell value write it to opposite cell
			f.Cells[x][oppositeCellY] = firstCellValue
		}
	}
}

func (f *Figure) Rotate(rotation int8, toLeft bool) {
	f.Init()
	if rotation <= 0 {
		return
	}

	for x := int8(0); x < f.SquareSize/2; x++ {
		Nx := f.SquareSize - 1 - x
		for y := x; y < Nx; y++ {
			// store current cell value
			cellValue := f.Cells[x][y]
			Ny := f.SquareSize - 1 - y

			if toLeft {
				f.Cells[x][y] = f.Cells[Ny][x]
				f.Cells[Ny][x] = f.Cells[Nx][Ny]
				f.Cells[Nx][Ny] = f.Cells[y][Nx]
				f.Cells[y][Nx] = cellValue
			} else {
				f.Cells[x][y] = f.Cells[y][Nx]
				f.Cells[y][Nx] = f.Cells[Nx][Ny]
				f.Cells[Nx][Ny] = f.Cells[Ny][x]
				f.Cells[Ny][x] = cellValue
			}
		}
	}

	f.Rotate(rotation-1, toLeft)
}
