import argparse

from egida.menu import menu
from egida.dsl import compile
from egida.config import config

# ========  Mode  ===================
def invalid_args():
    print("EGIDA Mode invalid -> [menu | compile | console]")

# ========  Main  ===================
def parse_args():
    parser = argparse.ArgumentParser(description='EGIDA')
    parser.add_argument("mode", help="EGIDA Mode [menu | compile | config]")
    args = parser.parse_args()
    return args

def main(args) -> None:
    if args.mode == "menu": menu(args)
    elif args.mode == "compile": compile(args)
    elif args.mode == "config": config(args)
    else: invalid_args()

if __name__ == "__main__":
    main(parse_args())