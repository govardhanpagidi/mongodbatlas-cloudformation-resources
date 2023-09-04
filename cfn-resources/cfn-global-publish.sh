#!/usr/bin/env bash

set -xe
set -o nounset

regions="${1:-}"
resource="${2:-}"
otherParams="${2:-}"



for region in ${regions}; do
  export AWS_DEFAULT_REGION="${region}"
  export AWS__REGION="${region}"
  ./cfn-publish.sh "${resource}" "${otherParams}"
done
