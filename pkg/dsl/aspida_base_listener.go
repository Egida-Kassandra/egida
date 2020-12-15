// Code generated from aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package dsl // aspida

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseaspidaListener is a complete listener for a parse tree produced by aspidaParser.
type BaseaspidaListener struct{}

var _ aspidaListener = &BaseaspidaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseaspidaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseaspidaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseaspidaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseaspidaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseaspidaListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseaspidaListener) ExitProgram(ctx *ProgramContext) {}

// EnterMain is called when production main is entered.
func (s *BaseaspidaListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseaspidaListener) ExitMain(ctx *MainContext) {}

// EnterHosts is called when production hosts is entered.
func (s *BaseaspidaListener) EnterHosts(ctx *HostsContext) {}

// ExitHosts is called when production hosts is exited.
func (s *BaseaspidaListener) ExitHosts(ctx *HostsContext) {}

// EnterTasks is called when production tasks is entered.
func (s *BaseaspidaListener) EnterTasks(ctx *TasksContext) {}

// ExitTasks is called when production tasks is exited.
func (s *BaseaspidaListener) ExitTasks(ctx *TasksContext) {}

// EnterVariables is called when production variables is entered.
func (s *BaseaspidaListener) EnterVariables(ctx *VariablesContext) {}

// ExitVariables is called when production variables is exited.
func (s *BaseaspidaListener) ExitVariables(ctx *VariablesContext) {}

// EnterMain_content is called when production main_content is entered.
func (s *BaseaspidaListener) EnterMain_content(ctx *Main_contentContext) {}

// ExitMain_content is called when production main_content is exited.
func (s *BaseaspidaListener) ExitMain_content(ctx *Main_contentContext) {}

// EnterMain_prop is called when production main_prop is entered.
func (s *BaseaspidaListener) EnterMain_prop(ctx *Main_propContext) {}

// ExitMain_prop is called when production main_prop is exited.
func (s *BaseaspidaListener) ExitMain_prop(ctx *Main_propContext) {}

// EnterName is called when production name is entered.
func (s *BaseaspidaListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseaspidaListener) ExitName(ctx *NameContext) {}

// EnterConnection is called when production connection is entered.
func (s *BaseaspidaListener) EnterConnection(ctx *ConnectionContext) {}

// ExitConnection is called when production connection is exited.
func (s *BaseaspidaListener) ExitConnection(ctx *ConnectionContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseaspidaListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseaspidaListener) ExitDescription(ctx *DescriptionContext) {}

// EnterHosts_list is called when production hosts_list is entered.
func (s *BaseaspidaListener) EnterHosts_list(ctx *Hosts_listContext) {}

// ExitHosts_list is called when production hosts_list is exited.
func (s *BaseaspidaListener) ExitHosts_list(ctx *Hosts_listContext) {}

// EnterHost is called when production host is entered.
func (s *BaseaspidaListener) EnterHost(ctx *HostContext) {}

// ExitHost is called when production host is exited.
func (s *BaseaspidaListener) ExitHost(ctx *HostContext) {}

// EnterTasks_content is called when production tasks_content is entered.
func (s *BaseaspidaListener) EnterTasks_content(ctx *Tasks_contentContext) {}

// ExitTasks_content is called when production tasks_content is exited.
func (s *BaseaspidaListener) ExitTasks_content(ctx *Tasks_contentContext) {}

// EnterTasks_prop is called when production tasks_prop is entered.
func (s *BaseaspidaListener) EnterTasks_prop(ctx *Tasks_propContext) {}

// ExitTasks_prop is called when production tasks_prop is exited.
func (s *BaseaspidaListener) ExitTasks_prop(ctx *Tasks_propContext) {}

// EnterSections is called when production sections is entered.
func (s *BaseaspidaListener) EnterSections(ctx *SectionsContext) {}

// ExitSections is called when production sections is exited.
func (s *BaseaspidaListener) ExitSections(ctx *SectionsContext) {}

// EnterPoints is called when production points is entered.
func (s *BaseaspidaListener) EnterPoints(ctx *PointsContext) {}

// ExitPoints is called when production points is exited.
func (s *BaseaspidaListener) ExitPoints(ctx *PointsContext) {}

// EnterControls is called when production controls is entered.
func (s *BaseaspidaListener) EnterControls(ctx *ControlsContext) {}

// ExitControls is called when production controls is exited.
func (s *BaseaspidaListener) ExitControls(ctx *ControlsContext) {}

// EnterExclusions is called when production exclusions is entered.
func (s *BaseaspidaListener) EnterExclusions(ctx *ExclusionsContext) {}

// ExitExclusions is called when production exclusions is exited.
func (s *BaseaspidaListener) ExitExclusions(ctx *ExclusionsContext) {}

// EnterVars_content is called when production vars_content is entered.
func (s *BaseaspidaListener) EnterVars_content(ctx *Vars_contentContext) {}

// ExitVars_content is called when production vars_content is exited.
func (s *BaseaspidaListener) ExitVars_content(ctx *Vars_contentContext) {}

// EnterVars_prop is called when production vars_prop is entered.
func (s *BaseaspidaListener) EnterVars_prop(ctx *Vars_propContext) {}

// ExitVars_prop is called when production vars_prop is exited.
func (s *BaseaspidaListener) ExitVars_prop(ctx *Vars_propContext) {}

// EnterValue is called when production value is entered.
func (s *BaseaspidaListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseaspidaListener) ExitValue(ctx *ValueContext) {}

// EnterStr_array is called when production str_array is entered.
func (s *BaseaspidaListener) EnterStr_array(ctx *Str_arrayContext) {}

// ExitStr_array is called when production str_array is exited.
func (s *BaseaspidaListener) ExitStr_array(ctx *Str_arrayContext) {}

// EnterCadena is called when production cadena is entered.
func (s *BaseaspidaListener) EnterCadena(ctx *CadenaContext) {}

// ExitCadena is called when production cadena is exited.
func (s *BaseaspidaListener) ExitCadena(ctx *CadenaContext) {}

// EnterArray is called when production array is entered.
func (s *BaseaspidaListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseaspidaListener) ExitArray(ctx *ArrayContext) {}

// EnterIp_v4 is called when production ip_v4 is entered.
func (s *BaseaspidaListener) EnterIp_v4(ctx *Ip_v4Context) {}

// ExitIp_v4 is called when production ip_v4 is exited.
func (s *BaseaspidaListener) ExitIp_v4(ctx *Ip_v4Context) {}

// EnterSingle_ip is called when production single_ip is entered.
func (s *BaseaspidaListener) EnterSingle_ip(ctx *Single_ipContext) {}

// ExitSingle_ip is called when production single_ip is exited.
func (s *BaseaspidaListener) ExitSingle_ip(ctx *Single_ipContext) {}

// EnterIp_range is called when production ip_range is entered.
func (s *BaseaspidaListener) EnterIp_range(ctx *Ip_rangeContext) {}

// ExitIp_range is called when production ip_range is exited.
func (s *BaseaspidaListener) ExitIp_range(ctx *Ip_rangeContext) {}

// EnterIp_v6 is called when production ip_v6 is entered.
func (s *BaseaspidaListener) EnterIp_v6(ctx *Ip_v6Context) {}

// ExitIp_v6 is called when production ip_v6 is exited.
func (s *BaseaspidaListener) ExitIp_v6(ctx *Ip_v6Context) {}

// EnterH16 is called when production h16 is entered.
func (s *BaseaspidaListener) EnterH16(ctx *H16Context) {}

// ExitH16 is called when production h16 is exited.
func (s *BaseaspidaListener) ExitH16(ctx *H16Context) {}

// EnterLs32 is called when production ls32 is entered.
func (s *BaseaspidaListener) EnterLs32(ctx *Ls32Context) {}

// ExitLs32 is called when production ls32 is exited.
func (s *BaseaspidaListener) ExitLs32(ctx *Ls32Context) {}

// EnterHexdigit is called when production hexdigit is entered.
func (s *BaseaspidaListener) EnterHexdigit(ctx *HexdigitContext) {}

// ExitHexdigit is called when production hexdigit is exited.
func (s *BaseaspidaListener) ExitHexdigit(ctx *HexdigitContext) {}
