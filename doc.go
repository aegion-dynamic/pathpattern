/*
Package pathpattern provides efficient URL path pattern matching with wildcard support.

It supports exact path matching and wildcard patterns,
with protection against path traversal attacks.

Patterns are preprocessed during initialization for better performance.

Basic usage:

    matcher := pathpattern.New([]string{
        "/public/*",
        "/api/health",
    })

    if matcher.Matches("/public/file.txt") {
        // handle matched path
    }

The library automatically handles path cleaning and normalization,
making it safe against path traversal attempts.
*/

package pathpattern