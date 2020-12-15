package codegen

import (
	"fmt"
	"github.com/antonioalfa22/egida/pkg/dsl"
)

// VarsVisitor
type VarsVisitor struct{
	dsl.BaseAspidaVisitor
	vars map[string]string
}

func NewVarsVisitor() dsl.AspidaVisitor {
	return &VarsVisitor{
		vars: make(map[string]string),
	}
}

func (v VarsVisitor) VisitProgram(ctx *dsl.ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v VarsVisitor) VisitMain(ctx *dsl.MainContext) interface{} {
	fmt.Println(ctx.GetText())
	return v.VisitChildren(ctx)
}