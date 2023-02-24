# TicTacToe multiplayer server

![License](https://img.shields.io/github/license/juho05/tictactoe-backend)
![Go version](https://img.shields.io/github/go-mod/go-version/juho05/tictactoe-backend)

A TicTacToe TCP multiplayer server.

## Features

- Unlimited 1v1 matches in parallel
- Winner/tie detection
- Disconnect handling

## Clients

- [CLI](https://github.com/juho05/tictactoe-cli) by *juho05*
- [WPF](https://github.com/Zersaeger/Tic-Tac-Toe-Multiplayer-Frontend) by *Zersaeger*

## Commands

### Sends
- `ping`
	- Expects a `pong` response
- `match-found:[xo]`, e.g. `match-found:x`
	- Is sent once a match is found
	- Contains *your* sign (`x` or `o`)
- `your-turn`
	- It's your turn (send a `click` message)
- `their-turn`
	- Your opponent needs to send a `click` message
- `opponent-disconnected`
	- Your opponent disconnected
- `board:000000000`, e.g. `board:120000000`
	- Is sent once the board status changes
	- Contains the content of all fields
		- Content represented using integers
			- 0: empty
			- 1: cross
			- 2: circle
		- Left to right, top to bottom
			- In this example the top left field is filled with a cross and the one to the right of it with a circle
- `winner:000`, e.g. `winner:048`
	- Is sent once you win
	- Contains the indices of the winning streak (in the example top-left to bottom-right)
- `loser:000`, e.g. `loser:048`
	- Is sent once you lose
	- Contains the indices of the winning streak (in the example top-left to bottom-right)
- `tie`
	- Is sent once all fields are filled without a winner
	
### Receives

- `ping`
	- Responds with `pong`

#### In a match

- `click:[0-8]`, e.g. `click:2`
	- Only allowed when it is your turn
	- Only allowed when the field is empty
	- Fill the field with the specified index with your sign
- `again`
	- Can be sent once the game is complete
	- If both parties send this command, the game will be restarted

## Setup

### Prerequisites

- [Go](https://go.dev/) 1.17+

### Cloning the repo

```sh
git clone https://github.com/juho05/tictactoe-backend.git
cd tictactoe-backend
```

### Building

```sh
go build -o tictactoe-backend ./cmd/main.go
```

### Running

```sh
./tictactoe-backend
```

## License

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

## Copyright

Copyright © 2022 Julian Hofmann
