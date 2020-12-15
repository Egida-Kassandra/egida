package dsl

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/antonioalfa22/egida/internal/dsl/codegen"
	"github.com/antonioalfa22/egida/pkg/dsl"
)

type TreeShapeListener struct {
	*dsl.BaseaspidaListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func ParseFile(file string)  {
	input, _ := antlr.NewFileStream(file)
	// Lexer
	lexer := dsl.NewaspidaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	// Parser
	p := dsl.NewaspidaParser(stream)
	// Listeners
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	// CodeGen
	varsVisitor := codegen.NewVarsVisitor()
	tree.Accept(varsVisitor)
}
