#!/bin/bash

set -eux

go install google.golang.org/protobuf/cmd/protoc-gen-go
go install github.com/mailru/easyjson/easyjson

protoc --go_out=. dropsonde-protocol/events/*.proto
mv github.com/cloudfoundry/sonde-go/events/*.pb.go events
rm -rf github.com

# generate easyjson marshalers/unmarshalers
# use -no_std_marshalers so that using the optimized easyjson marshalers is opt-in
easyjson -all -no_std_marshalers -pkg events
