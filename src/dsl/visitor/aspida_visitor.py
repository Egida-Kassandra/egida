# Generated from aspida.g4 by ANTLR 4.8
from antlr4 import *
from src.dsl.parser.aspida_parser import AspidaParser

# This class defines a complete generic visitor for a parse tree produced by aspidaParser.

class AspidaVisitor(ParseTreeVisitor):

    # Visit a parse tree produced by aspidaParser#program.
    def visitProgram(self, ctx:AspidaParser.ProgramContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#main.
    def visitMain(self, ctx:AspidaParser.MainContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#hosts.
    def visitHosts(self, ctx:AspidaParser.HostsContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#tasks.
    def visitTasks(self, ctx:AspidaParser.TasksContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#variables.
    def visitVariables(self, ctx:AspidaParser.VariablesContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#main_content.
    def visitMain_content(self, ctx:AspidaParser.Main_contentContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#main_prop.
    def visitMain_prop(self, ctx:AspidaParser.Main_propContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#name.
    def visitName(self, ctx:AspidaParser.NameContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#connection.
    def visitConnection(self, ctx:AspidaParser.ConnectionContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#description.
    def visitDescription(self, ctx:AspidaParser.DescriptionContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#hosts_list.
    def visitHosts_list(self, ctx:AspidaParser.Hosts_listContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#host.
    def visitHost(self, ctx:AspidaParser.HostContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#tasks_content.
    def visitTasks_content(self, ctx:AspidaParser.Tasks_contentContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#tasks_prop.
    def visitTasks_prop(self, ctx:AspidaParser.Tasks_propContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#sections.
    def visitSections(self, ctx:AspidaParser.SectionsContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#points.
    def visitPoints(self, ctx:AspidaParser.PointsContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#controls.
    def visitControls(self, ctx:AspidaParser.ControlsContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#exclusions.
    def visitExclusions(self, ctx:AspidaParser.ExclusionsContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#vars_content.
    def visitVars_content(self, ctx:AspidaParser.Vars_contentContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#vars_prop.
    def visitVars_prop(self, ctx:AspidaParser.Vars_propContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#value.
    def visitValue(self, ctx:AspidaParser.ValueContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#str_array.
    def visitStr_array(self, ctx:AspidaParser.Str_arrayContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#cadena.
    def visitCadena(self, ctx:AspidaParser.CadenaContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#array.
    def visitArray(self, ctx:AspidaParser.ArrayContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#ip_v4.
    def visitIp_v4(self, ctx:AspidaParser.Ip_v4Context):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#single_ip.
    def visitSingle_ip(self, ctx:AspidaParser.Single_ipContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#ip_range.
    def visitIp_range(self, ctx:AspidaParser.Ip_rangeContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#ip_v6.
    def visitIp_v6(self, ctx:AspidaParser.Ip_v6Context):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#h16.
    def visitH16(self, ctx:AspidaParser.H16Context):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#ls32.
    def visitLs32(self, ctx:AspidaParser.Ls32Context):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by aspidaParser#hexdigit.
    def visitHexdigit(self, ctx:AspidaParser.HexdigitContext):
        return self.visitChildren(ctx)



del AspidaParser