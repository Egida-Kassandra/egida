// Code generated from aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package dsl // aspida

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by aspidaParser.
type AspidaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by aspidaParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by aspidaParser#main.
	VisitMain(ctx *MainContext) interface{}

	// Visit a parse tree produced by aspidaParser#hosts.
	VisitHosts(ctx *HostsContext) interface{}

	// Visit a parse tree produced by aspidaParser#tasks.
	VisitTasks(ctx *TasksContext) interface{}

	// Visit a parse tree produced by aspidaParser#variables.
	VisitVariables(ctx *VariablesContext) interface{}

	// Visit a parse tree produced by aspidaParser#main_content.
	VisitMain_content(ctx *Main_contentContext) interface{}

	// Visit a parse tree produced by aspidaParser#main_prop.
	VisitMain_prop(ctx *Main_propContext) interface{}

	// Visit a parse tree produced by aspidaParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by aspidaParser#connection.
	VisitConnection(ctx *ConnectionContext) interface{}

	// Visit a parse tree produced by aspidaParser#description.
	VisitDescription(ctx *DescriptionContext) interface{}

	// Visit a parse tree produced by aspidaParser#hosts_list.
	VisitHosts_list(ctx *Hosts_listContext) interface{}

	// Visit a parse tree produced by aspidaParser#host.
	VisitHost(ctx *HostContext) interface{}

	// Visit a parse tree produced by aspidaParser#tasks_content.
	VisitTasks_content(ctx *Tasks_contentContext) interface{}

	// Visit a parse tree produced by aspidaParser#tasks_prop.
	VisitTasks_prop(ctx *Tasks_propContext) interface{}

	// Visit a parse tree produced by aspidaParser#sections.
	VisitSections(ctx *SectionsContext) interface{}

	// Visit a parse tree produced by aspidaParser#points.
	VisitPoints(ctx *PointsContext) interface{}

	// Visit a parse tree produced by aspidaParser#controls.
	VisitControls(ctx *ControlsContext) interface{}

	// Visit a parse tree produced by aspidaParser#exclusions.
	VisitExclusions(ctx *ExclusionsContext) interface{}

	// Visit a parse tree produced by aspidaParser#vars_content.
	VisitVars_content(ctx *Vars_contentContext) interface{}

	// Visit a parse tree produced by aspidaParser#vars_prop.
	VisitVars_prop(ctx *Vars_propContext) interface{}

	// Visit a parse tree produced by aspidaParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by aspidaParser#str_array.
	VisitStr_array(ctx *Str_arrayContext) interface{}

	// Visit a parse tree produced by aspidaParser#cadena.
	VisitCadena(ctx *CadenaContext) interface{}

	// Visit a parse tree produced by aspidaParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by aspidaParser#ip_v4.
	VisitIp_v4(ctx *Ip_v4Context) interface{}

	// Visit a parse tree produced by aspidaParser#single_ip.
	VisitSingle_ip(ctx *Single_ipContext) interface{}

	// Visit a parse tree produced by aspidaParser#ip_range.
	VisitIp_range(ctx *Ip_rangeContext) interface{}

	// Visit a parse tree produced by aspidaParser#ip_v6.
	VisitIp_v6(ctx *Ip_v6Context) interface{}

	// Visit a parse tree produced by aspidaParser#h16.
	VisitH16(ctx *H16Context) interface{}

	// Visit a parse tree produced by aspidaParser#ls32.
	VisitLs32(ctx *Ls32Context) interface{}

	// Visit a parse tree produced by aspidaParser#hexdigit.
	VisitHexdigit(ctx *HexdigitContext) interface{}
}
