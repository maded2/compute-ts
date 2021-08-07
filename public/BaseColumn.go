package public

type BaseColumn struct {
	Id   int64
	Name string
}

var _ Column = &BaseColumn{}

func NewValueCol(dt *DataTable, id int64, name string) *BaseColumn {
	col := &BaseColumn{Id: id, Name: name}
	dt.RegisterColumn(col)
	return col
}

func (col *BaseColumn) ID() int64 {
	return col.Id
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

func (col *BaseColumn) Evaluate(dt *DataTable) {
	return
}
