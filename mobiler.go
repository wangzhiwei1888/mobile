// mobiler 包用于进行一些手机相关的操作
//
// 这些操作有：
//
// 1. 验证一个字符串是否为手机号码，有严格模式与宽松模式两种，在严格模型下，字符串必须
// 与一个正常的手机号码完全一致，而在宽松模式下，手机号码可以存在四位打码的情况出现，即
// 允许手机号码为 139****0000。
//
// 2. 根据手机号码字符串取得该号码所对应的运营商、卡类型、所属的省市等基本号码信息
package mobiler

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// Mobiler 用于根据手机号码生成带地区号段的完整手机数据
type Mobiler struct {
	nis map[string]NetworkIdentifier
}

// LoadSegments 方法用于加载手机号码段数据
func (self *Mobiler) LoadSegments(path string) (err error) {
	log.Printf("加载号码段数据文件：%s。\n", path)
	segmentsFile, err := os.Open(path)
	defer segmentsFile.Close()
	if err != nil {
		log.Fatalf("无法载入号码段数据文件 \"%s\" \n", segmentsFile)
	}
	reader := bufio.NewReader(segmentsFile)
	self.nis = make(map[string]NetworkIdentifier)
	// 逐行读入记录
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		line = strings.TrimSpace(line)

		fields := strings.Split(line, ",")

		if len(fields) == 8 {
			identifier, ok := self.nis[fields[0]]
			if !ok {
				identifier = NetworkIdentifier{
					Code: fields[0],
					ISP:  fields[2],
					Type: fields[3],
				}
				identifier.Segments = make(map[string]Segment)
			}
			identifier.Segments[fields[1]] = Segment{
				Prefix:   fields[1],
				Province: fields[4],
				City:     fields[5],
				ZipCode:  fields[6],
				AreaCode: fields[7],
			}

			self.nis[fields[0]] = identifier
		}
	}
	return nil
}

// Mobile 方法根据一个手机号码字符串生成一个完成的手机格式
func (self *Mobiler) Mobile(mob string) (mobile Mobile) {
	if !IsMobile(mob, false) {
		return
	}
	mobile.Number = mob
	// 获取
	identifier, ok := self.nis[mob[:3]]
	if !ok {
		return
	}
	mobile.ISP = identifier.ISP
	mobile.Type = identifier.Type

	segment, ok := identifier.Segments[mob[:7]]
	if !ok {
		return
	}
	mobile.AreaCode = segment.AreaCode
	mobile.Province = segment.Province
	mobile.City = segment.City
	mobile.ZipCode = segment.ZipCode
	return
}
