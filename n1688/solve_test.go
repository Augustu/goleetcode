package n1688

import (
	"testing"

	"gotest.tools/assert"
)

type Sample struct {
	input  int
	output int
}

func TestNumberOfMatches(t *testing.T) {

	checks := []Sample{
		{
			input:  7,
			output: 6,
		},
		{
			input:  7,
			output: 6,
		},
	}

	for _, check := range checks {
		assert.Equal(t, numberOfMatches(check.input), check.output, "no equal")
	}

}
