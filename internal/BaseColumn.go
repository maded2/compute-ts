package internal

type BaseColumn struct {
	Name string
	DT   *DataTable
}

var _ Column = &BaseColumn{}

func NewValueCol(dt *DataTable, name string) *BaseColumn {
	col := &BaseColumn{}
	col.DT = dt
	col.Name = name
	dt.RegisterColumn(col)
	return col
}

func (col *BaseColumn) GetName() string {
	return col.Name
}
func (col *BaseColumn) Depends() []string {
	return nil
}

func (col *BaseColumn) MinimumValues() int {
	return 0
}
