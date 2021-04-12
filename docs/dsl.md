# Aspida Domain Specific Language

<!-- TABLE OF CONTENTS -->
## Table of Contents

* [Overview](#overview)
* [Your first Program](#your-first-program)
* [Blocks](#blocks)
  * [MAIN](#main)
  * [HOST](#host)
  * [TASKS](#tasks)
    * [sections](#sections)
    * [controls](#controls)
    * [points](#points)
    * [exclusions](#exclusions)
    * [tags](#tags)
  * [VARS](#vars)
* [Conditional Statements](#conditional-statements)
  * [If](#if)
  * [Else](#else)
  * [Elif](#elif)
* [Getting Host Info](#getting-host-info)


<!-- Overview -->

---

## Overview

Egida has a Domain Specific Language (DSL) called **Aspida** that allows customized hardening scripts that depend on the value of properties and attributes of the target system.

 For Egida to use a script written in Aspida we must use the **compile** option:

```commandline
sudo egida compile -f filename
```

> Example: `sudo egida compile -f /examples/example_cond.aspida`

**-f, --file**:  Aspida File to compile
> Example: --file /examples/example_cond.aspida

<!-- Your first Program -->

---

## Your First Program

To create our first script in Egida we have to create a new file. To improve the identification of these programs, the extension .aspida can be added to them.

An Aspida program is composed of four blocks:

* **MAIN**: Main program information, such as connection type (SSH or local)
* **HOST**: Name of the Egida Hosts group on which the script is to be applied
* **TASKS**: Main content of the Script, defines the behavior of the program (tasks to be performed),
* **VARS**: (Optional) Variables of the tasks to be executed

Example program:

```golang
MAIN : {
    name: "Hello World";
    connection: local;
    description: "Hello World description";
}

HOST : "localhost";

TASKS : {
    points : ["1.1", "1.2"];
    sections: ["1.1"];
    exclusions : ["1.3"];
}

VARS : {
    "user_ssh": "antonio";
    "nameservers": ["8.8.8.8", "8.8.4.4"];
}
```

<!-- Blocks -->

---

## Blocks

### MAIN

The main block may optionally contain the following attributes:

* **name**: Name of the script
* **description**: Description of the script
* **connection**: Method of connection to the target hosts (local or ssh)

```golang
MAIN : {
    name: "Hello World";
    connection: ssh;
    description: "Hello World description";
}
```

### HOST

The HOST block contains a single value that corresponds to the name of the Egida host group on which you want to act.

> If you want to add or remove a Host group you must use the Egida [add-group](index.md#add-group) option.

```golang
HOST : "bbdd_servers";
```

### TASKS

The TASKS block specifies which tasks are to be performed during script execution. 

The following actions can be defined:

#### **sections**

**CIS Benchmarks** sections. The complete list of sections is as follows:

* 1.1-Filesystem Configuration
* 1.2-Console Software Updates
* 1.3-Configure sudo
* 1.4-Filesystem Integrity
* 1.5-Secure Boot Settings
* 1.6-Additional Process Hardering
* 1.7-Mandatory Access Control
* 1.8-Warning Banners
* 1.9-Updates
* 2.1-Initd Services
* 2.2-Special Purpose Services
* 2.3-Service Clients
* 3.1-Network Parameters Host
* 3.2-Network Parameters Host and Router
* 3.3-TCP Wrappers
* 3.4-Uncommon Network Protocols
* 3.5-Firewall Configuration
* 3.6-Wireless
* 3.7-Disable IPv6
* 4.1-Configure System Accounting (auditd)
* 4.2-Configure logging
* 4.3-Ensure logrotate is configured
* 5.1-Configure cron
* 5.2-SSH Server configuration
* 5.3-Configure PAM
* 5.4-User accounts and environment
* 5.5-Ensure root login is restricted to system console
* 5.6-Ensure access to the su command is restricted
* 6.1-System file permisions
* 6.2-User and Group Settings

```golang
TASKS : {
    sections: ["1.1"];
}
```

#### **controls**

**CIS Benchmarks** controls. The complete list of controls is as follows:

* 2.-Inventory and Control of Software Assets
* 3.-Continuous Vulnerability Management
* 4.-Controlled Use of Administrative Privileges
* 5.-Secure Configuration for Hardware and Software on Mobile Devices, Laptops, Workstations and Servers
* 6.-Maintenance, Monitoring and Analysis of Audit Logs
* 8.-Malware Defenses
* 9.-Limitation and Control of Network Ports, Protocols and Services
* 13.-Data Protection
* 14.-Controlled Access Based on the Need to Know
* 16.-Account Monitoring and Control

```golang
TASKS : {
    controls: ["1.1"];
}
```

#### **points**

**CIS Benchmarks** points.

```golang
TASKS : {
    points: ["1.1.1.1"];
}
```

#### **exclusions**

**CIS Benchmarks** points exclusions.

```golang
TASKS : {
    exclusions: ["1.1.1.2"];
}
```

#### **tags**

Allows to execute all the tasks that contain a specific tag. 

> For example, it allows to execute all tasks related to SSH. 

```golang
TASKS : {
    tags: ["ssh"];
}
```

### VARS

The VARS block contains the list of values for the variables to be used during the hardening process.

```golang
VARS : {
    "user_ssh": "antonio";
    "password": {
        "max_days": 365;
        "min_days": 7;
        "warn_age": 7;
        "inactive": 30;
    };
    "nameservers": ["8.8.8.8", "8.8.4.4"];
}
```

In case a task is executed that requires a variable, if this variable has not been defined in the VARS block, its default value will be used.

The complete list of variables and their default values is shown below:

```yaml
# defaults file for cis

###############################################
# Values which modify the behaviour of the role
###############################################

run_all_level_1: true    # Whether Level 1 of the benchmark should be applied
run_all_level_2: true    # Whether Level 2 of the benchmark should be applied

# extras: true              # Check if want extras

cis_level_1_exclusions: []         # A list of Level 1 recommendations to exclude (i.e. ['1.1.1.1'])
cis_level_2_exclusions: []         # A list of Level 2 recommendations to exclude


###############################################
# Check specific values which can be overridden
###############################################

# ======== 1. Initial Setup ===================

# 1.3.2 AIDE cron settings

aide_cron:
  cron_user: root
  cron_file: /etc/crontab
  aide_job: '/usr/bin/aide.wrapper --check'
  aide_minute: 0
  aide_hour: 5
  aide_day: '*'
  aide_month: '*'
  aide_weekday: '*'

# 1.4.2 GRUB Password
grub_pass: antonio

# 1.4.3 root Password
root_pass: antonio

# ======== 3. Network configuration ===================

# 3.4.2 Host allow
host_allow:
  - "10.0.0.0/255.0.0.0"
  - "172.16.0.0/255.240.0.0"
  - "192.168.0.0/255.255.0.0"

# 3.5.2.1 UFW

ufw_ports_allow: ['22']
ufw_deny_outgoing: false

# ======== 4. Logging and auditing ===================
default_auditd: true  # Copy auditd template


# ======== 5. SSH Server Configuration ===================

sshd_access:
  ssh_port: 372
  allowusers: antonio
  # allowgroups: systems dba
  # denyusers:
  # denygroups:

# 5.3.1  Ensure password creation
pwquality:
  - key: 'minlen'
    value: '14'
  - key: 'dcredit'
    value: '-1'
  - key: 'ucredit'
    value: '-1'
  - key: 'ocredit'
    value: '-1'
  - key: 'lcredit'
    value: '-1'

# 5.4.1.1 Password

password:
  max_days: 365
  min_days: 7
  warn_age: 7
  inactive: 30


# ======== EXTRAS ===================
nameservers: [8.8.8.8, 8.8.4.4]
```

<!-- Conditional Statements -->

---

## Conditional Statements

### If

The TASKS block can include conditional statements that allow us to decide what to execute depending on the state of the target machine.

On the left side of the conditional statement must be the attribute to be measured and on the right side the value to be compared.

```golang
TASKS : {
    IF "services.ufw" == "stopped"{
        sections: ["1.1", "1.2"];
        exclusions : ["1.1.1.3"];
    }
}
```

### Else

```golang
TASKS : {
    IF "services.ufw" == "stopped"{
        sections: ["1.1", "1.2"];
        exclusions : ["1.1.1.3"];
    }
    ELSE {
        controls: ["9.4"];
    }
}
```

### Elif

```golang
TASKS : {
    IF "services.ufw" == "stopped"{
        sections: ["1.1", "1.2"];
        exclusions : ["1.1.1.3"];
    }
    ELIF "hardscores.lynis" <= 50 {
        points: ["1.1.1.5"];
    }
    ELIF "services.apache" == "STOPPED" {
        points: ["1.1.1.4"];
    }
    ELSE {
        controls: ["9.4"];
    }
}
```

<!-- Getting Host Info -->

---

## Getting Host Info

To use conditional statements, we must use machine values.

The following data are currently available:

* **services**: Provides a list of machine services. A particular service can be accessed with *services.service_name*

```golang
IF "services.ufw" == "stopped"
```

```golang
IF "services.ufw" == "running"
```

* **packages**: Provides a list of machine packages. A particular package can be accessed with *packages.service_name*

```golang
IF "packages.ssh" == "installed"
```

* **hardscores**: Provides the values of the available evaluators. Currently only **Lynis** is available. Returns a numeric value between 0 and 100

```golang
IF "hardscores.lynis" > 50
```

```golang
IF "hardscores.lynis" <= 50
```