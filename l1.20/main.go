package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	s, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	s = strings.TrimSuffix(s, "\n")

	l, r := 0, len(s)     // Указатели на концы уже перевёрнутых частей
	p1, p2 := 0, len(s)-1 // Указатели на концы переворачиваемых частей
	for p1 < p2 {
		if p1 == p2 {
			break
		}

		flag1 := false
		flag2 := false

		if s[p1] == ' ' {
			flag1 = true
		} else {
			p1++
		}

		if s[p2] == ' ' {
			flag2 = true
		} else {
			p2--
		}

		if flag1 && flag2 {
			s = s[:l] + s[p2+1:r] + s[p1:p2+1] + s[l:p1] + s[r:] // Все операции происходят со срезом байт s

			l, r = l+(r-p2), r-(l+p1+1)

			p1 = l
			p2 = r - 1
		}
	}

	fmt.Fprintln(out, s)
}

/* Примеры:
$ go run l1.20/main.go
one two three four five           // input
five four three two one           // output

$ go run l1.20/main.go
snow dog sun		              // input
sun dog snow		              // output
*/
