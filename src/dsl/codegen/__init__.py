from src.dsl.codegen.vars_visitor import VarsVisitor

def run_file(tree):
    # Visitor
    visitor = VarsVisitor()
    visitor.visit(tree)
    # Vars
    print("Name:",visitor.name)
    print("Description:",visitor.description)
    print("Connection:",visitor.connection)
    print("Hosts:",visitor.hosts)
    print("Vars:",visitor.vars)