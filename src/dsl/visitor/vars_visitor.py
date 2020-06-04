from src.dsl import AspidaVisitor


class VarsVisitor(AspidaVisitor):

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
        print("Name value:", value)

    def descriptioncontext(self, node):
        value = str(node.getChild(2)).replace('"', "")
        print("Description value:", value)

    def connectioncontext(self, node):
        value = str(node.getChild(2)).replace('"', "")
        print("Connection value:", value)

    def hostcontext(self, node):
        value = str(node.getText())
        print(value)

    def pointscontext(self, node):
        values = node.getChild(2).getChildren()
        points = [x.getText().replace('"', '') for x in values if x.__class__.__name__.lower() == "cadenacontext"]
        for x in points:
            print(x)

    def sectionscontext(self, node):
        values = node.getChild(2).getChildren()
        points = [x.getText().replace('"', '') for x in values if x.__class__.__name__.lower() == "cadenacontext"]
        for x in points:
            print(x)

    def controlscontext(self, node):
        values = node.getChild(2).getChildren()
        points = [x.getText().replace('"', '') for x in values if x.__class__.__name__.lower() == "cadenacontext"]
        for x in points:
            print(x)

    def exclusionscontext(self, node):
        values = node.getChild(2).getChildren()
        points = [x.getText().replace('"', '') for x in values if x.__class__.__name__.lower() == "cadenacontext"]
        for x in points:
            print(x)