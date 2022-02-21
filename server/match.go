package server

import "fmt"

type Match struct {
	clientA *Client
	clientB *Client
}

func NewMatch(clientA, clientB *Client) *Match {
	match := &Match{
		clientA: clientA,
		clientB: clientB,
	}

	clientA.match = match
	clientB.match = match

	err := clientA.send("match-found")
	if err != nil {
		return nil
	}

	err = clientB.send("match-found")
	if err != nil {
		return nil
	}

	fmt.Println("New match:", clientA.ip, "+", clientB.ip)

	return match
}
