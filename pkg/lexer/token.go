package lexer

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"
    
    // Identifiers + literals
    IDENT  = "IDENT"
    INT    = "INT"
    STRING = "STRING"
    
    // Operators
    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    ASTERISK = "*"
    SLASH    = "/"
    LT       = "<"
    GT       = ">"
    EQ       = "=="
    NOT_EQ   = "!="
    
    // Delimiters
    COMMA     = ","
    SEMICOLON = ";"
    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"
    
    // Keywords (Bahasa Madura)
    SANGO   = "SANGO"   // variabel
    MON     = "MON"     // if
    LAEN    = "LAEN"    // else
    SELAMA  = "SELAMA"  // while
    ULANG   = "ULANG"   // for
    MAREH   = "MAREH"   // return
    FUNGSI  = "FUNGSI"  // function
    BENDER  = "BENDER"  // true
    SALA    = "SALA"    // false
    TOLES   = "TOLES"   // print
    MACAH   = "MACAH"   // read
    DAFTAR  = "DAFTAR"  // array/list
    DAN     = "DAN"     // and
    ATAU    = "ATAU"    // or
    TIDAK   = "TIDAK"   // not
    MANDEK  = "MANDEK"  // break
    LANJOT  = "LANJOT"  // continue
    
    // Operators v0.3.0
    MODULO  = "%"       // modulo
)

var keywords = map[string]TokenType{
    "sango":  SANGO,
    "mon":    MON,
    "laen":   LAEN,
    "selama": SELAMA,
    "ulang":  ULANG,
    "mareh":  MAREH,
    "fungsi": FUNGSI,
    "bender": BENDER,
    "sala":   SALA,
    "toles":  TOLES,
    "macah":  MACAH,
    "daftar": DAFTAR,
    "dan":    DAN,
    "atau":   ATAU,
    "tidak":  TIDAK,
    "mandek": MANDEK,
    "lanjot": LANJOT,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
