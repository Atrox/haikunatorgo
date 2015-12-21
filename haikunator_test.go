package haikunator

import (
	"math/rand"
	"regexp"
	"testing"
)

func TestDefaultUse(t *testing.T) {
	haikunator := NewHaikunator()
	haiku := haikunator.Haikunate()

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
	haikunator.TokenHex = true

	haiku := haikunator.Haikunate()

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
	haikunator.TokenLength = 9

	haiku := haikunator.Haikunate()

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
	haikunator.TokenLength = 9
	haikunator.TokenHex = true

	haiku := haikunator.Haikunate()

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
	haiku1 := haikunator.Haikunate()
	haiku2 := haikunator.Haikunate()

	if haiku1 == haiku2 {
		t.Error(haiku1, " matches with ", haiku2)
	}
}

func TestDropsToken(t *testing.T) {
	haikunator := NewHaikunator()
	haikunator.TokenLength = 0

	haiku := haikunator.Haikunate()

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
	haikunator.Delimiter = "."

	haiku := haikunator.Haikunate()

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
	haikunator.Delimiter = " "
	haikunator.TokenLength = 0

	haiku := haikunator.Haikunate()

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
	haikunator.Delimiter = ""
	haikunator.TokenLength = 0

	haiku := haikunator.Haikunate()

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
	haikunator.TokenChars = "A"

	haiku := haikunator.Haikunate()

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
	haikunator.Adjectives = []string{"red"}
	haikunator.Nouns = []string{"reindeer"}

	haiku := haikunator.Haikunate()

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
	haikunator.Adjectives = []string{""}
	haikunator.Nouns = []string{""}

	haiku := haikunator.Haikunate()

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
	haikunator1.Random = rand.New(rand.NewSource(123))
	haiku1 := haikunator1.Haikunate()

	haikunator2 := NewHaikunator()
	haikunator2.Random = rand.New(rand.NewSource(123))
	haiku2 := haikunator2.Haikunate()

	if haiku1 != haiku2 {
		t.Error(haiku1, "does not match with ", haiku2)
	}
}
