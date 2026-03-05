package ast

import (
	"github.com/umid-stha/interpreter/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

/*
Program is the root node of our languages AST. This language is just a series of statements hence why it contains a list of statements.
Statements are a slice of AST.
*/
type Program struct {
	Statements []Statement
}

// implementing the node inteface
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

/*
LetStatements is a Statement with three fields: the token AST node is assoisted with in this case LET, Identifier and the expression which is assigned.
Example:
let x = 6/2 - 1;
in this case let being the token, x being the indentifier and 6/2-1 being expression which is a expression node
*/
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
