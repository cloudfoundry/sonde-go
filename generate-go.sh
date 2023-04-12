#!/bin/bash

set -eux

go install google.golang.org/protobuf/cmd/protoc-gen-go

rm -rf events/*.pb.go

protoc --go_out=. dropsonde-protocol/events/*.proto
mv github.com/cloudfoundry/sonde-go/events/*.pb.go events
rm -rf github.com
