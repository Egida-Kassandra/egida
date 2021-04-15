#!/bin/sh

set -e

mkdir -p /etc/ansible/roles

apt install ansible -y
apt install -y unzip
apt install sshpass -y

# ==============> Download & install egida-role-cis
wget https://github.com/egida-kassandra/egida-role-cis/releases/download/2.0.0/egida-role-cis.zip
unzip egida-role-cis.zip
mv egida-role-cis /etc/ansible/roles/egida-role-cis
rm egida-role-cis.zip

# ==============> Download & install egida-role-setup
wget https://github.com/egida-kassandra/egida-role-setup/releases/download/2.0.0/egida-role-setup.zip
unzip egida-role-setup.zip
mv egida-role-setup /etc/ansible/roles/egida-role-setup
rm egida-role-setup.zip

# Create egida vars location
mkdir /etc/egida
mkdir /etc/egida/vars

# ==============> Download & install egida
wget https://github.com/egida-kassandra/egida/releases/download/2.0.0/egida.zip
unzip egida.zip
cd build
chmod +x egida
chmod 777 ansible.cfg
mv egida /usr/local/bin/egida
mv hostsgroups /etc/egida/hostsgroups
mv vars_template.yml /etc/egida/vars/vars_template.yml
mv ansible.cfg /etc/ansible/ansible.cfg
echo "[local]" > /etc/ansible/hosts
echo "localhost" >> /etc/ansible/hosts
cd ..
rm -rf build
rm egida.zip