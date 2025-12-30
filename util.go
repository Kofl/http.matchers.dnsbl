// SPDX-FileCopyrightText: 2024 Gergely Nagy
// SPDX-FileContributor: Gergely Nagy
//
// SPDX-License-Identifier: EUPL-1.2

package dnsbl

import (
	"net"
	"strconv"
)

func reverseaddr(addr string) (arpa string, err error) {
	ip := net.ParseIP(addr)
	if ip == nil {
		return "", &net.DNSError{Err: "unrecognized address", Name: addr}
	}
	if ip.To4() != nil {
		return strconv.FormatUint(uint64(ip[15]), 10) + "." +
			strconv.FormatUint(uint64(ip[14]), 10) + "." +
			strconv.FormatUint(uint64(ip[13]), 10) + "." +
			strconv.FormatUint(uint64(ip[12]), 10), nil
	}
	buf := make([]byte, 0, len(ip)*4)
	for i := len(ip) - 1; i >= 0; i-- {
		v := ip[i]

		buf = append(buf, strconv.FormatUint(uint64(v&0xF), 16)[0])
		buf = append(buf, '.')
		buf = append(buf, strconv.FormatUint(uint64(v>>4), 16)[0])
		buf = append(buf, '.')
	}
	return string(buf[:len(buf)-1]), nil
}
