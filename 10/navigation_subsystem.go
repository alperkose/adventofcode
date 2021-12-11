package main

type NavigationSubSystem struct {
	chunks []string
}

func (nav *NavigationSubSystem) Accept(in string) {
	nav.chunks = append(nav.chunks, in)
}
