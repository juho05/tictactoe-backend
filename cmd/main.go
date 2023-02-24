package main

import (
	"fmt"

	"github.com/juho05/tictactoe-backend/server"
)

const port = 7531

func main() {
	fmt.Println("========== TicTacToe Server ==========")

	server := server.New()
	server.Listen(port)
}
