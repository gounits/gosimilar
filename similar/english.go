package similar

import "strings"

// 原理：根据英文的空格分割成单词
// 然后将单词转为数字进行比较
type english struct {
	w1, w2 string
}

func NewEnglishCompare(s1, s2 string) ISimilar {
	return &english{w1: s1, w2: s2}
}

func (e *english) Compare64() (c1 []uint8, c2 []uint8) {
	c1 = make([]uint8, 0, 64)
	c2 = make([]uint8, 0, 64)

	hashmap := make(map[string]uint8, 64)
	var count uint8 = 0

	flag := func(s string, c *[]uint8) {
		for _, word := range strings.Split(s, " ") {
			v, ok := hashmap[word]
			if !ok {
				hashmap[word] = count
				v = count
				count++
			}
			*c = append(*c, v)
		}
	}

	flag(e.w1, &c1)
	flag(e.w2, &c2)
	return
}
