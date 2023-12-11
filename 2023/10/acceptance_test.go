package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SampleInputForPart1(t *testing.T) {
	schematics := &Schematics{}
	schematics.Accept(".....")
	schematics.Accept(".S-7.")
	schematics.Accept(".|.|.")
	schematics.Accept(".L-J.")
	schematics.Accept(".....")

	assert.Equal(t, 4, schematics.SolvePt1())
	assert.Equal(t, 1, schematics.SolvePt2())

}

func Test_SampleInputForPart1_WithNonConnectingPipes(t *testing.T) {
	schematics := &Schematics{}
	schematics.Accept("7-F7-")
	schematics.Accept(".FJ|7")
	schematics.Accept("SJLL7")
	schematics.Accept("|F--J")
	schematics.Accept("LJ.LJ")

	assert.Equal(t, 8, schematics.SolvePt1())
}
func Test_MultiplePipesFromStaringPosition(t *testing.T) {
	schematics := &Schematics{}
	schematics.Accept("|F--7")
	schematics.Accept("LS-7|")
	schematics.Accept("||.||")
	schematics.Accept("|L-J|")
	schematics.Accept("L---J")

	assert.Equal(t, 4, schematics.SolvePt1())
}

func Test_SampleForPart2(t *testing.T) {
	schematics := &Schematics{}
	schematics.Accept("...........")
	schematics.Accept(".S-------7.")
	schematics.Accept(".|F-----7|.")
	schematics.Accept(".||.....||.")
	schematics.Accept(".||.....||.")
	schematics.Accept(".|L-7.F-J|.")
	schematics.Accept(".|..|.|..|.")
	schematics.Accept(".L--J.L--J.")
	schematics.Accept("...........")

	assert.Equal(t, 4, schematics.SolvePt2())

}
func Test_ComplexSampleForPart2(t *testing.T) {
	schematics := &Schematics{}
	schematics.Accept("FF7FSF7F7F7F7F7F---7")
	schematics.Accept("L|LJ||||||||||||F--J")
	schematics.Accept("FL-7LJLJ||||||LJL-77")
	schematics.Accept("F--JF--7||LJLJIF7FJ-")
	schematics.Accept("L---JF-JLJIIIIFJLJJ7")
	schematics.Accept("|F|F-JF---7IIIL7L|7|")
	schematics.Accept("|FFJF7L7F-JF7IIL---7")
	schematics.Accept("7-L-JL7||F7|L7F-7F7|")
	schematics.Accept("L.L7LFJ|||||FJL7||LJ")
	schematics.Accept("L7JLJL-JLJLJL--JLJ.L")

	assert.Equal(t, 10, schematics.SolvePt2())

}
