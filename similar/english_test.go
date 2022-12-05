package similar_test

import (
	"fmt"
	"github.com/gounits/gosimilar"
	"testing"
)

func TestEnglishSimilar(t *testing.T) {
	s1 := "I Love You"
	s2 := "I Like You"
	e := gosimilar.EnglishSimilar(s1, s2)
	score, _ := gosimilar.PhraseSimilar(e)
	if fmt.Sprintf("%.4f", score) != "0.6667" {
		t.Errorf("[NewEnglishCompare] s1:%s;s2:%s. result Error", s1, s2)
	}
}
