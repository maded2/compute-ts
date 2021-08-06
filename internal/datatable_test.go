package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDataTable_AddColumn(t *testing.T) {
	dt := &DataTable{}
	dt.Init()

	NewValue(dt, "col1")
	assert.Len(t, dt.table, 1)
	assert.Len(t, dt.headings, 1)
}

func TestDataTable_AddRow(t *testing.T) {
	dt := &DataTable{}
	dt.Init()

	NewValue(dt, "col1")
	for v := 1.0; v <= 10.0; v++ {
		dt.SetValue("col1", v)
	}
	assert.Len(t, dt.table[0], 10)
}

func TestDataTable_UniformityTest(t *testing.T) {
	dt := &DataTable{}
	dt.Init()

	NewValue(dt, "col1")
	NewValue(dt, "col2")
	for v := 1.0; v <= 11.0; v++ {
		dt.NewRow(time.Now())
	}
	for v := 1.0; v <= 10.0; v++ {
		dt.SetValue("col1", v)
	}
	for v := 1.0; v <= 11.0; v++ {
		dt.SetValue("col2", v)
	}
	assert.False(t, dt.UniformityTest())
	dt.SetValue("col1", 11)
	assert.True(t, dt.UniformityTest())
}
