from PyInquirer import prompt
from consolemenu import *
from src.menu.cis import sections, controls, all, points
from src.common.playbook import Playbook


class CISConsole:

    def __init__(self):
        self.questions = [
            {
                'type': 'list',
                'name': 'cis_mode',
                'message': 'Which execution mode do you want to use?',
                'choices': [
                    'All', 'Select CIS controls', 'Select CIS sections', 'Select CIS points'
                ],
                'filter': lambda val: val.lower().replace(' ', '_')
            }
        ]
        self.show_menu()

    def show_menu(self):
        answers = prompt(self.questions)
        exec('self.{}()'.format(answers['cis_mode']))

    def all(self):
        all.show_menu()
        playbook = Playbook('all')
        playbook.run()
        Screen().input('Press [Enter] to continue')

    def select_cis_sections(self):
        selected_sections = sections.show_menu()
        playbook = Playbook('sections', tags=selected_sections)
        playbook.run()
        Screen().input('Press [Enter] to continue')

    def select_cis_controls(self):
        selected_controls = controls.show_menu()
        playbook = Playbook('controls', tags=selected_controls)
        playbook.run()
        Screen().input('Press [Enter] to continue')

    def select_cis_points(self):
        selected_points = points.show_menu()
        playbook = Playbook('points', tags=selected_points)
        playbook.run()
        Screen().input('Press [Enter] to continue')
        