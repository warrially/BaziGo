# BaziGo
GO语言的八字库

## 使用范围

- 当前八字和农历数据支持公历 1800 年至 2299 年。
- `nSex=1` 表示男性，`nSex=0` 表示女性。
- 当前版本按输入的本地日期时间计算，不会自动进行时区、夏令时或真太阳时换算。调用方应先把出生时间转换到业务约定的时区。
- 旧接口 `GetBazi` 为兼容历史代码而保留，参数错误时返回 `nil`。新代码建议使用 `GetBaziE` 获取明确错误信息。

```go
pBazi, err := bazi.GetBaziE(1995, 6, 16, 19, 7, 0, bazi.SEX_FEMALE)
if err != nil {
	return err
}
fmt.Println(pBazi)
```

## 节气数据

节气运行数据位于 `internal/jieqidata/jieqi.bin`，通过 `go:embed` 编译进程序。引用本库的项目不需要复制数据文件，也不需要在运行时读取外部目录。

二进制记录保存的是从公元起点累计的秒数，不使用 Unix 时间戳。可维护的原始数据压缩在 `tools/jieqigen/jieqi.csv.gz`，修改原始数据后运行：

```bash
go generate ./...
go test ./...
```

生成工具会重新创建嵌入式二进制文件；测试会逐条核对全部节气记录。

作者只想保留版权，无任何使用或者发布限制。您只需要在您的发行版本中注明代码出处
https://github.com/warrially/BaziGo






八字部分参考的是三清宫命理
https://weibo.com/bazishequ

日历部分参考中国日历类
中国日历类（Chinese Calendar Class (CCC)）
版本：v0.1，JavaScript版本
版权所有 (C) 2002-2003 neweroica (wy25@mail.bnu.edu.cn)
联系方式： Email:  wy25@mail.bnu.edu.cn
QQ: 32460746

日历部分参考CNPACK
作者：刘啸 (liuxiao@cnpack.org)
周劲羽(zjy@cnpack.org)



出生日期新历：  1995 年 6 月 16 日   19 : 7 : 0
基本八字： 乙亥 壬午 戊寅 壬戌

#### 命盘解析 Tables

| 年 | 月 |日 |时 |
| :--------:   | :------:  | :---------:  |  :---------:  |
|**乙**(木)[官]|**壬**(水)[才]|**戊**(土)[主]|**壬**(水)[才]
|**亥**(水)[才杀]|**午**(火)[印劫]|**寅**(木)[杀卩比]|**戌**(土)[比伤印]
|山头火|               杨柳木|                  城墙土|                  大海水

所属节令：
芒种 1995 年 6 月 6 日   11 : 42 : 28
小暑 1995 年 7 月 7 日   22 : 1 : 0
大运： 癸未 甲申 乙酉 丙戌 丁亥 戊子 己丑 庚寅 辛卯 壬辰
起运时间 2002 年 5 月 25 日   7 : 7 : 0

```
package main

import (
	"flag"
	"fmt"

	bazi "github.com/warrially/BaziGo"
)

func main() {
	var nYear int
	var nMonth int
	var nDay int
	var nHour int
	var nMinute int
	var nSecond int
	var nSex int

	flag.IntVar(&nYear, "y", 1995, "-y=1995 ")
	flag.IntVar(&nMonth, "m", 6, "-m=6 ")
	flag.IntVar(&nDay, "d", 16, "-d=16 ")
	flag.IntVar(&nHour, "h", 19, "-h=19 ")
	flag.IntVar(&nMinute, "n", 7, "-n=7 ")
	flag.IntVar(&nSecond, "s", 0, "-s=0 ")
	flag.IntVar(&nSex, "x", 0, "-x=0  1是男0是女 ")

	flag.Parse() //解析输入的参数

	pBazi := bazi.GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex)
	fmt.Println(pBazi)
}
```

```
// HTTP 版本
package main

import (
	"fmt"
	"net/http"

	bazi "github.com/warrially/BaziGo"
)

type Myhandler struct{}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":7890", nil)
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pBazi := bazi.GetBazi(1995, 6, 16, 19, 7, 0, 0)
	fmt.Fprintln(w, pBazi.ToHTML())
}
```
