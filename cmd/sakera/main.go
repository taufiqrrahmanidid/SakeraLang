package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    
    "github.com/taufiqrrahmanidid/SakeraLang/pkg/evaluator"
)

func main() {
    // Jika ada argument, jalankan file
    if len(os.Args) > 1 {
        runFile(os.Args[1])
        return
    }
    
    // Jika tidak, jalankan REPL
    runREPL()
}

func runFile(filename string) {
    content, err := os.ReadFile(filename)
    if err != nil {
        fmt.Printf("Error: tidak bisa membuka file '%s'\n", filename)
        return
    }
    
    input := string(content)
    env := evaluator.NewEnvironment()
    
    // Cek apakah ada block statement atau function
    if strings.Contains(input, "{") || strings.Contains(input, "fungsi") {
        // Use advanced parser for block statements and functions
        evaluator.EvalProgram(input, env)
    } else {
        // Use simple line-by-line evaluation
        scanner := bufio.NewScanner(strings.NewReader(input))
        for scanner.Scan() {
            line := scanner.Text()
            line = strings.TrimSpace(line)
            
            // Skip empty lines and comments
            if line == "" || strings.HasPrefix(line, "//") {
                continue
            }
            
            evaluator.Eval(line, env)
        }
    }
}

func runREPL() {
    fmt.Println("=================================")
    fmt.Println("  Selamet Rabu e  SakeraLang v0.2.0")
    fmt.Println("  Bahasa Pemrograman Madura")
    fmt.Println("=================================")
    fmt.Println("Ketik 'keluar' untuk exit")
    fmt.Println()
    
    env := evaluator.NewEnvironment()
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        fmt.Print("sakera> ")
        scanned := scanner.Scan()
        if !scanned {
            return
        }
        
        line := scanner.Text()
        
        if line == "keluar" || line == "exit" {
            fmt.Println("Mator Sakalangkong!")
            break
        }
        
        if line == "" {
            continue
        }
        
        evaluator.Eval(line, env)
    }
}
