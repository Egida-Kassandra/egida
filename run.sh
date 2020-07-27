#!/bin/sh

set -e

# Install requirements
pip3 install -r requirements.txt

# Create egida vars location

mkdir /etc/egida
mkdir /etc/egida/lib
mkdir /etc/egida/custom
mkdir /etc/egida/common

# Move egida files
mv ./etc/common /etc/egida/
mv ./etc/custom /etc/egida/

# Install egida
python3 setup.py install