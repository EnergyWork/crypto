package crypto

import (
	"strconv"
	"strings"
)

func GetString(arr []string) string {
	return strings.Join(arr, "")
}

func AddByElement(seq *[]string, str []string) {
	*seq = append(*seq, str...)
	/*for _, el := range str {
		*seq = append(*seq, el)
	}*/
}

func Convolution(arr []string, step int) []string {
	var seq []string
	for i := 0; i < len(arr); {
		var counter = 0
		for j := i; j <= len(arr)-step; j += step {
			if GetString(arr[i:i+step]) == GetString(arr[j:j+step]) {
				counter += 1
			} else {
				break
			}
		}
		if step == 1 {
			if counter > 1 {
				seq = append(seq, strconv.Itoa(counter))
			}
			seq = append(seq, arr[i])
			if counter != 0 {
				i += step * counter
			} else {
				i += step
			}
		} else {
			if counter > 1 {
				seq = append(seq, strconv.Itoa(counter))
				seq = append(seq, "(")
				AddByElement(&seq, arr[i:i+step]) //seq = append(seq, GetString(arr[i:i+step]))
				seq = append(seq, ")")
				i += step * counter
			} else  {
				seq = append(seq, arr[i])
				i += 1
			}
		}
	}
	return seq
}

func Encrypt(arr []string) ([]string, error) {
	var seq = arr
	var step = 1
	for {
		if step <= len(seq) / 2 {
			seq = Convolution(seq, step)
		} else {
			break
		}
		step++
	}
	/*var result []string
	for step := 1; step <= (len(seq) / 2); step++ {
		var counter = 0
		seq1 := strings.Join(seq[0:0+step], "")
		for j := 0; j < len(seq) - step; j += step {
			seq2 := strings.Join(seq[j+step:j+2*step], "")
			if seq1 == seq2 {
				counter += 1
			}
		}
		if counter != 1 {
			result = append(result, strconv.Itoa(counter))
		}
		result = append(result, seq1)
	}*/
	return seq, nil
}
