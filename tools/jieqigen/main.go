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
	checkError(validateRecords(arrRecord))
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
			*pValue, err = strconv.Atoi(arrMatch[i+1])
			if err != nil {
				return nil, fmt.Errorf("invalid value at record %d: %w", len(arrRecord), err)
			}
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
	pFile, err := os.CreateTemp(filepath.Dir(strFilename), ".jieqi-*.tmp")
	if err != nil {
		return err
	}
	strTemp := pFile.Name()
	defer os.Remove(strTemp)
	pWriter := bufio.NewWriter(pFile)
	if _, err = pWriter.WriteString(STR_MAGIC); err != nil {
		return err
	}
	arrHeader := []any{N_FORMAT_VERSION, uint16(arrRecord[0].nYear), uint16(arrRecord[len(arrRecord)-1].nYear), uint32(len(arrRecord)), uint16(0)}
	for _, value := range arrHeader {
		if err = binary.Write(pWriter, binary.LittleEndian, value); err != nil {
			return err
		}
	}
	for _, stRecord := range arrRecord {
		if err = binary.Write(pWriter, binary.LittleEndian, getSecondsFromYearZero(stRecord)); err != nil {
			return err
		}
		if err = pWriter.WriteByte(byte(stRecord.nJieQi)); err != nil {
			return err
		}
	}
	if err = pWriter.Flush(); err != nil {
		return err
	}
	if err = pFile.Sync(); err != nil {
		return err
	}
	if err = pFile.Close(); err != nil {
		return err
	}
	return os.Rename(strTemp, strFilename)
}

func validateRecords(arrRecord []TJieQiRecord) error {
	if len(arrRecord) == 0 || len(arrRecord)%24 != 0 {
		return errors.New("jieqi record count must be a non-zero multiple of 24")
	}
	for i, stRecord := range arrRecord {
		if stRecord.nYear <= 0 || stRecord.nMonth < 1 || stRecord.nMonth > 12 || stRecord.nDay < 1 || stRecord.nDay > getMonthDays(stRecord.nYear, stRecord.nMonth) || stRecord.nHour < 0 || stRecord.nHour > 23 || stRecord.nMinute < 0 || stRecord.nMinute > 59 || stRecord.nSecond < 0 || stRecord.nSecond > 59 {
			return fmt.Errorf("invalid date at record %d", i)
		}
		if stRecord.nJieQi < 0 || stRecord.nJieQi >= 24 {
			return fmt.Errorf("invalid jieqi number at record %d", i)
		}
	}
	return nil
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
