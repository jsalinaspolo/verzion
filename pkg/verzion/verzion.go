package verzion

import (
	"fmt"
	"regexp"
	"strconv"
)

// Verzion is a semantic version, allowing suffix.
type Verzion struct {
	Major  int
	Minor  int
	Patch  int
	Suffix string
}

// FromString attempts to parse a Verzion from a string.
func FromString(s string) (Verzion, error) {
	regex := regexp.MustCompile(`^v?([0-9]+)\.([0-9]+)(\.([0-9]+))?(?:\+*\-*)([A-z0-9\-\.\+]*)`)
	if regex.Match([]byte(s)) == false {
		return Verzion{}, fmt.Errorf("'%s' is not a valid Zersion", s)
	}

	parts := regex.FindStringSubmatch(s)
	maj, _ := strconv.Atoi(parts[1])
	min, _ := strconv.Atoi(parts[2])
	patch, _ := strconv.Atoi(parts[4])

	return Verzion{
		Major:  maj,
		Minor:  min,
		Patch:  patch,
		Suffix: parts[5],
	}, nil
}
