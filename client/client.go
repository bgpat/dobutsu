package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/bgpat/dobutsu/shogi"
)

type Client struct {
	Conn   *net.Conn
	Reader *bufio.Reader
	Phase  string
	Board  *shogi.Board
	Player int
	Turn   int
	Count  map[string]int
	Queue  map[string]*shogi.Board
}

func New(host string, port int) (*Client, error) {
	var c Client
	err := c.Connect(host, port)
	c.Count = make(map[string]int)
	c.Phase = "connected"
	for err == nil && c.Phase != "" {
		err = c.Step()
	}
	return &c, err
}

func (c *Client) Connect(host string, port int) error {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	c.Conn = &conn
	c.Reader = bufio.NewReader(conn)
	return nil
}

func (c *Client) Read() string {
	data, err := c.Reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read: %+v\n", err)
	}
	log.Println("> " + data)
	return data
}

func (c *Client) Write(data string) {
	fmt.Fprintf(*c.Conn, data+"\n")
	log.Println("< " + data)
}

func (c *Client) Command(data string) string {
	c.Write(data)
	return c.Read()
}
