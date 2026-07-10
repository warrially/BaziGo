package bazi

// NormalizeNumber 将整数规范化到 [0, nMod) 范围。
func NormalizeNumber(nValue, nMod int) int {
	if nMod <= 0 {
		return 0
	}
	return ((nValue % nMod) + nMod) % nMod
}
