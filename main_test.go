package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Main", func() {
		g.It("Should show Hello world!", func() {
			message := HelloWorld()
			g.Assert(message).Equal("Hello, world!")
		})
	})
}