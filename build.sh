#!/bin/bash
#===============================================================================
#
#         FILE:  build.sh
#
#        USAGE:  build.sh
#
#  DESCRIPTION:  Bash script
#
#      OPTIONS:  ---
# REQUIREMENTS:  ---
#         BUGS:  ---
#        NOTES:  ---
#       AUTHOR:  jvzantvoort (John van Zantvoort), john@vanzantvoort.org
#      COMPANY:  JDC
#      CREATED:  2023-08-04
#
# Copyright (C) 2023 John van Zantvoort
#
#===============================================================================
C_SCRIPTPATH="$(readlink -f "$0")"
C_SCRIPTNAME="$(basename "$C_SCRIPTPATH" .sh)"
C_FACILITY="local0"

readonly C_SCRIPTPATH
readonly C_SCRIPTNAME
readonly C_FACILITY

declare -xr LANG="C"

function logging()
{
  local priority="$1"; shift
  logger -p "${C_FACILITY}.${priority}" -i -s -t "${C_SCRIPTNAME}" -- "${priority} $*"
}

function logging_err()   { logging "err" "$@";   }
function logging_info()  { logging "info" "$@";  }
function logging_warn()  { logging "warn" "$@";  }
function logging_debug() { logging "debug" "$@"; }

function script_exit()
{
  local string="$1"
  local retv="${2:-0}"
  if [ "$retv" = "0" ]
  then
    logging_info "$string"
  else
    logging_err "$string"
  fi
  exit "$retv"
}

function pathmunge()
{
  [ -d "$1" ] || return

  if echo "$PATH" | grep -E -q "(^|:)$1($|:)"
  then
    return
  fi

  if [ "$2" = "after" ]
  then
      PATH=$PATH:$1
  else
      PATH=$1:$PATH
  fi
}

function gitroot() { git rev-parse --show-toplevel; }
function listsubdirs()
{
  local root

  root="$(gitroot)"
  find "${root}" -maxdepth 1 -mindepth 1 -type d \
    -not -name vendor \
    -not -name cmd \
    -not -name ref | \
    while read -r target
    do
      find "$target" -name '*.go' -printf "%h\n"
    done|sort -u
}

function apidocsdir() { printf "%s/docs/api" "$(gitroot)"; }

function buildapidoc()
{
  local root
  local outputdir

  root="$(gitroot)"
  outputdir="$(apidocsdir)"

  pushd "${root}" >/dev/null 2>&1 || exit 1
  listsubdirs | sed "s,${root}\/,," | while read -r dirn
  do
    godocdown "${dirn}" > "${outputdir}/${dirn}.md"
  done
  popd >/dev/null 2>&1 || exit 2

}

function usage()
{
  local indent
  indent="${C_SCRIPTNAME//?/ }    "
  printf "%s [help|doc]\n\n" "${C_SCRIPTNAME}.sh"
  printf "%shelp - print help\n" "${indent}"
  printf "%sdoc  - generate API doc\n" "${indent}"
  printf "\n\n"
  exit
}

#------------------------------------------------------------------------------#
#                                    Main                                      #
#------------------------------------------------------------------------------#

ACTION="${1:-help}"

case "${ACTION}" in
  doc)
    buildapidoc
    ;;
  help) usage ;;
  *) usage ;;
esac

#------------------------------------------------------------------------------#
#                                  The End                                     #
#------------------------------------------------------------------------------#
