package mannkendall

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

type Trending string

const (
	Increasing Trending = "increasing"
	Descending Trending = "descending"
	None       Trending = "none"
)

type MannKendall struct {
	Trend Trending
	H     bool
	P     float64
	Z     float64
}

// MannKendall performs the Mann-Kendall test on the data slice x.
// Returns the trend ("increasing", "decreasing", or "no trend"),
// the boolean result of the test, the p-value, and the test statistic z.
func Test(x []float64, alpha float64) MannKendall {
	n := len(x)

	// Calculate S.
	s := float64(0)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if x[i] > x[j] {
				s++
			} else if x[i] < x[j] {
				s--
			}
		}
	}

	// Calculate the unique values.
	uniq := make([]float64, 0, n)
	m := make(map[float64]int)
	for _, val := range x {
		_, ok := m[val]
		if !ok {
			uniq = append(uniq, val)
			m[val] = 1
		} else {
			m[val]++
		}
	}
	g := len(uniq)

	// Calculate the variance.
	var v float64
	if n == g { // No ties
		v = float64(n*(n-1)*(2*n+5)) / 18
	} else { // There are ties
		var tp int
		for _, c := range m {
			tp += c * (c - 1) * (2*c + 5)
		}
		v = float64(n*(n-1)*(2*n+5)-tp) / 18
	}

	// Calculate the test statistic z.
	var z float64
	if s > 0 {
		z = (s - 1) / math.Sqrt(v)
	} else if s < 0 {
		z = (s + 1) / math.Sqrt(v)
	} else {
		z = 0
	}

	// Calculate the p-value.
	p := 2 * (1 - distuv.UnitNormal.CDF(math.Abs(z)))
	// Determine if trend is present.
	h := math.Abs(z) > distuv.UnitNormal.Quantile(1-alpha/2)

	var trend Trending
	if (z < 0) && h {
		trend = Descending
	} else if (z > 0) && h {
		trend = Increasing
	} else {
		trend = None
	}

	return MannKendall{
		Trend: trend,
		H:     h,
		P:     p,
		Z:     z,
	}
}
