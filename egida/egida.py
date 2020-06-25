import argparse

from egida.menu import menu
from egida.dsl import compile
from egida.config import config
from egida.info import info


# ========  Mode  ===================
def invalid_args():
    print("EGIDA Mode invalid -> [menu | compile | console | info]")


# ========  Main  ===================
def parse_args():
    parser = argparse.ArgumentParser(description='EGIDA')
    parser.add_argument("mode", help="EGIDA Mode [menu | compile | config | info]")
    parser.add_argument("--file", help="ASPIDA file")
    parser.add_argument("-g", "--group", help="Host group")
    parser.add_argument('-H','--hosts', nargs='+', help='List of hosts, ej: 192.128.2.1 localhost 129.1.1.1')
    parser.add_argument('-a', '--audit', help='Audit hosts with lynis', default=False, action='store_true')
    args = parser.parse_args()
    return args


def main() -> None:
    args = parse_args()
    if args.mode == "menu":
        menu()
    elif args.mode == "compile":
        compile(args)
    elif args.mode == "config":
        config(args)
    elif args.mode == "info":
        info(args)
    else:
        invalid_args()


