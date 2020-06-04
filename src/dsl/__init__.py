from antlr4 import *

from src.dsl.listeners.aspida_listener import AspidaListener
from src.dsl.parser.aspida_lexer import AspidaLexer
from src.dsl.parser.aspida_parser import AspidaParser
from src.dsl.visitor.aspida_visitor import AspidaVisitor
from src.dsl.visitor.vars_visitor import VarsVisitor


def compile(args):
    filename = args.file
    if filename is not None:
        parse(filename)
    else:
        print("Missing .aspida file -> egida compile file.aspida")


def parse(filename):
    input_stream = FileStream(filename)
    lexer = AspidaLexer(input_stream)
    stream = CommonTokenStream(lexer)
    parser = AspidaParser(stream)
    tree = parser.program()
    # Listener
    walker = ParseTreeWalker()
    listener = AspidaListener()
    walker.walk(listener, tree)
    # Visitor
    visitor = VarsVisitor()
    visitor.visit(tree)
    # Vars
    print(visitor.vars)
