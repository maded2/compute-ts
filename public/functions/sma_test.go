package functions

import (
	"compute-ts/public"
	"testing"
	"time"
)

func TestSma(t *testing.T) {
	dt := public.NewDataTable()
	public.NewValueCol(dt, "Open")
	public.NewValueCol(dt, "Close")
	NewSimpleMovingAverage(dt, "CloseSma1", 3, "Close")
	NewSimpleMovingAverage(dt, "CloseSma2", 3, "CloseSma1")
	NewSimpleMovingAverage(dt, "OpenSma1", 3, "Open")
	NewSimpleMovingAverage(dt, "OpenCloseSma", 3, "OpenSma1", "CloseSma2")
	dt.FinishRegistration()

	for i := 1; i < 20; i++ {
		dt.NewRow(time.Now())
		dt.SetColValue("Open", float64(i))
		dt.SetColValue("Close", float64(i))
		dt.ComputeRow()
	}
	t.Log(dt.Dump())

	t.Log(dt.Dot)
}
