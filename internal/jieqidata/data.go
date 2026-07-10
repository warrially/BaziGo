package jieqidata

import (
	_ "embed"
	"encoding/binary"
	"errors"
)

const (
	N_HEADER_SIZE = 16
	N_RECORD_SIZE = 9
)

//go:generate go run ../../tools/jieqigen
//go:embed jieqi.bin
var pData []byte

var ErrInvalidData = errors.New("invalid embedded jieqi data")

func YearRange() (int, int, error) {
	if len(pData) < N_HEADER_SIZE || string(pData[:4]) != "BZJQ" {
		return 0, 0, ErrInvalidData
	}
	return int(binary.LittleEndian.Uint16(pData[6:8])), int(binary.LittleEndian.Uint16(pData[8:10])), nil
}

func Count() int {
	if len(pData) < N_HEADER_SIZE {
		return 0
	}
	return int(binary.LittleEndian.Uint32(pData[10:14]))
}

// Record 返回从公元起点累计的秒数和节气编号。
func Record(nIndex int) (int64, byte, error) {
	if nIndex < 0 || nIndex >= Count() {
		return 0, 0, ErrInvalidData
	}
	nOffset := N_HEADER_SIZE + nIndex*N_RECORD_SIZE
	if nOffset+N_RECORD_SIZE > len(pData) {
		return 0, 0, ErrInvalidData
	}
	return int64(binary.LittleEndian.Uint64(pData[nOffset : nOffset+8])), pData[nOffset+8], nil
}
