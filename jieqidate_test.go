package bazi

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/warrially/BaziGo/internal/jieqidata"
)

func TestCompressedJieQiData(t *testing.T) {
	pFile, err := os.Open("tools/jieqigen/jieqi.csv.gz")
	if err != nil {
		t.Fatal(err)
	}
	defer pFile.Close()
	pGzip, err := gzip.NewReader(pFile)
	if err != nil {
		t.Fatal(err)
	}
	defer pGzip.Close()

	nIndex := 0
	for {
		var nYear, nMonth, nDay, nHour, nMinute, nSecond, nJieQi int
		_, err = fmt.Fscanf(pGzip, "%d,%d,%d,%d,%d,%d,%d\n", &nYear, &nMonth, &nDay, &nHour, &nMinute, &nSecond, &nJieQi)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("record %d: %v", nIndex, err)
		}
		nSeconds, nDataJieQi, err := jieqidata.Record(nIndex)
		if err != nil {
			t.Fatalf("record %d: %v", nIndex, err)
		}
		pSolarDate := NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond)
		if pSolarDate.Get64TimeStamp() != nSeconds || byte(nJieQi) != nDataJieQi {
			t.Fatalf("record %d mismatch", nIndex)
		}
		nIndex++
	}
	if nIndex != jieqidata.Count() {
		t.Fatalf("record count = %d, want %d", nIndex, jieqidata.Count())
	}
}
