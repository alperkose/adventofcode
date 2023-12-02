package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AcceptsBinaryString(t *testing.T) {
	testCases := []struct {
		input    string
		expected uint32
	}{
		{"001", uint32(1)},
		{"010", uint32(2)},
		{"011", uint32(3)},
		{"11011", uint32(27)},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			container := ByteContainer{}
			container.Accept(tC.input)
			assert.Equal(t, []uint32{tC.expected}, container.Values())
			assert.Equal(t, len(tC.input), container.Digits())
		})
	}
}
