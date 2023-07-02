package pronto

// FrequencyToHZ converts the prontohex format frequency into cycles per second.
//
//nolint:gomnd // constants from prontohex format.
func FrequencyToHZ(freq Frequency) uint32 {
	return uint32(1000000 / (float64(freq) * 0.241246))
}

const (
	// sonyButtonCodeLength is the number of bits used for the remote button code.
	sonyButtonCodeLength = 7
)

//nolint:gochecknoglobals,gomnd // constant for decoding Sony device and remote codes.
var sonyBPTrue = &BurstPair{
	a: 0x0030,
	b: 0x0018,
}

func SonyDeviceCode(c Code) int32 {
	bps := c.GetBurstPairB()
	if len(bps) < sonyButtonCodeLength+1 {
		return -1
	}

	dev := int32(0)

	for idx := 8; idx < len(bps); idx++ {
		if bps[idx].Equal(sonyBPTrue) {
			dev += 1 << (idx - sonyButtonCodeLength - 1)
		}
	}

	return dev
}

func SonyRemoteButton(c Code) int16 {
	bps := c.GetBurstPairB()
	if len(bps) < sonyButtonCodeLength+1 {
		return -1
	}

	btn := int16(0)

	for idx := 0; idx <= sonyButtonCodeLength; idx++ {
		if bps[idx+1].Equal(sonyBPTrue) {
			btn += 1 << idx
		}
	}

	return btn
}
