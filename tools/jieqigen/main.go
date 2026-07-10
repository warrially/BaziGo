package main

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const (
	STR_MAGIC        = "BZJQ"
	N_FORMAT_VERSION = uint16(1)
	N_HEADER_SIZE    = 16
	N_RECORD_SIZE    = 9
)

type TJieQiRecord struct {
	nYear, nMonth, nDay     int
	nHour, nMinute, nSecond int
	nJieQi                  int
}

func main() {
	strImport := flag.String("import", "", "从旧版 jieqidate.go 导入并生成压缩源数据")
	flag.Parse()

	strRoot, err := findProjectRoot()
	checkError(err)
	strSource := filepath.Join(strRoot, "tools", "jieqigen", "jieqi.csv.gz")
	strOutput := filepath.Join(strRoot, "internal", "jieqidata", "jieqi.bin")

	var arrRecord []TJieQiRecord
	if *strImport != "" {
		arrRecord, err = readOldGoSource(*strImport)
		checkError(err)
		checkError(writeCompressedSource(strSource, arrRecord))
	} else {
		arrRecord, err = readCompressedSource(strSource)
		checkError(err)
	}
	checkError(writeBinary(strOutput, arrRecord))
	fmt.Printf("generated %s: %d records\n", strOutput, len(arrRecord))
}

func findProjectRoot() (string, error) {
	strDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err = os.Stat(filepath.Join(strDir, "go.mod")); err == nil {
			return strDir, nil
		}
		strParent := filepath.Dir(strDir)
		if strParent == strDir {
			return "", errors.New("go.mod not found")
		}
		strDir = strParent
	}
}

func readOldGoSource(strFilename string) ([]TJieQiRecord, error) {
	pFile, err := os.Open(strFilename)
	if err != nil {
		return nil, err
	}
	defer pFile.Close()
	pRegexp := regexp.MustCompile(`^\s*\{(\d+),\s*(\d+),\s*(\d+),\s*(\d+),\s*(\d+),\s*(\d+),\s*(\d+)\},`)
	arrRecord := make([]TJieQiRecord, 0, 55000)
	pScanner := bufio.NewScanner(pFile)
	for pScanner.Scan() {
		arrMatch := pRegexp.FindStringSubmatch(pScanner.Text())
		if len(arrMatch) != 8 {
			continue
		}
		var stRecord TJieQiRecord
		arrValue := []*int{&stRecord.nYear, &stRecord.nMonth, &stRecord.nDay, &stRecord.nHour, &stRecord.nMinute, &stRecord.nSecond, &stRecord.nJieQi}
		for i, pValue := range arrValue {
			*pValue, _ = strconv.Atoi(arrMatch[i+1])
		}
		arrRecord = append(arrRecord, stRecord)
	}
	if err = pScanner.Err(); err != nil {
		return nil, err
	}
	if len(arrRecord) == 0 {
		return nil, errors.New("no jieqi records found")
	}
	return arrRecord, nil
}

func writeCompressedSource(strFilename string, arrRecord []TJieQiRecord) error {
	if err := os.MkdirAll(filepath.Dir(strFilename), 0755); err != nil {
		return err
	}
	pFile, err := os.Create(strFilename)
	if err != nil {
		return err
	}
	defer pFile.Close()
	pGzip := gzip.NewWriter(pFile)
	for _, stRecord := range arrRecord {
		fmt.Fprintf(pGzip, "%d,%d,%d,%d,%d,%d,%d\n", stRecord.nYear, stRecord.nMonth, stRecord.nDay, stRecord.nHour, stRecord.nMinute, stRecord.nSecond, stRecord.nJieQi)
	}
	return pGzip.Close()
}

func readCompressedSource(strFilename string) ([]TJieQiRecord, error) {
	pFile, err := os.Open(strFilename)
	if err != nil {
		return nil, err
	}
	defer pFile.Close()
	pGzip, err := gzip.NewReader(pFile)
	if err != nil {
		return nil, err
	}
	defer pGzip.Close()
	arrRecord := make([]TJieQiRecord, 0, 55000)
	for {
		var stRecord TJieQiRecord
		_, err = fmt.Fscanf(pGzip, "%d,%d,%d,%d,%d,%d,%d\n", &stRecord.nYear, &stRecord.nMonth, &stRecord.nDay, &stRecord.nHour, &stRecord.nMinute, &stRecord.nSecond, &stRecord.nJieQi)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		arrRecord = append(arrRecord, stRecord)
	}
	return arrRecord, nil
}

func writeBinary(strFilename string, arrRecord []TJieQiRecord) error {
	if len(arrRecord) == 0 {
		return errors.New("empty jieqi records")
	}
	if err := os.MkdirAll(filepath.Dir(strFilename), 0755); err != nil {
		return err
	}
	pFile, err := os.Create(strFilename)
	if err != nil {
		return err
	}
	defer pFile.Close()
	pWriter := bufio.NewWriter(pFile)
	pWriter.WriteString(STR_MAGIC)
	binary.Write(pWriter, binary.LittleEndian, N_FORMAT_VERSION)
	binary.Write(pWriter, binary.LittleEndian, uint16(arrRecord[0].nYear))
	binary.Write(pWriter, binary.LittleEndian, uint16(arrRecord[len(arrRecord)-1].nYear))
	binary.Write(pWriter, binary.LittleEndian, uint32(len(arrRecord)))
	binary.Write(pWriter, binary.LittleEndian, uint16(0))
	for _, stRecord := range arrRecord {
		binary.Write(pWriter, binary.LittleEndian, getSecondsFromYearZero(stRecord))
		pWriter.WriteByte(byte(stRecord.nJieQi))
	}
	return pWriter.Flush()
}

func getSecondsFromYearZero(stRecord TJieQiRecord) int64 {
	nDays := (stRecord.nYear-1)*365 + stRecord.nDay
	for i := 1; i < stRecord.nMonth; i++ {
		nDays += getMonthDays(stRecord.nYear, i)
	}
	nLeapDays := (stRecord.nYear - 1) / 4
	isGregorian := stRecord.nYear > 1582 || (stRecord.nYear == 1582 && (stRecord.nMonth > 10 || (stRecord.nMonth == 10 && stRecord.nDay >= 15)))
	if isGregorian {
		nLeapDays -= 10
		if stRecord.nYear > 1700 {
			nLeapDays -= 1 + (stRecord.nYear-1701)/100
			nLeapDays += (stRecord.nYear - 1601) / 400
		}
	}
	return int64(nDays+nLeapDays)*86400 + int64(stRecord.nHour*3600+stRecord.nMinute*60+stRecord.nSecond)
}

func getMonthDays(nYear, nMonth int) int {
	if nMonth == 2 {
		if nYear < 1582 && nYear%4 == 0 {
			return 29
		}
		if nYear >= 1582 && nYear%4 == 0 && (nYear%100 != 0 || nYear%400 == 0) {
			return 29
		}
		return 28
	}
	if nMonth == 4 || nMonth == 6 || nMonth == 9 || nMonth == 11 {
		return 30
	}
	return 31
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
