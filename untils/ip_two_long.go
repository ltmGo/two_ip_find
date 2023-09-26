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

// LongToIp 转成 ip
func LongToIp(long uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, long)
	return ip.String()
}
