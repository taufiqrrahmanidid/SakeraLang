package parser

import (
    //"fmt"
    "strconv"
    
    "github.com/taufiqrrahmanidid/SakeraLang/pkg/lexer"
)

type Parser struct {
    lexer   *lexer.Lexer
    curToken  lexer.Token
    peekToken lexer.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{lexer: l}
    p.nextToken()
    p.nextToken()
    return p
}

func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseStatement() Statement {
    switch p.curToken.Type {
    case lexer.SANGO:
        return p.parseVarStatement()
    case lexer.TOLES:
        return p.parsePrintStatement()
    case lexer.MON:
        return p.parseIfStatement()
    case lexer.SELAMA:
        return p.parseWhileStatement()
    case lexer.ULANG:
        return p.parseForStatement()
    case lexer.FUNGSI:
        return p.parseFunctionStatement()
    case lexer.MAREH:
        return p.parseReturnStatement()
    case lexer.IDENT:
        // Check if this is a function call
        if p.peekToken.Type == lexer.LPAREN {
            return p.parseExpressionStatement()
        }
        return nil
    default:
        return nil
    }
}

func (p *Parser) parseVarStatement() *VarStatement {
    stmt := &VarStatement{Token: p.curToken}
    
    p.nextToken()
    stmt.Name = p.curToken.Literal
    //fmt.Printf("DEBUG parseVarStatement: Variable name = '%s', curToken = %s, peekToken = %s\n", 
               //stmt.Name, p.curToken.Literal, p.peekToken.Literal)
    
    p.nextToken() // skip identifier
    //fmt.Printf("DEBUG parseVarStatement: After skip ident, curToken = %s, peekToken = %s\n", 
               //p.curToken.Literal, p.peekToken.Literal)
    
    if p.curToken.Type != lexer.ASSIGN {
        return stmt
    }
    
    p.nextToken() // skip '='
    //fmt.Printf("DEBUG parseVarStatement: After skip =, curToken = %s (type=%s), peekToken = %s\n", 
               //p.curToken.Literal, p.curToken.Type, p.peekToken.Literal)
    
    // Parse the value expression (which might be a function call)
    stmt.Value = p.parseExpression()
    //fmt.Printf("DEBUG parseVarStatement: Parsed expression type = %T\n", stmt.Value)
    
    return stmt
}

func (p *Parser) parsePrintStatement() *PrintStatement {
    stmt := &PrintStatement{Token: p.curToken}
    
    p.nextToken()
    stmt.Value = p.parseExpression()
    
    return stmt
}

func (p *Parser) parseIfStatement() *IfStatement {
    stmt := &IfStatement{Token: p.curToken}
    
    p.nextToken()
    stmt.Condition = p.parseExpression()
    
    if p.peekToken.Type == lexer.LBRACE {
        p.nextToken()
        stmt.Consequence = p.parseBlockStatement()
    }
    
    if p.peekToken.Type == lexer.LAEN {
        p.nextToken()
        p.nextToken()
        if p.curToken.Type == lexer.LBRACE {
            stmt.Alternative = p.parseBlockStatement()
        }
    }
    
    return stmt
}

func (p *Parser) parseBlockStatement() *BlockStatement {
    block := &BlockStatement{Token: p.curToken}
    block.Statements = []Statement{}
    
    p.nextToken()
    
    for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
        stmt := p.ParseStatement()
        if stmt != nil {
            block.Statements = append(block.Statements, stmt)
        }
        p.nextToken()
    }
    
    return block
}

func (p *Parser) parseExpression() Expression {
    var left Expression
    
    switch p.curToken.Type {
    case lexer.IDENT:
        // Check if this is a function call
        if p.peekToken.Type == lexer.LPAREN {
            left = p.parseFunctionCall()
        } else {
            left = &Identifier{Token: p.curToken, Value: p.curToken.Literal}
        }
    case lexer.INT:
        value, _ := strconv.Atoi(p.curToken.Literal)
        left = &IntegerLiteral{Token: p.curToken, Value: value}
    case lexer.STRING:
        left = &StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
    case lexer.BENDER:
        left = &Boolean{Token: p.curToken, Value: true}
    case lexer.SALA:
        left = &Boolean{Token: p.curToken, Value: false}
    default:
        return nil
    }
    
    // Check for infix operators
    if p.peekToken.Type == lexer.PLUS || p.peekToken.Type == lexer.MINUS ||
        p.peekToken.Type == lexer.ASTERISK || p.peekToken.Type == lexer.SLASH ||
        p.peekToken.Type == lexer.LT || p.peekToken.Type == lexer.GT ||
        p.peekToken.Type == lexer.EQ || p.peekToken.Type == lexer.NOT_EQ {
        
        p.nextToken()
        operator := p.curToken.Literal
        p.nextToken()
        
        right := p.parseExpression()
        
        return &InfixExpression{
            Token:    p.curToken,
            Left:     left,
            Operator: operator,
            Right:    right,
        }
    }
    
    return left
}

func (p *Parser) parseFunctionCall() *FunctionCall {
    call := &FunctionCall{
        Token: p.curToken,
        Name:  p.curToken.Literal,
    }
    
    p.nextToken() // skip name
    p.nextToken() // skip '('
    
    call.Arguments = []Expression{}
    
    for p.curToken.Type != lexer.RPAREN && p.curToken.Type != lexer.EOF {
        arg := p.parseExpression()
        if arg != nil {
            call.Arguments = append(call.Arguments, arg)
        }
        
        if p.peekToken.Type == lexer.COMMA {
            p.nextToken() // current arg
            p.nextToken() // skip comma
        } else {
            p.nextToken()
        }
    }
    
    return call
}

func (p *Parser) CurToken() lexer.Token {
    return p.curToken
}

func (p *Parser) NextToken() {
    p.nextToken()
}

func (p *Parser) parseWhileStatement() *WhileStatement {
    stmt := &WhileStatement{Token: p.curToken}
    
    p.nextToken()
    stmt.Condition = p.parseExpression()
    
    if p.peekToken.Type == lexer.LBRACE {
        p.nextToken()
        stmt.Body = p.parseBlockStatement()
    }
    
    return stmt
}

func (p *Parser) parseForStatement() *ForStatement {
    stmt := &ForStatement{Token: p.curToken}
    
    // For simple: ulang i = 0; i < 5; i = i + 1 { ... }
    // Skip 'ulang'
    p.nextToken()
    
    // Parse init (variable declaration)
    stmt.Init = p.parseVarStatement()
    
    // Skip semicolon
    if p.peekToken.Type == lexer.SEMICOLON {
        p.nextToken()
    }
    p.nextToken()
    
    // Parse condition
    stmt.Condition = p.parseExpression()
    
    // Skip semicolon
    if p.peekToken.Type == lexer.SEMICOLON {
        p.nextToken()
    }
    p.nextToken()
    
    // Parse increment
    stmt.Increment = p.parseVarStatement()
    
    // Parse body
    if p.peekToken.Type == lexer.LBRACE {
        p.nextToken()
        stmt.Body = p.parseBlockStatement()
    }
    
    return stmt
}

func (p *Parser) parseFunctionStatement() *FunctionStatement {
    stmt := &FunctionStatement{Token: p.curToken}
    
    // Get function name
    p.nextToken()
    stmt.Name = p.curToken.Literal
    
    // Parse parameters
    if p.peekToken.Type == lexer.LPAREN {
        p.nextToken() // skip name
        p.nextToken() // skip '('
        
        stmt.Parameters = []string{}
        
        for p.curToken.Type != lexer.RPAREN && p.curToken.Type != lexer.EOF {
            if p.curToken.Type == lexer.IDENT {
                stmt.Parameters = append(stmt.Parameters, p.curToken.Literal)
            }
            p.nextToken()
            
            // Skip comma
            if p.curToken.Type == lexer.COMMA {
                p.nextToken()
            }
        }
    }
    
    // Parse body
    if p.peekToken.Type == lexer.LBRACE {
        p.nextToken()
        stmt.Body = p.parseBlockStatement()
        // parseBlockStatement already positions us at RBRACE
        // We need to stay at RBRACE so the main loop can skip it
    }
    
    return stmt
}

func (p *Parser) parseReturnStatement() *ReturnStatement {
    stmt := &ReturnStatement{Token: p.curToken}
    
    p.nextToken()
    stmt.Value = p.parseExpression()
    
    return stmt
}

func (p *Parser) parseExpressionStatement() Statement {
    // Create a wrapper to treat expression as statement
    expr := p.parseExpression()
    
    // Wrap in a print-like statement for function calls
    if call, ok := expr.(*FunctionCall); ok {
        return &ExpressionStatement{Expression: call}
    }
    
    return nil
}
