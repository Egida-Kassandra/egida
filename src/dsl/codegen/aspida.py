

class Aspida:

    def __init__(self,
                 name,
                 description,
                 connection,
                 hosts = [],
                 points = [],
                 sections = [],
                 controls = [],
                 exclusions = [],
                 vars = {}):
        self.name = name
        self.description = description
        self.connection = connection
        self.hosts = hosts
        self.points = points
        self.sections = sections
        self.controls = controls
        self.exclusions = exclusions
        self.vars = vars
