package crypto

import (
	"strconv"
	"strings"
)

func Convolution(arr []string) []string {
	var seq []string
	for i := 0; i < len(arr); i++ {
		var counter = 0
		for j := i; j < len(arr) && arr[j] == arr[i]; j++ {
			counter += 1
		}
		if counter != 1 {
			seq = append(seq, strconv.Itoa(counter))
		}
		seq = append(seq, arr[i])
		i += counter-1
	}
	return seq
}

func Encrypt(arr []string) ([]string, error) {
	seq := Convolution(arr)
	var result []string
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
	}
	return seq, nil
}
