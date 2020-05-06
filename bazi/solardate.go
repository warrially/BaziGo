package bazi

import (
	"log"

	"github.com/warrially/BaziGo/Days"
)

// NewSolarDate 创建一个新历时间
func NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond int) *TSolarDate {

	// 把具体时间实例化出来
	pDate := &TSolarDate{
		Year:   nYear,   // 年
		Month:  nMonth,  // 月
		Day:    nDay,    // 日
		Hour:   nHour,   // 时
		Minute: nMinute, // 分
		Second: nSecond, // 秒
	}

	if !pDate.GetDateIsValid() {
		log.Println("无效的日期", nYear, nMonth, nDay)
		return nil
	}

	// 计算64位时间戳值
	pDate.Gen64TimeStamp()

	return pDate
}

// NewSolarDateFrom64TimeStamp 从64位时间戳反推日期
func NewSolarDateFrom64TimeStamp(nTimeStamp int64) *TSolarDate {
	pDate := &TSolarDate{}
	pDate.TimeStamp = nTimeStamp

	// 计算出年份
	pDate.GetYearFrom64TimeStamp()
	// 计算月份
	pDate.GetMonthFrom64TimeStamp()
	// 计算其他参数
	pDate.GetDayTimeFrom64TimeStamp()

	return pDate
}

// TSolarDate 日期
type TSolarDate struct {
	Year      int   // 年
	Month     int   // 月
	Day       int   // 日
	Hour      int   // 时
	Minute    int   // 分
	Second    int   // 秒
	TimeStamp int64 // 时间戳
}

// Gen64TimeStamp 生成64位时间戳
func (self *TSolarDate) Gen64TimeStamp() int64 {
	nAllDays := Days.GetAllDays(self.Year, self.Month, self.Day) // 先获取公元原点的日数
	nResult := int64(nAllDays)
	nResult *= 24 * 60 * 60 // 天数换成秒

	//再计算出秒数
	nResult += int64(self.Hour) * 60 * 60
	nResult += int64(self.Minute) * 60
	nResult += int64(self.Second)

	self.TimeStamp = nResult
	return nResult
}

// GetYearFrom64TimeStamp 从64位时间戳反推年
func (self *TSolarDate) GetYearFrom64TimeStamp() *TSolarDate {
	// 准备进行二分法
	nLow := 0
	nHigh := 3001

	for {
		nMid := (nLow + nHigh) / 2

		// 拿到中间年的数据
		v := NewSolarDate(nMid, 1, 1, 0, 0, 0).Gen64TimeStamp()

		if v <= self.TimeStamp {
			nLow = nMid
		} else {
			nHigh = nMid
		}

		if nHigh == nLow+1 {
			break
		}
	}
	self.Year = nLow
	return self
}

// GetMonthFrom64TimeStamp 从64位时间戳反推月,
func (self *TSolarDate) GetMonthFrom64TimeStamp() {
	// 这里开始特殊处理
	for i := 1; i <= 11; i++ {
		if self.TimeStamp < NewSolarDate(self.Year, i+1, 1, 0, 0, 0).Gen64TimeStamp() {
			self.Month = i
			return
		}
	}
	self.Month = 12
}

// GetDayTimeFrom64TimeStamp 从64位时间戳反推其他参数
func (self *TSolarDate) GetDayTimeFrom64TimeStamp() {
	nTimeStamp := self.TimeStamp - NewSolarDate(self.Year, self.Month, 1, 0, 0, 0).Gen64TimeStamp()

	// 计算日
	self.Day = int(nTimeStamp / (24 * 60 * 60))
	// 扣掉日
	nTimeStamp -= int64(self.Day) * 24 * 60 * 60

	self.Day++ // 因为每个月的天数是从1开始的, 所以这里需要补1天
	if self.Year == 1582 && self.Month == 10 && self.Day >= 5 {
		self.Day += 10 // 1582 年需要补10天
	}
	self.Hour = int(nTimeStamp / (60 * 60))
	nTimeStamp -= int64(self.Hour) * 60 * 60
	self.Minute = int(nTimeStamp / 60)
	nTimeStamp -= int64(self.Minute) * 60
	self.Second = int(nTimeStamp)
}

// GetMonthDays 取本月天数，不考虑 1582 年 10 月的特殊情况
func (self *TSolarDate) GetMonthDays(nYear, nMonth int) int {
	switch nMonth {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2: // 闰年
		if self.GetIsLeapYear(nYear) {
			return 29
		} else {
			return 28
		}
	}
	return 0
}

// GetIsLeapYear 返回某公历是否闰年
func (self *TSolarDate) GetIsLeapYear(nYear int) bool {
	if self.GetCalendarType(nYear, 1, 1) == ctGregorian {
		return (nYear%4 == 0) && ((nYear%100 != 0) || (nYear%400 == 0))
	} else if nYear >= 0 {
		return nYear%4 == 0
	} else { // 需要独立判断公元前的原因是没有公元 0 年
		return (nYear-3)%4 == 0
	}
}

const (
	ctInvalid   = iota //非法，
	ctJulian           //儒略，
	ctGregorian        //格利高里
)

// 根据公历日期判断当时历法
func (self *TSolarDate) GetCalendarType(nYear, nMonth, nDay int) int {
	if !self.GetDateIsValid(nYear, nMonth, nDay) {
		return ctInvalid
	}
	if nYear > 1582 {
		return ctGregorian
	} else if nYear < 1582 {
		return ctJulian
	} else if nMonth < 10 {
		return ctJulian
	} else if (nMonth == 10) && (nDay <= 4) {
		return ctJulian
	} else if (nMonth == 10) && (nDay <= 14) {
		return ctInvalid
	} else {
		return ctGregorian
	}
	// 在现在通行的历法记载上，全世界居然有十天没有任何人出生过，也没有任何人死亡过，也没有发生过大大小小值得纪念的人或事。这就是1582年10月5日至10月14日。格里奥，提出了公历历法。这个历法被罗马教皇格里高利十三世采纳了。那么误差的十天怎么办？罗马教皇格里高利十三世下令，把1582年10月4日的后一天改为10月15日，这样误差的十天没有了，历史上也就无影无踪地消失了十天，当然史书上也就没有这十天的记载了。“格里高利公历”一直沿用到今天。
}

// GetDateIsValid 返回公历日期是否合法
func (self *TSolarDate) GetDateIsValid(nYear, nMonth, nDay int) bool {
	// 没有公元0年
	if nYear == 0 {
		return false
	}

	// 1月开始, 12月结束
	if nMonth < 1 || nMonth > 12 {
		return false
	}

	// 1号开始, 获取每个月有多少天结束
	if nDay < 1 || nDay > self.GetMonthDays(nYear, nMonth) {
		return false
	}

	// 1582 年的特殊情况
	if nYear != 1582 {
		return true
	}
	if nMonth != 10 {
		return true
	}
	//
	if nDay < 5 || nDay > 14 {
		return true
	}

	return false
}
