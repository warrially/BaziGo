package JieQi

import (
	. "github.com/warrially/BaziGo/Common"
	_ "github.com/warrially/BaziGo/Days"
)

type TJieQi struct {
	DateList []TDate
}

var m_MapJieQi map[int]TJieQi

func init() {
	m_MapJieQi = make(map[int]TJieQi)

	for i := range JIE_QI_LIST {
		// 将数据转存成TDate类型
		var date TDate
		date.Year = JIE_QI_LIST[i][0]
		date.Month = JIE_QI_LIST[i][1]
		date.Day = JIE_QI_LIST[i][2]
		date.Hour = JIE_QI_LIST[i][3]
		date.Minute = JIE_QI_LIST[i][4]
		date.Second = JIE_QI_LIST[i][5]
		date.JieQi = JIE_QI_LIST[i][6]

		pMap := m_MapJieQi[date.Year]
		pMap.DateList = append(pMap.DateList, date)
		m_MapJieQi[date.Year] = pMap
	}
}

// // 获取某个日期的节气, 并且前后两个节气
func GetJieQi(date TDate) (TDate, TDate) {

	if (date.Year < 31) || (date.Year > 2300) {
		return date, date
	}
	// 一口气获取3年的所有节气
	pMap1, _ := m_MapJieQi[date.Year-1]
	pMap2, _ := m_MapJieQi[date.Year]
	pMap3, _ := m_MapJieQi[date.Year+1]

	// 上一个日期
	var pLastDate TDate

	for i := range pMap1.DateList {
		// 必须是节, 中气是不行
		if pMap1.DateList[i].JieQi%2 == 0 {
			// 找到最后的那个节气
			if CompareDate(date, pMap1.DateList[i]) <= 1 {
				pLastDate = pMap1.DateList[i]
			} else {
				return pLastDate, pMap1.DateList[i]
			}
		}
	}
	for i := range pMap2.DateList {
		// 必须是节, 中气是不行
		if pMap2.DateList[i].JieQi%2 == 0 {
			// 找到最后的那个节气
			if CompareDate(date, pMap2.DateList[i]) <= 1 {
				pLastDate = pMap2.DateList[i]
			} else {
				return pLastDate, pMap2.DateList[i]
			}
		}
	}
	for i := range pMap3.DateList {
		// 必须是节, 中气是不行
		if pMap3.DateList[i].JieQi%2 == 0 {
			// 找到最后的那个节气
			if CompareDate(date, pMap3.DateList[i]) <= 1 {
				pLastDate = pMap3.DateList[i]
			} else {
				return pLastDate, pMap3.DateList[i]
			}
		}
	}

	// 错误
	return date, date
}

// 日期比较, 返回1  1大, 返回2 2大  返回0 相等
func CompareDate(date1 TDate, date2 TDate) int {
	if date1.Year > date2.Year {
		return 1
	}
	if date1.Year < date2.Year {
		return 2
	}

	if date1.Month > date2.Month {
		return 1
	}
	if date1.Month < date2.Month {
		return 2
	}

	if date1.Day > date2.Day {
		return 1
	}
	if date1.Day < date2.Day {
		return 2
	}

	if date1.Hour > date2.Hour {
		return 1
	}
	if date1.Hour < date2.Hour {
		return 2
	}

	if date1.Minute > date2.Minute {
		return 1
	}
	if date1.Minute < date2.Minute {
		return 2
	}

	if date1.Second > date2.Second {
		return 1
	}
	if date1.Second < date2.Second {
		return 2
	}

	return 0
}
