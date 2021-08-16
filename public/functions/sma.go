package functions

import (
	"compute-ts/public"
	"gonum.org/v1/gonum/stat"
)

type SimpleMovingAverage struct {
	public.DerivedColumn
}

func NewSimpleMovingAverage(dt *public.DataTable, name string, lookback int, dependsOn ...string) *SimpleMovingAverage {
	col := &SimpleMovingAverage{}
	col.Name = name
	col.DependOn = dependsOn
	col.DependentOffset = lookback
	col.Id = dt.RegisterColumn(col)
	return col
}

func (col *SimpleMovingAverage) Evaluate(dt *public.DataTable, rowIndex int, previousValues []float64) {
	x := dt.GetColumn(col.DependOn[0])
	if len(x) >= col.StartValues {
		dt.SetColValue(col.Name, stat.Mean(x[rowIndex-col.DependentOffset:], nil))
	}
}
