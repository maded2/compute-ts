package public

import "gonum.org/v1/gonum/graph"

type Column interface {
	graph.Node
	GetName() string
	Depends() []string
	MinimumValues() int
	Evaluate(dt *DataTable, rowIndex int, previousValues []float64)
	SetStartValue(offset int)
	GetStartValue() int
}
