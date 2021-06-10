package crypto

import (
	"fmt"
	"strconv"
	"strings"
)

func Convolution(arr []string) []string {
	var seq []string
	for i := 0; i < len(arr); i++ {
		var counter int = 0
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
	fmt.Println(strings.Join(seq, ""))
	return seq, nil
}
