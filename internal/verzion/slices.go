package verzion

// Slice is a sortable slice of Verzions.
type Slice []Verzion

// Swap swaps two Verzions in a Slice.
func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Len gives the length of a Slice.
func (s Slice) Len() int {
	return len(s)
}

//// Less returns true if the Verzion at s[i] is less than s[j]
func (s Slice) Less(i, j int) bool {
	return s[i].Less(s[j])
}
