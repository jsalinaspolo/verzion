package verzion

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Verzion is a semantic version, allowing suffix.
type Verzion struct {
	Major    int
	Minor    int
	Patch    int
	Metadata string
}

// Zero is the Zero Verzion
var Zero Verzion

// Less returns true if the receiver Verzion is less than a given Verzion.
func (v Verzion) Less(cmp Verzion) bool {
	if v.Major == cmp.Major {
		if v.Minor == cmp.Minor {
			return v.Patch <= cmp.Patch
		}
		return v.Minor <= cmp.Minor
	}
	return v.Major <= cmp.Major
}

// Equal checks if a version is equal to the receiver.
func (v Verzion) Equal(cmp Verzion) bool {
	if v.Major == cmp.Major && v.Minor == cmp.Minor && v.Patch == cmp.Patch {
		return true
	}

	return false
}

// String prints the Verzion to a string.
func (z Verzion) String() string {
	s := fmt.Sprintf("%d.%d.%d", z.Major, z.Minor, z.Patch)
	if len(strings.TrimSpace(z.Metadata)) > 0 {
		return s + "+" + z.Metadata
	}
	return s
}

// FromFile parses a VERSION file into a Verzion.
func FromFile(path string) (Verzion, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Verzion{}, err
	}

	ver, err := FromString(strings.TrimSpace(string(content)))
	if err != nil {
		return Verzion{}, err
	}

	return ver, nil
}

// FromString attempts to parse a Verzion from a string.
func FromString(s string) (Verzion, error) {
	regex := regexp.MustCompile(`^v?([0-9]+)\.([0-9]+)(\.([0-9]+))?(?:\+*\-*)([A-z0-9\-\.\+]*)`)
	if regex.Match([]byte(s)) == false {
		return Verzion{}, fmt.Errorf("'%s' is not a valid Verzion", s)
	}

	parts := regex.FindStringSubmatch(s)
	maj, _ := strconv.Atoi(parts[1])
	min, _ := strconv.Atoi(parts[2])
	patch, _ := strconv.Atoi(parts[4])

	return Verzion{
		Major:    maj,
		Minor:    min,
		Patch:    patch,
		Metadata: parts[5],
	}, nil
}
