package pathpattern

import (
	"path"
	"strings"
)

// preprocessed URL path pattern
type Pattern struct {
	Original  string   // the original pattern string
	Prefix    string   // pattern with wildcard removed (if any)
	IsWildCard bool     // whether this pattern ends with "*"
}

type PathMatcher struct {
	patterns []Pattern
}

func New(paths []string) *PathMatcher {
	patterns := make([]Pattern, len(paths))
	
	for i, p := range paths {
		patterns[i] = createPattern(p)
	}
	
	return &PathMatcher{patterns: patterns}
}

func createPattern(p string) Pattern {
	cleanPath := path.Clean(p)
	isWildcard := strings.HasSuffix(cleanPath, "*")
	prefix := cleanPath
	
	if isWildcard {
		prefix = strings.TrimSuffix(cleanPath, "*")
	}
	
	return Pattern{
		Original:  cleanPath,
		Prefix:    prefix,
		IsWildCard: isWildcard,
	}
}

func (pm *PathMatcher) Matches(requestPath string) bool {
	cleanPath := path.Clean(requestPath)
	
	for _, pattern := range pm.patterns {
		if pattern.IsWildCard {
			if strings.HasPrefix(cleanPath, pattern.Prefix) {
				return true
			}
		} else {
			if cleanPath == pattern.Original {
				return true
			}
		}
	}
	
	return false
}

func (pm *PathMatcher) AddPattern(pattern string) {
	pm.patterns = append(pm.patterns, createPattern(pattern))
}