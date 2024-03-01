package main

import (
	"fmt"
	"strings"
)

var (
	alphabet  = "abcdefghijklmnopqrstuvwxyz}{"
	alphabet2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ}{"
)

func main() {
	ciphertext := "Ybjufbqjycbq{ Mca!ub'i Tqo, cfywyhbq{{og }bcmb qi Ybjuhfbqjycbgq{ Mcf}ybw Mcaub'i Tqo, yi su{uhrfqjutg qfckbt jxu mcf{t ulufo ouqf cb ntiohmqop{wnuizkpg."

	for i := 1; i < 28; i++ {
		apply(ciphertext, i)
	}

}

func apply(ciphertext string, n int) {
	text := []byte{}
	for _, r := range ciphertext {
		w := shift(byte(r), n)
		text = append(text, w)
	}
	fmt.Println(string(text))
}

func shift(r byte, n int) byte {
	if strings.Contains(alphabet, string(r)) {
		r = r - 'a'
		tmp := int(r) + n
		if tmp >= len(alphabet) {
			return alphabet[tmp-len(alphabet)]
		} else {
			return alphabet[tmp]
		}
	} else if strings.Contains(alphabet2, string(r)) {
		r = r - 'A'
		tmp := int(r) + n
		if tmp >= len(alphabet) {
			return alphabet2[tmp-len(alphabet)]
		} else {
			return alphabet2[tmp]
		}
	}
	return r
}
