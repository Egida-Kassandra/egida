#!/bin/bash

###################################
#
#              EGIDA
#
###################################


# ============ CONSTS ===============
set -e
BASEDIR=$(dirname "$0")

python "$BASEDIR/egida" "$@"