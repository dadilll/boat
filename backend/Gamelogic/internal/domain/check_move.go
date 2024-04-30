package domain

import (
	"github.com/Dadil/boat/backend/Gamelogic/internal/models"
)

// CheckHit проверяет попадание в корабль
func CheckHit(board *models.GameBoard, x, y int) bool {
	if x < 0 || y < 0 || x >= board.Width || y >= board.Height {
		// Проверка, что координаты находятся в пределах игрового поля
		return false
	}
	return board.Cells[y][x] == models.CellStateShip
}

// ProcessMove обрабатывает ход игрока и обновляет состояние игры
func ProcessMove(board *models.GameBoard, x, y int) models.CellState {
	if CheckHit(board, x, y) {
		// Если попадание
		board.Cells[y][x] = models.CellStateHit
		// Дополнительные действия, например, проверка условий победы
		return models.CellStateEmpty
	}
	// Если промах
	board.Cells[y][x] = models.CellStateMiss
	return models.CellStateMiss
}
