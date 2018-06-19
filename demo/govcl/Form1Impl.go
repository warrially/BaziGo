// 在这里写你的事件

package main

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	// "github.com/ying32/govcl/vcl/types"
	"github.com/warrially/BaziGo"
)

func (self *TForm1) OnFormCreate(sender vcl.IObject) {
	self.MonthCalendar1.SetOnClick(func(sender vcl.IObject) {
		self.calcBazi()
	})
}

func (self *TForm1) calcBazi() {
	fmt.Println("点击了", self.MonthCalendar1.Date())

	dt := self.MonthCalendar1.Date()
	tm := self.DateTimePicker1.Time()
	sex := 1
	if self.CheckBox1.Checked() {
		sex = 0
	}

	// 年月日时分秒 性别(1男0女)
	bazi := BaziGo.GetBazi(
		int(dt.Year()),
		int(dt.Month()),
		int(dt.Day()),
		int(tm.Hour()),
		int(tm.Minute()),
		int(tm.Second()),
		sex)

	// fmt.Println(bazi)

	// BaziGo.PrintBazi(bazi)

	self.Label1.SetCaption(bazi.SiZhu.YearZhu.Gan.Str)
	self.Label2.SetCaption(bazi.SiZhu.MonthZhu.Gan.Str)
	self.Label3.SetCaption(bazi.SiZhu.DayZhu.Gan.Str)
	self.Label4.SetCaption(bazi.SiZhu.HourZhu.Gan.Str)

	self.Label5.SetCaption(bazi.SiZhu.YearZhu.Zhi.Str)
	self.Label6.SetCaption(bazi.SiZhu.MonthZhu.Zhi.Str)
	self.Label7.SetCaption(bazi.SiZhu.DayZhu.Zhi.Str)
	self.Label8.SetCaption(bazi.SiZhu.HourZhu.Zhi.Str)
}
