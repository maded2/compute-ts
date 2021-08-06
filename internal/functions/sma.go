package functions

import (
	"compute-ts/internal"
	"gonum.org/v1/gonum/stat"
)

type SimpleMovingAverage struct {
	internal.DerivedColumn
}

func NewSimpleMovingAverage(dt *internal.DataTable, name string, lookback int, dependsOn ...string) *SimpleMovingAverage {
	col := &SimpleMovingAverage{}
	col.DT = dt
	col.Name = name
	col.DependOn = dependsOn
	col.DependentOffset = lookback
	dt.RegisterColumn(col)
	return col
}

func (col *SimpleMovingAverage) Exec() {
	x := col.DT.GetColumn(col.DependOn[0])
	if len(x) > col.StartValues {
		col.DT.SetValue(col.Name, stat.Mean(x, nil))
	}
}
