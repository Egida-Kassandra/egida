from src.common.playbook import Playbook

def menu(args):
    playbook = Playbook("Menu", tags=["rule_2.1.1", "rule_2.1.2"])
    playbook.run()
