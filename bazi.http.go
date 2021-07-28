package bazi

import "fmt"

// ToHTML 转换成网页形式
func (self *TBazi) ToHTML() string {
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
	strHTML += `<div class="div_flex div_column">`
	strHTML += `<div>`
	strHTML += `<p>`
	strHTML += fmt.Sprintf(`<font color="gray" size="2">%d年</font>`, self.pSolarDate.nYear)
	strHTML += `</p>`
	strHTML += `<p>`
	strHTML += `<font color="gray" size="2">一九八六</font>`
	strHTML += `</p>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += `<div>`

	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, self.SiZhu().YearZhu().Gan().ToWuXing().Color(), self.SiZhu().YearZhu().Gan().String())
	strHTML += `</div>`
	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<font color="%s">%s</font>`, self.SiZhu().YearZhu().Gan().ToWuXing().Color(), self.SiZhu().YearZhu().ShiShen().String())
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, self.SiZhu().YearZhu().Zhi().ToWuXing().Color(), self.SiZhu().YearZhu().Zhi().String())
	strHTML += `<div class="div_column">`
	for i := 0; i < self.SiZhu().YearZhu().CangGan().Size(); i++ {
		strHTML += fmt.Sprintf(`<font color="%s">%s</font>`,
			self.SiZhu().YearZhu().CangGan().Gan(i).ToWuXing().Color(),
			self.SiZhu().YearZhu().CangGan().ShiShen(i).String())
	}
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_flex div_column">`
	strHTML += `<div>`
	strHTML += `<p>`
	strHTML += `<font color="gray" size="2">09月</font>`
	strHTML += `</p>`
	strHTML += `<p>`
	strHTML += `<font color="gray" size="2">八月</font>`
	strHTML += `</p>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, self.SiZhu().MonthZhu().Gan().ToWuXing().Color(), self.SiZhu().MonthZhu().Gan().String())
	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<font color="%s">%s</font>`, self.SiZhu().MonthZhu().Gan().ToWuXing().Color(), self.SiZhu().MonthZhu().ShiShen().String())
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, self.SiZhu().MonthZhu().Zhi().ToWuXing().Color(), self.SiZhu().MonthZhu().Zhi().String())
	strHTML += `<div class="div_column">`
	for i := 0; i < self.SiZhu().MonthZhu().CangGan().Size(); i++ {
		strHTML += fmt.Sprintf(`<font color="%s">%s</font>`,
			self.SiZhu().MonthZhu().CangGan().Gan(i).ToWuXing().Color(),
			self.SiZhu().MonthZhu().CangGan().ShiShen(i).String())
	}

	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_flex div_column">`
	strHTML += `<div>`
	strHTML += `<p>`
	strHTML += `<font color="gray" size="2">22日</font>`
	strHTML += `</p>`
	strHTML += `<p>`
	strHTML += `<font color="gray" size="2">十九日</font>`
	strHTML += `</p>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, self.SiZhu().DayZhu().Gan().ToWuXing().Color(), self.SiZhu().DayZhu().Gan().String())
	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<font color="%s">%s</font>`, self.SiZhu().DayZhu().Gan().ToWuXing().Color(), self.SiZhu().DayZhu().ShiShen().String())
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, self.SiZhu().DayZhu().Zhi().ToWuXing().Color(), self.SiZhu().DayZhu().Zhi().String())
	strHTML += `<div class="div_column">`
	for i := 0; i < self.SiZhu().DayZhu().CangGan().Size(); i++ {
		strHTML += fmt.Sprintf(`<font color="%s">%s</font>`,
			self.SiZhu().DayZhu().CangGan().Gan(i).ToWuXing().Color(),
			self.SiZhu().DayZhu().CangGan().ShiShen(i).String())
	}
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_flex div_column">`
	strHTML += `<div>`
	strHTML += `<p>`
	strHTML += `<font color="gray" size="2">12时</font>`
	strHTML += `</p>`
	strHTML += `<p>`
	strHTML += `<font color="gray" size="2">午时</font>`
	strHTML += `</p>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, self.SiZhu().HourZhu().Gan().ToWuXing().Color(), self.SiZhu().HourZhu().Gan().String())
	strHTML += `<div>`
	strHTML += fmt.Sprintf(`<font color="%s">%s</font>`, self.SiZhu().HourZhu().Gan().ToWuXing().Color(), self.SiZhu().HourZhu().ShiShen().String())

	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `<div class="div_row" style="padding: 10px;">`
	strHTML += fmt.Sprintf(`<font color="%s" size="8">%s</font>`, self.SiZhu().HourZhu().Zhi().ToWuXing().Color(), self.SiZhu().HourZhu().Zhi().String())
	strHTML += `<div class="div_column">`
	for i := 0; i < self.SiZhu().HourZhu().CangGan().Size(); i++ {
		strHTML += fmt.Sprintf(`<font color="%s">%s</font>`,
			self.SiZhu().HourZhu().CangGan().Gan(i).ToWuXing().Color(),
			self.SiZhu().HourZhu().CangGan().ShiShen(i).String())
	}
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `</div>`
	strHTML += `</div>`

	strHTML += `</div> </body>`

	strHTML += `</html>`

	return strHTML
}
