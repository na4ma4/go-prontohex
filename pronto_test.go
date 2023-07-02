package pronto_test

import (
	"reflect"
	"testing"

	pronto "github.com/na4ma4/go-prontohex"
)

//nolint:gocognit // test code
func TestFromString(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		want       string
		wantErr    bool
		wantCommas bool
	}{
		{
			name:    "Basic Test",
			input:   basicTestCode,
			wantErr: false,
		},
		{
			name:    "Sony Basic Test",
			input:   sonyBasicTestCode,
			wantErr: false,
		},
		{
			name:    "Minimum Valid Pronto Hex",
			input:   minimumValidProntoCode,
			wantErr: false,
		},
		{
			name:    "Sony Basic Test with Commas",
			input:   sonyBasicTestCodeCommas,
			want:    sonyBasicTestCode,
			wantErr: false,
		},
		{
			name:       "Sony Basic Test with Commas returning Commas",
			input:      sonyBasicTestCodeCommas,
			want:       sonyBasicTestCodeCommas,
			wantErr:    false,
			wantCommas: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pronto.FromString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				t.Error("FromString() returned nil code")
			}

			var gotStr string
			if tt.wantCommas {
				gotStr = got.String()
			} else {
				gotStr = got.StringWithSpaces()
			}

			if tt.want != "" {
				if !reflect.DeepEqual(gotStr, tt.want) {
					t.Errorf("FromString() = %v, want %v", gotStr, tt.want)
				}
			} else {
				if !reflect.DeepEqual(gotStr, tt.input) {
					t.Errorf("FromString() = %v, want %v", gotStr, tt.input)
				}
			}
		})
	}
}

func TestHexCode_EqualDetail(t *testing.T) {
	tests := []struct {
		name  string
		left  pronto.Code
		right pronto.Code
		want  bool
	}{
		{
			name:  "Sony Basic Test",
			left:  prontoFromStringTest(t, sonyBasicTestCode),
			right: prontoFromStringTest(t, sonyBasicTestCodeCommas),
			want:  true,
		},
		{
			name:  "Different Codes",
			left:  prontoFromStringTest(t, basicTestCode),
			right: prontoFromStringTest(t, sonyBasicTestCodeCommas),
			want:  false,
		},
		{
			name:  "Different Codes",
			left:  prontoFromStringTest(t, sonyBasicTestCodeCommas),
			right: prontoFromStringTest(t, sonyBasicTestCodeMinorChange),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, msg := tt.left.EqualDetail(tt.right)
			if got != tt.want {
				t.Errorf("HexCode.EqualDetail() got = %v, want = %v, msg = %s", got, tt.want, msg)
			}

			got1 := tt.left.Equal(tt.right)
			if got1 != tt.want {
				t.Errorf("HexCode.Equal() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
