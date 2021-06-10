package crypto

import (
	"errors"
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
			} else {
				return -1, errors.New("Неверная последовательность скобок")
			}
		}
	}
	return -1, errors.New("Неверная последовательность скобок")
}

func Decrypt(arr []string) ([]string, error) {
	var seq []string
	//var num int
	for i := 0; i < len(arr); i++ {
		if arr[i] == "(" {
			j, err := GetIndexLastBracket(arr[i:], i)
			if err != nil {
				return nil, errors.New("Неверная последовательность скобок")
			}
			if n, serr := strconv.Atoi(arr[i-1]); serr == nil {
				tmp, decerr := Decrypt(arr[i+1:j])
				if decerr != nil {
					return nil, errors.New("Неверная последовательность скобок")
				}
				for k := 0; k < n; k++ {
					seq = append(seq, strings.Join(tmp, ""))
				}
			} else {
				if decarr, decerr := Decrypt(arr[i+1:j]); decerr == nil {
					decoded := strings.Join(decarr, "")
					seq = append(seq, decoded)
				} else {
					return nil, errors.New("Неверная последовательность скобок")
				}
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
	return seq, nil
}
