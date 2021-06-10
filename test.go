package main

import (
	"errors"
	"fmt"
	. "github.com/EnergyWork/go-stack"
	"regexp"
	"strconv"
	"strings"
)

func GetIndexLastBracket(arr []string, i int) (int, error) {
	var stack Stack
	for l := 0; l < len(arr); l++ {
		if arr[l] == "(" {
			stack.Push(arr[l])
		} else if arr[l] == ")" {
			if el, err := stack.Top(); err == nil {
				if el == "(" {
					stack.Pop()
					if len(stack) == 0 {
						return l + i, nil
					}
				}
			}
		}
	}
	return -1, errors.New("Неверная последовательность скобок")
}

func Decrypt(arr []string) []string {
	var seq []string
	//var num int
	for i := 0; i < len(arr); i++ {
		if arr[i] == "(" {
			j, err := GetIndexLastBracket(arr[i:], i)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			if n, err := strconv.Atoi(arr[i-1]); err == nil {
				tmp := Decrypt(arr[i+1:j])
				for k := 0; k < n; k++ {
					seq = append(seq, strings.Join(tmp, ""))
				}
			} else {
				decoded := strings.Join(Decrypt(arr[i+1:j]), "")
				seq = append(seq, decoded)
			}
			i += j - i - 1
		} else {
			re := regexp.MustCompile(`[A-Za-z]`)
			if re.MatchString(arr[i]) {
				if i > 0 {
					if n, err := strconv.Atoi(arr[i-1]); err == nil {
						for k := 0; k < n; k++ {
							seq = append(seq, arr[i])
						}
					} else {
						seq = append(seq, arr[i])
					}
				} else {
					seq = append(seq, arr[i])
				}
			}
		}
	}
	return seq
}

func debug_decrypt() {
	s := "2a(c2e)"
	t := strings.Split(s, "")
	decoded := Decrypt(t)
	fmt.Println(strings.Join(decoded, ""))
}
