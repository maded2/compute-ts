package public

import "gonum.org/v1/gonum/graph"

type Column interface {
	graph.Node
	GetName() string
	Depends() []string
	MinimumValues() int
	Evaluate(dt *DataTable)
	SetStartValue(offset int)
	GetStartValue() int
}
