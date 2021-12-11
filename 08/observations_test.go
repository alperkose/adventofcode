package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParsesEntry(t *testing.T) {
	entry := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	obs := &Observations{}
	obs.Accept(entry)

	assert.Equal(t, []Entry{
		{
			[]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
			[]string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
		}},
		obs.Entries(),
	)
}
