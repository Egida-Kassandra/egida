#!/bin/sh

set -e

# ==============> Download & install egida-role-cis
wget https://github.com/antonioalfa22/egida-role-cis/releases/download/2.0.0/egida-role-cis.zip
unzip egida-role-cis.zip
mv egida-role-cis /etc/ansible/roles/egida-role-cis
rm egida-role-cis.zip

# ==============> Download & install egida-role-setup
wget https://github.com/antonioalfa22/egida-role-setup/releases/download/2.0.0/egida-role-setup.zip
unzip egida-role-setup.zip
mv egida-role-setup-master /etc/ansible/roles/egida-role-setup
rm egida-role-setup.zip

# Create egida vars location
mkdir /etc/egida
mkdir /etc/egida/vars

# ==============> Download & install egida
wget https://github.com/antonioalfa22/egida/releases/download/2.0.0/egida.zip
unzip egida.zip
chmod +x egida
mv egida /usr/local/bin/egida
mv hostsgroups /etc/egida/hostsgroups
mv vars_template.yml /etc/egida/vars/vars_template.yml
echo "[local]" > /etc/ansible/hosts
echo "localhost" >> /etc/ansible/hosts
