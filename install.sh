#!/bin/bash

###################################
#
#              EGIDA
#
###################################


# ============ CONSTS ===============
set -e
BASEDIR=$(dirname "$0")

pip install -r "$BASEDIR/requirements.txt"
