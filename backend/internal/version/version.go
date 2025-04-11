package version

import (
	"fmt"
	"time"
)

var (
	// ReleaseVersion
	ReleaseVersion string
	// BuildTS
	BuildTS string
	// GitHash
	GitHash string
	// GitBranch
	GitBranch string
	// GoVersion
	GoVersion string
)

func Version() string {
	return fmt.Sprintf("ReleaseVersion=%s  BuildTS=%s GitHash=%s GitBranch=%s GoVersion=%s",
		ReleaseVersion, BuildTS, GitHash, GitBranch, GoVersion)
}

func WebVersion() string {
	ts, err := time.Parse("2006-01-02 15:04:05", BuildTS)
	if err != nil {
		return ReleaseVersion
	}
	tsStr := ts.Format("20060102")
	return fmt.Sprintf("%s-%s", ReleaseVersion, tsStr)
}
