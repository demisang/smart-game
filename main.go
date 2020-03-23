package main

import (
	"encoding/json"
	"net/http"
	"smartGame/game"
	"strconv"
)

func main() {
	board := game.Board{}

	cell := [4][4]int8{
		{0, 1, 1, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	pinkFigure := game.Figure{TypeId: game.PinkCell, Color: game.FigureColor{HEX: "red"}, Cells: cell}
	cell = [4][4]int8{
		{0, 1, 1, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	//redFigure := game.Figure{TypeId: game.RedCell, Color: color.RGB(158, 0, 0), Cells: cell}
	//cell = [4][4]int8{
	//	{1, 1, 1, 1},
	//	{0, 1, 0, 0},
	//	{0, 0, 0, 0},
	//	{0, 0, 0, 0},
	//}
	/*cell = [4][4]int8{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}*/
	//yellowFigure := game.Figure{TypeId: game.YellowCell, Color: color.RGB(247, 247, 7), Cells: cell}

	board.AddFigure(pinkFigure, 0, false, 0, 0)
	//board.AddFigure(redFigure, 0, false, 0, 0)
	//board.AddFigure(yellowFigure, 0, true, 3, 0)
	//board.AddFigure(pinkFigure, 0, false, 6, 0)

	board.PrintToScreen()

	startService(board)
}

func startService(board game.Board) {
	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		params := map[string]string{}
		json.NewDecoder(r.Body).Decode(&params)

		figureId, _ := strconv.ParseInt(params["figureId"], 10, 8)
		figure := board.GetFigureById(int8(figureId))

		board.Move(figure, params["direction"])

		w.Write(board.ToJson())
	})

	http.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		w.Write(board.ToJson())
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(":8770", nil)
}
