//go:build tools
// +build tools

package tools

// for more about tools.go please visit
// https://go.dev/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

import (
	_ "github.com/hexdigest/gowrap/cmd/gowrap"
)
