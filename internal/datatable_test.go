package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDataTable_AddColumn(t *testing.T) {
	dt := DataTable{}
	dt.AddColumn("col1")
	assert.Len(t, dt.table, 1)
	assert.Len(t, dt.headings, 1)
}

func TestDataTable_AddRow(t *testing.T) {
	dt := DataTable{}
	dt.AddColumn("col1")
	for v := 1.0; v <= 10.0; v++ {
		dt.AddRow("col1", v)
	}
	assert.Len(t, dt.table[0], 10)
}

func TestDataTable_UniformityTest(t *testing.T) {
	dt := DataTable{}
	dt.AddColumn("col1")
	dt.AddColumn("col2")
	for v := 1.0; v <= 11.0; v++ {
		dt.AddTimestamp(time.Now())
	}
	for v := 1.0; v <= 10.0; v++ {
		dt.AddRow("col1", v)
	}
	for v := 1.0; v <= 11.0; v++ {
		dt.AddRow("col2", v)
	}
	assert.False(t, dt.UniformityTest())
	dt.AddRow("col1", 11)
	assert.True(t, dt.UniformityTest())
}
