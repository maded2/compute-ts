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
	NewSimpleMovingAverage(dt, 5, "CloseSma", 3, "Close")
	dt.FinishRegistration()

}
