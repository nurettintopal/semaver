package semaver

import (
	"fmt"
	"regexp"
	"strconv"
)

// MAJOR.MINOR.PATCH[-PRERELEASE][+BUILD]
type Version struct {
	major, minor, patch int
	prerelease          string
	build               string
}

func ParseVersion(v string) (Version, error) {
	regex := `^(\d+)\.(\d+)\.(\d+)(?:-([a-zA-Z0-9.-]+))?(?:\+([a-zA-Z0-9.-]+))?$`
	re := regexp.MustCompile(regex)

	match := re.FindStringSubmatch(v)
	if match == nil {
		return Version{}, fmt.Errorf("invalid version format")
	}

	major, _ := strconv.Atoi(match[1])
	minor, _ := strconv.Atoi(match[2])
	patch, _ := strconv.Atoi(match[3])

	prerelease := match[4]
	build := match[5]

	return Version{major, minor, patch, prerelease, build}, nil
}

func CompareVersions(v1, operator, v2 string) (bool, error) {
	version1, err := ParseVersion(v1)
	if err != nil {
		return false, err
	}
	version2, err := ParseVersion(v2)
	if err != nil {
		return false, err
	}

	switch operator {
	case "=":
		return version1.Equals(version2), nil
	case "<":
		return version1.LessThan(version2), nil
	case ">":
		return version1.GreaterThan(version2), nil
	case "<=":
		return version1.LessThanOrEqual(version2), nil
	case ">=":
		return version1.GreaterThanOrEqual(version2), nil
	default:
		return false, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func (v Version) Equals(other Version) bool {
	return v.major == other.major && v.minor == other.minor && v.patch == other.patch && v.prerelease == other.prerelease
}

func (v Version) LessThan(other Version) bool {
	if v.major < other.major {
		return true
	}
	if v.major > other.major {
		return false
	}
	if v.minor < other.minor {
		return true
	}
	if v.minor > other.minor {
		return false
	}
	if v.patch < other.patch {
		return true
	}
	if v.patch > other.patch {
		return false
	}
	if v.prerelease != "" && other.prerelease == "" {
		return true
	}
	if v.prerelease == "" && other.prerelease != "" {
		return false
	}
	if v.prerelease < other.prerelease {
		return true
	}
	if v.prerelease > other.prerelease {
		return false
	}
	return false
}

func (v Version) GreaterThan(other Version) bool {
	return !v.LessThan(other) && !v.Equals(other)
}

func (v Version) LessThanOrEqual(other Version) bool {
	return v.LessThan(other) || v.Equals(other)
}

func (v Version) GreaterThanOrEqual(other Version) bool {
	return v.GreaterThan(other) || v.Equals(other)
}
