from egida.dsl.codegen.aspida import Aspida
from egida.dsl.codegen.vars_visitor import VarsVisitor

def run_file(tree):
    # Visitor
    visitor = VarsVisitor()
    visitor.visit(tree)
    # Aspida
    aspida = Aspida(
        visitor.name,
        visitor.description,
        visitor.connection,
        visitor.hosts,
        visitor.points,
        visitor.sections,
        visitor.controls,
        visitor.exclusions,
        visitor.vars
    )
