#!/bin/sh

CGO_ENABLED=0 go build --ldflags="-s -w --buildid="
strip pwdc
