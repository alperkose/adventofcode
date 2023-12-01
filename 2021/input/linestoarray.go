package input

import (
	"bufio"
	"io"
)

type Container interface {
	Accept(string)
}

func Parse(source io.Reader, c Container) {
	buffer := bufio.NewReader(source)
	for {
		line, _, err := buffer.ReadLine()
		if err == io.EOF {
			break
		}
		lineStr := string(line)
		c.Accept(lineStr)
	}
}
