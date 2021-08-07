package functions

import (
	"compute-ts/public"
	"testing"
)

func TestSma(t *testing.T) {
	dt := public.NewDataTable()
	public.NewValueCol(dt, 1, "Open")
	public.NewValueCol(dt, 2, "Close")
	public.NewValueCol(dt, 3, "High")
	public.NewValueCol(dt, 4, "Low")
	NewSimpleMovingAverage(dt, 5, "CloseSma1", 3, "Close")
	NewSimpleMovingAverage(dt, 6, "CloseSma2", 3, "CloseSma1")
	NewSimpleMovingAverage(dt, 7, "OpenSma1", 3, "Open")
	NewSimpleMovingAverage(dt, 8, "OpenCloseSma", 3, "OpenSma1", "CloseSma2")
	dt.FinishRegistration()

}
