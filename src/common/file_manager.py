import os
from PyInquirer import prompt
from jinja2 import Template


class FileManager:

    def __init__(self):
        self.path = ''
        self.vars_file = ''
        self.hosts_file = ''
        self.vars = []
        self.hosts = []

    def scan_custom_dir(self):
        with os.scandir('custom/') as entries:
            var_files = [x.name.split('_')[1] for x in entries if x.name.split('_')[0] == 'vars']
            questions = [
                {
                    'type': 'list',
                    'name': 'vars_file',
                    'message': 'Which vars file do you want to use?',
                    'choices': var_files,
                    'filter': lambda val: val.lower().replace(' ', '_')
                }
            ]
            answers = prompt(questions)
            self.vars_file = answers['vars_file']
        with os.scandir('custom/') as entries:
            hosts_files = [x.name.split('_')[1] for x in entries if x.name.split('_')[0] == 'hosts']
            questions = [
                {
                    'type': 'list',
                    'name': 'hosts_file',
                    'message': 'Which hosts file do you want to use?',
                    'choices': hosts_files,
                    'filter': lambda val: val.lower().replace(' ', '_')
                }
            ]
            answers = prompt(questions)
            self.hosts_file = answers['hosts_file']

    def read_vars_files(self):
        with open(r'custom/vars_{}'.format(self.vars_file)) as var_file:
            for line in var_file.readlines():
                self.vars.append('    {}'.format(line))

    def read_hosts_files(self):
        with open(r'custom/hosts_{}'.format(self.vars_file)) as host_file:
            for line in host_file.readlines():
                self.vars.append('- {}'.format(line))


    def write_playbook(self):
        cur_path = os.path.abspath(os.path.dirname(__file__))
        with open(os.path.join(cur_path, "playbook-custom.yml"), 'w+') as file:
            with open(os.path.join(cur_path, 'playbook.yml.j2'), 'r') as template:
                t = Template(template.read())
                file.write(t.render(vars=''.join(self.vars), hosts='- localhost', connection='ssh'))

    def create_vars_and_hosts(self):
        self.scan_custom_dir()
        self.read_vars_files()
        self.read_hosts_files()
        self.write_playbook()
