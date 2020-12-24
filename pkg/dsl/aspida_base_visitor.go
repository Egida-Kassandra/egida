package dsl

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// BaseAspidaVisitor
type BaseAspidaVisitor struct{
	*antlr.BaseParseTreeVisitor
}

func (v BaseAspidaVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitMain(ctx *MainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitHosts(ctx *HostsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitTasks(ctx *TasksContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitVariables(ctx *VariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitMain_content(ctx *Main_contentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitMain_prop(ctx *Main_propContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitConnection(ctx *ConnectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitHosts_list(ctx *Hosts_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitHost(ctx *HostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitTasks_content(ctx *Tasks_contentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitTasks_prop(ctx *Tasks_propContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitSections(ctx *SectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitPoints(ctx *PointsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitControls(ctx *ControlsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitExclusions(ctx *ExclusionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitVars_content(ctx *Vars_contentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitVars_prop(ctx *Vars_propContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitStr_array(ctx *Str_arrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitCadena(ctx *CadenaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitIp_v4(ctx *Ip_v4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitSingle_ip(ctx *Single_ipContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitIp_range(ctx *Ip_rangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitIp_v6(ctx *Ip_v6Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitH16(ctx *H16Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitLs32(ctx *Ls32Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v BaseAspidaVisitor) VisitHexdigit(ctx *HexdigitContext) interface{} {
	return v.VisitChildren(ctx)
}
