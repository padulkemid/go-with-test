package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	PlayerPrompt = "Please enter the number of players: "
	SillyWord    = "silly n"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Texas
}

func (c *CLI) readLine() string {
	c.in.Scan()

	return c.in.Text()
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	numOfPlayerIn := c.readLine()
	numOfPlayers, err := strconv.Atoi(strings.Trim(numOfPlayerIn, "\n"))
	if err != nil {
		fmt.Fprint(c.out, SillyWord)
		return
	}

	c.game.Start(numOfPlayers, c.out)

	i := c.readLine()
	ew := extractWinner(i)

	c.game.Finish(ew)
}

func extractWinner(i string) string {
	return strings.Replace(i, " wins!", "", 1)
}

func NewCli(in io.Reader, out io.Writer, game Texas) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}
