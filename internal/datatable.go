package internal

import (
	"sort"
	"time"
)

type DataTable struct {
	columns    []Column
	headings   map[string]int
	timestamps []time.Time
	table      [][]float64
}

func (dt *DataTable) RegisterColumn(col Column) {
	dt.columns = append(dt.columns, col)
}

func (dt *DataTable) Init() {
	sort.Slice(dt.columns, func(i, j int) bool {
		return dt.columns[i].MinimumValues() < dt.columns[j].MinimumValues()
	})
}

func (dt *DataTable) AddTimestamp(t time.Time) {
	dt.timestamps = append(dt.timestamps, t)
}

func (dt *DataTable) AddColumn(name string) {
	if dt.headings == nil {
		dt.headings = map[string]int{}
	}
	dt.headings[name] = len(dt.table)
	dt.table = append(dt.table, []float64{})
}

func (dt *DataTable) AddRow(name string, value float64) {
	if colIdx, found := dt.headings[name]; found {
		dt.table[colIdx] = append(dt.table[colIdx], value)
	}
}

func (dt *DataTable) UniformityTest() bool {
	for _, idx := range dt.headings {
		if len(dt.table[idx]) != len(dt.timestamps) {
			return false
		}
	}
	return true
}
