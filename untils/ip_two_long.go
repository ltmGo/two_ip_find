package untils

import (
	"bytes"
	"encoding/binary"
	"net"
)

func IpTwoLong(ip string) uint32 {
	var long uint32
	err := binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	if err != nil {
		return 0
	}
	return long
}
