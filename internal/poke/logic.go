package poke

import "fmt"

const (
	BLUE_VERSION  = "blue"
	GREEN_VERSION = "green"
)

type versionOpts struct {
	Name    string
	MaxID   int
	MinID   int
	IsValid bool
}

func retrieveOptsFromVersion(version string) (versionOpts, error) {
	versions := map[string]versionOpts{
		BLUE_VERSION: {
			Name:    BLUE_VERSION,
			MaxID:   151,
			MinID:   1,
			IsValid: true,
		},
		GREEN_VERSION: {
			Name:    BLUE_VERSION,
			MaxID:   649,
			MinID:   494,
			IsValid: true,
		},
	}

	opts, ok := versions[version]
	if !ok {
		return versionOpts{}, fmt.Errorf("%s is not a valid version", version)
	}

	return opts, nil
}
