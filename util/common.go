package util

import (
	"math/rand"
	"strings"
	"time"
)

func IToPI(i int64) *int64 {
	return &i
}

func RandomTest(f, t int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(t-f+1) + f
}

func ToLower(s string) string {
	return strings.ToLower(s)
}
