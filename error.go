package bazi

import "errors"

const (
	MIN_BAZI_YEAR = 1800
	MAX_BAZI_YEAR = 2299

	SEX_FEMALE = 0
	SEX_MALE   = 1
)

var (
	ErrInvalidDate       = errors.New("invalid solar date")
	ErrInvalidLunarDate  = errors.New("invalid lunar date")
	ErrInvalidSex        = errors.New("invalid sex, only 0(female) or 1(male) is supported")
	ErrDateOutOfRange    = errors.New("date is outside the supported bazi range")
	ErrInvalidBaziSource = errors.New("solar date and lunar date are both nil")
	ErrIndexOutOfRange   = errors.New("index is out of range")
)

// CheckBaziInput 检查八字入口参数。
func CheckBaziInput(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) error {
	if nSex != SEX_FEMALE && nSex != SEX_MALE {
		return ErrInvalidSex
	}
	if nYear < MIN_BAZI_YEAR || nYear > MAX_BAZI_YEAR {
		return ErrDateOutOfRange
	}
	if NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond) == nil {
		return ErrInvalidDate
	}
	return nil
}
