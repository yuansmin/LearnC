package stack

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	cases := map[string]string{
		"/home/":                "/home",
		"/../":                  "/",
		"/home//foo/":           "/home/foo",
		"/a/./b/../../c/":       "/c",
		"/a/../../b/../c//.//":  "/c",
		"/a//b////c/d//././/..": "/a/b/c",
		"/...":                  "/...",
		"/.../":                 "/...",
		"/...///":               "/...",
		"/..hidden":             "/..hidden",
	}
	for k, v := range cases {
		if simplifyPath(k) != v {
			t.Errorf("%s , %s, got: %s", k, v, simplifyPath(k))
		}
	}
}
