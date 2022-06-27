//go:build tools
// +build tools

package main

import (
	_ "github.com/mailru/easyjson/easyjson"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
