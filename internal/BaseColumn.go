package internal

type BaseColumn struct {
	Name string
}

var _ Column = &BaseColumn{}

func (col *BaseColumn) Depends() []string {
	return nil
}

func (col *BaseColumn) MinimumValues() int {
	return 0
}
