from src.dsl import AspidaVisitor


class VarsVisitor(AspidaVisitor):

    def __init__(self):
        self.actual_var = None
        self.actual_var_child = 0

        self.name = ""
        self.description = ""
        self.connection = ""

        self.hosts = []

        self.points = []
        self.sections = []
        self.controls = []
        self.exclusions = []

        self.vars = {}

    def visitChildren(self, node):
        self.get_values(node)
        result = self.defaultResult()
        n = node.getChildCount()
        for i in range(n):
            if not self.shouldVisitNextChild(node, result):
                return result
            c = node.getChild(i)
            childResult = c.accept(self)
            result = self.aggregateResult(result, childResult)
        return result

    def get_values(self, node):
        name_class = node.__class__.__name__.lower()
        if name_class in dir(self):
            self.__getattribute__(name_class)(node)

    def namecontext(self, node):
        value = str(node.getChild(2)).replace('"', "")
        self.name = value

    def descriptioncontext(self, node):
        value = str(node.getChild(2)).replace('"', "")
        self.description = value

    def connectioncontext(self, node):
        value = str(node.getChild(2)).replace('"', "")
        self.connection = value

    def hostcontext(self, node):
        value = str(node.getText())
        self.hosts.append(value)

    def pointscontext(self, node):
        values = node.getChild(2).getChildren()
        points = [x.getText().replace('"', '') for x in values if x.__class__.__name__.lower() == "cadenacontext"]
        for x in points:
            self.points.append(x)

    def sectionscontext(self, node):
        values = node.getChild(2).getChildren()
        sections = [x.getText().replace('"', '') for x in values if x.__class__.__name__.lower() == "cadenacontext"]
        for x in sections:
            self.sections.append(x)

    def controlscontext(self, node):
        values = node.getChild(2).getChildren()
        controls = [x.getText().replace('"', '') for x in values if x.__class__.__name__.lower() == "cadenacontext"]
        for x in controls:
            self.controls.append(x)

    def exclusionscontext(self, node):
        values = node.getChild(2).getChildren()
        exclusions = [x.getText().replace('"', '') for x in values if x.__class__.__name__.lower() == "cadenacontext"]
        for x in exclusions:
            self.exclusions.append(x)

    def vars_propcontext(self, node):
        prop = node.getChild(0).getText().replace('"', '')
        value = self.get_vars_value(node.getChild(2))
        if value != '{' and self.actual_var is None:
            self.vars[prop] = value
        elif value == '{':
            self.actual_var_child = node.getChildCount()-2
            self.vars[prop] = {}
            self.actual_var = prop
        elif value != '{' and self.actual_var is not None and self.actual_var_child > 0:
            self.vars[self.actual_var][prop] = value
            self.actual_var_child = self.actual_var_child - 1
        else:
            self.actual_var = None
            self.actual_var_child = 0
            self.vars[prop] = value


    def get_vars_value(self, value):
        if value.getChildCount() > 0 and value.getChild(0).__class__.__name__.lower() == "arraycontext":
            n = value.getChild(0).getChildCount()
            values = []
            for i in range(n):
                val = value.getChild(0).getChild(i).getText()
                if val != '[' and val != ',' and val != ']':
                    values.append(val.replace('"', ''))
            return values
        else:
            return value.getText().replace('"', '')