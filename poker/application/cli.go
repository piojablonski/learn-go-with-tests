package application

import (
	"fmt"
	"io"
)

func (s *ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %s", s.Amount, s.ScheduledAt)
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{in, out, game}
}

type CLI struct {
	in   io.Reader
	out  io.Writer
	game Game
}

const PlayersPrompt = "Podaj ilość graczy: "

func (c *CLI) PlayPoker() error {
	var (
		winner      string
		noOfPlayers int
	)

	fmt.Fprint(c.out, PlayersPrompt)
	fmt.Fscanln(c.in, &noOfPlayers)
	//if err != nil {
	//	return fmt.Errorf("playing poker, problem while scanning number of players, %w", err)
	//}

	c.game.StartGame(noOfPlayers)

	//cli.scheduleAllAlerts(noOfPlayers)
	fmt.Fscanln(c.in, &winner)
	err := c.game.Finish(winner)

	//if err != nil {
	//	return fmt.Errorf("playing poker, problem while scanning name, %w", err)
	//}
	if err != nil {
		return fmt.Errorf("playing poker, problem recording name, %w", err)
	}
	return nil
}
