from antlr4 import *

from egida.dsl.listeners.aspida_listener import AspidaListener
from egida.dsl.parser.aspida_lexer import AspidaLexer
from egida.dsl.parser.aspida_parser import AspidaParser
from egida.dsl.codegen import run_file


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
    # Codegen
    run_file(tree)
