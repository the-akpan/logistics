package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func ContainsString(seq []string, el string) string {
	for _, value := range seq {
		if strings.EqualFold(value, el) {
			return value
		}
	}
	return ""
}

func GenerateNumericString(length int) string {
	rand.Seed(time.Now().Unix())
	nums := make([]string, length)

	for i := 0; i < int(length); i++ {
		nums = append(nums, strconv.Itoa(rand.Intn(10)))
	}
	return strings.Join(nums, "")
}
