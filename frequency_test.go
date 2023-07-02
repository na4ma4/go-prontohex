package pronto_test

import (
	"testing"

	pronto "github.com/na4ma4/go-prontohex"
)

func TestFrequency_EqualDetail(t *testing.T) {
	btf := pronto.NewFrequency(40000)
	tests := []struct {
		name   string
		target uint16
		want   bool
	}{

		{target: 35600, want: false, name: "-11%"},
		{target: 36000, want: false, name: "-10%"},
		{target: 36400, want: false, name: "-9%"},
		{target: 36800, want: false, name: "-8%"},
		{target: 37200, want: false, name: "-7%"},
		{target: 37600, want: false, name: "-6%"},
		{target: 38000 - 1, want: false, name: "-5%-1"},
		{target: 38000, want: false, name: "-5%"},
		{target: 38000 + 1, want: true, name: "-5%+1"},
		{target: 38400, want: true, name: "-4%"},
		{target: 38800, want: true, name: "-3%"},
		{target: 39200, want: true, name: "-2%"},
		{target: 39600, want: true, name: "-1%"},

		{target: 40000, want: true, name: "Exact Match"},
		{target: 40400, want: true, name: "+1%"},
		{target: 40800, want: true, name: "+2%"},
		{target: 41200, want: true, name: "+3%"},
		{target: 41600, want: true, name: "+4%"},
		{target: 42000 - 1, want: true, name: "+5%-1"},
		{target: 42000, want: false, name: "+5%"},
		{target: 42000 + 1, want: false, name: "+5%+1"},
		{target: 42400, want: false, name: "+6%"},
		{target: 42800, want: false, name: "+7%"},
		{target: 43200, want: false, name: "+8%"},
		{target: 43600, want: false, name: "+9%"},
		{target: 44000, want: false, name: "+10%"},
		{target: 44400, want: false, name: "+11%"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, msg := btf.EqualDetail(pronto.Frequency(tt.target))
			if got != tt.want {
				t.Errorf("Frequency.EqualDetail() got = %v, want %v, msg = %s", got, tt.want, msg)
			}
		})
	}
}

func TestFrequency_EqualDetail_SmallFrequency(t *testing.T) {
	btf := pronto.NewFrequency(10)
	tests := []struct {
		name   string
		target uint16
		want   bool
	}{
		{target: 0, want: false, name: "-10%"},
		{target: 1, want: false, name: "-9%"},
		{target: 2, want: false, name: "-8%"},
		{target: 3, want: false, name: "-7%"},
		{target: 4, want: false, name: "-6%"},
		{target: 5, want: false, name: "-5%"},
		{target: 6, want: false, name: "-4%"},
		{target: 7, want: true, name: "-3%"},
		{target: 8, want: true, name: "-2%"},
		{target: 9, want: true, name: "-1%"},
		{target: 10, want: true, name: "Exact Match"},
		{target: 11, want: true, name: "+1%"},
		{target: 12, want: true, name: "+2%"},
		{target: 13, want: true, name: "+3%"},
		{target: 14, want: false, name: "+4%"},
		{target: 15, want: false, name: "+5%"},
		{target: 16, want: false, name: "+6%"},
		{target: 17, want: false, name: "+7%"},
		{target: 18, want: false, name: "+8%"},
		{target: 19, want: false, name: "+9%"},
		{target: 20, want: false, name: "+10%"},
		{target: 21, want: false, name: "+11%"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, msg := btf.EqualDetail(pronto.Frequency(tt.target))
			if got != tt.want {
				t.Errorf("Frequency.EqualDetail() got = %v, want %v, msg = %s", got, tt.want, msg)
			}
		})
	}
}

func TestFrequency_Equal(t *testing.T) {
	btf := pronto.NewFrequency(40000)
	tests := []struct {
		name   string
		target uint16
		want   bool
	}{

		{target: 35600, want: false, name: "-11%"},
		{target: 36000, want: false, name: "-10%"},
		{target: 36400, want: false, name: "-9%"},
		{target: 36800, want: false, name: "-8%"},
		{target: 37200, want: false, name: "-7%"},
		{target: 37600, want: false, name: "-6%"},
		{target: 38000 - 1, want: false, name: "-5%-1"},
		{target: 38000, want: false, name: "-5%"},
		{target: 38000 + 1, want: true, name: "-5%+1"},
		{target: 38400, want: true, name: "-4%"},
		{target: 38800, want: true, name: "-3%"},
		{target: 39200, want: true, name: "-2%"},
		{target: 39600, want: true, name: "-1%"},

		{target: 40000, want: true, name: "Exact Match"},
		{target: 40400, want: true, name: "+1%"},
		{target: 40800, want: true, name: "+2%"},
		{target: 41200, want: true, name: "+3%"},
		{target: 41600, want: true, name: "+4%"},
		{target: 42000 - 1, want: true, name: "+5%-1"},
		{target: 42000, want: false, name: "+5%"},
		{target: 42000 + 1, want: false, name: "+5%+1"},
		{target: 42400, want: false, name: "+6%"},
		{target: 42800, want: false, name: "+7%"},
		{target: 43200, want: false, name: "+8%"},
		{target: 43600, want: false, name: "+9%"},
		{target: 44000, want: false, name: "+10%"},
		{target: 44400, want: false, name: "+11%"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := btf.Equal(pronto.Frequency(tt.target))
			if got != tt.want {
				t.Errorf("Frequency.EqualDetail() for target = %d, got = %v, want %v", tt.target, got, tt.want)
			}
		})
	}
}
