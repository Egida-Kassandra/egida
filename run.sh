#!/bin/sh

set -e

# Get role cis
git clone https://github.com/antonioalfa22/egida-role-cis
mv egida-role-cis /etc/ansible/roles

# Install requirements
pip3 install -r requirements.txt

# Create egida vars location

mkdir /etc/egida
mkdir /etc/egida/lib
mkdir /etc/egida/custom
mkdir /etc/egida/common

# Move egida files
mv ./custom /etc/egida/

# Install egida
python3 setup.py install
