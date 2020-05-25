from PyInquirer import prompt, Separator


def show_menu():
    questions = [
        {
            'type': 'checkbox',
            'message': 'Select CIS Sections',
            'name': 'sections',
            'choices': [
                Separator('------------> 1. Initial Setup'),
                {'name': '1.1-Filesystem Configuration'},
                {'name': '1.2-Console Software Updates'},
                {'name': '1.3-Configure sudo'},
                {'name': '1.4-Filesystem Integrity'},
                {'name': '1.5-Secure Boot Settings'},
                {'name': '1.6-Additional Process Hardering'},
                {'name': '1.7-Mandatory Access Control'},
                {'name': '1.8-Warning Banners'},
                {'name': '1.9-Updates'},

                Separator('------------> 2. Services'),
                {'name': '2.1-Initd Services'},
                {'name': '2.2-Special Purpose Services'},
                {'name': '2.3-Service Clients'},

                Separator('------------> 3. Network Configuration'),
                {'name': '3.1-Network Parameters Host'},
                {'name': '3.2-Network Parameters Host and Router'},
                {'name': '3.3-TCP Wrappers'},
                {'name': '3.4-Uncommon Network Protocols'},
                {'name': '3.5-Firewall Configuration'},
                {'name': '3.6-Wireless'},
                {'name': '3.7-Disable IPv6'},

                Separator('------------> 4. Logging and Auditing'),
                {'name': '4.1-Configure System Accounting (auditd)'},
                {'name': '4.2-Configure logging'},
                {'name': '4.3-Ensure logrotate is configured'},

                Separator('------------> 5. Access Authentication and Authorization'),
                {'name': '5.1-Configure cron'},
                {'name': '5.2-SSH Server configuration'},
                {'name': '5.3-Configure PAM'},
                {'name': '5.4-User accounts and environment'},
                {'name': '5.5-Ensure root login is restricted to system console'},
                {'name': '5.6-Ensure access to the su command is restricted'},

                Separator('------------> 6. System maintenance'),
                {'name': '6.1-System file permisions'},
                {'name': '6.2-User and Group Settings'},
            ]
        }
    ]
    answers = prompt(questions)
    sections = ['section_{}'.format(x.split('-')[0]) for x in answers['sections']]
    return sections
