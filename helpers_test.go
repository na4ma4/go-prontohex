package pronto_test

import (
	"testing"

	pronto "github.com/na4ma4/go-prontohex"
)

func TestFrequencyToHZ(t *testing.T) {
	tests := []struct {
		name string
		freq pronto.Frequency
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "Common Sony Frequency",
			freq: 0x0067,
			want: 40244,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pronto.FrequencyToHZ(tt.freq); got != tt.want {
				t.Errorf("FrequencyToHZ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func prontoFromStringTest(t *testing.T, in string) pronto.Code {
	c, err := pronto.FromString(in)
	if err != nil {
		t.Fatalf("unable to parse prontohex string into code object: %s", err)
	}

	return c
}

func TestSonyDeviceCode(t *testing.T) {
	tests := []struct {
		name string
		code pronto.Code
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "Valid Sony IR Code",
			code: prontoFromStringTest(t, sonyBasicTestCode),
			want: 2362,
		},
		{
			name: "Not Valid Sony IR Code",
			code: prontoFromStringTest(t, basicTestCode),
			want: 0,
		},
		{
			name: "Not Enough Burst Pairs to Decode",
			code: prontoFromStringTest(t, minimumValidProntoCode),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pronto.SonyDeviceCode(tt.code); got != tt.want {
				t.Errorf("SonyDeviceCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSonyRemoteButton(t *testing.T) {
	tests := []struct {
		name string
		code pronto.Code
		want int16
	}{
		// TODO: Add test cases.
		{
			name: "Valid Sony IR Code",
			code: prontoFromStringTest(t, sonyBasicTestCode),
			want: 46,
		},
		{
			name: "Not Valid Sony IR Code",
			code: prontoFromStringTest(t, basicTestCode),
			want: 0,
		},
		{
			name: "Not Enough Burst Pairs to Decode",
			code: prontoFromStringTest(t, minimumValidProntoCode),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pronto.SonyRemoteButton(tt.code); got != tt.want {
				t.Errorf("SonyRemoteButton() = %v, want %v", got, tt.want)
			}
		})
	}
}
