package similar_test

import (
	"github.com/gounits/gosimilar"
	"testing"
)

func TestEnglishDictSimilar(t *testing.T) {
	s1 := "I Love You"
	s2 := "I Like You"
	e := gosimilar.EnglishDictSimilar(s1, s2, map[string]string{"Love": "Like"})
	score, _ := gosimilar.PhraseSimilar(e)
	if score != 1 {
		t.Errorf("[NewEnglishDictCompare] s1:%s;s2:%s. result Error", s1, s2)
	}
}
