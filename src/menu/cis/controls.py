from PyInquirer import prompt, Separator


def show_menu():
    questions = [
        {
            'type': 'checkbox',
            'message': 'Select CIS Controls',
            'name': 'controls',
            'choices': [
                Separator('------------> BASIC'),
                {'name': '1.-Inventory and Control of Hardware Assets'},
                {'name': '2.-Inventory and Control of Software Assets'},
                {'name': '3.-Continuous Vulnerability Management'},
                {'name': '4.-Controlled Use of Administrative Privileges'},
                {'name': '5.-Secure Configuration for Hardware and Software on Mobile Devices, ' +
                         ' Laptops, Workstations and Servers'},
                {'name': '6.-Maintenance, Monitoring and Analysis of Audit Logs'},

                Separator('------------> FOUNDATIONAL'),
                {'name': '7.-Email and Web Browser Protections'},
                {'name': '8.-Malware Defenses'},
                {'name': '9.-Limitation and Control of Network Ports, Protocols and Services'},
                {'name': '10.-Data Recovery Capabilities'},
                {'name': '11.-Secure Configuration for Network Devices, such as Firewalls, Routers and Switches'},
                {'name': '12.-Boundary Defense'},
                {'name': '13.-Data Protection'},
                {'name': '14.-Controlled Access Based on the Need to Know'},
                {'name': '15.-Wireless Access Control'},
                {'name': '16.-Account Monitoring and Control'},

                Separator('------------> ORGANIZATIONAL'),
                {'name': '17.-Implement a Security Awareness and Training Program'},
                {'name': '18.-Application Software Security'},
                {'name': '19.-Incident Response and Management'},
                {'name': '20.-Penetration Tests and Red Team Exercises'},
            ]
        }
    ]
    answers = prompt(questions)
    return answers['controls']
