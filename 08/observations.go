package main

import "strings"

type Observations struct {
	entries []Entry
}

func (o *Observations) Accept(in string) {
	signalsNdigits := strings.Split(in, " ")
	signals := signalsNdigits[0:10]
	digits := signalsNdigits[11:]
	o.entries = append(o.entries, Entry{signals, digits})
}

func (o *Observations) Entries() []Entry {
	return o.entries
}
