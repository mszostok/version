#!/usr/bin/env bash

# standard bash error handling
set -o nounset # treat unset variables as an error and exit immediately.
set -o errexit # exit immediately when a command fails.
set -E         # needs to be set if we want the ERR trap

readonly YELLOW='\033[1;33m'
readonly NC='\033[0m' # No Color

CURRENT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
REPO_ROOT_DIR=$(cd "${CURRENT_DIR}/../.." && pwd)
readonly CURRENT_DIR
readonly REPO_ROOT_DIR

setup() {
	cd "$CURRENT_DIR" || true
	profile=$1
	echo -e "\033]50;SetProfile=$profile\a"
	osascript resize_window.scpt
	export KUBECONFIG=''
	clear
}

capture() {
	program=$1
	ver=$2
	clear

	cd "$REPO_ROOT_DIR/example" || exit
	go install -ldflags "-X 'github.com/mszostok/version.buildDate=$(date)' -X 'github.com/mszostok/version.version=v0.6.0'" ./$program
	cd "$HOME" || exit

	# shellcheck disable=SC2059
	printf "▲ ${YELLOW}$program${NC} $ver\n"

	# shellcheck disable=SC2086
	$program $ver

	local filename="${REPO_ROOT_DIR}/docs/assets/examples/screen-$program-${ver// /_}.png"
	rm -f "$filename" || true

	# only term: screencapture -ol$(osascript -e 'tell app "iTerm" to id of window 1') test.png
	screencapture -x -R0,25,1285,650 "$filename"
}

main() {
	# Big
	setup "version-cmd"

	capture "plain" ""
	capture "cobra" "version"
	capture "printer" ""
	capture "printer" "-oyaml"

	capture "custom-formatting" ""

	# Medium
	setup "custom-pretty-cmd"
	sleep 1
	capture "upgrade-notice-sub-cmd" "version check"

	capture "upgrade-notice-box" "version"

	capture "upgrade-notice" "version"
	capture "upgrade-notice" "version -ojson"

	capture "custom-layout" ""
	capture "custom-renderer" ""

	# Small
	setup "help-cmd"
	sleep 1

	capture "cobra" "version -h"
}

main
