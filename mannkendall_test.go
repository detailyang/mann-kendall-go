package mannkendall

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMannKendall(t *testing.T) {
	alpha := float64(0.05)
	x := []float64{1.1, 1.1, 1.1, 1.1, 1.1}
	mk := Test(x, alpha)
	assert.Equal(t, None, mk.Trend)

	x = []float64{1.1, 1.2, 1.5, 1.9, 2.1}
	mk = Test(x, alpha)
	assert.Equal(t, Increasing, mk.Trend)

	x = []float64{1.1, -1.2, -1.5, -1.9, -2.1}
	mk = Test(x, alpha)
	assert.Equal(t, Descending, mk.Trend)
}
