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

const (
	PlayersPrompt              = "Podaj ilość graczy: "
	WrongNumberOfPlayers       = "Podałeś niepoprawną wartość, wpisz cyfrę."
	ErrFinishingGameWrongInput = "Nie zrozumiałem tego co do mnie mówisz."
)

func (c *CLI) PlayPoker() error {
	var (
		noOfPlayers int
	)

	fmt.Fprintln(c.out, PlayersPrompt)
	_, err := fmt.Fscanln(c.in, &noOfPlayers)
	if err != nil {
		if err.Error() == "expected integer" {
			fmt.Fprintln(c.out, WrongNumberOfPlayers)
			return nil
		} else {
			return fmt.Errorf("playing poker, problem while scanning number of players, %w",
				err)
		}
	}
	c.game.StartGame(noOfPlayers)

	//for {
	var (
		winner, command string
	)
	fmt.Fscanln(c.in, &winner, &command)

	if len(winner) > 0 && command == "wins" {
		err = c.game.Finish(winner)
		if err != nil {
			return fmt.Errorf("playing poker, problem recording name, %w", err)
		}
		//break
	} else {
		fmt.Fprintln(c.out, ErrFinishingGameWrongInput)
	}
	//if err != nil {
	//	return fmt.Errorf("playing poker, problem while scanning name, %w", err)
	//}
	//}

	return nil
}
