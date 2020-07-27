import os
from PyInquirer import prompt
from jinja2 import Template


class FileManager:

    def __init__(self):
        self.path = ''
        self.vars_file = ''
        self.hosts_file = ''
        self.vars = []
        self.hosts = ''

    def scan_custom_dir(self):
        with os.scandir('/etc/egida/custom/') as entries:
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
        hosts_groups = []
        with open(r'/etc/egida/custom/hosts') as host_file:
            for line in host_file.readlines():
                hosts_groups.append(line)
        questions = [
            {
                'type': 'list',
                'name': 'hosts',
                'message': 'Which hosts group do you want to use?',
                'choices': hosts_groups,
                'filter': lambda val: val.lower().replace(' ', '_')
            }
        ]
        answers = prompt(questions)
        self.hosts = '- {}'.format(answers['hosts'])

    def read_vars_files(self):
        with open(r'/etc/egida/custom/vars_{}'.format(self.vars_file)) as var_file:
            for line in var_file.readlines():
                self.vars.append('    {}'.format(line))

    def write_playbook(self):
        cur_path = os.path.abspath(os.path.dirname(__file__))
        with open(os.path.join(cur_path, "playbook-custom.yml"), 'w+') as file:
            with open(os.path.join(cur_path, 'playbook.yml.j2'), 'r') as template:
                t = Template(template.read())
                file.write(t.render(vars=''.join(self.vars), hosts=self.hosts, connection='ssh'))

    def create_vars_and_hosts(self):
        self.scan_custom_dir()
        self.read_vars_files()
        self.write_playbook()
