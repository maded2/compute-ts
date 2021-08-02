package internal

type DerivedColumn struct {
	Name        string
	DependOn    []string
	StartValues int
}

var _ Column = &DerivedColumn{}

func (col *DerivedColumn) Depends() []string {
	return col.DependOn
}

func (col *DerivedColumn) MinimumValues() int {
	return col.StartValues
}
