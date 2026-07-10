package bazi

import (
	"errors"
	"testing"
)

func TestGetBaziReadmeSample(t *testing.T) {
	pBazi, err := GetBaziE(1995, 6, 16, 19, 7, 0, SEX_FEMALE)
	if err != nil {
		t.Fatalf("GetBaziE() error = %v", err)
	}

	strSiZhu := pBazi.SiZhu().YearZhu().String() + " " +
		pBazi.SiZhu().MonthZhu().String() + " " +
		pBazi.SiZhu().DayZhu().String() + " " +
		pBazi.SiZhu().HourZhu().String()
	if strSiZhu != "乙亥 壬午 戊寅 壬戌" {
		t.Fatalf("SiZhu = %q, want %q", strSiZhu, "乙亥 壬午 戊寅 壬戌")
	}
}

func TestGetBaziRejectsInvalidInput(t *testing.T) {
	arrTests := []struct {
		strName string
		nYear   int
		nMonth  int
		nDay    int
		nSex    int
		errWant error
	}{
		{"year too small", MIN_BAZI_YEAR - 1, 1, 1, SEX_MALE, ErrDateOutOfRange},
		{"year too large", MAX_BAZI_YEAR + 1, 1, 1, SEX_MALE, ErrDateOutOfRange},
		{"invalid date", 1995, 2, 29, SEX_MALE, ErrInvalidDate},
		{"invalid sex", 1995, 6, 16, 2, ErrInvalidSex},
	}

	for _, stTest := range arrTests {
		t.Run(stTest.strName, func(t *testing.T) {
			pBazi, err := GetBaziE(stTest.nYear, stTest.nMonth, stTest.nDay, 0, 0, 0, stTest.nSex)
			if pBazi != nil {
				t.Fatal("GetBaziE() returned a non-nil result")
			}
			if !errors.Is(err, stTest.errWant) {
				t.Fatalf("GetBaziE() error = %v, want %v", err, stTest.errWant)
			}
		})
	}
}

func TestDaYunRejectsInvalidIndex(t *testing.T) {
	pBazi, err := GetBaziE(1995, 6, 16, 19, 7, 0, SEX_FEMALE)
	if err != nil {
		t.Fatal(err)
	}
	if pBazi.DaYun().Zhu(-1) != nil || pBazi.DaYun().Zhu(12) != nil {
		t.Fatal("Zhu() must return nil for an invalid index")
	}
	if pBazi.DaYun().Age(-1) != -1 || pBazi.DaYun().Age(12) != -1 {
		t.Fatal("Age() must return -1 for an invalid index")
	}
}
