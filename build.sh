#!/bin/bash

C_SCRIPTPATH="$(readlink -f "$0")"
C_SCRIPTDIR="$(dirname "$C_SCRIPTPATH")"
C_THEMEPACK_URL="https://github.com/jimeh/tmux-themepack.git"
C_THEMEPACK_DIR="$(basename "${C_THEMEPACK_URL}" .git)"

readonly C_SCRIPTPATH
readonly C_SCRIPTDIR
readonly C_THEMEPACK_URL
readonly C_THEMEPACK_DIR

function mkstaging_area()
{
  local template
  local retv

  template="/tmp/goproj.XXXXXXXX"
  retv="0"

  STAGING_AREA="$(mktemp -d ${template})"
  retv="$?"
  return "${retv}"
}

mkstaging_area || exit 1
[[ -n "${STAGING_AREA}" ]] || exit 2
[[ -d "${STAGING_AREA}" ]] || exit 3

pushd "${STAGING_AREA}" >/dev/null 2>&1 || exit 4

git clone "${C_THEMEPACK_URL}" || exit 5

pushd "${C_THEMEPACK_DIR}" >/dev/null 2>&1 || exit 6

find src/ -name '*.tmuxtheme' | while read -r src
do
  dest="${C_SCRIPTDIR}/templates/tmux/${src##src/}"
  bin/build-theme "${src}" "${dest}"
done

popd >/dev/null 2>&1 || exit 7
popd >/dev/null 2>&1 || exit 8

rm -rf "${STAGING_AREA}" || exit 9
