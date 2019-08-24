package stat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcSum(t *testing.T) {
	assert.Equal(t, CalcSum(1, 2), 3)
}
