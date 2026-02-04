package evaluator

import (

    "bufio"

    "fmt"

    "os"

    "strconv"

    "strings"

    

    "github.com/taufiqrrahmanidid/SakeraLang/pkg/lexer"

)

type Environment struct {

    variables map[string]interface{}

}

func NewEnvironment() *Environment {

    return &Environment{

        variables: make(map[string]interface{}),

    }

}

func Eval(input string, env *Environment) {

    l := lexer.New(input)

    tokens := []lexer.Token{}

    

    for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {

        tokens = append(tokens, tok)

    }

    

    if len(tokens) == 0 {

        return

    }

    

    // Handle SANGO (variabel)

    if tokens[0].Type == lexer.SANGO {

        if len(tokens) >= 4 && tokens[2].Type == lexer.ASSIGN {

            varName := tokens[1].Literal

            value := evaluateExpression(tokens[3:], env)

            env.variables[varName] = value

            return

        }

    }

    

    // Handle TOLES (print)

    if tokens[0].Type == lexer.TOLES {

        result := evaluateExpression(tokens[1:], env)

        fmt.Println(result)

        return

    }

    

    // Handle MACAH (read input)

    if tokens[0].Type == lexer.MACAH {

        if len(tokens) >= 2 {

            varName := tokens[1].Literal

            reader := bufio.NewReader(os.Stdin)

            fmt.Print("Input: ")

            text, _ := reader.ReadString('\n')

            text = strings.TrimSpace(text)

            env.variables[varName] = text

            return

        }

    }

}

func evaluateExpression(tokens []lexer.Token, env *Environment) interface{} {

    if len(tokens) == 0 {

        return nil

    }

    

    // Simple expression evaluation

    if len(tokens) == 1 {

        tok := tokens[0]

        switch tok.Type {

        case lexer.INT:

            val, _ := strconv.Atoi(tok.Literal)

            return val

        case lexer.STRING:

            return tok.Literal

        case lexer.BENDER:

            return true

        case lexer.SALA:

            return false

        case lexer.IDENT:

            if val, ok := env.variables[tok.Literal]; ok {

                return val

            }

            return tok.Literal

        }

    }

    

    // Handle binary operations

    if len(tokens) >= 3 {

        left := evaluateExpression([]lexer.Token{tokens[0]}, env)

        operator := tokens[1]

        right := evaluateExpression(tokens[2:], env)

        

        leftInt, leftIsInt := left.(int)

        rightInt, rightIsInt := right.(int)

        

        if leftIsInt && rightIsInt {

            switch operator.Type {

            case lexer.PLUS:

                return leftInt + rightInt

            case lexer.MINUS:

                return leftInt - rightInt

            case lexer.ASTERISK:

                return leftInt * rightInt

            case lexer.SLASH:

                if rightInt != 0 {

                    return leftInt / rightInt

                }

            case lexer.LT:

                return leftInt < rightInt

            case lexer.GT:

                return leftInt > rightInt

            case lexer.EQ:

                return leftInt == rightInt

            }

        }

        

        // String concatenation

        if operator.Type == lexer.PLUS {

            return fmt.Sprintf("%v%v", left, right)

        }

    }

    

    return tokens[0].Literal

}
