package models

// GameState представляет состояние игры
type GameState struct {
	Board         *GameBoard
	CurrentPlayer int // Идентификатор текущего игрока (например, 0 или 1)
	// Другие поля, если необходимо добавить дополнительные данные о состоянии игры
}

// NewGameState создает новое состояние игры с заданным размером игрового поля
func NewGameState(width, height int) *GameState {
	board := NewGameBoard(width, height)
	return &GameState{
		Board:         board,
		CurrentPlayer: 0, // Игрок 0 начинает первым
	}
}

// NextPlayer переключает текущего игрока на следующего
func (gs *GameState) NextPlayer() {
	gs.CurrentPlayer = (gs.CurrentPlayer + 1) % 2 // Переключение между игроками 0 и 1
}
