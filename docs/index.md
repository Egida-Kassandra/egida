# Egida Project

<!-- PROJECT SHIELDS -->
[![EGIDA VERSION](https://img.shields.io/badge/egida-v1.0.0-blue?style=for-the-badge&logo=ansible&color=ff69b4)](https://github.com/antonioalfa22/egida)
[![GitHub license](https://img.shields.io/badge/license-Apache-blue?style=for-the-badge)](https://github.com/antonioalfa22/egida/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/badge/release-v.1.0.0-yellowgreen?style=for-the-badge)](https://github.com/antonioalfa22/egida/releases)

<!-- PROJECT LOGO -->

<br />
<div align="center">
  <a href="https://github.com/antonioalfa22/egida">
    <img src="img/logo.png" alt="Logo" width="180" height="180">
  </a>

  <p align="center">
    <br />
    <a href="https://github.com/antonioalfa22/egida">View Source</a>
    ·
    <a href="https://github.com/antonioalfa22/egida/issues">Report Bug</a>
    ·
    <a href="https://github.com/antonioalfa22/egida/issues">Request Feature</a>
  </p>
</div>


<!-- TABLE OF CONTENTS -->
## Table of Contents

* [Overview](#overview)
* [Installation](#installation)
  * [Prerequisites](#prerequisites)
  * [Download and install](#download-and-install)
* [Getting Started](#getting-started)
  * [Environment SetUp](#environment-setup)
  * [Add Host](#add-host)
  * [Variables](#variables)
* [Hardening](#hardening)
  * [All CIS Benchmarks](#all-cis-benchmarks)
  * [CIS Points](#cis-points)
  * [CIS Sections](#cis-sections)
  * [CIS Controls](#cis-controls)
* [Getting Info](#getting-info)
  * [Lynis Score](#lynis-score)
  * [Machine Info](#machine-info)
* [License](#license)
* [Contact](#contact)

<!-- Overview -->
---
## Overview

The Egida project is a server orchestration system that allows to perform and deploy security configurations 
(custom control lists) over a machine infrastructure. These security configurations can shield and protect those 
servers by implementing the desired security measures depending on the server profile. Controls are sourced from the 
[CIS Benchmarks](https://www.cisecurity.org/cis-benchmarks/), and we also need to obtain system information about each 
of the deployed servers to ensure proper deployment.

To achieve that, Egida is built using a microservices-based architecture composed of the following three modules:


- **egida**: This is the main module, in charge of providing the communication interfaces with the user, as well as the 
         process of the specific domain language called Aspida. Using this module, the user can define 
         the different profiles to work with and the actions to be performed.
         
- **egida-roles**: This module contains the definition of the Ansible roles that contains the actions corresponding to 
        the security controls that are defined for each profile that a machine may have assigned. These actions can be 
        either hardening operations (CIS Benchmarks) or setup actions to prepare that machine so it can be 
        used correctly by Egida.
        
- **egida-api**: The functionality of this module is to provide information of each machine to be used by Egida. 
        This information can be varied: the services that are currently running, the operating system version or the 
        score obtained with the [Lynis](https://cisofy.com/lynis/) tool… any kind of information that we determine it is 
        interesting to better deploy any security control. 


![Egida Network](img/esquema.png)

<!-- Installation -->
---
## Installation

### Prerequisites

Egida v1.0.0 requires the following software to be installed on the master node:

> Currently, Egida v1.0.0 needs an Ubuntu 18.04 LTS OS.

- **Ansible >2.8**: [Install Ansible](https://docs.ansible.com/ansible/latest/installation_guide/index.html)

```commandline
sudo apt update
sudo apt install software-properties-common
sudo apt-add-repository --yes --update ppa:ansible/ansible
sudo apt install ansible
```

- **Python 3.x**: [Install Python 3](https://www.python.org/downloads/)
```commandline
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository ppa:deadsnakes/ppa
sudo apt install python3.7
```

> At this point, Python 3.7 is installed on your Ubuntu system and ready to be used. You can verify it by typing
> `python3.7 --version`

### Download and install

Download and install

<!-- Getting Started -->
---
## Getting Started

Getting Started

### Environment SetUp

Environment SetUp

### Add Host

Add Host

### Variables

Variables


<!-- Hardening -->
---
## Hardening

Hardening

### All CIS Benchmarks
All

### CIS Points
CIS Points

### CIS Sections
CIS Sections

### CIS Controls
CIS Controls


<!-- Getting info -->
---
## Getting Info

Getting Info

### Lynis Score
Lynis Score

### Machine Info
Machine Info


<!-- LICENSE -->
---
## License

Distributed under the Apache 2.0 License. See `LICENSE` for more information.

<!-- CONTACT -->
---
## Contact

Authors:

* [Antonio Payá González](https://antoniopg.tk)
* [Alba Cotarelo Tuñón](https://antoniopg.tk)
* [Jose Manuel Redondo Lopez](http://orcid.org/0000-0002-0939-0186)

Project Link: [https://github.com/antonioalfa22/egida](https://github.com/antonioalfa22/egida)


