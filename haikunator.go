// Package haikunator generates Heroku-like random names to use in your go applications
package haikunator

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
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

// NewHaikunator creates a new Haikunator with all default options
func NewHaikunator() Haikunator {
	return Haikunator{
		Adjectives: []string{
			"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark",
			"summer", "icy", "delicate", "quiet", "white", "cool", "spring", "winter",
			"patient", "twilight", "dawn", "crimson", "wispy", "weathered", "blue",
			"billowing", "broken", "cold", "damp", "falling", "frosty", "green",
			"long", "late", "lingering", "bold", "little", "morning", "muddy", "old",
			"red", "rough", "still", "small", "sparkling", "throbbing", "shy",
			"wandering", "withered", "wild", "black", "young", "holy", "solitary",
			"fragrant", "aged", "snowy", "proud", "floral", "restless", "divine",
			"polished", "ancient", "purple", "lively", "nameless", "lucky", "odd", "tiny",
			"free", "dry", "yellow", "orange", "gentle", "tight", "super", "royal", "broad",
			"steep", "flat", "square", "round", "mute", "noisy", "hushy", "raspy", "soft",
			"shrill", "rapid", "sweet", "curly", "calm", "jolly", "fancy", "plain", "shinny",
		},
		Nouns: []string{
			"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning",
			"snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter",
			"forest", "hill", "cloud", "meadow", "sun", "glade", "bird", "brook",
			"butterfly", "bush", "dew", "dust", "field", "fire", "flower", "firefly",
			"feather", "grass", "haze", "mountain", "night", "pond", "darkness",
			"snowflake", "silence", "sound", "sky", "shape", "surf", "thunder",
			"violet", "water", "wildflower", "wave", "water", "resonance", "sun",
			"wood", "dream", "cherry", "tree", "fog", "frost", "voice", "paper",
			"frog", "smoke", "star", "atom", "band", "bar", "base", "block", "boat",
			"term", "credit", "art", "fashion", "truth", "disk", "math", "unit", "cell",
			"scene", "heart", "recipe", "union", "limit", "bread", "toast", "bonus",
			"lab", "mud", "mode", "poetry", "tooth", "hall", "king", "queen", "lion", "tiger",
			"penguin", "kiwi", "cake", "mouse", "rice", "coke", "hola", "salad", "hat",
		},
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

	adjective := h.Adjectives[h.Random.Intn(len(h.Adjectives))]
	noun := h.Nouns[h.Random.Intn(len(h.Nouns))]
	var buffer bytes.Buffer

	for i := 0; i < h.TokenLength; i++ {
		buffer.WriteByte(h.TokenChars[h.Random.Intn(len(h.TokenChars))])
	}

	token := buffer.String()
	sections := deleteEmpty([]string{adjective, noun, token})
	return strings.Join(sections, h.Delimiter)
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
