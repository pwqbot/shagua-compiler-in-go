package eval

import (
	"compiler/ast"
	"compiler/object"
	"compiler/token"

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
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		defer EvalLogger.Close(EvalLogger.Log("Eval PrefixExpression"))
		return evalPrefixExpression(node)
	case *ast.InfixExpression:
		defer EvalLogger.Close(EvalLogger.Log("Eval InfixExpression"))
		return evalInfixExpression(node)
	case *ast.IfExpreesion:
		defer EvalLogger.Close(EvalLogger.Log("Eval IfExpression"))
		return evalIfExpression(node)
	case *ast.BlockStatement:
		defer EvalLogger.Close(EvalLogger.Log("Eval BlockStatements"))
		return evalStatements(node.Statements)
	default:
		glog.Fatal("node type not match")
	}
	return nil
}

func evalIfExpression(node *ast.IfExpreesion) object.Object {
	return nil
}

func evalInfixExpression(node *ast.InfixExpression) object.Object {
	leftObj := Eval(node.Left)
	rightObj := Eval(node.Right)
	switch {
	case leftObj.Type() == object.INTEGER_OBJ && rightObj.Type() == object.INTEGER_OBJ:
		return evalIntegerExpression(node.Token.Type, leftObj, rightObj)
	case leftObj.Type() == object.BOOLEAN_OBJ && rightObj.Type() == object.BOOLEAN_OBJ:
		return evalBooleanExpression(node.Token.Type, leftObj, rightObj)
	default:
		glog.Fatal("token not found", node.Token.Type)
		return object.NULL
	}
}

func evalBooleanExpression(op token.TokenType, leftObj object.Object, rightObj object.Object) object.Object {
	l, _ := leftObj.(*object.Boolean)
	r, _ := rightObj.(*object.Boolean)
	switch op {
	case token.EQ:
		return nativeBoolToBooleanObject(l == r)
	case token.GE:
		return nativeBoolToBooleanObject(false)
	case token.GT:
		return nativeBoolToBooleanObject(false)
	case token.LE:
		return nativeBoolToBooleanObject(false)
	case token.LT:
		return nativeBoolToBooleanObject(false)
	default:
		glog.Fatal("boolean expression do not support ", op)
	}
	return nil
}

func evalIntegerExpression(op token.TokenType, leftObj object.Object, rightObj object.Object) object.Object {
	l, _ := leftObj.(*object.Integer)
	r, _ := rightObj.(*object.Integer)
	switch op {
	case token.MINUS:
		value := l.Value - r.Value
		return &object.Integer{Value: value}
	case token.PLUS:
		value := l.Value + r.Value
		return &object.Integer{Value: value}
	case token.MULTI:
		value := l.Value * r.Value
		return &object.Integer{Value: value}
	case token.DIVIDE:
		value := l.Value / r.Value
		return &object.Integer{Value: value}
	case token.EQ:
		return nativeBoolToBooleanObject(l.Value == r.Value)
	case token.GE:
		return nativeBoolToBooleanObject(l.Value >= r.Value)
	case token.GT:
		return nativeBoolToBooleanObject(l.Value > r.Value)
	case token.LE:
		return nativeBoolToBooleanObject(l.Value <= r.Value)
	case token.LT:
		return nativeBoolToBooleanObject(l.Value < r.Value)
	default:
		glog.Fatal("Unknown Operator type")
		return nil
	}
}

func evalPrefixExpression(node *ast.PrefixExpression) object.Object {
	rightObj := Eval(node.Right)
	switch node.Token.Type {
	case token.BANG:
		return evalBangExpression(rightObj)
	case token.MINUS:
		return evalMinusExpression(rightObj)
	}
	return object.NULL
}

func evalMinusExpression(rightObj object.Object) object.Object {
	obj, ok := rightObj.(*object.Integer)
	if !ok {
		return object.NULL
	}
	return &object.Integer{Value: -obj.Value}
}

func evalBangExpression(rightObj object.Object) object.Object {
	switch rightObj {
	case object.TRUE:
		return object.FALSE
	case object.FALSE:
		return object.TRUE
	case object.NULL:
		return object.TRUE
	default:
		return object.FALSE
	}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	switch input {
	case true:
		return object.TRUE
	case false:
		return object.FALSE
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
