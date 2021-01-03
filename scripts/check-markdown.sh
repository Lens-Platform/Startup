#!/usr/bin/env bash
set -euo pipefail

# check-markdown.sh
#
# SUMMARY
#
#   Checks the markdown format within the Blackspace repo.
#   This ensures that markdown is consistent and easy to read across the
#   entire Blackspace repo.

scripts/node_modules/.bin/markdownlint \
  --config scripts/.markdownlintrc \
  --ignore scripts/node_modules \
  .
