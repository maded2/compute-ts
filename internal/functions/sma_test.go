package functions

import "testing"

func TestSma(t *testing.T) {
	sma := &SimpleMovingAverage{}
	sma.DependentOffset = 4
	sma.DependOn = []string{"Close"}
}
