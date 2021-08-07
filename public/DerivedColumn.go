package public

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
	return col.DependentOffset
}

func (col *DerivedColumn) SetStartValue(offset int) {
	col.StartValues = offset
}

func (col *DerivedColumn) GetStartValue() int {
	return col.StartValues
}
