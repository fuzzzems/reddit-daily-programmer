package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsConsonant(t *testing.T) {

	assert.Equal(t, "Jojagog totalolaror Rorövovarorsospoproråkoketot!", Robberify("Jag talar Rövarspråket!"))
	assert.Equal(t, "I'mom sospopeakokinongog Rorobobboberor'sos lolanongoguagoge!", Robberify("I'm speaking Robber's language!"))
}
