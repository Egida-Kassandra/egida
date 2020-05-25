from PyInquirer import prompt, Separator


def show_menu():
    questions = [
        {
            'type': 'checkbox',
            'message': 'Select CIS Controls',
            'name': 'controls',
            'choices': [
                Separator('------------> BASIC'),
                # {'name': '1.-Inventory and Control of Hardware Assets'},
                {'name': '2.-Inventory and Control of Software Assets'},
                {'name': '3.-Continuous Vulnerability Management'},
                {'name': '4.-Controlled Use of Administrative Privileges'},
                {'name': '5.-Secure Configuration for Hardware and Software on Mobile Devices, ' +
                         ' Laptops, Workstations and Servers'},
                {'name': '6.-Maintenance, Monitoring and Analysis of Audit Logs'},

                Separator('------------> FOUNDATIONAL'),
                # {'name': '7.-Email and Web Browser Protections'},
                {'name': '8.-Malware Defenses'},
                {'name': '9.-Limitation and Control of Network Ports, Protocols and Services'},
                # {'name': '10.-Data Recovery Capabilities'},
                # {'name': '11.-Secure Configuration for Network Devices, such as Firewalls, Routers and Switches'},
                # {'name': '12.-Boundary Defense'},
                {'name': '13.-Data Protection'},
                {'name': '14.-Controlled Access Based on the Need to Know'},
                # {'name': '15.-Wireless Access Control'},
                {'name': '16.-Account Monitoring and Control'},
            ]
        }
    ]
    answers = prompt(questions)
    controls = [x.split('.')[0] for x in answers['controls']]
    return get_controls(controls)


def get_controls(controls):
    result = []
    all_controls = {
        "2": ['control_2.6'],
        "3": ['control_3.4', 'control_3.5'],
        "4": ['control_4.3', 'control_4.4', 'control_4.5', 'control_4.8', 'control_4.9'],
        "5": ['control_5.1', 'control_5.5'],
        "6": ['control_6', 'control_6.1', 'control_6.2', 'control_6.3'],
        "8": ['control_8.3'],
        "9": ['control_9.2', 'control_9.4'],
        "13": ['control_13'],
        "14": ['control_14.4', 'control_14.6', 'control_14.9'],
        "16": ['control_16', 'control_16.3', 'control_16.4', 'control_16.5', 'control_16.7',
               'control_16.11', 'control_16.13'],
    }
    for ctr in controls:
        if ctr in all_controls.keys():
            result.extend(all_controls[ctr])
    return result