#!/usr/bin/env bash
set -euo pipefail

fly -t bosh-ecosystem set-pipeline -p bosh-davcli -c ci/pipeline.yml
