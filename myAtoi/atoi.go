package main

import "fmt"

func main() {
	fmt.Println(myAtoi("  0000000000012345678"))
}

const (
	minInt32 = -1 << 31
	maxInt32 = (1 << 31) - 1
)

func myAtoi(s string) int {
	//check if empty string
	if len(s) > 0 {
		//get number from string 
		s = getNum(s)
		if len(s) > 0 {
			//check and remember number's sign
			sign := 1
			if s[0] == '-' {
				sign = -1
				s = s[1:]
			} else if s[0] == '+' {
				s = s[1:]
			}

			//atoi algo
			res := 0
			for _, n := range s {
				if isNum(n) {
					res += int(n - '0')
					//clamp to int32 bounds
					if res < minInt32 || res > maxInt32 {
						if sign == 1 {
							return maxInt32
						} else {
							return minInt32
						}
					}
					res *= 10
				} else {
					return 0
				}
			}
			return res * sign / 10
		}
		return 0
	}
	return 0
}

func getNum(s string) string {
	startIdx := -1
	signs := 0
	for i, r := range s {
		if startIdx == -1 && isNumSign(r) { //number started
			startIdx = i
		} else if startIdx != -1 && !isNum(r) { //number ended
			return string(s[startIdx:i])
		} else if startIdx != -1 && acceptable(r) { //counting signs, if more that 1 return empty
			if r == '-' || r == '+' {
				signs += 1
				if signs > 1 {
					return ""
				}
			}
		} else if startIdx == -1 && !acceptable(r) { //check if some chars are before number
			return ""
		}
	}
	if startIdx != -1 {
		return string(s[startIdx:])
	}
	return ""
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}
func acceptable(r rune) bool {
	return r == ' ' || r == '-' || r == '+'
}
func isNumSign(r rune) bool {
	return isNum(r) || r == '-' || r == '+'
}
