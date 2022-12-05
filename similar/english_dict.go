package similar

import "strings"

type englishDict struct {
	dict map[string]string
}

func NewEnglishDictCompare(s1, s2 string, dict map[string]string) ISimilar {
	e := englishDict{dict: dict}
	s1 = e.replace(s1)
	s2 = e.replace(s2)
	return &english{w1: s1, w2: s2}
}

// 字符串根据字典去替换成新的字符串
func (e *englishDict) replace(s1 string) string {
	value := strings.Split(s1, " ")
	for index, word := range value {
		if v, ok := e.dict[word]; ok {
			value[index] = v
		}
	}
	return strings.Join(value, " ")
}
