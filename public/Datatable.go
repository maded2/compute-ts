package public

import (
	"fmt"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"log"
	"math"
	"time"
)

type DataTable struct {
	columns     []Column
	headings    map[string]int
	timestamps  []time.Time
	table       [][]float64
	sortedNodes []Column
}

func (dt *DataTable) RegisterColumn(col Column) {
	dt.headings[col.GetName()] = len(dt.table)
	dt.table = append(dt.table, []float64{})
	dt.columns = append(dt.columns, col)
}

func NewDataTable() (dt *DataTable) {
	dt = &DataTable{
		headings: map[string]int{},
	}
	return dt
}

func (dt *DataTable) FinishRegistration() (err error) {
	nodeMap := map[string]Column{}
	graph := simple.NewDirectedGraph()
	for _, col := range dt.columns {
		graph.AddNode(col)
		nodeMap[col.GetName()] = col
	}
	for _, col := range dt.columns {
		for _, parent := range col.Depends() {
			if parentCol, found := nodeMap[parent]; found {
				graph.SetEdge(simple.Edge{F: parentCol, T: col})
			}
		}
	}
	s := ""
	if sortedNodes, err := topo.Sort(graph); err == nil {
		offset := 0
		for _, node := range sortedNodes {
			col := node.(Column)
			offset += col.MinimumValues()
			col.SetStartValue(offset)
			dt.sortedNodes = append(dt.sortedNodes, col)
			if len(s) > 0 {
				s += " -> "
			}
			s += fmt.Sprintf("%s[%d]", col.GetName(), col.GetStartValue())
		}
		log.Println(s)
	}
	return
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
