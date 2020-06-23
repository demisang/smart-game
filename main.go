package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smartGame/game"
	"strconv"
)

func main() {
	board := game.Board{}
	initDefaultFigures(&board)
	//board.PrintToScreen()

	startService(&board)
}

func startService(board *game.Board) {
	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		// map[direction:2, figureId:9, rotate: 1]
		params := map[string]string{}
		json.NewDecoder(r.Body).Decode(&params)

		figureId, _ := strconv.ParseInt(params["figureId"], 10, 8)
		figure := board.GetFigureById(int8(figureId))

		fmt.Println("\n", params)

		// Move
		direction, _ := strconv.ParseInt(params["direction"], 10, 8)
		if direction > 0 {
			board.Move(figure, int8(direction))
		}

		// Rotate
		rotate, _ := strconv.ParseInt(params["rotate"], 10, 8)
		flip, _ := strconv.ParseBool(params["flip"])
		if rotate > 0 || flip {
			board.Rotate(figure, int8(rotate), flip)
		}
		fmt.Println(params, figureId, figure.TypeId, rotate, flip)

		w.Write(board.ToJson())
	})

	http.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		w.Write(board.ToJson())
	})

	http.HandleFunc("/solve", func(w http.ResponseWriter, r *http.Request) {
		solver := game.Solver{Board: *board}

		solvedBoard, isSolved := solver.Run(*board)
		if !isSolved {
			w.WriteHeader(400)
			return
		}
		w.Write(solvedBoard.ToJson())
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(":8770", nil)
}

func initDefaultFigures(board *game.Board) {
	cell := [4][4]int8{}

	// pink
	cell = [4][4]int8{
		{1, 1, 1, 0},
		{0, 0, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	pinkFigure := game.Figure{TypeId: game.PinkCell, Color: game.FigureColor{HEX: "#f542c2"}, Cells: cell}
	board.AddFigure(pinkFigure, 0, false, 0, 0)

	// blue
	cell = [4][4]int8{
		{1, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	blueFigure := game.Figure{TypeId: game.BlueCell, Color: game.FigureColor{HEX: "#1928fa"}, Cells: cell}
	board.AddFigure(blueFigure, 0, false, 3, 0)

	// Light Blue
	cell = [4][4]int8{
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	lightBlueFigure := game.Figure{TypeId: game.LightBlueCell, Color: game.FigureColor{HEX: "#6daebf"}, Cells: cell}
	board.AddFigure(lightBlueFigure, 0, false, 6, 0)

	// Yellow
	cell = [4][4]int8{
		{1, 1, 1, 1},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	yellowFigure := game.Figure{TypeId: game.YellowCell, Color: game.FigureColor{HEX: "#e3d510"}, Cells: cell}
	board.AddFigure(yellowFigure, 0, false, 7, 0)

	// Green Blue
	cell = [4][4]int8{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}
	greenBlueFigure := game.Figure{TypeId: game.GreenBlueCell, Color: game.FigureColor{HEX: "#17e6c3"}, Cells: cell}
	board.AddFigure(greenBlueFigure, 0, false, 0, 1)

	// Orange
	cell = [4][4]int8{
		{0, 0, 1, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}
	orangeFigure := game.Figure{TypeId: game.OrangeCell, Color: game.FigureColor{HEX: "#f79f1b"}, Cells: cell}
	board.AddFigure(orangeFigure, 0, false, 4-2, 1)

	//// Red
	cell = [4][4]int8{
		{1, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
	}
	redFigure := game.Figure{TypeId: game.RedCell, Color: game.FigureColor{HEX: "#ed2c09"}, Cells: cell}
	board.AddFigure(redFigure, 0, false, 9, 1)

	// Purple
	//cell = [4][4]int8{
	//	{1, 1, 0, 0},
	//	{0, 1, 1, 0},
	//	{0, 0, 1, 0},
	//	{0, 0, 0, 0},
	//}
	//purpleFigure := game.Figure{TypeId: game.PurpleCell, Color: game.FigureColor{HEX: "#730d59"}, Cells: cell}
	//board.AddFigure(purpleFigure, 0, false, 6, 2)

	// Dark Blue
	cell = [4][4]int8{
		{1, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}
	darkBlueFigure := game.Figure{TypeId: game.DarkBlueCell, Color: game.FigureColor{HEX: "#160175"}, Cells: cell}
	board.AddFigure(darkBlueFigure, 0, false, 8, 2)

	// Green
	cell = [4][4]int8{
		{1, 0, 1, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	greenFigure := game.Figure{TypeId: game.GreenCell, Color: game.FigureColor{HEX: "#029916"}, Cells: cell}
	board.AddFigure(greenFigure, 0, false, 0, 3)

	// Dark Red
	cell = [4][4]int8{
		{0, 1, 1, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	darkRedFigure := game.Figure{TypeId: game.DarkRedCell, Color: game.FigureColor{HEX: "#700c01"}, Cells: cell}
	//board.AddFigure(darkRedFigure, 0, false, 3, 3)
	board.AddFreeFigure(darkRedFigure)

	// Cold Green
	cell = [4][4]int8{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	coldGreenFigure := game.Figure{TypeId: game.ColdGreenCell, Color: game.FigureColor{HEX: "#00704d"}, Cells: cell}
	//board.AddFigure(coldGreenFigure, 0, false, 5, 3)
	board.AddFreeFigure(coldGreenFigure)
}