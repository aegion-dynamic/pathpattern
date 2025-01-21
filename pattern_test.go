package pathpattern

import "testing"

func TestPatternMatching(t *testing.T) {
	tests := []struct {
		name     string
		patterns []string
		path     string
		want     bool
	}{
		{
			name:     "exact match",
			patterns: []string{"/api/health"},
			path:     "/api/health",
			want:     true,
		},
		{
			name:     "wildcard match",
			patterns: []string{"/public/*"},
			path:     "/public/file.txt",
			want:     true,
		},
		{
			name:     "no match",
			patterns: []string{"/public/*"},
			path:     "/private/file.txt",
			want:     false,
		},
		{
			name:     "path traversal attempt",
			patterns: []string{"/public/*"},
			path:     "/public/../private/secret",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := New(tt.patterns)
			got := matcher.Matches(tt.path)
			if got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddPattern(t *testing.T) {
	matcher := New([]string{"/api/*"})
	
	// initial pattern should match
	if !matcher.Matches("/api/users") {
		t.Error("Expected /api/users to match initially")
	}
	
	// new pattern shouldn't match yet
	if matcher.Matches("/public/file.txt") {
		t.Error("Expected /public/file.txt to not match initially")
	}
	
	// add new pattern
	matcher.AddPattern("/public/*")
	
	// now it should match
	if !matcher.Matches("/public/file.txt") {
		t.Error("Expected /public/file.txt to match after adding pattern")
	}
}