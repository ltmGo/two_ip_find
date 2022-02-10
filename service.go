package two_ip_find

import (
	"bufio"
	"errors"
	"github.com/ltmGo/two_ip_find/i_qurey"
	"github.com/ltmGo/two_ip_find/untils"
	"io"
	"os"
	"sort"
)

var ErrorIpRangeNotFound = errors.New("ip range not found")

type IpRange struct {
	Begin uint32
	End   uint32
	Data  []byte
}

//IpService 具体的实现方法
type IpService struct {
	ipList []*IpRange
}

func MakeIpService() *IpService {
	return &IpService{
		ipList: make([]*IpRange, 0, 10000),
	}
}

func (i *IpService) openIpFile(filePath string) (error, io.Reader) {
	reader, err := os.Open(filePath)
	return err, reader
}

//LoadFileToIp 加载ip到内存
func (i *IpService) LoadFileToIp (r i_qurey.InterfaceRuleIp, filePath string) error {
	err, f := i.openIpFile(filePath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		begin, end, dataByte := r.LoadIpRule(line)
		if begin > end {
			continue
		}
		ir := &IpRange{
			Begin: begin,
			End:   end,
			Data:  dataByte,
		}
		i.ipList = append(i.ipList, ir)
	}
	//进行排序
	sort.Slice(i.ipList, func(k, j int) bool {
		return i.ipList[k].Begin < i.ipList[j].Begin
	})
	return nil
}

//ReLoadFileToIp 重新加载到内存
func (i *IpService) ReLoadFileToIp (r i_qurey.InterfaceRuleIp, filePath string) error {
	i.ipList = make([]*IpRange, 0, 10000)
	err := i.LoadFileToIp(r, filePath)
	return err
}

//FindIp 查找ip
func (i *IpService) FindIp(ip string) (*IpRange, error) {
	ir, err := i.getIpRange(ip)
	if err != nil {
		return nil, err
	}
	return ir, nil
}

func (i *IpService) length() int {
	return len(i.ipList)
}

func (i *IpService) getIpRange(ip string) (*IpRange, error) {
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