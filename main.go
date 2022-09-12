package main

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	n := len(words)
	res := []string{}
	l := 0
	for i, j := 0, 0; i < n; {
		if j == n {
			line := layout(words, i, j-1, j-i-1, maxWidth)
			res = append(res, line)
			l = 0
			i = j
		} else if l+len(words[j])+j-i > maxWidth {
			line := layout(words, i, j-1, maxWidth-l, maxWidth)
			res = append(res, line)
			l = 0
			i = j
		} else {
			l = l + len(words[j])
			j++
		}
	}
	return res
}

func layout(words []string, first int, last int, spaces int, maxWidth int) string {
	var line strings.Builder
	for i := first; i < last; i++ {
		line.WriteString(words[i])
		for s := (spaces + last - i - 1) / (last - i); s > 0; s-- {
			line.WriteByte(' ')
			spaces--
		}
	}
	line.WriteString(words[last])
	for i := line.Len(); i < maxWidth; i++ {
		line.WriteByte(' ')
	}
	return line.String()
}
