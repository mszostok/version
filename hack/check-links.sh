#!/usr/bin/env bash
# standard bash error handling
set -o nounset # treat unset variables as an error and exit immediately.
set -o errexit # exit immediately when a command fails.
set -E         # needs to be set if we want the ERR trap

CURRENT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly CURRENT_DIR

check::install() {
  npm install --location=global markdown-link-check
}

check::links() {
  find . \
    -type f -regex '.*\.md' \
    -print0 \
    | xargs -n 1 -0 markdown-link-check -q -c "${CURRENT_DIR}/.mlc.config.json"
}


main() {
  check::install
  check::links
}

main
