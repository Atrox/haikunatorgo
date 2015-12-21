package haikunator

import (
	"math/rand"
	"regexp"
	"testing"
)

func TestDefaultUse(t *testing.T) {
	haikunator := NewHaikunator()
	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("((?:[a-z][a-z]+))(-)((?:[a-z][a-z]+))(-)(\\d{4})$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestHexUse(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.tokenHex = true

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("((?:[a-z][a-z]+))(-)((?:[a-z][a-z]+))(-)(.{4})$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestDigitsUse(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.tokenLength = 9

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("((?:[a-z][a-z]+))(-)((?:[a-z][a-z]+))(-)(\\d{9})$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestDigitsAsHexUse(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.tokenLength = 9
	haikunator.tokenHex = true

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("((?:[a-z][a-z]+))(-)((?:[a-z][a-z]+))(-)(.{9})$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestWontReturnSameForSubsequentCalls(t *testing.T) {
	haikunator := NewHaikunator()
	haiku1 := haikunator.haikunate()
	haiku2 := haikunator.haikunate()

	if haiku1 == haiku2 {
		t.Error(haiku1, " matches with ", haiku2)
	}
}

func TestDropsToken(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.tokenLength = 0

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("((?:[a-z][a-z]+))(-)((?:[a-z][a-z]+))$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestPermitsOptionalDelimiter(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.delimiter = "."

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("((?:[a-z][a-z]+))(\\.)((?:[a-z][a-z]+))(\\.)(\\d+)$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestSpaceDelimiterAndNoToken(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.delimiter = " "
	haikunator.tokenLength = 0

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("((?:[a-z][a-z]+))( )((?:[a-z][a-z]+))$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestOneSingleWord(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.delimiter = ""
	haikunator.tokenLength = 0

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("((?:[a-z][a-z]+))$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestCustomChars(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.tokenChars = "A"

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("((?:[a-z][a-z]+))(-)((?:[a-z][a-z]+))(-)(AAAA)$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestCustomAdjectivesAndNouns(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.adjectives = []string{"red"}
	haikunator.nouns = []string{"reindeer"}

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("(red)(-)(reindeer)(-)(\\d{4})$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestRemoveAdjectivesAndNouns(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.adjectives = []string{""}
	haikunator.nouns = []string{""}

	haiku := haikunator.haikunate()

	matched, err := regexp.MatchString("(\\d{4})$", haiku)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Regex did not match with: ", haiku)
	}
}

func TestCustomRandom(t *testing.T) {
	haikunator1 := NewHaikunator()
	haikunator1.random = rand.New(rand.NewSource(123))
	haiku1 := haikunator1.haikunate()

	haikunator2 := NewHaikunator()
	haikunator2.random = rand.New(rand.NewSource(123))
	haiku2 := haikunator2.haikunate()

	if haiku1 != haiku2 {
		t.Error(haiku1, "does not match with ", haiku2)
	}
}
