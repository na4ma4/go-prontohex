package pronto

type BurstPairSequence []*BurstPair

func (bps *BurstPairSequence) Append(bp *BurstPair) {
	*bps = append(*bps, bp)
}

func (bps BurstPairSequence) EqualDetail(e BurstPairSequence) (bool, string) {
	if len(bps) != len(e) {
		return false, "different length sequences"
	}

	for idx := 0; idx < len(bps); idx++ {
		if v, msg := bps[idx].EqualDetail(e[idx]); !v {
			return false, msg
		}
	}

	return true, "OK"
}

func (bps BurstPairSequence) Equal(e BurstPairSequence) bool {
	v, _ := bps.EqualDetail(e)
	return v
}

func (bps BurstPairSequence) Get(idx int) *BurstPair {
	if idx < len(bps) {
		return bps[idx]
	}

	return nil
}
