package template

// Int generates a random integer using the full range of the randomizer.
func (nr *Randomizer) Int() int {
	return nr.Rand.Int()
}
