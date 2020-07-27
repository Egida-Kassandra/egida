import argparse

from egida.menu import menu
from egida.dsl import compile
from egida.config import config
from egida.info import info
from egida.common import set_connection_mode


# ========  Mode  ===================
def invalid_args():
    print("EGIDA Mode invalid -> [menu | compile | console | info]")


# ========  Main  ===================
def parse_args():
    parser = argparse.ArgumentParser(description='EGIDA')
    parser.add_argument("mode", help="EGIDA Mode [menu | compile | config | info]")
    parser.add_argument("--file", help="ASPIDA file")
    parser.add_argument("-g", "--group", help="Host group")
    parser.add_argument("-c", "--connection", help="Connection type (default local): local | ssh")
    parser.add_argument('-a', '--audit', help='Audit hosts with lynis', default=False, action='store_true')
    args = parser.parse_args()
    return args

def main():
    args = parse_args()
    connection_mode = args.connection if args.connection is "ssh" else "local"
    set_connection_mode(connection_mode)
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
