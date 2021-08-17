package version

import "fmt"

const version = "0.0.0"

func Version() string {
	return fmt.Sprintf("v%s", version)
}
