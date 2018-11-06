package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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
	if index >= len(in) {
		return "", -1
	}

	if index+1 == len(in) {
		return string(in[index]), -1
	}

	if in[index] == '^' {
		if len(in) == index+2 {
			return in[index : index+2], -1
		}
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

	dat, err := ioutil.ReadFile("in.txt")
	if err != nil {
		panic(err)
	}

	input := string(dat)
	fmt.Println(input)

	input = tokenize(input)

	var token string
	index := 0

	keyIndex := rune(48)
	dict := make(map[string]struct{})
	fullDict := make(map[string]string)

	token, index = nextToken(input, index)

	for index != -1 {

		_, ok := dict[token]
		if ok {
			delete(dict, token)
			newKey := "^" + string(keyIndex)
			dict[newKey] = struct{}{}
			keyIndex++
			if keyIndex == '^' {
				keyIndex++
			}

			// Change tokens to newKey
			subToken, subIndex := "", 0
			outstring := ""

			for subIndex != -1 {
				subToken, subIndex = nextToken(input, subIndex)
				if subToken != "" {
					if subToken == token {
						outstring += newKey
						if subToken[0] == '^' {
							if len(subToken) == 4 {
								subIndex += 2
							} else {
								subIndex++
							}
						} else {
							if subToken[1] == '^' {
								subIndex += 2
							} else {
								subIndex++
							}
						}
					} else {
						if subToken[0] == '^' {
							outstring += string(subToken[0:2])
						} else {
							outstring += string(subToken[0])
						}
					}
				}
			}

			input = outstring
			//fmt.Println("-> ", newKey, ":", token, "\n", input)
			dict = make(map[string]struct{})
			fullDict[newKey] = token
			index = 0
		} else {
			dict[token] = struct{}{}
		}

		//fmt.Println(token + " : " + strconv.Itoa(index))
		token, index = nextToken(input, index)
	}

	// Enforce rule utility

	/*
		Loop1:
			for {
			Loop2:
				for key, val := range fullDict {
					tk1, inx1 := nextToken(input, 0)
					for inx1 != -1 {
						if tk1 != "" {
							if tk1 == key {
								continue Loop2
							}
						}
						tk1, inx1 = nextToken(input, inx1)
					}

					for skey, sval := range fullDict {
						tk1, inx1 := nextToken(sval, 0)
						for inx1 != -1 {
							if tk1 != "" {
								if tk1 == key {
									continue Loop2
								}
							}
						}
					}
					continue Loop1
				}
				break
			}
	*/

	fmt.Println("Output: " + input)
	for key, val := range fullDict {
		fmt.Println(key + " ~ " + strings.Replace(val, "\n", "⏎", -1))
	}
}
