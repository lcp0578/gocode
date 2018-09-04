package main

import (
	"fmt"
	strs "strings"
)

var PL = fmt.Println

func main() {
	PL("Contains:", strs.Contains("lcpeng", "lc"))
	PL("Count:", strs.Count("lcpenglcp", "lc"))
	PL("HasPrefix:", strs.HasPrefix("lcpeng", "lcp"))
	PL("HashSuffix:", strs.HasSuffix("lcpeng", "eng"))
	PL("Index:", strs.Index("lcpeng", "h"))
	PL("Join:", strs.Join([]string{"lcp", "eng"}, "-"))
	PL("Repeat:", strs.Repeat("lcp", 3))
	PL("Replace:", strs.Replace("lcpengcc", "c", "C", 1))
	PL("Replace:", strs.Replace("lcpengcc", "c", "C", -1))
	PL("Split:", strs.Split("l-c-p-e-n-g", "-"))
	PL("ToLower:", strs.ToLower("LCpenG"))
	PL("ToUpper:", strs.ToUpper("lcpEnG"))

	PL("len", len("lcpeng"))
	PL("char:", "lcpeng"[2])
}
