package evaluator

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    
    "github.com/taufiqrrahmanidid/SakeraLang/pkg/lexer"
    "github.com/taufiqrrahmanidid/SakeraLang/pkg/parser"
)

type Function struct {
    Parameters []string
    Body       *parser.BlockStatement
    Env        *Environment
}

type ReturnValue struct {
    Value interface{}
}

type Environment struct {
    variables map[string]interface{}
    functions map[string]*Function
}

func NewEnvironment() *Environment {
    return &Environment{
        variables: make(map[string]interface{}),
        functions: make(map[string]*Function),
    }
}

// Simple line-based evaluation (backward compatible)
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
            reader := bufio.NewReader(os.Stdin)
            text, _ := reader.ReadString('\n')
            text = strings.TrimSpace(text)
            
            // Try to parse as integer
            if num, err := strconv.Atoi(text); err == nil {
                env.variables[tokens[1].Literal] = num
            } else {
                env.variables[tokens[1].Literal] = text
            }
            return
        }
    }
    
    // Handle MON (if statement) - simplified version
    if tokens[0].Type == lexer.MON {
        evaluateSimpleIf(tokens, env)
        return
    }
}

// Evaluate simple if statement (single line)
func evaluateSimpleIf(tokens []lexer.Token, env *Environment) {
    // Find condition (between MON and LBRACE or end of line)
    conditionTokens := []lexer.Token{}
    i := 1
    
    for i < len(tokens) && tokens[i].Type != lexer.LBRACE {
        conditionTokens = append(conditionTokens, tokens[i])
        i++
    }
    
    // Evaluate condition
    result := evaluateExpression(conditionTokens, env)
    
    // Check if condition is true
    isTruthy := false
    switch v := result.(type) {
    case bool:
        isTruthy = v
    case int:
        isTruthy = v != 0
    case string:
        isTruthy = v != ""
    }
    
    if isTruthy {
        fmt.Println("Kondisi benar (true)")
    } else {
        fmt.Println("Kondisi salah (false)")
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
        
        leftInt, leftIsInt := toInt(left)
        rightInt, rightIsInt := toInt(right)
        
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
            case lexer.NOT_EQ:
                return leftInt != rightInt
            }
        }
        
        // String concatenation
        if operator.Type == lexer.PLUS {
            return fmt.Sprintf("%v%v", left, right)
        }
        
        // String/Boolean comparison
        if operator.Type == lexer.EQ {
            return fmt.Sprintf("%v", left) == fmt.Sprintf("%v", right)
        }
        if operator.Type == lexer.NOT_EQ {
            return fmt.Sprintf("%v", left) != fmt.Sprintf("%v", right)
        }
    }
    if len(tokens) >= 3 {
	    for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == lexer.DAN || tokens[i].Type == lexer.ATAU {
		    left := evaluateExpression(tokens[:i], env)
		    operator := tokens[i]
		    right := evaluateExpression(tokens[i+1:], env)
		    
		    leftBool := isTruthy(left)
		    rightBool := isTruthy(right)
		    
		    if operator.Type == lexer.DAN {
		        return leftBool && rightBool
		    } else if operator.Type == lexer.ATAU {
		        return leftBool || rightBool
		    }
		}
	    }
	}
	// Handle TIDAK (NOT)
	if len(tokens) > 0 && tokens[0].Type == lexer.TIDAK {
	    value := evaluateExpression(tokens[1:], env)
	    return !isTruthy(value)
	}
    
    return tokens[0].Literal
}

func toInt(v interface{}) (int, bool) {
    switch val := v.(type) {
    case int:
        return val, true
    case string:
        if i, err := strconv.Atoi(val); err == nil {
            return i, true
        }
    }
    return 0, false
}

func EvalProgram(input string, env *Environment) {
    l := lexer.New(input)
    p := parser.New(l)
    
    for p.CurToken().Type != lexer.EOF {
        stmt := p.ParseStatement()
        if stmt != nil {
            evalStatement(stmt, env)
        }
        
        // Move to next token
        p.NextToken()
        
        // Skip any remaining braces or semicolons
        for p.CurToken().Type == lexer.RBRACE || p.CurToken().Type == lexer.SEMICOLON {
            p.NextToken()
        }
    }
}

func evalStatement(stmt parser.Statement, env *Environment) interface{} {
    switch node := stmt.(type) {
    case *parser.VarStatement:
    	value := evalExpression(node.Value, env)
    	//fmt.Printf("DEBUG VarStatement: Setting '%s' = %v (type: %T)\n", node.Name, value, value)
    	env.variables[node.Name] = value
    	//fmt.Printf("DEBUG VarStatement: Verify '%s' in env = %v\n", node.Name, env.variables[node.Name])
    	return value
        
    case *parser.PrintStatement:
        value := evalExpression(node.Value, env)
        fmt.Println(value)
        return value
        
    case *parser.ExpressionStatement:
        return evalExpression(node.Expression, env)
        
    case *parser.IfStatement:
        return evalIfStatement(node, env)
        
    case *parser.WhileStatement:
        return evalWhileStatement(node, env)
        
    case *parser.ForStatement:
        return evalForStatement(node, env)
        
    case *parser.FunctionStatement:
        fn := &Function{
            Parameters: node.Parameters,
            Body:       node.Body,
            Env:        env,
        }
        env.functions[node.Name] = fn
        return fn
        
    case *parser.ReturnStatement:
        value := evalExpression(node.Value, env)
        return &ReturnValue{Value: value}
        
    case *parser.BlockStatement:
        return evalBlockStatement(node, env)
    }
    
    return nil
}

func evalIfStatement(node *parser.IfStatement, env *Environment) interface{} {
    condition := evalExpression(node.Condition, env)
    
    if isTruthy(condition) {
        return evalBlockStatement(node.Consequence, env)
    } else if node.Alternative != nil {
        return evalBlockStatement(node.Alternative, env)
    }
    
    return nil
}

func evalWhileStatement(node *parser.WhileStatement, env *Environment) interface{} {
    var result interface{}
    
    for {
        condition := evalExpression(node.Condition, env)
        
        if !isTruthy(condition) {
            break
        }
        
        result = evalBlockStatement(node.Body, env)
    }
    
    return result
}

func evalForStatement(node *parser.ForStatement, env *Environment) interface{} {
    var result interface{}
    
    // Execute init
    if node.Init != nil {
        evalStatement(node.Init, env)
    }
    
    // Loop
    for {
        // Check condition
        if node.Condition != nil {
            condition := evalExpression(node.Condition, env)
            if !isTruthy(condition) {
                break
            }
        }
        
        // Execute body
        result = evalBlockStatement(node.Body, env)
        
        // Execute increment
        if node.Increment != nil {
            evalStatement(node.Increment, env)
        }
    }
    
    return result
}

func evalBlockStatement(block *parser.BlockStatement, env *Environment) interface{} {
    var result interface{}
    
    for _, stmt := range block.Statements {
        result = evalStatement(stmt, env)
        
        // Check for return statement
        if _, ok := result.(*ReturnValue); ok {
            return result
        }
    }
    
    return result
}

func evalExpression(expr parser.Expression, env *Environment) interface{} {
    switch node := expr.(type) {
    case *parser.IntegerLiteral:
        return node.Value
        
    case *parser.StringLiteral:
        return node.Value
        
    case *parser.Boolean:
        return node.Value
        
    case *parser.Identifier:
    	if val, ok := env.variables[node.Value]; ok {
    		//fmt.Printf("DEBUG Identifier: Found '%s' = %v\n", node.Value, val)
    		return val
    	}
    	//fmt.Printf("DEBUG Identifier: NOT FOUND '%s', returning literal\n", node.Value)
    	return node.Value
        
    case *parser.InfixExpression:
        left := evalExpression(node.Left, env)
        right := evalExpression(node.Right, env)
        return evalInfixExpression(node.Operator, left, right)
        
    case *parser.FunctionCall:
        return evalFunctionCall(node, env)
    }
    
    return nil
}

func evalInfixExpression(operator string, left, right interface{}) interface{} {
    // Handle logical operators first
    if operator == "dan" {
        return isTruthy(left) && isTruthy(right)
    }
    if operator == "atau" {
        return isTruthy(left) || isTruthy(right)
    }
    if operator == "tidak" {
        return !isTruthy(right)
    }
    
    leftInt, leftIsInt := toInt(left)
    rightInt, rightIsInt := toInt(right)
    
    if leftIsInt && rightIsInt {
        switch operator {
        case "+":
            return leftInt + rightInt
        case "-":
            return leftInt - rightInt
        case "*":
            return leftInt * rightInt
        case "/":
            if rightInt != 0 {
                return leftInt / rightInt
            }
            return 0
        case "%":
            if rightInt != 0 {
                return leftInt % rightInt
            }
            return 0
        case "<":
            return leftInt < rightInt
        case ">":
            return leftInt > rightInt
        case "==":
            return leftInt == rightInt
        case "!=":
            return leftInt != rightInt
        }
    }
    
    // String operations
    if operator == "+" {
        return fmt.Sprintf("%v%v", left, right)
    }
    
    if operator == "==" {
        return fmt.Sprintf("%v", left) == fmt.Sprintf("%v", right)
    }
    
    if operator == "!=" {
        return fmt.Sprintf("%v", left) != fmt.Sprintf("%v", right)
    }
    
    return nil
}

func isTruthy(obj interface{}) bool {
    switch v := obj.(type) {
    case bool:
        return v
    case int:
        return v != 0
    case string:
        return v != ""
    case nil:
        return false
    default:
        return true
    }
}

func evalFunctionCall(call *parser.FunctionCall, env *Environment) interface{} {
    fn, ok := env.functions[call.Name]
    if !ok {
        //fmt.Printf("DEBUG: Function '%s' not found\n", call.Name)
        return nil
    }
    
    // Evaluate arguments
    args := []interface{}{}
    for _, arg := range call.Arguments {
        val := evalExpression(arg, env)
        args = append(args, val)
    }
    
    // Create new environment for function
    fnEnv := &Environment{
        variables: make(map[string]interface{}),
        functions: env.functions,
    }
    
    // Bind parameters to arguments
    for i, param := range fn.Parameters {
        if i < len(args) {
            fnEnv.variables[param] = args[i]
        }
    }
    
    // Execute function body
    result := evalBlockStatement(fn.Body, fnEnv)
    
    //fmt.Printf("DEBUG: Function '%s' result: %v (type: %T)\n", call.Name, result, result)
    
    // Handle return value
    if returnVal, ok := result.(*ReturnValue); ok {
        //fmt.Printf("DEBUG: Returning value: %v\n", returnVal.Value)
        return returnVal.Value
    }
    
    //fmt.Printf("DEBUG: No ReturnValue, returning raw result\n")
    return result
}

// Public wrapper for evalStatement (for debugging)
func EvalStatement(stmt parser.Statement, env *Environment) interface{} {
    return evalStatement(stmt, env)
}
