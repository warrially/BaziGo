package bazi

import "fmt"

// ToHTML 转换成网页形式
func (m *TBazi) ToHTML() string {
	strHTML := `<!DOCTYPE html>`
	strHTML += `<html>`
	strHTML += `<head>`
	strHTML += `<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">`
	strHTML += `<meta name="viewport" content="width=device-width,initial-scale=1.0">`
	strHTML += `<title>三清宫八字</title>`
	strHTML += `<style type="text/css"> .container { display: -webkit-box; display: -ms-flexbox; display: flex; width: 100%; height: 100%; -webkit-box-orient: vertical; -webkit-box-direction: normal; -ms-flex-direction: column; flex-direction: column; } .div_flex { text-align: center; -webkit-box-flex: 1; -ms-flex: 1; flex: 1; } .div_column { -webkit-box-orient: vertical; -webkit-box-direction: normal; -ms-flex-direction: column; flex-direction: column; -webkit-box-align: center; -ms-flex-align: center; align-items: center; display: -webkit-box; display: -ms-flexbox; display: flex; } .div_row { -webkit-box-align: center; -ms-flex-align: center; align-items: center; display: -webkit-box; display: -ms-flexbox; display: flex; -webkit-box-orient: horizontal; -webkit-box-direction: normal; -ms-flex-direction: row; flex-direction: row; } </style>`
	strHTML += `</head>`

	strHTML += `<body> <div>`

	strHTML += `<div class="div_row">`
	// ---------------------------------------------------------------------------------------------------------
	strHTML += `<div class="div_flex div_column">`

	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<p><font color="gray" size="2">%d年</font></p>`, m.Date().Year())
	strHTML += fmt.Sprintf(`<p><font color="gray" size="2">%s</font></p>`, m.LunarDate().Year())
	strHTML += `</div>`

	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, m.SiZhu().YearZhu().Gan().ToWuXing().Color(), m.SiZhu().YearZhu().Gan().String())
	strHTML += `</div>`
	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<font color="%s">%s</font>`, m.SiZhu().YearZhu().Gan().ToWuXing().Color(), m.SiZhu().YearZhu().ShiShen().String())
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, m.SiZhu().YearZhu().Zhi().ToWuXing().Color(), m.SiZhu().YearZhu().Zhi().String())
	strHTML += `<div class="div_column">`
	for i := 0; i < m.SiZhu().YearZhu().CangGan().Size(); i++ {
		strHTML += fmt.Sprintf(`<font color="%s">%s</font>`,
			m.SiZhu().YearZhu().CangGan().Gan(i).ToWuXing().Color(),
			m.SiZhu().YearZhu().CangGan().ShiShen(i).String())
	}
	strHTML += `</div></div></div>`
	// ---------------------------------------------------------------------------------------------------------
	strHTML += `<div class="div_flex div_column">`
	strHTML += `<div>`
	strHTML += `<p>`
	strHTML += fmt.Sprintf(`<font color="gray" size="2">%02d月</font>`, m.pSolarDate.nMonth)
	strHTML += `</p>`
	strHTML += `<p>`
	strHTML += `<font color="gray" size="2">八月</font>`
	strHTML += `</p>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, m.SiZhu().MonthZhu().Gan().ToWuXing().Color(), m.SiZhu().MonthZhu().Gan().String())
	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<font color="%s">%s</font>`, m.SiZhu().MonthZhu().Gan().ToWuXing().Color(), m.SiZhu().MonthZhu().ShiShen().String())
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, m.SiZhu().MonthZhu().Zhi().ToWuXing().Color(), m.SiZhu().MonthZhu().Zhi().String())
	strHTML += `<div class="div_column">`
	for i := 0; i < m.SiZhu().MonthZhu().CangGan().Size(); i++ {
		strHTML += fmt.Sprintf(`<font color="%s">%s</font>`,
			m.SiZhu().MonthZhu().CangGan().Gan(i).ToWuXing().Color(),
			m.SiZhu().MonthZhu().CangGan().ShiShen(i).String())
	}

	strHTML += `</div></div></div>`
	// ---------------------------------------------------------------------------------------------------------
	strHTML += `<div class="div_flex div_column">`
	strHTML += `<div>`
	strHTML += `<p>`
	strHTML += fmt.Sprintf(`<font color="gray" size="2">%d日</font>`, m.Date().Day())
	strHTML += `</p>`
	strHTML += `<p>`
	strHTML += fmt.Sprintf(`<font color="gray" size="2">%s日</font>`, m.LunarDate().Day())
	strHTML += `</p>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, m.SiZhu().DayZhu().Gan().ToWuXing().Color(), m.SiZhu().DayZhu().Gan().String())
	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<font color="%s">%s</font>`, m.SiZhu().DayZhu().Gan().ToWuXing().Color(), "主")
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, m.SiZhu().DayZhu().Zhi().ToWuXing().Color(), m.SiZhu().DayZhu().Zhi().String())
	strHTML += `<div class="div_column">`
	for i := 0; i < m.SiZhu().DayZhu().CangGan().Size(); i++ {
		strHTML += fmt.Sprintf(`<font color="%s">%s</font>`,
			m.SiZhu().DayZhu().CangGan().Gan(i).ToWuXing().Color(),
			m.SiZhu().DayZhu().CangGan().ShiShen(i).String())
	}
	strHTML += `</div></div></div>`
	// ---------------------------------------------------------------------------------------------------------
	strHTML += `<div class="div_flex div_column">`
	strHTML += `<div>`
	strHTML += `<p>`
	strHTML += fmt.Sprintf(`<font color="gray" size="2">%d时</font>`, m.Date().Hour())
	strHTML += `</p>`
	strHTML += `<p>`
	fmt.Println(m.LunarDate())
	strHTML += fmt.Sprintf(`<font color="gray" size="2">%s</font>`, m.LunarDate().Hour())
	strHTML += `</p>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, m.SiZhu().HourZhu().Gan().ToWuXing().Color(), m.SiZhu().HourZhu().Gan().String())
	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<font color="%s">%s</font>`, m.SiZhu().HourZhu().Gan().ToWuXing().Color(), m.SiZhu().HourZhu().ShiShen().String())

	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, m.SiZhu().HourZhu().Zhi().ToWuXing().Color(), m.SiZhu().HourZhu().Zhi().String())
	strHTML += `<div class="div_column">`
	for i := 0; i < m.SiZhu().HourZhu().CangGan().Size(); i++ {
		strHTML += fmt.Sprintf(`<font color="%s">%s</font>`,
			m.SiZhu().HourZhu().CangGan().Gan(i).ToWuXing().Color(),
			m.SiZhu().HourZhu().CangGan().ShiShen(i).String())
	}
	strHTML += `</div></div></div>`
	strHTML += `</div>`
	strHTML += `</div>`

	// ---------------------------------------------------------------------------------------------------------
	strHTML += `<div style="background-color: rgb(238, 238, 238); height: 5px;"></div>`

	// ---------------------------------------------------------------------------------------------------------
	strHTML += `<div style="margin: 10px;">`
	strHTML += `大运`
	strHTML += `<font color="green">庚子</font>`
	strHTML += `</div>`

	// ---------------------------------------------------------------------------------------------------------
	strHTML += `<div style="background-color: lightgray;">`
	strHTML += `<div class="div_row">`

	for i := 0; i < 5; i++ {
		strHTML += `<div class="div_column div_flex" style="background-color: white; margin: 1px;">`
		strHTML += `<div style="margin: 10px;">`
		strHTML += fmt.Sprintf(`<font color="gray" size="2">%d</font>`, m.DaYun().Age(i)+m.Date().Year())
		strHTML += `</div>`
		strHTML += `<div style="margin: 3px;">`
		strHTML += fmt.Sprintf(`<font color="%s" size="5">%s</font>`, m.DaYun().Zhu(i).Gan().ToWuXing().Color(), m.DaYun().Zhu(i).Gan().String())
		strHTML += `</div>`
		strHTML += `<div style="margin: 3px;">`
		strHTML += fmt.Sprintf(`<font color="%s" size="5">%s</font>`, m.DaYun().Zhu(i).Zhi().ToWuXing().Color(), m.DaYun().Zhu(i).Zhi().String())
		strHTML += `</div>`
		strHTML += `<div style="margin: 10px;">`
		strHTML += fmt.Sprintf(`<font color="gray" size="2">%d</font>`, m.DaYun().Age(i))
		strHTML += `</div>`
		strHTML += `</div>`
	}
	strHTML += `</div>`
	strHTML += `</div>`

	// ---------------------------------------------------------------------------------------------------------
	strHTML += `<div style="background-color: lightgray;">`
	strHTML += `<div class="div_row">`

	for i := 5; i < 10; i++ {
		strHTML += `<div class="div_column div_flex" style="background-color: white; margin: 1px;">`
		strHTML += `<div style="margin: 10px;">`
		strHTML += fmt.Sprintf(`<font color="gray" size="2">%d</font>`, m.DaYun().Age(i)+m.Date().Year())
		strHTML += `</div>`
		strHTML += `<div style="margin: 3px;">`
		strHTML += fmt.Sprintf(`<font color="%s" size="5">%s</font>`, m.DaYun().Zhu(i).Gan().ToWuXing().Color(), m.DaYun().Zhu(i).Gan().String())
		strHTML += `</div>`
		strHTML += `<div style="margin: 3px;">`
		strHTML += fmt.Sprintf(`<font color="%s" size="5">%s</font>`, m.DaYun().Zhu(i).Zhi().ToWuXing().Color(), m.DaYun().Zhu(i).Zhi().String())
		strHTML += `</div>`
		strHTML += `<div style="margin: 10px;">`
		strHTML += fmt.Sprintf(`<font color="gray" size="2">%d</font>`, m.DaYun().Age(i))
		strHTML += `</div>`
		strHTML += `</div>`
	}

	strHTML += `</div>`
	strHTML += `</div>`
	// ---------------------------------------------------------------------------------------------------------

	strHTML += `</div>`

	strHTML += `</body>`
	strHTML += `</html>`

	return strHTML
}
