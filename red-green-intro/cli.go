package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

func (c *CLI) readLine() string  {
  c.in.Scan()

  return c.in.Text()
}

func (c *CLI) PlayPoker() {
  i := c.readLine()
	ew := extractWinner(i)

	c.store.RecordWin(ew)
}

func extractWinner(i string) string {
	return strings.Replace(i, " wins!", "", 1)
}

func NewCli(store PlayerStore, in io.Reader) *CLI  {
  return &CLI{
    store,
    bufio.NewScanner(in),
  }
}
