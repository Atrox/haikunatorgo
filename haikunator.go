// Package haikunator generates Heroku-like random names to use in your go applications
package haikunator

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

// A Haikunator represents all options needed to use haikunate()
type Haikunator struct {
	Adjectives  []string
	Nouns       []string
	Delimiter   string
	TokenLength int
	TokenHex    bool
	TokenChars  string
	Random      *rand.Rand
}

// Adjectives used by haikunator
var adjectives = []string{
	"aged", "ancient", "autumn", "billowing", "bitter", "black", "blue", "bold",
	"broad", "broken", "calm", "cold", "cool", "crimson", "curly", "damp",
	"dark", "dawn", "delicate", "divine", "dry", "empty", "falling", "fancy",
	"flat", "floral", "fragrant", "frosty", "gentle", "green", "hidden", "holy",
	"icy", "jolly", "late", "lingering", "little", "lively", "long", "lucky",
	"misty", "morning", "muddy", "mute", "nameless", "noisy", "odd", "old",
	"orange", "patient", "plain", "polished", "proud", "purple", "quiet", "rapid",
	"raspy", "red", "restless", "rough", "round", "royal", "shiny", "shrill",
	"shy", "silent", "small", "snowy", "soft", "solitary", "sparkling", "spring",
	"square", "steep", "still", "summer", "super", "sweet", "throbbing", "tight",
	"tiny", "twilight", "wandering", "weathered", "white", "wild", "winter", "wispy",
	"withered", "yellow", "young",
}

// Nouns used by haikunator
var nouns = []string{
	"art", "band", "bar", "base", "bird", "block", "boat", "bonus",
	"bread", "breeze", "brook", "bush", "butterfly", "cake", "cell", "cherry",
	"cloud", "credit", "darkness", "dawn", "dew", "disk", "dream", "dust",
	"feather", "field", "fire", "firefly", "flower", "fog", "forest", "frog",
	"frost", "glade", "glitter", "grass", "hall", "hat", "haze", "heart",
	"hill", "king", "lab", "lake", "leaf", "limit", "math", "meadow",
	"mode", "moon", "morning", "mountain", "mouse", "mud", "night", "paper",
	"pine", "poetry", "pond", "queen", "rain", "recipe", "resonance", "rice",
	"river", "salad", "scene", "sea", "shadow", "shape", "silence", "sky",
	"smoke", "snow", "snowflake", "sound", "star", "sun", "sun", "sunset",
	"surf", "term", "thunder", "tooth", "tree", "truth", "union", "unit",
	"violet", "voice", "water", "water", "waterfall", "wave", "wildflower", "wind",
	"wood",
}

// NewHaikunator creates a new Haikunator with all default options
func NewHaikunator() Haikunator {
	return Haikunator{
		Adjectives:  adjectives,
		Nouns:       nouns,
		Delimiter:   "-",
		TokenLength: 4,
		TokenHex:    false,
		TokenChars:  "0123456789",
		Random:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Haikunate generates a random Heroku-like string
func (h *Haikunator) Haikunate() string {
	if h.TokenHex {
		h.TokenChars = "0123456789abcdef"
	}

	adjective := h.randomString(h.Adjectives)
	noun := h.randomString(h.Nouns)

	var token string

	if len(h.TokenChars) > 0 {
		var buffer bytes.Buffer

		for i := 0; i < h.TokenLength; i++ {
			randomIndex := h.Random.Intn(utf8.RuneCountInString(h.TokenChars))
			buffer.WriteRune([]rune(h.TokenChars)[randomIndex])
		}

		token = buffer.String()
	}

	sections := deleteEmpty([]string{adjective, noun, token})
	return strings.Join(sections, h.Delimiter)
}

// Get random string from string array
func (h *Haikunator) randomString(s []string) string {
	size := len(s)

	if size <= 0 { // Random.Intn panics otherwise
		return ""
	}

	return s[h.Random.Intn(size)]
}

// Deletes empty strings from string array
func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
