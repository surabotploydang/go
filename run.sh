#!/usr/bin/env bash

# Export env file var
set -a
[ -f .env ] && . .env
set +a

# Version pattern 1.0.0-build20210120
export VERSION="$(git describe --always --long --dirty)"
export BUILD_TIME=$(date)
export BUILD_COMMIT="$(git rev-parse --short HEAD)"
export MAIN_GO_FILE=./cmd/admin/main.go


# Check input args
if [ $# -eq 0 ]; then
  echo "Require running arguments"
  echo "Usage: " "$0" " {app} {app arguments}"
  echo "Available App: server, web, export"
  echo "example: " "$0" " server start"
  exit 1
fi

# shellcheck disable=SC2145
echo "Running server with arguments: ${@:1}"
go run -ldflags "-X 'main.ProgramName=$APP_NAME' -X 'main.Version=$VERSION' -X 'main.BuildTime=$BUILD_TIME' -X 'main.BuildCommit=$BUILD_COMMIT'" "${MAIN_GO_FILE}" "${@:1}"
