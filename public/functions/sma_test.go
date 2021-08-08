package functions

import (
	"compute-ts/public"
	"testing"
	"time"
)

func TestSma(t *testing.T) {
	dt := public.NewDataTable()
	public.NewValueCol(dt, 1, "Open")
	public.NewValueCol(dt, 2, "Close")
	NewSimpleMovingAverage(dt, 3, "CloseSma1", 3, "Close")
	NewSimpleMovingAverage(dt, 4, "CloseSma2", 3, "CloseSma1")
	NewSimpleMovingAverage(dt, 5, "OpenSma1", 3, "Open")
	NewSimpleMovingAverage(dt, 6, "OpenCloseSma", 3, "OpenSma1", "CloseSma2")
	dt.FinishRegistration()

	for i := 1; i < 10; i++ {
		dt.NewRow(time.Now())
		dt.SetColValue("Open", float64(i))
		dt.SetColValue("Close", float64(i))
		dt.ComputeRow()
	}
	t.Log(dt.Dump())
}
