package template_test

import (
	"testing"

	number "github.com/go-random/template"
	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Int()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized signed integer should not be 0")
}
