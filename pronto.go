package pronto

import (
	"fmt"
	"strings"
)

type Code interface {
	StringWithSpaces() string
	String() string
	InfoDump() string
	Equal(Code) bool
	EqualDetail(Code) (bool, string)

	GetSource() uint16
	SetSource(source uint16)

	GetFrequency() Frequency
	SetFrequency(freq Frequency)

	GetBurstPairA() BurstPairSequence
	AppendBurstPairA(bp *BurstPair)

	GetBurstPairB() BurstPairSequence
	AppendBurstPairB(bp *BurstPair)
}

type HexCode struct {
	source                 uint16
	freq                   Frequency
	burstPairA, burstPairB BurstPairSequence
}

func New() Code {
	return &HexCode{
		burstPairA: BurstPairSequence{},
		burstPairB: BurstPairSequence{},
	}
}

func FromString(in string) (Code, error) {
	var (
		lenA, lenB   uint16
		wordA, wordB uint16
		err          error
	)

	c := &HexCode{}

	rd := newSequenceReader(strings.NewReader(in))
	c.source, err = rd.ReadSequence()
	if err != nil {
		return c, err
	}

	c.freq, err = NewFrequencyPTError(rd.ReadSequence())
	if err != nil {
		return c, err
	}

	if lenA, err = rd.ReadSequence(); err != nil {
		return c, fmt.Errorf("%w: invalid sequence A length", err)
	}
	c.burstPairA = make(BurstPairSequence, lenA)

	if lenB, err = rd.ReadSequence(); err != nil {
		return c, fmt.Errorf("%w: invalid sequence B length", err)
	}
	c.burstPairB = make(BurstPairSequence, lenB)

	for idxA := 0; idxA < int(lenA); idxA++ {
		wordA, err = rd.ReadSequence()
		if err != nil {
			return c, fmt.Errorf("%w: sequence read failed wordA/idxA[%d/%d]", err, idxA, lenA)
		}

		wordB, err = rd.ReadSequence()
		if err != nil {
			return c, fmt.Errorf("%w: sequence read failed wordB/idxA[%d/%d]", err, idxA, lenA)
		}

		c.burstPairA[idxA] = &BurstPair{
			a: wordA,
			b: wordB,
		}
	}

	for idxB := 0; idxB < int(lenB); idxB++ {
		wordA, err = rd.ReadSequence()
		if err != nil {
			return c, fmt.Errorf("%w: sequence read failed wordA/idxB[%d/%d]", err, idxB, lenB)
		}

		wordB, err = rd.ReadSequence()
		if err != nil {
			return c, fmt.Errorf("%w: sequence read failed wordB/idxB[%d/%d]", err, idxB, lenB)
		}

		c.burstPairB[idxB] = &BurstPair{
			a: wordA,
			b: wordB,
		}
	}

	return c, nil
}

func (c *HexCode) StringWithSpaces() string {
	return c.string(false)
}

func (c *HexCode) String() string {
	return c.string(true)
}

func (c *HexCode) stringBurstPair(commas, leading bool, bps BurstPairSequence) string {
	var (
		sb strings.Builder
	)

	for _, bp := range bps {
		if commas {
			sb.WriteString(fmt.Sprintf(",%s", bp.String()))
		} else {
			sb.WriteString(fmt.Sprintf(" %s", bp.StringWithSpaces()))
		}
	}

	if len(bps) > 0 && !leading {
		return sb.String()[1:]
	}

	return sb.String()
}

func (c *HexCode) string(commas bool) string {
	var sb strings.Builder
	hdr := "%04x %04x %04x %04x"
	if commas {
		hdr = "%04x,%04x,%04x,%04x"
	}
	sb.WriteString(fmt.Sprintf(
		hdr,
		c.source,
		c.freq,
		len(c.burstPairA),
		len(c.burstPairB),
	))
	sb.WriteString(c.stringBurstPair(commas, true, c.burstPairA))
	sb.WriteString(c.stringBurstPair(commas, true, c.burstPairB))

	return sb.String()
}

func (c *HexCode) InfoDump() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Source: %04x (%d)\n", c.source, c.source))
	sb.WriteString(fmt.Sprintf("Frequency: %04x (%d) [%d Hz]\n", c.freq, c.freq, FrequencyToHZ(c.freq)))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf(
		"Burst Pair Sequence 1[%d]:\n",
		len(c.burstPairA),
	))
	for _, bp := range c.burstPairA {
		sb.WriteString(fmt.Sprintf("\t%s\n", bp.StringWithSpaces()))
	}
	sb.WriteString(fmt.Sprintf(
		"Burst Pair Sequence 2[%d]:\n",
		len(c.burstPairB),
	))
	for _, bp := range c.burstPairB {
		sb.WriteString(fmt.Sprintf("\t%s\n", bp.StringWithSpaces()))
	}

	return sb.String()
}

func (c *HexCode) EqualDetail(e Code) (bool, string) {
	if v, msg := c.freq.EqualDetail(e.GetFrequency()); !v {
		return false, msg
	}

	if v, msg := c.burstPairA.EqualDetail(e.GetBurstPairA()); !v {
		return false, fmt.Sprintf("burst pair sequence 1: %s", msg)
	}

	if v, msg := c.burstPairB.EqualDetail(e.GetBurstPairB()); !v {
		return false, fmt.Sprintf("burst pair sequence 2: %s", msg)
	}

	return true, "OK"
}

func (c *HexCode) Equal(e Code) bool {
	v, _ := c.EqualDetail(e)
	return v
}

func (c *HexCode) GetBurstPairA() BurstPairSequence {
	return c.burstPairA
}

// func (c *HexCode) GetBurstPairAIndex(idx int) *BurstPair {
// 	if idx < len(c.burstPairA) {
// 		return c.burstPairA[idx]
// 	}

// 	return nil
// }

func (c *HexCode) AppendBurstPairA(bp *BurstPair) {
	c.burstPairA = append(c.burstPairA, bp)
}

func (c *HexCode) GetBurstPairB() BurstPairSequence {
	return c.burstPairB
}

// func (c *HexCode) GetBurstPairBIndex(idx int) *BurstPair {
// 	if idx < len(c.burstPairB) {
// 		return c.burstPairB[idx]
// 	}

// 	return nil
// }

func (c *HexCode) AppendBurstPairB(bp *BurstPair) {
	c.burstPairB = append(c.burstPairB, bp)
}

func (c *HexCode) GetFrequency() Frequency {
	return c.freq
}

func (c *HexCode) SetFrequency(freq Frequency) {
	c.freq = freq
}

func (c *HexCode) GetSource() uint16 {
	return c.source
}

func (c *HexCode) SetSource(source uint16) {
	c.source = source
}
