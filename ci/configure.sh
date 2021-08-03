#!/bin/bash

fly -t bosh-ecosystem set-pipeline -p bosh-davcli -c ci/pipeline.yml \
  -l <(lpass show -G "davcli concourse secrets" --notes) \
  -l <(lpass show --notes "pivotal-tracker-resource-keys")
