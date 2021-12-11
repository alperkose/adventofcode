package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EntryCalculatesDigitValues(t *testing.T) {
	testCases := []struct {
		entry         Entry
		expectedValue int
	}{
		{
			entry: Entry{
				[]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
				[]string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
			},
			expectedValue: 5353,
		},
		{
			entry: Entry{
				[]string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"},
				[]string{"fcgedb", "cgb", "dgebacf", "gc"},
			},
			expectedValue: 9781,
		},
		{
			entry: Entry{
				[]string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"},
				[]string{"cg", "cg", "fdcagb", "cbg"},
			},
			expectedValue: 1197,
		},
		{
			entry: Entry{
				[]string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"},
				[]string{"efabcd", "cedba", "gadfec", "cb"},
			},
			expectedValue: 9361,
		},
		{
			entry: Entry{
				[]string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"},
				[]string{"gecf", "egdcabf", "bgf", "bfgea"},
			},
			expectedValue: 4873,
		},
		{
			entry: Entry{
				[]string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"},
				[]string{"gebdcfa", "ecba", "ca", "fadegcb"},
			},
			expectedValue: 8418,
		},
		{
			entry: Entry{
				[]string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"},
				[]string{"efabcd", "fegcbd", "gadfec", "cb"},
			},
			expectedValue: 9061,
		},
		{
			entry: Entry{
				[]string{"efcagd", "dbcef", "bgda", "gbf", "dfaegb", "cfegab", "bfged", "bg", "agcfbde", "afegd"},
				[]string{"bfagce", "cefdb", "fgeda", "aegfcd"},
			},
			expectedValue: 256,
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintln(tC.entry), func(t *testing.T) {
			assert.Equal(t, tC.expectedValue, tC.entry.DigitValues())
		})
	}
}
