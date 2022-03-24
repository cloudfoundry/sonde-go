//go:build tools
// +build tools

package main

import (
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/mailru/easyjson/easyjson"
)
