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
	adjectives  []string
	nouns       []string
	delimiter   string
	tokenLength int
	tokenHex    bool
	tokenChars  string
	random      *rand.Rand
}

// NewHaikunator creates a new Haikunator with all default options
func NewHaikunator() Haikunator {
	return Haikunator{
		adjectives: []string{
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
		nouns: []string{
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
		delimiter:   "-",
		tokenLength: 4,
		tokenHex:    false,
		tokenChars:  "0123456789",
		random:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (h *Haikunator) haikunate() string {
	if h.tokenHex {
		h.tokenChars = "0123456789abcdef"
	}

	adjective := h.adjectives[h.random.Intn(len(h.adjectives))]
	noun := h.nouns[h.random.Intn(len(h.nouns))]
	var buffer bytes.Buffer

	for i := 0; i < h.tokenLength; i++ {
		buffer.WriteByte(h.tokenChars[h.random.Intn(len(h.tokenChars))])
	}

	token := buffer.String()
	sections := deleteEmpty([]string{adjective, noun, token})
	return strings.Join(sections, h.delimiter)
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
