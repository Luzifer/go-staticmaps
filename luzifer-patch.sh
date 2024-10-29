#!/usr/bin/env bash
set -euo pipefail

# Patch import path
find . -name 'go.mod' -or -name '*.go' | xargs sed -i 's@github.com/flopp/go-staticmaps@github.com/Luzifer/go-staticmaps@'
# Tidy up code-files after patch
find . -name '*.go' | xargs gofumpt -w

# Update deps and clean mod-file
go get -u && go mod tidy
