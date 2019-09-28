package scope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	A string `readScope:"read:A" writeScope:"write:A"`
	B string `readScope:"read:B" writeScope:"write:B"`
}

func TestOperationAllowed(t *testing.T) {
	assert := assert.New(t)

	output := testStruct{
		A: "A",
		B: "B",
	}

	FilterRead(&output, []string{"read:A"})
	assert.Zero(output.B)
	assert.Equal("A", output.A)

}
