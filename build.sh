#!/usr/bin/env bash
# Cross-compiles For-MAC.exe (Windows x64) from Linux/macOS, or builds natively on Windows (Git Bash).
set -euo pipefail
cd "$(dirname "$0")"
# Regenerate icon/manifest/version resources only if go-winres is installed;
# the committed rsrc_windows_amd64.syso is used otherwise.
if command -v go-winres >/dev/null 2>&1; then
  go-winres make --arch amd64
fi
go test ./internal/...
mkdir -p dist
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
  go build -trimpath -ldflags "-s -w -H windowsgui" -o dist/For-MAC.exe .
echo "Built dist/For-MAC.exe"
