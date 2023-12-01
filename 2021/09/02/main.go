package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	cmds := Parse(os.Stdin)
	x, d := RunCommands(cmds, 0, 0)
	fmt.Println(x, d, x*d)
	x, d = RunCommandsWithAim(cmds, 0, 0, 0)
	fmt.Println(x, d, x*d)
}

type Command struct {
	Direction string
	Distance  int
}

func Parse(source io.Reader) []Command {
	buffer := bufio.NewReader(source)
	cmds := []Command{}
	for {
		line, _, err := buffer.ReadLine()
		if err == io.EOF {
			break
		}
		lineStr := string(line)
		if len(strings.Trim(lineStr, " ")) == 0 {
			continue
		}
		parts := strings.Split(lineStr, " ")
		mvm, _ := strconv.Atoi(parts[1])
		cmds = append(cmds, Command{parts[0], mvm})
	}
	return cmds
}

func RunCommands(cmds []Command, initX, initDepth int) (int, int) {
	x, d := initX, initDepth
	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		if cmd.Direction == "forward" {
			x += cmd.Distance
		} else if cmd.Direction == "down" {
			d += cmd.Distance
		} else if cmd.Direction == "up" {
			d -= cmd.Distance
		}
	}
	return x, d
}

func RunCommandsWithAim(cmds []Command, initX, initDepth, initAim int) (int, int) {
	x, d, aim := initX, initDepth, initAim
	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		if cmd.Direction == "forward" {
			x += cmd.Distance
			d += aim * cmd.Distance
		} else if cmd.Direction == "down" {
			aim += cmd.Distance
		} else if cmd.Direction == "up" {
			aim -= cmd.Distance
		}
	}
	return x, d
}
