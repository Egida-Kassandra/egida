#!/bin/sh


set -e

# ==============> Download & install egida-role-cis
wget https://github.com/antonioalfa22/egida-role-cis/releases/download/1.0.3/egida-role-cis.zip
unzip egida-role-cis.zip -d /etc/ansible/roles
rm egida-role-cis.zip

# ==============> Download & install egida
wget https://github.com/antonioalfa22/egida/releases/download/1.0.3/egida.zip
unzip egida.zip
cd egida

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