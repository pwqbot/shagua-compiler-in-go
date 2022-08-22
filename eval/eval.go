package eval

import (
	"compiler/ast"
	"compiler/object"

	"github.com/golang/glog"
)

type Logger struct {
	indent string
}

func (l *Logger) Log(s string) interface{} {
	glog.Error(l.indent, s, "\n")
	l.indent += "    "
	return nil
}

func (l *Logger) Close(a interface{}) {
	l.indent = l.indent[4:]
}

var EvalLogger Logger = Logger{}

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		defer EvalLogger.Close(EvalLogger.Log("Eval Program"))
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		defer EvalLogger.Close(EvalLogger.Log("Eval ExpressionStatement"))
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		defer EvalLogger.Close(EvalLogger.Log("Eval IntegerLiteral"))
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		defer EvalLogger.Close(EvalLogger.Log("Eval Boolean"))
		return &object.Boolean{Value: node.Value}
	default:
		glog.Fatal("node type not match")
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var r object.Object
	for _, stmt := range stmts {
		r = Eval(stmt)
	}
	return r
}
