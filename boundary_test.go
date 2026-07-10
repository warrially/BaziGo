package bazi

import (
	"errors"
	"testing"

	"github.com/warrially/BaziGo/internal/jieqidata"
)

func TestLiChunBoundary(t *testing.T) {
	nMinYear, _, err := jieqidata.YearRange()
	if err != nil {
		t.Fatal(err)
	}
	nYear := 1995
	nSeconds, _, err := jieqidata.Record((nYear-nMinYear)*24 + 2)
	if err != nil {
		t.Fatal(err)
	}
	if nGot := GetLiChunYear(NewSolarDateFrom64TimeStamp(nSeconds - 1)); nGot != nYear-1 {
		t.Fatalf("before LiChun = %d", nGot)
	}
	if nGot := GetLiChunYear(NewSolarDateFrom64TimeStamp(nSeconds)); nGot != nYear-1 {
		t.Fatalf("at LiChun = %d", nGot)
	}
	if nGot := GetLiChunYear(NewSolarDateFrom64TimeStamp(nSeconds + 1)); nGot != nYear {
		t.Fatalf("after LiChun = %d", nGot)
	}
}

func TestSafeConstructors(t *testing.T) {
	if _, err := NewSolarDateE(1995, 2, 29, 0, 0, 0); !errors.Is(err, ErrInvalidDate) {
		t.Fatalf("NewSolarDateE error = %v", err)
	}
	if _, err := NewLunarDateE(1799, 1, 1, 0, 0, 0); !errors.Is(err, ErrInvalidLunarDate) {
		t.Fatalf("NewLunarDateE error = %v", err)
	}
	if _, err := NewBaziE(nil, SEX_FEMALE); !errors.Is(err, ErrInvalidDate) {
		t.Fatalf("NewBaziE error = %v", err)
	}
	if _, err := NewBaziFromLunarDateE(nil, SEX_FEMALE); !errors.Is(err, ErrInvalidLunarDate) {
		t.Fatalf("NewBaziFromLunarDateE error = %v", err)
	}
}

func TestNegativeNumberNormalization(t *testing.T) {
	if NewGan(-1).String() != "癸" {
		t.Fatal("NewGan(-1)")
	}
	if NewZhi(-1).String() != "亥" {
		t.Fatal("NewZhi(-1)")
	}
	if NewGanZhi(-1).String() != "癸亥" {
		t.Fatal("NewGanZhi(-1)")
	}
	if NewWuXing(-1).String() != "土" {
		t.Fatal("NewWuXing(-1)")
	}
	if NewYinYang(-1).String() != "阳" {
		t.Fatal("NewYinYang(-1)")
	}
	if NewNaYin(-1).String() != "大海水" {
		t.Fatal("NewNaYin(-1)")
	}
}

func TestSolarLunarRoundTrip(t *testing.T) {
	pSolarDate := NewSolarDate(1995, 6, 16, 19, 7, 0)
	pLunarDate := pSolarDate.ToLunarDate()
	if pLunarDate == nil {
		t.Fatal("ToLunarDate returned nil")
	}
	pResult := pLunarDate.ToSolarDate()
	if pResult == nil || pResult.Get64TimeStamp() != pSolarDate.Get64TimeStamp() {
		t.Fatalf("round trip = %v", pResult)
	}
}
