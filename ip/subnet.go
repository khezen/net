package ip

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

// SubnetIPV4 - takes an ip v4 as string("192.168.220.254") and a subnet as int (17)
// and returns a the corresponding subnet as string ("192.168.92.0/17")
func SubnetIPV4(ipv4 string, mask int) (subnet string, err error) {
	if mask < 0 || mask > 32 {
		return "", fmt.Errorf("subnet mask out of bound:%d", mask)
	}
	// parse ip into 4 fragments
	var (
		ipFragments = make([]uint8, 0, 4)
		fragmentStr string
		fragmentInt int
		i           int
	)
	for i = 0; i < 15; i += 4 {
		fragmentStr = ipv4[i : i+3]
		fragmentInt, err = strconv.Atoi(fragmentStr)
		if err != nil {
			return "", err
		}
		ipFragments = append(ipFragments, uint8(fragmentInt))
	}
	// override fragments accordingly to subnet mask
	var (
		bitsLeftToTruncate = 32 - mask
		fragmentFilter     uint8
		fragIndex          = 3
	)
	for bitsLeftToTruncate > 0 && fragIndex >= 0 {
		if bitsLeftToTruncate >= 8 {
			fragmentFilter = 0
			bitsLeftToTruncate -= 8
		} else {
			fragmentFilter = 255 - uint8(math.Pow(2, float64(bitsLeftToTruncate)))
			bitsLeftToTruncate = 0
		}
		ipFragments[fragIndex] &= fragmentFilter
		fragIndex--
	}
	// format subnet and return result
	subnetBuf := bytes.NewBuffer([]byte{})
	for i = 0; i < 3; i++ {
		subnetBuf.WriteString(strconv.Itoa(int(ipFragments[i])))
		subnetBuf.WriteRune('.')
	}
	subnetBuf.WriteString(strconv.Itoa(int(ipFragments[3])))
	subnetBuf.WriteRune('/')
	subnetBuf.WriteString(strconv.Itoa(mask))
	return subnetBuf.String(), nil
}
