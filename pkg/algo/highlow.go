package algo

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// kata https://www.codewars.com/kata/554b4ac871d6813a03000035

func HighLow(in string) string {
	if len(in) == 0 {
		return "0 0"
	}
	arrStr := strings.Fields(in)
	arrInt := make([]int, len(arrStr))
	for _, v := range arrStr {
		digit, _ := strconv.Atoi(v)
		arrInt = append(arrInt, digit)
	}

	sort.Ints(arrInt)

	max := arrInt[len(arrInt)-1]
	min := arrInt[0]

	return fmt.Sprintf("%d %d", max, min)
}
