package Common

// 日期
type TDate struct {
	Year   int // 年
	Month  int // 月
	Day    int // 日
	Hour   int // 时
	Minute int // 分
	Second int // 秒
	JieQi  int // 节气
}

// 柱子
type TZhu struct {
}

// 四柱
type TSiZhu struct {
	YearZhu  TZhu // 年柱
	MonthZhu TZhu // 月柱
	DayZhu   TZhu // 日柱
	HourZhu  TZhu // 时柱
}

var JIE_QI_STR = [24]string{
	"立春", // 节气  Beginning of Spring   0
	"雨水", // 中气  Rain Water            1
	"惊蛰", // 节气  Waking of Insects     2
	"春分", // 中气  March Equinox         3
	"清明", // 节气  Pure Brightness       4
	"谷雨", // 中气  Grain Rain            5
	"立夏", // 节气  Beginning of Summer   6
	"小满", // 中气  Grain Full            7
	"芒种", // 节气  Grain in Ear          8
	"夏至", // 中气  Summer Solstice       9
	"小暑", // 节气  Slight Heat           10
	"大暑", // 中气  Great Heat            11
	"立秋", // 节气  Beginning of Autumn   12
	"处暑", // 中气  Limit of Heat         13
	"白露", // 节气  White Dew             14
	"秋分", // 中气  September Equinox     15
	"寒露", // 节气  Cold Dew              16
	"霜降", // 中气  Descent of Frost      17
	"立冬", // 节气  Beginning of Winter   18
	"小雪", // 中气  Slight Snow           19
	"大雪", // 节气  Great Snow            20
	"冬至", // 中气  Winter Solstice       21
	"小寒", // 节气  Slight Cold           22，这是一公历年中的第一个节气
	"大寒"} // 中气  Great Cold            23
