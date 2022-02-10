package i_qurey

import (
	"github.com/ltmGo/two_ip_find/untils"
	"strings"
)

const (
	ipRangeFieldCount = 3
)

type InterfaceRuleIp interface {
	LoadIpRule(text string) (uint32, uint32, []byte) //加载ip文件到内存中，提供用户，重写方法
}

type DefaultRule struct {
}

func MakeDefaultRule() DefaultRule {
	return DefaultRule{}
}

func (d DefaultRule) LoadIpRule(line string) (begin uint32, end uint32, data []byte) {
	line = strings.Trim(line, "\t")
	item := strings.SplitN(line, "\t", ipRangeFieldCount)
	if len(item) != ipRangeFieldCount {
		return
	}
	begin = untils.IpTwoLong(item[0])
	end = untils.IpTwoLong(item[1])
	data = []byte(strings.ReplaceAll(item[2], "\t", "_"))
	return
}
