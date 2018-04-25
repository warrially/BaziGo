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
	GanZhi    int    // 干支0-59 对应 甲子到癸亥
	GanZhiStr string // 干支的实际字符串
	Gan       int    // 天干
	GanStr    string // 天干实际字符串
	Zhi       int    // 地支
	ZhiStr    string // 地支实际字符串
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

// 60 干支
var GAN_ZHI_STR = [60]string{
	"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
	"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
	"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
	"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
	"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
	"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"}

// 从数字获得天干地支名, 0-59
func GetGanZhiFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 60) {
		return GAN_ZHI_STR[nValue]
	}
	return "未知"
}

// {* 天干字符串，Heavenly Stems}
var TIAN_GAN_STR = [10]string{
	"甲", "乙", "丙", "丁", "戊",
	"己", "庚", "辛", "壬", "癸"}

// 从数字获得天干名, 0-9
func GetTianGanFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 10) {
		return TIAN_GAN_STR[AValue]
	}
	return ""
}

// {* 地支字符串，Earthly Branches}
var DI_ZHI_STR = [12]string{
	"子", "丑", "寅", "卯",
	"辰", "巳", "午", "未",
	"申", "酉", "戌", "亥"}

// 从数字获得地支名, 0-11
func GetDiZhiFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 12) {
		return DI_ZHI_STR[nValue]
	}
	return ""
}
