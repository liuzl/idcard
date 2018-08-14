package idcard

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/liuzl/goutil"
)

var (
	startDay, _ = time.Parse("20060102", "18491001")
	idcardLen   = 18
	ratioVal    = 11

	//身份证号码前17位数分别乘以不同的系数
	ratio = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	//将这17位数字和系数相乘的结果相加,对11进行求余，得出身份证最后一个字符
	matchMap = map[int]rune{
		0:  '1',
		1:  '0',
		2:  'X',
		3:  '9',
		4:  '8',
		5:  '7',
		6:  '6',
		7:  '5',
		8:  '4',
		9:  '3',
		10: '2'}
)

type IDCardInfo struct {
	ID       string    `json:"id"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	Province string    `json:"province"`
	City     string    `json:"city"`
	County   string    `json:"county"`
}

func Parse(id string) (*IDCardInfo, error) {
	id = strings.ToUpper(id)
	last := []rune(id)[17]
	if len(id) != idcardLen || !goutil.StringIs(id[:17], unicode.IsDigit) ||
		(!unicode.IsDigit(last) && last != 'X') {
		return nil, fmt.Errorf("format not correct")
	}

	t, err := time.Parse("20060102", id[6:14])
	if err != nil || t.After(time.Now()) || t.Before(startDay) {
		return nil, fmt.Errorf("birthday %s not correct", id[6:14])
	}

	total := 0
	for i := 0; i < 17; i++ {
		total += ratio[i] * int(id[i]-'0')
	}
	if matchMap[total%ratioVal] != last {
		return nil, fmt.Errorf("checksum not correct")
	}

	gender := "male"
	if int(id[16]-'0')%2 == 0 {
		gender = "female"
	}

	county := CodeMap[id[:6]]
	if county == "" {
		return nil, fmt.Errorf("areacode not correct")
	}

	return &IDCardInfo{id, t, gender, CodeMap[id[:2]+"0000"], CodeMap[id[:4]+"00"], county}, nil
}

func VerifyBirthday(day string) bool {
	if len(day) != 8 || !goutil.StringIs(day, unicode.IsDigit) {
		return false
	}
	t, err := time.Parse("20060102", day)
	if err != nil || t.After(time.Now()) || t.Before(startDay) {
		return false
	}
	return true
}

func Verify(id string) bool {
	if len(id) != idcardLen || !goutil.StringIs(id[:17], unicode.IsDigit) {
		return false
	}
	id = strings.ToUpper(id)
	last := []rune(id)[17]
	if !unicode.IsDigit(last) && last != 'X' {
		return false
	}
	if !VerifyBirthday(id[6:14]) {
		return false
	}
	total := 0
	for i, c := range id {
		if i == 17 {
			continue
		}
		total += ratio[i] * int(c-'0')
	}
	if matchMap[total%ratioVal] != last {
		return false
	}
	return true
}
