package input_test

import (
	"testing"

	"github.com/alperkose/adventofcode2021/input"
	"github.com/stretchr/testify/assert"
)

func Test_IntContainer_AcceptsNumberString(t *testing.T) {
	intContainer := &input.IntContainer{}

	intContainer.Accept("123")

	assert.Equal(t, []int{123}, intContainer.Values())
}
func Test_IntContainer_DoesNotAcceptNonNumberString(t *testing.T) {
	intContainer := &input.IntContainer{}

	intContainer.Accept("123.5")

	assert.Empty(t, intContainer.Values())
}
