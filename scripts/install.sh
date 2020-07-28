#!/bin/sh

set -e

# ==============> Download & install egida-role-cis
wget https://github.com/antonioalfa22/egida-role-cis/releases/download/1.0.3/egida-role-cis.zip
unzip egida-role-cis.zip
mv egida-role-cis-master /etc/ansible/roles/egida-role-cis
rm egida-role-cis.zip

# ==============> Download & install egida-role-setup
wget https://github.com/antonioalfa22/egida-role-setup/releases/download/1.0.3/egida-role-setup.zip
unzip egida-role-setup.zip
mv egida-role-setup-master /etc/ansible/roles/egida-role-setup
rm egida-role-setup.zip

# ==============> Download & install egida
wget https://github.com/antonioalfa22/egida/releases/download/1.0.3/egida.zip
unzip egida.zip
cd egida-master

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