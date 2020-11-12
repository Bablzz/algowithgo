package algo

import "strconv"

func Looksay(s string) (rs string) {
	var (
		cb  byte
		inc int
		i   int
	)

	cb = s[0]
	inc = 1
	for i = 1; i < len(s); i++ {
		var db byte
		db = s[i]
		if db == cb {
			inc++
			continue
		}
		rs = rs + strconv.Itoa(inc) + string(cb)
		cb = db
		inc = 1
	}

	return rs + strconv.Itoa(inc) + string(cb)

}
