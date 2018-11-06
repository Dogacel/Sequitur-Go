package main

import (
	"fmt"
)

func tokenize(in string) (out string) {
	for _, val := range in {
		if val == '^' {
			out += "^^"
		} else {
			out += string(val)
		}
	}

	return out
}

func nextToken(in string, index int) (string, int) {
	if index+2 > len(in) {
		return "", -1
	}

	if in[index] == '^' {
		if in[index+2] == '^' {
			return in[index : index+4], index + 2
		}
		return in[index : index+3], index + 2
	}

	if in[index+1] == '^' {
		return in[index : index+3], index + 1
	}

	return in[index : index+2], index + 1
}

func main() {
	var input string
	fmt.Scanln(&input)

	input = tokenize(input)

	var token string
	index := 0

	keyIndex := rune(0)
	dict := make(map[string]struct{})

	token, index = nextToken(input, index)

	for index != -1 {

		_, ok := dict[token]
		if ok {
			delete(dict, token)
			newKey := "^" + string(keyIndex)
			dict[newKey] = struct{}{}

			// Change tokens to newKey

			index = 0
		} else {
			dict[token] = struct{}{}
		}

		// fmt.Println(token + " : " + strconv.Itoa(index))
		token, index = nextToken(input, index)
	}

}
