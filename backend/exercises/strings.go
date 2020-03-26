package exercises

import (
	"khanek/exercise-generator/core/math"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var polishAlphabet []rune = []rune{'a', 'ą', 'b', 'c', 'd', 'e', 'ę', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'ó', 'p', 'r', 's', 'ś', 't', 'u', 'w', 'y', 'z', 'ź', 'ż'}

func addSpaces(word string) string {
	return strings.Join(strings.Split(word, ""), " ")
}

func mask(s string, percent float64) string {
	n := int(float64(len(s)) * percent)
	m := '_'
	out := []rune(s)
	maxN := 0
	for _, r := range out {
		if unicode.IsLetter(r) {
			maxN++
		}
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	n = math.Min(n, maxN)
	for n > 0 {
		randIndex := r.Intn(len(out))
		if out[randIndex] == m || !unicode.IsLetter(out[randIndex]) {
			continue
		}
		out[randIndex] = m
		n = n - 1
	}
	return string(out)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func shuffleWord(s string) string {
	sep := ""
	vals := strings.Split(strings.ToUpper(s), sep)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range vals {
		j := r.Intn(i + 1)
		vals[i], vals[j] = vals[j], vals[i]
	}
	return strings.Join(vals, sep)
}

func shuffle(s string) string {
	sep := " "
	vals := strings.Split(strings.ToUpper(s), sep)
	for i, val := range vals {
		vals[i] = shuffleWord(val)
	}
	return strings.Join(vals, sep)
}

func indexOf(runes []rune, search rune) int {
	for i, v := range runes {
		if v == search {
			return i
		}
	}
	return -1
}

func mod(i int, n int) int {
	return ((i % n) + n) % n
}

func ceasar(s string, shift int) string {
	runes := []rune(strings.ToLower(s))
	for i, r := range runes {
		pos := indexOf(polishAlphabet, r)
		if pos == -1 {
			continue
		}
		runes[i] = polishAlphabet[mod(pos+shift, len(polishAlphabet))]
	}
	return strings.ToUpper(string(runes))
}
