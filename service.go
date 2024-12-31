package two_ip_find

import (
	"bufio"
	"errors"
	"github.com/ltmGo/two_ip_find/i_qurey"
	"github.com/ltmGo/two_ip_find/ip_range"
	"github.com/ltmGo/two_ip_find/untils"
	"io"
	"os"
	"sort"
)

var ErrorIpRangeNotFound = errors.New("ip range not found")

// IpService 具体的实现方法
type IpService struct {
	ipList []*ip_range.IpRange
}

func MakeIpService() *IpService {
	return &IpService{
		ipList: make([]*ip_range.IpRange, 0, 10000),
	}
}

func (i *IpService) openIpFile(filePath string) (error, io.Reader) {
	reader, err := os.Open(filePath)
	return err, reader
}

// LoadFileToIp 加载ip到内存
func (i *IpService) LoadFileToIp(r i_qurey.InterfaceRuleIp, filePath ...string) error {
	if len(filePath) == 0 {
		return errors.New("file path is empty")
	}
	for _, k := range filePath {
		err, f := i.openIpFile(k)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			ir := r.LoadIpRule(line)
			for _, j := range ir {
				if j.Begin > j.End || j.Begin == 0 || j.End == 0 {
					continue
				} else {
					i.ipList = append(i.ipList, j)
				}
			}
		}
	}
	//进行排序
	sort.Slice(i.ipList, func(k, j int) bool {
		return i.ipList[k].Begin < i.ipList[j].Begin
	})
	return nil
}

// ReLoadFileToIp 重新加载到内存
func (i *IpService) ReLoadFileToIp(r i_qurey.InterfaceRuleIp, filePath ...string) error {
	i.ipList = make([]*ip_range.IpRange, 0, 10000)
	err := i.LoadFileToIp(r, filePath...)
	return err
}

// FindIp 查找ip
func (i *IpService) FindIp(ip string) (*ip_range.IpRange, error) {
	ir, err := i.getIpRange(ip)
	if err != nil {
		return nil, err
	}
	return ir, nil
}

func (i *IpService) length() int {
	return len(i.ipList)
}

func (i *IpService) getIpRange(ip string) (*ip_range.IpRange, error) {
	var low, high = 0, i.length() - 1
	ipDt := i.ipList
	il := untils.IpTwoLong(ip)
	if il <= 0 {
		return nil, ErrorIpRangeNotFound
	}
	for low <= high {
		middle := (high-low)/2 + low
		ir := ipDt[middle]
		if il >= ir.Begin && il <= ir.End {
			return ir, nil
		} else if il < ir.Begin {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return nil, ErrorIpRangeNotFound
}
