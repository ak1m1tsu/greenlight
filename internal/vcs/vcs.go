package vcs

import (
	"fmt"
	"runtime/debug"
)

func Version() string {
	var (
		revision string
		modified bool
		time string
	)
	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value
			case "vcs.modified":
				if s.Value == "true" {
					modified = true
				}
			case "vcs.time":
				time = s.Value
			}
		}
	}

	if modified {
		return fmt.Sprintf("%s-%s-dirty", time, revision)
	}

	return fmt.Sprintf("%s-%s", time, revision)
}
