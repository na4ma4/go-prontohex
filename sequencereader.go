package pronto

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type sequenceReader struct {
	rd  *bufio.Reader
	seq []byte
	spc []byte
}

func newSequenceReader(r io.Reader) *sequenceReader {
	return &sequenceReader{
		rd:  bufio.NewReader(r),
		seq: make([]byte, burstPairSequenceLength),
		spc: make([]byte, 1),
	}
}

var ErrReadSequenceFailed = errors.New("read sequence failed")

func (sr *sequenceReader) ReadSequence() (uint16, error) {
	var (
		p   int64
		n   int
		pk  []byte
		err error
	)
	n, err = sr.rd.Read(sr.seq)
	if n != 4 || err != nil {
		return 0, fmt.Errorf("%w: invalid sequence", ErrReadSequenceFailed)
	}

	for pk, err = sr.rd.Peek(1); err == nil && sr.acceptableSpaces(pk); pk, err = sr.rd.Peek(1) {
		n, err = sr.rd.Read(sr.spc)
		if n > 1 {
			return 0, fmt.Errorf("%w: sequence space read failed", ErrReadSequenceFailed)
		}
		if err != nil {
			break
		}
	}

	p, err = strconv.ParseInt(string(sr.seq), 16, 16)
	if err != nil {
		return 0, fmt.Errorf("%w: unable to parse sequence (%s)", ErrReadSequenceFailed, sr.seq)
	}

	return uint16(p), nil
}

func (sr *sequenceReader) acceptableSpaces(in []byte) bool {
	if len(in) != 1 {
		return false
	}

	switch in[0] {
	case ' ', ',': // spaces and commas are ok.
		return true
	case '\n', '\r': // carriage returns are ok.
		return true
	case '\t': // tabs are ok.
		return true
	}

	return false
}
