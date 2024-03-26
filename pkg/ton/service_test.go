package ton

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestTon(t *testing.T) {
	d1 := decimal.NewFromInt(3)
	d2 := decimal.NewFromInt(1000000)
	d3 := d1.Mul(d2)
	t.Logf(d3.String())
}
