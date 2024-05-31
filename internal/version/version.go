package version

import (
	_ "embed"
)

//go:generate bash get_version.sh
//go:embed version.txt
var version string

// Function is an action, depends on when it is called.
func GetGrdnrVersion() string {
	if version == "" {
		return "No version found, anarchy time!"
	}
	return version
}
