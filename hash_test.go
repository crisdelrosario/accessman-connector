package accessman

import (
	"regexp"
	"testing"
)

func TestUtilMD5(t *testing.T) {
	value := "/client"
	expecting := "a80250ca9753cea7eb0becb4cc4a5c15"
	regex := "[a-f0-9]{32}"
	hashString := Hash{value}.MD5()

	if got, want := len(hashString), 32; got != want {
		t.Errorf("Length was `%v`, expecting `%v`", got, want)
	}

	if got, want := hashString, expecting; got != want {
		t.Errorf("Value was `%v`, expecting `%v`", got, want)
	}

	match, _ := regexp.MatchString(regex, hashString)

	if got, want := match, true; got != want {

		t.Errorf("Regex matching `%s` was `%v`, expecting `%v`", regex, got, want)
	}
}

func TestUtilSHA1(t *testing.T) {
	value := "/client"
	expecting := "15a5f701524b9430ea1c1ee08841bfd846c1b11d"
	regex := "[a-f0-9]{40}"

	hashString := Hash{value}.SHA1()

	if got, want := len(hashString), 40; got != want {
		t.Errorf("Length was `%v`, expecting `%v`", got, want)
	}

	if got, want := hashString, expecting; got != want {
		t.Errorf("Value was `%v`, expecting `%v`", got, want)
	}

	match, _ := regexp.MatchString(regex, hashString)

	if got, want := match, true; got != want {

		t.Errorf("Regex matching `%s` was `%v`, expecting `%v`", regex, got, want)
	}
}
