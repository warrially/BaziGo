package bazi

import (
	"fmt"

	"github.com/warrially/BaziGo/internal/jieqidata"
)

// TJieQiDate 节气。
type TJieQiDate struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
	JieQi  TJieQi
}

// ToSolarDate 节气日期转换成普通日期。
func (m *TJieQiDate) ToSolarDate() *TSolarDate {
	if m == nil {
		return nil
	}
	return NewSolarDate(m.Year, m.Month, m.Day, m.Hour, m.Minute, m.Second)
}

// GetJieQiDate 获取某个日期前后相邻的两个“节”。
func GetJieQiDate(pSolarDate *TSolarDate) (*TJieQiDate, *TJieQiDate) {
	if pSolarDate == nil {
		return nil, nil
	}

	nMinYear, nMaxYear, err := jieqidata.YearRange()
	if err != nil || pSolarDate.Year() <= nMinYear || pSolarDate.Year() >= nMaxYear {
		return nil, nil
	}

	// 数据从 nMinYear 年开始，每年固定包含 24 个节气。
	nIndex := (pSolarDate.Year() - nMinYear) * 24
	nSeconds := pSolarDate.Get64TimeStamp()
	for i := nIndex - 1; i <= nIndex+25; i++ {
		nJieQiSeconds, nJieQi, err := jieqidata.Record(i)
		if err != nil {
			return nil, nil
		}
		stJieQi := TJieQi(nJieQi)
		if stJieQi.IsJie() && nSeconds < nJieQiSeconds {
			return newJieQiDateFromRecord(i - 2), newJieQiDateFromRecord(i)
		}
	}
	return nil, nil
}

func newJieQiDateFromRecord(nIndex int) *TJieQiDate {
	nSeconds, nJieQi, err := jieqidata.Record(nIndex)
	if err != nil {
		return nil
	}
	pSolarDate := NewSolarDateFrom64TimeStamp(nSeconds)
	if pSolarDate == nil {
		return nil
	}
	return &TJieQiDate{
		Year:   pSolarDate.Year(),
		Month:  pSolarDate.Month(),
		Day:    pSolarDate.Day(),
		Hour:   pSolarDate.Hour(),
		Minute: pSolarDate.Minute(),
		Second: pSolarDate.Second(),
		JieQi:  TJieQi(nJieQi),
	}
}

func (m *TJieQiDate) String() string {
	return fmt.Sprintf("节气:%d年%02d月%02d日 %02d:%02d:%02d (%v)",
		m.Year, m.Month, m.Day, m.Hour, m.Minute, m.Second, m.JieQi.String())
}
