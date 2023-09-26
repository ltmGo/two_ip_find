package i_qurey

import (
	"github.com/ltmGo/two_ip_find/ip_range"
	"github.com/ltmGo/two_ip_find/untils"
	"strings"
)

const (
	ipRangeFieldCount = 3
)

type InterfaceRuleIp interface {
	LoadIpRule(text string) []*ip_range.IpRange //加载ip文件到内存中，提供用户，重写方法
}

type DefaultRule struct {
}

func MakeDefaultRule() DefaultRule {
	return DefaultRule{}
}

func (d DefaultRule) LoadIpRule(line string) (list []*ip_range.IpRange) {
	line = strings.Trim(line, "\t")
	item := strings.SplitN(line, "\t", ipRangeFieldCount)
	list = make([]*ip_range.IpRange, 0)
	if len(item) != ipRangeFieldCount {
		return
	}
	r := &ip_range.IpRange{}
	r.Begin = untils.IpTwoLong(item[0])
	r.End = untils.IpTwoLong(item[1])
	r.Data = []byte(strings.ReplaceAll(item[2], "\t", "_"))
	return
}
