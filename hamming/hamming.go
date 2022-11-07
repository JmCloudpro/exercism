package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	dnsa := strings.Split(a, "")
	dnsb := strings.Split(b, "")

	dnsalen := len(dnsa)
	dnsblen := len(dnsb)

	if dnsalen != dnsblen {

		err := errors.New("Different Size")
		return 0, err
	}

	dst := 0
	for i := 0; i < dnsalen; i++ {
		if dnsa[i] != dnsb[i] {
			dst++
		}

	}

	return dst, nil
}
