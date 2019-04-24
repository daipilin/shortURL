package utils

//进制转换
type Convertor struct {
	numToCharMap map[uint]string
	charToNumMap map[string]uint
}
func NewConvertor() *Convertor {
	convertor := new(Convertor)

	convertor.numToCharMap = map[uint]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
		10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j",
		20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t",
		30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: "A", 37: "B", 38: "C", 39: "D",
		40: "E", 41: "F", 42: "G", 43: "H", 44: "I", 45: "J", 46: "K", 47: "L", 48: "M", 49: "N",
		50: "O", 51: "P", 52: "Q", 53: "R", 54: "S", 55: "T", 56: "U", 57: "V", 58: "W", 59: "X",
		60: "Y", 61: "Z"}

	convertor.charToNumMap = make(map[string]uint)
	for key, value := range convertor.numToCharMap {
		convertor.charToNumMap[value] = key
	}

	return convertor
}
func (convertor *Convertor) ConvertToString(num uint) string {
	newNumStr := ""
	for num != 0 {
		newNumStr = convertor.numToCharMap[num % 62] + newNumStr
		num = num / 62
	}
	return newNumStr
}
func (convertor *Convertor) ConvertToNum(str string) uint {
	var num uint = 0
	for i := 0; i < len(str); i++ {
		num = convertor.charToNumMap[str[i:i + 1]] + num*62
	}
	return num
}