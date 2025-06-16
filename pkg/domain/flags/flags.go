package flags

import (
	"errors"
	"strings"
)

// Flags represents the command-line flags
type Flags struct {
	LongFormat bool // -l: use long listing format Maybe invert this to be non list format ?????
	All        bool // -a: show hidden files
}

// ParseFlags parses the command-line arguments and returns the flags and target path
func ParseFlags(args []string) (*Flags, string, error) {
	flags := &Flags{}
	var targetPath string

	if len(args) == 0 {
		return flags, ".", nil
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flagStr := arg[1:]
			for _, char := range flagStr {
				switch char {
				case 'l':
					flags.LongFormat = true
				case 'a':
					flags.All = true
				default:
					return nil, "", errors.New("invalid flag: -" + string(char))
				}
			}
		} else {
			if targetPath == "" {
				targetPath = arg
			} else {
				return nil, "", errors.New("multiple paths specified")
			}
		}
	}

	if targetPath == "" {
		targetPath = "."
	}

	return flags, targetPath, nil
}

// GetTargetPath returns the target path from the arguments
func GetTargetPath(args []string) string {
	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		return "."
	}
	return args[0]
}

