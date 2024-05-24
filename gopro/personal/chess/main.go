package main

import (
	"fmt"
	"math"
)

type ChessPiece struct {
	PieceType string
	Color     string
}

type Movable interface {
	Move(srcRow, srcCol, destRow, destCol int, cb *ChessBoard) bool
	GetColor() string
	GetPieceType() string
}

type ChessBoard struct {
	board [8][8]Movable
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{}
	board.ResetBoard()
	return board
}

func (cb *ChessBoard) ResetBoard() {
	for row := 0; row < 8; row++ {
		cb.board[1][row] = &Pawn{ChessPiece{"P", "White"}}
		cb.board[6][row] = &Pawn{ChessPiece{"P", "Black"}}
	}

	// White pieces
	cb.board[0][0] = &Rook{ChessPiece{"R", "White"}}
	cb.board[0][1] = &Knight{ChessPiece{"N", "White"}}
	cb.board[0][2] = &Bishop{ChessPiece{"B", "White"}}
	cb.board[0][3] = &Queen{ChessPiece{"Q", "White"}}
	cb.board[0][4] = &King{ChessPiece{"K", "White"}}
	cb.board[0][5] = &Bishop{ChessPiece{"B", "White"}}
	cb.board[0][6] = &Knight{ChessPiece{"N", "White"}}
	cb.board[0][7] = &Rook{ChessPiece{"R", "White"}}

	// Black pieces
	cb.board[7][0] = &Rook{ChessPiece{"R", "Black"}}
	cb.board[7][1] = &Knight{ChessPiece{"N", "Black"}}
	cb.board[7][2] = &Bishop{ChessPiece{"B", "Black"}}
	cb.board[7][3] = &Queen{ChessPiece{"Q", "Black"}}
	cb.board[7][4] = &King{ChessPiece{"K", "Black"}}
	cb.board[7][5] = &Bishop{ChessPiece{"B", "Black"}}
	cb.board[7][6] = &Knight{ChessPiece{"N", "Black"}}
	cb.board[7][7] = &Rook{ChessPiece{"R", "Black"}}
}

func (cb *ChessBoard) DisplayBoard() {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if cb.board[row][col] != nil {
				fmt.Printf("%s%s ", cb.board[row][col].GetColor()[:1], cb.board[row][col].GetPieceType())
			} else {
				fmt.Print(" -  ")
			}
		}
		fmt.Println()
	}
}

func (cb *ChessBoard) MovePiece(srcRow, srcCol, destRow, destCol int) bool {
	if cb.board[srcRow][srcCol] != nil {
		if cb.board[srcRow][srcCol].Move(srcRow, srcCol, destRow, destCol, cb) {
			cb.board[destRow][destCol] = cb.board[srcRow][srcCol]
			cb.board[srcRow][srcCol] = nil
			return true
		}
	}
	return false
}

type Pawn struct {
	ChessPiece
}

func (p *Pawn) Move(srcRow, srcCol, destRow, destCol int, cb *ChessBoard) bool {
	// Implement pawn-specific movement logic
	if p.Color == "White" {
		if destRow == srcRow-1 && destCol == srcCol && cb.board[destRow][destCol] == nil {
			return true
		}
		if srcRow == 6 && destRow == srcRow-2 && destCol == srcCol && cb.board[destRow][destCol] == nil {
			return true
		}
		if destRow == srcRow-1 && math.Abs(float64(destCol-srcCol)) == 1 && cb.board[destRow][destCol] != nil && cb.board[destRow][destCol].GetColor() != p.Color {
			return true
		}
	} else if p.Color == "Black" {
		if destRow == srcRow+1 && destCol == srcCol && cb.board[destRow][destCol] == nil {
			return true
		}
		if srcRow == 1 && destRow == srcRow+2 && destCol == srcCol && cb.board[destRow][destCol] == nil {
			return true
		}
		if destRow == srcRow+1 && math.Abs(float64(destCol-srcCol)) == 1 && cb.board[destRow][destCol] != nil && cb.board[destRow][destCol].GetColor() != p.Color {
			return true
		}
	}
	return false
}

func (p *Pawn) GetColor() string {
	return p.Color
}

func (p *Pawn) GetPieceType() string {
	return p.PieceType
}

type Bishop struct {
	ChessPiece
}

func (b *Bishop) Move(srcRow, srcCol, destRow, destCol int, cb *ChessBoard) bool {
	// Implement bishop-specific movement logic
	if math.Float64frombits(math.Float64bits(float64(srcRow-destRow))&^(1<<63)) == math.Abs(float64(srcCol-destCol)) {
		return true
	}
	return false
}

func (b *Bishop) GetColor() string {
	return b.Color
}

func (b *Bishop) GetPieceType() string {
	return b.PieceType
}

type King struct {
	ChessPiece
}

func (k *King) Move(srcRow, srcCol, destRow, destCol int, cb *ChessBoard) bool {
	// Implement king-specific movement logic
	if math.Abs(float64(srcRow-destRow)) <= 1 && math.Abs(float64(srcCol-destCol)) <= 1 {
		return true
	}
	return false
}

func (k *King) GetColor() string {
	return k.Color
}

func (k *King) GetPieceType() string {
	return k.PieceType
}

type Queen struct {
	ChessPiece
}

func (q *Queen) Move(srcRow, srcCol, destRow, destCol int, cb *ChessBoard) bool {
	// Implement queen-specific movement logic
	if srcRow == destRow || srcCol == destCol || math.Abs(float64(srcRow-destRow)) == math.Abs(float64(srcCol-destCol)) {
		return true
	}
	return false
}

func (q *Queen) GetColor() string {
	return q.Color
}

func (q *Queen) GetPieceType() string {
	return q.PieceType
}

type Knight struct {
	ChessPiece
}

func (k *Knight) Move(srcRow, srcCol, destRow, destCol int, cb *ChessBoard) bool {
	// Implement knight-specific movement logic
	if (math.Abs(float64(srcRow-destRow)) == 2 && math.Abs(float64(srcCol-destCol)) == 1) || (math.Abs(float64(srcRow-destRow)) == 1 && math.Abs(float64(srcCol-destCol)) == 2) {
		return true
	}
	return false
}

func (k *Knight) GetColor() string {
	return k.Color
}

func (k *Knight) GetPieceType() string {
	return k.PieceType
}

type Rook struct {
	ChessPiece
}

func (r *Rook) Move(srcRow, srcCol, destRow, destCol int, cb *ChessBoard) bool {
	// Implement rook-specific movement logic
	if srcRow == destRow || srcCol == destCol {
		return true
	}
	return false
}

func (r *Rook) GetColor() string {
	return r.Color
}

func (r *Rook) GetPieceType() string {
	return r.PieceType
}

func main() {
	board := NewChessBoard()

	for {
		fmt.Println("Current board:")
		board.DisplayBoard()
		fmt.Print("Enter your move (e.g., '1 0 to 2 0'): ")
		var srcRow, srcCol, destRow, destCol int
		_, _ = fmt.Scanf("%d %d to %d %d", &srcRow, &srcCol, &destRow, &destCol)

		if !board.MovePiece(srcRow, srcCol, destRow, destCol) {
			fmt.Println("Invalid move.")
		}
	}
}
