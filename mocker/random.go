package mocker

import (
	"log"
	"math/rand"
	"strconv"
)

const (
	letterBytes = "abcdefghijklmnoqrstuvwyxz"
)

// returnRandom returns an array of values matching the type of each
// format specifier in the string.
func returnRandom(s string) []interface{} {
	var randomValues []interface{}
	for i := 0; i < len(s); i++ {
		if s[i] == '%' {
			if i+1 < len(s) {
				size, i := getSize(i+1, s)

				switch s[i] {
				case 'd':
					digits := 1
					for j := 0; j < size; j++ {
						digits *= 10
					}

					randomValues = append(randomValues, rand.Int()%digits)
				case 's':
					randomValues = append(randomValues, randomString(size))
				}
			}
		}
	}
	return randomValues
}

// getSize iterates through a format string and returns the requested size
// along with the incremented index. Default size is five if no size is specified.
// Eg. getSize(0, "234d") returns 3, 234.
func getSize(i int, s string) (int, int) {
	size := 0
	for s[i] >= '0' && s[i] <= '9' {
		value, err := strconv.Atoi(string(s[i]))
		if err != nil {
			log.Fatal(err)
		}
		size = size*10 + value

		i++
		if i >= len(s) {
			break
		}
	}

	if size == 0 {
		return 5, i
	}
	return size, i
}

// randomString produces a random lowercase string of length n.
func randomString(length int) string {
	if length < 0 {
		return ""
	}

	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(bytes)
}
