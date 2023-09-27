#!/usr/bin/env bash
set -euo pipefail

my_dir="$( cd "$(dirname "${0}")" && pwd )"
release_dir="$( cd "${my_dir}" && cd ../.. && pwd )"
workspace_dir="$( cd "${release_dir}" && cd .. && pwd )"

# inputs
semver_dir="${workspace_dir}/version-semver"

# outputs
output_dir="${workspace_dir}/out"

get_semver() {
  cat "${semver_dir}/number"
}

make_binname() {
  GOOS="${GOOS:-$(go env GOOS)}"
  GOARCH="${GOARCH:-$(go env GOARCH)}"
  local binname
  binname="davcli-$(get_semver)-${GOOS}-${GOARCH}"

  if [ "${GOOS}" = "windows" ]; then
    echo "${binname}.exe"
  else
    echo "${binname}"
  fi
}

pushd "${release_dir}" > /dev/null
  VERSION=$(get_semver)
  BINNAME=$(make_binname)
  export VERSION
  export BINNAME

  bin/build

  mv "out/${BINNAME}" "${output_dir}/"
popd > /dev/null
