import os

from egida.common.command import run_command
from egida.common.file_manager import FileManager


class Playbook:

    def __init__(self, name, tags=None):
        if tags is None:
            tags = []
        self.name = name
        self.tags = tags

    def get_str_tags(self):
        if len(self.tags) == 0:
            return ""
        else:
            tags_str = ' --tags "' + ','.join(self.tags) + '"'
            return tags_str

    def run(self):
        fm = FileManager()
        fm.create_vars_and_hosts()
        cur_path = os.path.abspath(os.path.dirname(__file__))
        path = os.path.join(cur_path, "playbook-custom.yml")
        command_str = 'ansible-playbook ' + path + self.get_str_tags()
        run_command(command_str)


    def run_aspida(self, aspida):
        pass