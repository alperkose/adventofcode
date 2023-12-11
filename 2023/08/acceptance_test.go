package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SampleInputForPart1(t *testing.T) {
	network := NewNetwork()
	network.Accept("RL")
	network.Accept("")
	network.Accept("AAA = (BBB, CCC)")
	network.Accept("BBB = (DDD, EEE)")
	network.Accept("CCC = (ZZZ, GGG)")
	network.Accept("DDD = (DDD, DDD)")
	network.Accept("EEE = (EEE, EEE)")
	network.Accept("GGG = (GGG, GGG)")
	network.Accept("ZZZ = (ZZZ, ZZZ)")

	assert.Equal(t, 2, network.SolvePt1())
}

func Test_SecondSampleInputForPart1(t *testing.T) {
	network := NewNetwork()
	network.Accept("LLR")
	network.Accept("")
	network.Accept("AAA = (BBB, BBB)")
	network.Accept("BBB = (AAA, ZZZ)")
	network.Accept("ZZZ = (ZZZ, ZZZ)")

	assert.Equal(t, 6, network.SolvePt1())
}

func Test_SampleInputForPart2(t *testing.T) {
	network := NewNetwork()
	network.Accept("LR")
	network.Accept("")
	network.Accept("11A = (11B, XXX)")
	network.Accept("11B = (XXX, 11Z)")
	network.Accept("11Z = (11B, XXX)")
	network.Accept("22A = (22B, XXX)")
	network.Accept("22B = (22C, 22C)")
	network.Accept("22C = (22Z, 22Z)")
	network.Accept("22Z = (22B, 22B)")
	network.Accept("XXX = (XXX, XXX)")

	assert.Equal(t, 6, network.SolvePt2())
}
