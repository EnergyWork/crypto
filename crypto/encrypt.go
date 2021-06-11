package crypto

import (
	"strconv"
	"strings"
)

func GetString(arr []string) string {
	return strings.Join(arr, "")
}

func AddByElement(seq *[]string, str []string) {
	for _, el := range str {
		*seq = append(*seq, el)
	}
}

func Convolution(arr []string, step int) []string {
	var seq []string
	for i := 0; i < len(arr); {
		var counter = 0
		for j := i; j < (len(arr)) && (GetString(arr[j:j+step]) == GetString(arr[i:i+step])); j += step {
			counter += 1
		}
		if step == 1 {
			if counter > 1{
				seq = append(seq, strconv.Itoa(counter))
			}
			seq = append(seq, arr[i])
			i += step * counter
		} else {
			if i !=0 {
 				seq = seq[:len(seq)-1]
			}
			if counter > 1 {
				seq = append(seq, strconv.Itoa(counter))
				seq = append(seq, "(")
				AddByElement(&seq, arr[i:i+step]) //seq = append(seq, GetString(arr[i:i+step]))
				seq = append(seq, ")")
				i += step * counter - 1
			} else  {
				AddByElement(&seq, arr[i:i+step]) //seq = append(seq, GetString(arr[i:i+step]))
				i += step * counter - 1
			}
		}
	}
	if step > 1 {
		//AddByElement(&seq, arr[len(arr)-step:])
	}
	return seq
}

func Encrypt(arr []string) ([]string, error) {
	var seq = arr
	var step = 1
	for {
		if step <= len(seq) {
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
