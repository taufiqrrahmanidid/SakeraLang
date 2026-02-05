package parser

import "github.com/taufiqrrahmanidid/SakeraLang/pkg/lexer"

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

// Variable Declaration
type VarStatement struct {
    Token lexer.Token
    Name  string
    Value Expression
}

func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }

// Print Statement
type PrintStatement struct {
    Token lexer.Token
    Value Expression
}

func (ps *PrintStatement) statementNode()       {}
func (ps *PrintStatement) TokenLiteral() string { return ps.Token.Literal }

// If Statement
type IfStatement struct {
    Token       lexer.Token
    Condition   Expression
    Consequence *BlockStatement
    Alternative *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }

// Block Statement
type BlockStatement struct {
    Token      lexer.Token
    Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

// Expressions
type Identifier struct {
    Token lexer.Token
    Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type IntegerLiteral struct {
    Token lexer.Token
    Value int
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

type StringLiteral struct {
    Token lexer.Token
    Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

type Boolean struct {
    Token lexer.Token
    Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// Function Statement
type FunctionStatement struct {
    Token      lexer.Token
    Name       string
    Parameters []string
    Body       *BlockStatement
}

func (fs *FunctionStatement) statementNode()       {}
func (fs *FunctionStatement) TokenLiteral() string { return fs.Token.Literal }

// Return Statement
type ReturnStatement struct {
    Token lexer.Token
    Value Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Expression Statement (for function calls as statements)
type ExpressionStatement struct {
    Token      lexer.Token
    Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// Function Call Expression
type FunctionCall struct {
    Token     lexer.Token
    Name      string
    Arguments []Expression
}

func (fc *FunctionCall) expressionNode()      {}
func (fc *FunctionCall) TokenLiteral() string { return fc.Token.Literal }

type InfixExpression struct {
    Token    lexer.Token
    Left     Expression
    Operator string
    Right    Expression
}

// While Statement
type WhileStatement struct {
    Token     lexer.Token
    Condition Expression
    Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }

type ForStatement struct {
    Token      lexer.Token
    Init       Statement
    Condition  Expression
    Increment  Statement
    Body       *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
