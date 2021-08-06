package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDataTable_AddColumn(t *testing.T) {
	dt := &DataTable{}
	dt.Init()

	NewValueCol(dt, "col1")
	assert.Len(t, dt.table, 1)
	assert.Len(t, dt.headings, 1)
}

func TestDataTable_AddRow(t *testing.T) {
	dt := &DataTable{}
	dt.Init()

	NewValueCol(dt, "col1")
	for v := 1.0; v <= 10.0; v++ {
		dt.NewRow(time.Now())
		dt.SetColValue("col1", v)
	}
	assert.Len(t, dt.table[0], 10)
}

func TestDataTable_UniformityTest(t *testing.T) {
	dt := &DataTable{}
	dt.Init()

	NewValueCol(dt, "col1")
	NewValueCol(dt, "col2")
	for v := 1.0; v <= 10.0; v++ {
		dt.NewRow(time.Now())
		dt.SetColValue("col1", v)
		dt.SetColValue("col2", v)
	}
	assert.True(t, dt.UniformityTest())
}
