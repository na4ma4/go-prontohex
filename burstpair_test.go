package pronto_test

import (
	"testing"

	pronto "github.com/na4ma4/go-prontohex"
)

func TestBurstPair_String(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Spaces",
			input: "0018 0018",
			want:  "0018,0018",
		},
		{
			name:  "Commas",
			input: "0018,0018",
			want:  "0018,0018",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := pronto.NewBurstPairFromString(tt.input)
			if err != nil {
				t.Errorf("NewBurstPairFromString() error %v", err)
			}
			if got := b.String(); got != tt.want {
				t.Errorf("BurstPair.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBurstPair_String_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Invalid number in both parts of pair",
			input: "001 018",
		},
		{
			name:  "Invalid number in second part of pair",
			input: "0018 018",
		},
		{
			name:  "Invalid number in first part of pair",
			input: "001 0180",
		},
		{
			name:  "Invalid null character as spacer",
			input: "0010\x000180",
		},
		{
			name:  "Not valid hex",
			input: "0010 018G",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := pronto.NewBurstPairFromString(tt.input)
			if err == nil {
				t.Errorf("NewBurstPairFromString() expected error with input = %s", tt.input)
			}
			if b != nil {
				t.Error("NewBurstPairFromString() should return b as nil when returning an error")
			}
		})
	}
}

func TestBurstPair_String_InvalidValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Invalid number of sequences, only first two used",
			input: "0000 0000 0000",
			want:  "0000,0000",
		},
		{
			name:  "Carriage returns in string",
			input: "0200\r\n0230",
			want:  "0200,0230",
		},
		{
			name:  "Tab in string",
			input: "0400\t0430",
			want:  "0400,0430",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := pronto.NewBurstPairFromString(tt.input)
			if err != nil {
				t.Errorf("NewBurstPairFromString() error %v", err)
			}
			if got := b.String(); got != tt.want {
				t.Errorf("BurstPair.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBurstPair_EqualDetail(t *testing.T) {
	tests := []struct {
		name  string
		left  string
		right string
		want  bool
	}{
		// TODO: Add test cases.
		{
			name:  "Equal Pairs",
			left:  "0000 0000",
			right: "0000 0000",
			want:  true,
		},
		{
			name:  "Equal Pairs Inside Deviation",
			left:  "0000 0100",
			right: "0000 0101",
			want:  true,
		},
		{
			name:  "Not Equal Pairs",
			left:  "0000 0100",
			right: "0000 0200",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bpLeft, err := pronto.NewBurstPairFromString(tt.left)
			if err != nil {
				t.Errorf("NewBurstPairFromString(left) returned error for input = %s, err = %v", tt.left, err)
			}

			bpRight, err := pronto.NewBurstPairFromString(tt.right)
			if err != nil {
				t.Errorf("NewBurstPairFromString(right) returned error for input = %s, err = %v", tt.left, err)
			}

			got, msg := bpLeft.EqualDetail(bpRight)
			if got != tt.want {
				t.Errorf("BurstPair.EqualDetail() got = %v, want = %v, msg = %s", got, tt.want, msg)
			}

			got1 := bpRight.Equal(bpLeft)
			if got != got1 {
				t.Errorf("BurstPair.EqualDetail() got = %v returned different to BurstPair.Equal() got = %v", got, got1)
			}
		})
	}
}
