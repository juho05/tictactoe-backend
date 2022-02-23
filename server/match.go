package server

import (
	"fmt"
	"strconv"
	"strings"
)

type cellState int

const (
	cellEmpty  cellState = 0
	cellCross  cellState = 1
	cellCircle cellState = 2
)

type board [9]cellState

func (b board) String() string {
	return strings.Trim(strings.ReplaceAll(fmt.Sprint([9]cellState(b)), " ", ""), "[]")
}

type Match struct {
	clientCross  *Client
	clientCircle *Client

	currentPlayerId string

	board board
}

func NewMatch(clientCross, clientCircle *Client) *Match {
	match := &Match{
		clientCross:  clientCross,
		clientCircle: clientCircle,

		currentPlayerId: clientCross.id,
	}

	clientCross.match = match
	clientCircle.match = match

	return match
}

func (m *Match) begin() {
	err := m.clientCross.send("match-found:x")
	if err != nil {
		m.terminate()
		return
	}

	err = m.clientCircle.send("match-found:o")
	if err != nil {
		m.terminate()
		return
	}

	err = m.clientCross.send("your-turn")
	if err != nil {
		m.terminate()
		return
	}

	err = m.clientCircle.send("their-turn")
	if err != nil {
		m.terminate()
		return
	}

	fmt.Println("Started new match:", m.clientCross.ip, "+", m.clientCircle.ip)
}

func (m *Match) handleCommand(client *Client, command string) {
	if client.id != m.currentPlayerId {
		fmt.Printf("Player %s tried to execute an action even though it is not their turn.", client.ip)
		return
	}

	if strings.HasPrefix(command, "click:") {
		parts := strings.Split(command, ":")
		if len(parts) != 2 {
			invalidCommand(client.ip, command, "expected value after ':'")
			return
		}

		index, err := strconv.Atoi(parts[1])
		if err != nil || index < 0 || index > 8 {
			invalidCommand(client.ip, command, "expected index between 0-8 after ':'")
			return
		}

		if m.board[index] != 0 {
			fmt.Printf("Player %s tried to click on a non-empty field (%d)", client.ip, index)
			return
		}

		if client.id == m.clientCross.id {
			m.board[index] = cellCross
		} else {
			m.board[index] = cellCircle
		}

		if !m.checkWon() {
			m.switchTurns()
		}
	}
}

func (m *Match) switchTurns() {
	err := m.sendBoard()
	if err != nil {
		m.terminate()
	}

	if m.currentPlayerId == m.clientCross.id {
		m.currentPlayerId = m.clientCircle.id
		m.clientCircle.send("your-turn")
		m.clientCross.send("their-turn")
	} else {
		m.currentPlayerId = m.clientCross.id
		m.clientCross.send("your-turn")
		m.clientCircle.send("their-turn")
	}
}

func (m *Match) checkWon() bool {
	for i := 0; i < 3; i++ {
		// top to bottom
		if m.board[0+i] != cellEmpty && m.board[0+i] == m.board[3+i] && m.board[0+i] == m.board[6+i] {
			m.complete(m.board[0+i], fmt.Sprintf("%d%d%d", 0+i, 3+i, 6+i))
			return true
		}

		// left to right
		if m.board[i*3+0] != cellEmpty && m.board[i*3+0] == m.board[i*3+1] && m.board[i*3+0] == m.board[i*3+2] {
			m.complete(m.board[i*3+0], fmt.Sprintf("%d%d%d", i*3+0, i*3+1, i*3+2))
			return true
		}
	}

	// top left to bottom right
	if m.board[0] != cellEmpty && m.board[0] == m.board[4] && m.board[0] == m.board[8] {
		m.complete(m.board[0], fmt.Sprintf("%d%d%d", 0, 4, 8))
		return true
	}

	// top right to bottom left
	if m.board[2] != cellEmpty && m.board[2] == m.board[4] && m.board[2] == m.board[6] {
		m.complete(m.board[2], fmt.Sprintf("%d%d%d", 2, 4, 6))
		return true
	}

	return false
}

func (m *Match) complete(state cellState, indices string) {
	m.sendBoard()
	if state == cellCross {
		m.clientCross.send("winner:" + indices)
		m.clientCircle.send("loser:" + indices)
	} else {
		m.clientCircle.send("winner:" + indices)
		m.clientCross.send("loser:" + indices)
	}
}

func (m *Match) disconnect(client *Client) {
	if client == m.clientCross {
		m.clientCircle.send("opponent-disconnected")
	} else {
		m.clientCross.send("opponent-disconnected")
	}
	m.terminate()
}

func (m *Match) terminate() {
	panic("TODO")
}

func (m *Match) sendBoard() error {
	err := m.send("board:" + m.board.String())
	if err != nil {
		fmt.Println("Failed to broadcast the new board state:", err)
	}
	return err
}

func (m *Match) send(text string) error {
	err := m.clientCross.send(text)
	if err != nil {
		m.clientCircle.send("opponent-disconnected")
		m.terminate()
		return err
	}

	err = m.clientCircle.send(text)
	if err != nil {
		m.clientCross.send("opponent-disconnected")
		m.terminate()
		return err
	}

	return nil
}

func invalidCommand(clientIP, command, errorMsg string) {
	fmt.Printf("Client %s sent an invalid command '%s': %s\n", clientIP, command, errorMsg)
}
