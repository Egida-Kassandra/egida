from consolemenu import *
from consolemenu.items import *
from consolemenu.format import *
from consolemenu.menu_component import Dimension

from egida.menu.cis import CISConsole


class Console:

    def show_menu(self):
        thin = Dimension(width=100, height=100)
        menu_format = MenuFormatBuilder(max_dimension=thin)
        menu_format.set_border_style_type(MenuBorderStyleType.DOUBLE_LINE_OUTER_LIGHT_INNER_BORDER)

        menu_format.set_title_align('center')
        menu_format.set_prologue_text_align('center')
        menu_format.show_prologue_bottom_border(True)

        # Create the root menu
        menu = ConsoleMenu("EGIDA", prologue_text="Subtitle")
        menu.formatter = menu_format
        self.set_menu_items(menu)

        # Show the menu
        menu.show(True)

    def set_menu_items(self, menu):
        menu.append_item(FunctionItem("CIS Benchmarks", self.action, args=['cis_benchmarks']))
        menu.append_item(FunctionItem("LAMP (soon)", self.action, args=['lamp']))
        menu.append_item(FunctionItem("LEMP (soon)", self.action, args=['lemp']))

    def action(self, name):
        exec('self.{}()'.format(name))

    def cis_benchmarks(self):
        CISConsole()

    def lamp(self):
        print("Lamp")
        Screen().input('Press [Enter] to continue')

    def lemp(self):
        print("Lemp")
        Screen().input('Press [Enter] to continue')
