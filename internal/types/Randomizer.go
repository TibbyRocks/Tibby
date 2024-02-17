package types

import (
	"math/rand"
	"slices"

	"github.com/tibbyrocks/tibby/internal/utils"
)

type Randomizer struct {
	items []string
}

// Appends a string to the randomizer slice
func (r *Randomizer) Append(appendableItems ...string) {
	r.items = append(r.items, appendableItems...)
}

/*
This function fills the randomizer's items slice with the contents of filepath. filepath points to a file with 1 item per line.
The clear parameter lets you decide whether to clear the slice before filling it (rather than append to it)
*/
func (r *Randomizer) Fill(filepath string, clear bool) {
	if clear {
		r.clear()
	}
	content := utils.FileToSlice(filepath)
	for _, c := range content {
		r.Append(c)
	}
}

func (destRand *Randomizer) Combine(srcRand ...*Randomizer) {
	for _, appendableRand := range srcRand {
		destRand.items = slices.Concat(destRand.items, appendableRand.items)
	}

}

// Clears the randomizer slice
func (r *Randomizer) clear() {
	clear(r.items)
}

// Produces a random item from the Randomizer slice.
// Examples show the use of rand.Seed but as of Go 1.20 the runtime seeds the generator automatically.
func (r Randomizer) Random() string {
	n := rand.Intn(len(r.items))
	return r.items[n]
}
