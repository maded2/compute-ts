package functions

import (
	"compute-ts/public"
	"gonum.org/v1/gonum/stat"
)

type SimpleMovingAverage struct {
	public.DerivedColumn
}

func NewSimpleMovingAverage(dt *public.DataTable, id int64, name string, lookback int, dependsOn ...string) *SimpleMovingAverage {
	col := &SimpleMovingAverage{}
	col.Id = id
	col.Name = name
	col.DependOn = dependsOn
	col.DependentOffset = lookback
	dt.RegisterColumn(col)
	return col
}

func (col *SimpleMovingAverage) Evaluate(dt *public.DataTable) {
	x := dt.GetColumn(col.DependOn[0])
	if len(x) > col.StartValues {
		dt.SetColValue(col.Name, stat.Mean(x, nil))
	}
}
