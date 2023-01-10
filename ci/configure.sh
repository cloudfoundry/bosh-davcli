#!/usr/bin/env bash
set -euo pipefail

if [[ $(lpass status -q; echo $?) != 0 ]]; then
  echo "Login with lpass first"
  exit 1
fi

fly -t bosh-ecosystem set-pipeline -p bosh-davcli -c ci/pipeline.yml \
  -l <(lpass show -G "davcli concourse secrets" --notes)
