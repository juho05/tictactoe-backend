# TicTacToe multiplayer server

![License](https://img.shields.io/github/license/Bananenpro/tictactoe-backend)
![Go version](https://img.shields.io/github/go-mod/go-version/Bananenpro/tictactoe-backend)

A simple TicTacToe multiplayer server.

## Features

- Unlimited 1v1 matches in parallel
- Winner/tie detection
- Disconnect handling

## Clients

- [tictactoe-cli](https://github.com/Bananenpro/tictactoe-cli)

## Setup

### Prerequisites

- [Go](https://go.dev/) 1.17+

### Cloning the repo

```sh
git clone https://github.com/Bananenpro/tictactoe-backend.git
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
