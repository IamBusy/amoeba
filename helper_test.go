package amoeba

import (
	"testing"
)

func TestParseInclude(t *testing.T) {
	str := "measurements.(tags,fields)"
	res := SplitAttrs(str)
	if res[0] != "measurements.(tags,fields)" {
		t.Fatal(res)
	}


	str = "group.(user.(roles,operations),manager),maintainers"
	res = SplitAttrs(str)
	if res[0] != "group.(user.(roles,operations),manager)" {
		t.Fatal(res)
	}
	if res[1] != "maintainers" {
		t.Fatal(res)
	}

}

func TestParseAttrs(t *testing.T) {
	str := "measurements.(tags,fields)"
	first, rest := ParseAttrs(str)
	if first != "measurements" {
		t.Fatal(first)
	}
	if rest != "tags,fields" {
		t.Fatal(rest)
	}


	str = "group.(user.(roles,operations),manager)"
	first,rest = ParseAttrs(str)
	if first != "group" {
		t.Fatal(first)
	}
	if rest != "user.(roles,operations),manager" {
		t.Fatal(rest)
	}
}
