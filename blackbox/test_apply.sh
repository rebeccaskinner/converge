#!/usr/bin/env bash
set -eo pipefail

ROOT=$(pwd)
SOURCE=${1:-"$ROOT"/samples/basic.hcl}

TMP=$(mktemp -d -t converge.apply.XXXXXXXXXX)
function finish {
    rm -rf "$TMP"
}
trap finish EXIT

pushd "$TMP" 2>/dev/null

"$ROOT"/converge apply --local -p "filename=test.txt" -p "message=x" "$SOURCE"

if [ ! -f test.txt ]; then
    echo "test.txt doesn't exist"
    exit 1
fi

if [[ "$(cat test.txt)" != "x" ]]; then
    echo "test.txt doesn't have the right content"
    echo "has '$(cat test.txt)', want 'x'"
    exit 1
fi

popd
