package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

type Client struct {
	Conn   *net.Conn
	Reader *bufio.Reader
}

func (c *Client) Connect(host string, port int) (string, error) {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		return "", err
	}
	c.Conn = &conn
	c.Reader = bufio.NewReader(conn)
	return c.Read(), nil
}

func (c *Client) Read() string {
	data, err := c.Reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read: %+v\n", err)
	}
	return data
}

func (c *Client) Write(data string) {
	fmt.Fprintf(*c.Conn, data+"\n")
}

func (c *Client) Command(data string) string {
	c.Write(data)
	return c.Read()
}
