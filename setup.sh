#!/bin/bash

go run $(pwd)/cmd/prepare/main.go block
go run $(pwd)/cmd/prepare/main.go mute
go run $(pwd)/cmd/prepare/main.go follower
go run $(pwd)/cmd/prepare/main.go followee