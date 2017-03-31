package haikunator

import (
	"math/rand"
	"reflect"
	"regexp"
	"testing"
)

// haikunateTest contains everything needed for the general tests
type haikunateTest struct {
	Name       string
	Regex      string
	Parameters m
}

// shorter name
type m map[string]interface{}

func TestGeneralUsage(t *testing.T) {
	tests := []haikunateTest{
		// basic
		{"default usage", `[a-z]+-[a-z]+-[0-9]{4}$`, m{}},

		// token
		{"tokenhex", `[a-z]+-[a-z]+-[0-f]{4}$`, m{"TokenHex": true}},
		{"tokenlength", `[a-z]+-[a-z]+-[0-9]{9}$`, m{"TokenLength": 9}},
		{"tokenlength and tokenhex", `[a-z]+-[a-z]+-[0-f]{9}$`, m{"TokenLength": 9, "TokenHex": true}},
		{"zero tokenlength", `[a-z]+-[a-z]+$`, m{"TokenLength": 0}},

		// delimiter
		{"delimiter", `[a-z]+.[a-z]+.[0-9]{4}$`, m{"Delimiter": "."}},
		{"delimiter and tokenlength", `[a-z]+ [a-z]+`, m{"Delimiter": " ", "TokenLength": 0}},
		{"empty delimiter and tokenlength", `[a-z]+$`, m{"Delimiter": "", "TokenLength": 0}},

		// token chars
		{"tokenchars", `[a-z]+-[a-z]+-[x-z]{4}$`, m{"TokenChars": "xyz"}},

		// adjectives and nouns
		{"adjectives and nouns", `adjective-noun-\d{4}$`, m{"Adjectives": []string{"adjective"}, "Nouns": []string{"noun"}}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			h := New()

			// set specified parameters with reflection
			val := reflect.ValueOf(h).Elem()
			for key, value := range test.Parameters {
				val.FieldByName(key).Set(reflect.ValueOf(value))
			}

			haikunate := h.Haikunate()
			matched, err := regexp.MatchString(test.Regex, haikunate)
			if err != nil {
				t.Error(err)
			}
			if !matched {
				t.Error("Regex did not match with: ", haikunate)
			}
		})
	}
}

func TestWontReturnSameForSubsequentCalls(t *testing.T) {
	tests := []*Haikunator{New(), New()}

	for _, h1 := range tests {
		for _, h2 := range tests {
			v1 := h1.Haikunate()
			v2 := h2.Haikunate()
			if v1 == v2 {
				t.Error("Should not return same result", v1, v2)
			}
		}
	}
}

func TestReturnsSameForSameSeed(t *testing.T) {
	var seed int64 = 1001

	h1 := New()
	h1.Random = rand.New(rand.NewSource(seed))

	h2 := New()
	h2.Random = rand.New(rand.NewSource(seed))

	if h1.Haikunate() != h2.Haikunate() {
		t.Error("Sould return same")
	}
	if h1.Haikunate() != h2.Haikunate() {
		t.Error("Should return same")
	}

}

func TestZeroLengthOptionsPanic(t *testing.T) {
	h := New()
	h.Adjectives = make([]string, 0)
	h.Nouns = make([]string, 0)
	h.TokenChars = ""

	h.Haikunate() // should not panic when generating random numbers
}
