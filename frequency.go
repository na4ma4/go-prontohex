package pronto

import "fmt"

type Frequency uint16

func NewFrequency(fq uint16) Frequency {
	return Frequency(fq)
}

func NewFrequencyPTError(fq uint16, err error) (Frequency, error) {
	return Frequency(fq), err
}

func (f Frequency) EqualDetail(e Frequency) (bool, string) {
	cfq := float64(f)
	efq := float64(e)
	fpc := cfq * allowedSignalDeviation
	if fpc < minimumSignalDeviation {
		fpc = minimumSignalDeviation
	}

	if !((cfq+fpc > efq) && (cfq-fpc < efq)) {
		return false, fmt.Sprintf("frequency outside allowed difference (%d !~ %d)", f, e)
	}

	return true, "OK"
}

func (f Frequency) Equal(e Frequency) bool {
	v, _ := f.EqualDetail(e)
	return v
}
