package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/google/uuid"
)

type Client struct {
	id     string
	con    net.Conn
	ip     string
	match  *Match
	server *Server
}

func NewClient(server *Server, con net.Conn) *Client {
	return &Client{
		id:     uuid.NewString(),
		con:    con,
		ip:     con.RemoteAddr().String(),
		match:  nil,
		server: server,
	}
}

func (c *Client) handleConnection() {
	fmt.Println("New client connected:", c.con.RemoteAddr())
	defer fmt.Printf("Client %s disconnected.\n", c.ip)
	defer c.con.Close()

	scanner := bufio.NewScanner(c.con)
	for scanner.Scan() {
		text := strings.TrimSpace(strings.ToLower(scanner.Text()))
		fmt.Printf("Received '%s' from %s.\n", text, c.con.RemoteAddr())
		c.handleCommand(text)
	}

	if c.match != nil {
		c.match.disconnect(c)
	}
}

func (c *Client) send(text string) error {
	_, err := fmt.Fprintf(c.con, text+"\n")
	if err != nil {
		fmt.Printf("Failed to send '%s' to %s: '%s'\n", text, c.ip, err)
		return err
	}
	fmt.Printf("Sent '%s' to '%s'.\n", text, c.ip)
	return nil
}

func (c *Client) handleCommand(command string) {
	switch command {
	case "ping":
		c.send("pong")
	default:
		if c.match != nil {
			c.match.handleCommand(c, command)
		} else {
			fmt.Printf("Client %s sent an invalid command: %s\n", c.ip, command)
		}
	}
}
