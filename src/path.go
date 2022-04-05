// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

import (
	"path/filepath"
	"strings"
)

// HasPathPrefix reports whether the slash-separated path s
// begins with the elements in prefix.
func HasPathPrefix(s, prefix string) bool {
	if len(s) == len(prefix) {
		return s == prefix
	}
	if prefix == "" {
		return true
	}
	if len(s) > len(prefix) {
		if prefix[len(prefix)-1] == '/' || s[len(prefix)] == '/' {
			return s[:len(prefix)] == prefix
		}
	}
	return false
}

// HasFilePathPrefix reports whether the filesystem path s
// begins with the elements in prefix.
func HasFilePathPrefix(s, prefix string) bool {
	sv := strings.ToUpper(filepath.VolumeName(s))
	pv := strings.ToUpper(filepath.VolumeName(prefix))
	s = s[len(sv):]
	prefix = prefix[len(pv):]
	switch {
	default:
		return false
	case sv != pv:
		return false
	case len(s) == len(prefix):
		return s == prefix
	case prefix == "":
		return true
	case len(s) > len(prefix):
		if prefix[len(prefix)-1] == filepath.Separator {
			return strings.HasPrefix(s, prefix)
		}
		return s[len(prefix)] == filepath.Separator && s[:len(prefix)] == prefix
	}
}

// TrimFilePathPrefix returns s without the leading path elements in prefix.
// If s does not start with prefix (HasFilePathPrefix with the same arguments
// returns false), TrimFilePathPrefix returns s. If s equals prefix,
// TrimFilePathPrefix returns "".
func TrimFilePathPrefix(s, prefix string) string {
	if !HasFilePathPrefix(s, prefix) {
		return s
	}
	if len(s) == len(prefix) {
		return ""
	}
	trimmed := s[len(prefix):]
	if trimmed[0] != filepath.Separator {
		// executed when either the prefix is root
		// or it includes a trailing separator.
		return trimmed
	}
	return trimmed[1:]
}