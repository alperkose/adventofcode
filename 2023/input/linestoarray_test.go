package input_test

import (
	"strings"
	"testing"

	"github.com/alperkose/adventofcode/2023/input"
	"github.com/stretchr/testify/assert"
)

type TestContainer struct {
	items []string
}

func (tc *TestContainer) Accept(in string) {
	tc.items = append(tc.items, in)
}

func Test_ParseWithFunction(t *testing.T) {
	sampleInput := "123\n456\n789"
	testContainer := &TestContainer{}

	input.Parse(strings.NewReader(sampleInput), testContainer)

	assert.Equal(t, []string{"123", "456", "789"}, testContainer.items)
}
