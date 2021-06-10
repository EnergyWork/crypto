package crypto

import (
	"fmt"
	"strconv"
	"strings"
)

func Convolution(arr []string) []string {
	var seq []string
	for i := 0; i < len(arr)-1; i++ {
		var counter int = 0
		for j := i+1; j <= len(arr) && arr[j] == arr[i]; j++ {
			counter += 1
		}
		seq = append(seq, strconv.Itoa(counter))
		seq = append(seq, arr[i])
	}
	return seq
}

func Encrypt(arr []string) ([]string, error) {
	seq := Convolution(arr)
	fmt.Println(strings.Join(seq, ""))
	return seq, nil
}
