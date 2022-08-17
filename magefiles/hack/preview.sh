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
readonly ASSET_EXAMPLES_DIR="${REPO_ROOT_DIR}/docs/assets/examples"

setup() {
	cd "$CURRENT_DIR" || true
	profile=$1
	echo -e "\033]50;SetProfile=$profile\a"
	osascript preview_window.scpt
	export KUBECONFIG=''
	clear
}

capture() {
	cmd=$1
	filename=$2
	clear

	# shellcheck disable=SC2059
	printf "▲ ${YELLOW}gimme${NC} version\n"

	$cmd

	rm -f "$filename"

	# shellcheck disable=SC2155
	local filepath="${ASSET_EXAMPLES_DIR}/$filename"
	screencapture -ol$(osascript -e 'tell app "iTerm" to id of window 1') "$filepath"
}

main() {
	setup "custom-pretty-cmd"
	sleep 1

	sudo gdate -s "-4 months"
	capture "${REPO_ROOT_DIR}/magefiles/preview" "screen-preview-0.png"
	sudo gdate -s "+4 months"
	capture "${REPO_ROOT_DIR}/magefiles/hack/note.sh" "screen-preview-1.png"
	capture "${REPO_ROOT_DIR}/magefiles/preview" "screen-preview-2.png"

}

main
