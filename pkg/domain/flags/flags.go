package flags

import (
	"errors"
	"strings"
)

// Flags represents the command-line flags
type Flags struct {
	LongFormat bool // -l: use long listing format
	All        bool // -a: show hidden files
	HumanReadable bool // -h: human-readable sizes
}

// ParseFlags parses the command-line arguments and returns the flags and target path
func ParseFlags(args []string) (*Flags, string, error) {
	flags := &Flags{}
	var targetPath string

	// If no args, return default flags and current directory
	if len(args) == 0 {
		return flags, ".", nil
	}

	// Process all arguments
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			// Process flags
			flagStr := arg[1:]
			for _, char := range flagStr {
				switch char {
				case 'l':
					flags.LongFormat = true
				case 'a':
					flags.All = true
				case 'h':
					flags.HumanReadable = true
				default:
					return nil, "", errors.New("invalid flag: -" + string(char))
				}
			}
		} else {
			// This is a path argument
			if targetPath == "" {
				targetPath = arg
			} else {
				return nil, "", errors.New("multiple paths specified")
			}
		}
	}

	// If no path was specified, use current directory
	if targetPath == "" {
		targetPath = "."
	}

	return flags, targetPath, nil
}

// GetTargetPath returns the target path from the arguments
func GetTargetPath(args []string) string {
	// If no args or first arg is a flag, use current directory
	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		return "."
	}
	return args[0]
} 