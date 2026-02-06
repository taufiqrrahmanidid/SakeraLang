package lexer

type Lexer struct {
    input        string
    position     int
    readPosition int
    ch           byte
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

func (l *Lexer) NextToken() Token {
    var tok Token
    
    l.skipWhitespace()
    l.skipComment()
    l.skipWhitespace()  
    
    switch l.ch {
    case '=':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            tok = Token{Type: EQ, Literal: string(ch) + string(l.ch)}
        } else {
            tok = newToken(ASSIGN, l.ch)
        }
    case '+':
        tok = newToken(PLUS, l.ch)
    case '-':
        tok = newToken(MINUS, l.ch)
    case '*':
        tok = newToken(ASTERISK, l.ch)
    case '/':
        tok = newToken(SLASH, l.ch)
    case '<':
        tok = newToken(LT, l.ch)
    case '>':
        tok = newToken(GT, l.ch)
    case '!':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            tok = Token{Type: NOT_EQ, Literal: string(ch) + string(l.ch)}
        } else {
            tok = newToken(ILLEGAL, l.ch)
        }
    case ';':
        tok = newToken(SEMICOLON, l.ch)
    case ',':
        tok = newToken(COMMA, l.ch)
    case '(':
        tok = newToken(LPAREN, l.ch)
    case ')':
        tok = newToken(RPAREN, l.ch)
    case '{':
        tok = newToken(LBRACE, l.ch)
    case '}':
        tok = newToken(RBRACE, l.ch)
    case '"':
        tok.Type = STRING
        tok.Literal = l.readString()
    case 0:
        tok.Literal = ""
        tok.Type = EOF
    case '%':
        tok = newToken(MODULO, l.ch)
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = LookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Type = INT
            tok.Literal = l.readNumber()
            return tok
        } else {
            tok = newToken(ILLEGAL, l.ch)
        }
    }
    
    l.readChar()
    return tok
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func (l *Lexer) skipComment() {
    if l.ch == '/' && l.peekChar() == '/' {
        // Skip sampai akhir baris
        for l.ch != '\n' && l.ch != 0 {
            l.readChar()
        }
    }
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    
    // First character must be letter or underscore
    if isLetter(l.ch) {
        l.readChar()
    }
    
    // Subsequent characters can be letters, digits, or underscores
    for isLetter(l.ch) || isDigit(l.ch) {
        l.readChar()
    }
    
    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) readString() string {
    position := l.position + 1
    for {
        l.readChar()
        if l.ch == '"' || l.ch == 0 {
            break
        }
    }
    return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    }
    return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func newToken(tokenType TokenType, ch byte) Token {
    return Token{Type: tokenType, Literal: string(ch)}
}
