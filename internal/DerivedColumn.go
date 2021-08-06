package internal

type DerivedColumn struct {
	BaseColumn
	DependOn        []string
	StartValues     int
	DependentOffset int
}

var _ Column = &DerivedColumn{}

func (col *DerivedColumn) Depends() []string {
	return col.DependOn
}

func (col *DerivedColumn) MinimumValues() int {
	return col.StartValues
}
