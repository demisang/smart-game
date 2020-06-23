package game

type Solver struct {
	Board       Board
	FreeFigures []Figure
}

func (s *Solver) Run(b Board) (Board, bool) {
	s.Board = b
	s.FreeFigures = b.FreeFigures
	solved := false
	for _, figure := range s.FreeFigures {
		if solved {
			break
		}
		for {
			if s.IsCompleted() {
				solved = true
				break
			}

			s.tryingToAddFigure(figure)
		}
	}

	return s.Board, solved
}

func (s *Solver) tryingToAddFigure(f Figure) bool {
	for flip := int8(0); flip <= 1; flip++ {
		for rotation := int8(0); rotation <= 4; rotation++ {
			for x := int8(0); x < BoardX; x++ {
				for y := int8(0); y < BoardY; y++ {
					err := s.Board.AddFigure(f, rotation, flip == 1, x, y)
					if err != nil {
						return false
					}
				}
			}
		}
	}

	return true
}

func (s *Solver) IsCompleted() bool {
	if len(s.FreeFigures) != 0 {
		return false // all figures should be used
	}

	for x := int8(0); x < BoardX; x++ {
		for y := int8(0); y < BoardY; y++ {
			if s.Board.Cells[x][y] == EmptyCell {
				return false // at least one cell is empty - fails
			}
		}
	}

	return true
}
