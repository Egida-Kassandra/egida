#!/bin/sh

set -e

# ==============> Update egida-role-cis
rm -rf /etc/ansible/rolesegida-role-cis
wget https://github.com/antonioalfa22/egida-role-cis/releases/download/1.0.3/egida-role-cis.zip
unzip egida-role-cis.zip
mv egida-role-cis-master /etc/ansible/roles/egida-role-cis
rm egida-role-cis.zip
