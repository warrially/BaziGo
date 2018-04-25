package JieQi

import (
	. "github.com/warrially/BaziGo/Common"
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
	// log.Println(m_MapJieQi)
}

// // 获取某个日期的节气, 并且返回月份
func GetJieQi(date TDate) int {

	if (date.Year < 31) || (date.Year > 2300) {
		return date.Month - 1
	}

	// 1 根据年份 和 月份来获取两个大概可能的日期
	pMap, _ := m_MapJieQi[date.Year]
	var Result int

	for i := range pMap.DateList {
		if i == 0 {
			// 初始值
			Result = pMap.DateList[i].JieQi - 1
			if Result < 0 {
				// 防止负数, 基本上不会出现
				Result += 24
			}
		}
		// 找到最后的那个节气
		if CompareDate(date, pMap.DateList[i]) <= 1 {
			Result = pMap.DateList[i].JieQi
		} else {
			return Result
		}
	}

	return Result
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
