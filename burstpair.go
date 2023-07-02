package pronto

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidBurstPairString = errors.New("invalid burst pair string")

type BurstPair struct {
	a, b uint16
}

func NewBurstPairFromString(in string) (*BurstPair, error) {
	rd := newSequenceReader(strings.NewReader(in))
	a, err := rd.ReadSequence()
	if err != nil {
		return nil, fmt.Errorf("%w: burst pair A failed to parse", err)
	}

	b, err := rd.ReadSequence()
	if err != nil {
		return nil, fmt.Errorf("%w: burst pair B failed to parse", err)
	}

	return &BurstPair{
		a: a,
		b: b,
	}, nil
}

func (bp *BurstPair) EqualDetail(e *BurstPair) (bool, string) {
	bpaf := float64(bp.a)
	eaf := float64(e.a)
	apc := bpaf * allowedSignalDeviation
	if apc < minimumSignalDeviation {
		apc = minimumSignalDeviation
	}

	if (bpaf+apc > eaf) && (bpaf-apc < eaf) {
		bpbf := float64(bp.b)
		ebf := float64(e.b)
		bpc := float64(bp.b) * allowedSignalDeviation
		if bpc < minimumSignalDeviation {
			bpc = minimumSignalDeviation
		}

		if (bpbf+bpc > ebf) && (bpbf-bpc < ebf) {
			return true, "OK"
		}
	}

	return false, fmt.Sprintf("pair is outside tollerance: '%s' !~ '%s'", bp.String(), e.String())
}

func (bp *BurstPair) Equal(e *BurstPair) bool {
	v, _ := bp.EqualDetail(e)
	return v
}

func (bp *BurstPair) String() string {
	return bp.string(true)
}

func (bp *BurstPair) StringWithSpaces() string {
	return bp.string(false)
}

func (bp *BurstPair) string(commas bool) string {
	if commas {
		return fmt.Sprintf("%04x,%04x", bp.a, bp.b)
	}

	return fmt.Sprintf("%04x %04x", bp.a, bp.b)
}
