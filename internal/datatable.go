package internal

import (
	"math"
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
	dt.headings[col.GetName()] = len(dt.table)
	dt.table = append(dt.table, []float64{})
	dt.columns = append(dt.columns, col)
}

func (dt *DataTable) Init() {
	dt.headings = map[string]int{}
	sort.Slice(dt.columns, func(i, j int) bool {
		return dt.columns[i].MinimumValues() < dt.columns[j].MinimumValues()
	})
}

func (dt *DataTable) NewRow(timestamp time.Time) {
	dt.timestamps = append(dt.timestamps, timestamp)
	for _, colIdx := range dt.headings {
		dt.table[colIdx] = append(dt.table[colIdx], math.NaN())
	}
}

func (dt *DataTable) SetColValue(name string, value float64) {
	if colIdx, found := dt.headings[name]; found {
		col := dt.table[colIdx]
		col[len(col)-1] = value
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

func (dt *DataTable) GetColumn(name string) []float64 {
	if idx, found := dt.headings[name]; found {
		return dt.table[idx]
	} else {
		return nil
	}
}
