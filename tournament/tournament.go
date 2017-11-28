package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Championship struct {
	teams      map[string]*Team
	teamsnames []string
}

type victories int
type defeats int
type ties int
type points int
type matches int

var victory string = "win"
var draw string = "draw"
var defeat string = "loss"

type Team struct {
	name string
	v    victories
	d    defeats
	t    ties
	p    points
	m    matches
}

func AddGame(c *Championship, game []string) {
	switch game[2] {
	case victory:
		if _, ok := c.teams[game[0]]; ok {
			c.teams[game[0]].p += 3
			c.teams[game[0]].m += 1
			c.teams[game[0]].v += 1
		} else {
			aTeam := Team{game[0], 1, 0, 0, 3, 1}
			c.teams[game[0]] = &aTeam
			c.teamsnames = append(c.teamsnames, game[0])
		}
		if _, ok := c.teams[game[1]]; ok {
			c.teams[game[1]].d += 1
			c.teams[game[1]].m += 1
		} else {
			aTeam := Team{game[1], 0, 1, 0, 0, 1}
			c.teams[game[1]] = &aTeam
			c.teamsnames = append(c.teamsnames, game[1])
		}

		break
	case draw:
		if _, ok := c.teams[game[0]]; ok {
			c.teams[game[0]].p += 1
			c.teams[game[0]].m += 1
			c.teams[game[0]].t += 1
		} else {
			aTeam := Team{game[0], 0, 0, 1, 1, 1}
			c.teams[game[0]] = &aTeam
			c.teamsnames = append(c.teamsnames, game[0])
		}
		if _, ok := c.teams[game[1]]; ok {
			c.teams[game[1]].p += 1
			c.teams[game[1]].m += 1
			c.teams[game[1]].t += 1
		} else {
			aTeam := Team{game[1], 0, 0, 1, 1, 1}
			c.teams[game[1]] = &aTeam
			c.teamsnames = append(c.teamsnames, game[1])
		}

		break
	case defeat:
		if _, ok := c.teams[game[0]]; ok {
			c.teams[game[0]].m += 1
			c.teams[game[0]].d += 1
		} else {
			aTeam := Team{game[0], 0, 1, 0, 0, 1}
			c.teams[game[0]] = &aTeam
			c.teamsnames = append(c.teamsnames, game[0])
		}
		if _, ok := c.teams[game[1]]; ok {
			c.teams[game[1]].v += 1
			c.teams[game[1]].m += 1
			c.teams[game[1]].p += 3
		} else {
			aTeam := Team{game[1], 1, 0, 0, 3, 1}
			c.teams[game[1]] = &aTeam
			c.teamsnames = append(c.teamsnames, game[1])
		}
		break
	}

}

//validate if input is correct
func IsInputCorrect(s string) (bool, error) {
	//if is comment
	if strings.ContainsRune(s, 35) {
		return false, nil

		//if is just a newline
	} else if s == string('\n') {

		return false, nil
	} else if s == "" {

		return false, nil
	} else {
		if len(strings.Split(s, ";")) != 3 {
			return false, fmt.Errorf("Incorrect len")
		} else if strings.Split(s, ";")[2] == victory || strings.Split(s, ";")[2] == draw || strings.Split(s, ";")[2] == defeat {

			return true, nil
		} else {

			return false, fmt.Errorf("Incorrect Input")
		}
	}

}

//sort method
func (c *Championship) Len() int {
	return len(c.teams)
}

//sorts the teams by points, then aphabetically
func (c *Championship) Less(i, j int) bool {
	if c.teams[c.teamsnames[i]].p == c.teams[c.teamsnames[j]].p {
		return len(c.teamsnames[i]) < len(c.teamsnames[j])
	} else {
		return c.teams[c.teamsnames[i]].p > c.teams[c.teamsnames[j]].p
	}
}

//sort method
func (c *Championship) Swap(i, j int) {
	c.teamsnames[i], c.teamsnames[j] = c.teamsnames[j], c.teamsnames[i]
}

//output passed to writer
func CreateOutputTable(writer io.Writer, c *Championship) {
	header := "Team                           | MP |  W |  D |  L |  P\n"
	writer.Write([]byte(header))

	for _, v := range c.teamsnames {
		row := fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", c.teams[v].name, c.teams[v].m, c.teams[v].v, c.teams[v].t, c.teams[v].d, c.teams[v].p)
		writer.Write([]byte(row))
	}
}

func Tally(reader io.Reader, writer io.Writer) error {

	c := Championship{}
	c.teams = make(map[string]*Team)
	c.teamsnames = make([]string, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		proceed, err := IsInputCorrect(line)
		if err != nil {
			return err
		}
		if proceed {
			game := strings.Split(line, ";")

			AddGame(&c, game)
		}
	}
	sort.Sort(&c)
	CreateOutputTable(writer, &c)

	return nil
}
