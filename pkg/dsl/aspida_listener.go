// Code generated from aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package dsl // aspida

import "github.com/antlr/antlr4/runtime/Go/antlr"

// aspidaListener is a complete listener for a parse tree produced by aspidaParser.
type aspidaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterHosts is called when entering the hosts production.
	EnterHosts(c *HostsContext)

	// EnterTasks is called when entering the tasks production.
	EnterTasks(c *TasksContext)

	// EnterVariables is called when entering the variables production.
	EnterVariables(c *VariablesContext)

	// EnterMain_content is called when entering the main_content production.
	EnterMain_content(c *Main_contentContext)

	// EnterMain_prop is called when entering the main_prop production.
	EnterMain_prop(c *Main_propContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterConnection is called when entering the connection production.
	EnterConnection(c *ConnectionContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterHosts_list is called when entering the hosts_list production.
	EnterHosts_list(c *Hosts_listContext)

	// EnterHost is called when entering the host production.
	EnterHost(c *HostContext)

	// EnterTasks_content is called when entering the tasks_content production.
	EnterTasks_content(c *Tasks_contentContext)

	// EnterTasks_prop is called when entering the tasks_prop production.
	EnterTasks_prop(c *Tasks_propContext)

	// EnterSections is called when entering the sections production.
	EnterSections(c *SectionsContext)

	// EnterPoints is called when entering the points production.
	EnterPoints(c *PointsContext)

	// EnterControls is called when entering the controls production.
	EnterControls(c *ControlsContext)

	// EnterExclusions is called when entering the exclusions production.
	EnterExclusions(c *ExclusionsContext)

	// EnterVars_content is called when entering the vars_content production.
	EnterVars_content(c *Vars_contentContext)

	// EnterVars_prop is called when entering the vars_prop production.
	EnterVars_prop(c *Vars_propContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterStr_array is called when entering the str_array production.
	EnterStr_array(c *Str_arrayContext)

	// EnterCadena is called when entering the cadena production.
	EnterCadena(c *CadenaContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterIp_v4 is called when entering the ip_v4 production.
	EnterIp_v4(c *Ip_v4Context)

	// EnterSingle_ip is called when entering the single_ip production.
	EnterSingle_ip(c *Single_ipContext)

	// EnterIp_range is called when entering the ip_range production.
	EnterIp_range(c *Ip_rangeContext)

	// EnterIp_v6 is called when entering the ip_v6 production.
	EnterIp_v6(c *Ip_v6Context)

	// EnterH16 is called when entering the h16 production.
	EnterH16(c *H16Context)

	// EnterLs32 is called when entering the ls32 production.
	EnterLs32(c *Ls32Context)

	// EnterHexdigit is called when entering the hexdigit production.
	EnterHexdigit(c *HexdigitContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitHosts is called when exiting the hosts production.
	ExitHosts(c *HostsContext)

	// ExitTasks is called when exiting the tasks production.
	ExitTasks(c *TasksContext)

	// ExitVariables is called when exiting the variables production.
	ExitVariables(c *VariablesContext)

	// ExitMain_content is called when exiting the main_content production.
	ExitMain_content(c *Main_contentContext)

	// ExitMain_prop is called when exiting the main_prop production.
	ExitMain_prop(c *Main_propContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitConnection is called when exiting the connection production.
	ExitConnection(c *ConnectionContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitHosts_list is called when exiting the hosts_list production.
	ExitHosts_list(c *Hosts_listContext)

	// ExitHost is called when exiting the host production.
	ExitHost(c *HostContext)

	// ExitTasks_content is called when exiting the tasks_content production.
	ExitTasks_content(c *Tasks_contentContext)

	// ExitTasks_prop is called when exiting the tasks_prop production.
	ExitTasks_prop(c *Tasks_propContext)

	// ExitSections is called when exiting the sections production.
	ExitSections(c *SectionsContext)

	// ExitPoints is called when exiting the points production.
	ExitPoints(c *PointsContext)

	// ExitControls is called when exiting the controls production.
	ExitControls(c *ControlsContext)

	// ExitExclusions is called when exiting the exclusions production.
	ExitExclusions(c *ExclusionsContext)

	// ExitVars_content is called when exiting the vars_content production.
	ExitVars_content(c *Vars_contentContext)

	// ExitVars_prop is called when exiting the vars_prop production.
	ExitVars_prop(c *Vars_propContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitStr_array is called when exiting the str_array production.
	ExitStr_array(c *Str_arrayContext)

	// ExitCadena is called when exiting the cadena production.
	ExitCadena(c *CadenaContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitIp_v4 is called when exiting the ip_v4 production.
	ExitIp_v4(c *Ip_v4Context)

	// ExitSingle_ip is called when exiting the single_ip production.
	ExitSingle_ip(c *Single_ipContext)

	// ExitIp_range is called when exiting the ip_range production.
	ExitIp_range(c *Ip_rangeContext)

	// ExitIp_v6 is called when exiting the ip_v6 production.
	ExitIp_v6(c *Ip_v6Context)

	// ExitH16 is called when exiting the h16 production.
	ExitH16(c *H16Context)

	// ExitLs32 is called when exiting the ls32 production.
	ExitLs32(c *Ls32Context)

	// ExitHexdigit is called when exiting the hexdigit production.
	ExitHexdigit(c *HexdigitContext)
}
