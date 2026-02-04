package main

import (
    "bufio"
    "fmt"
    "os"
    
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
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("Error: tidak bisa membuka file '%s'\n", filename)
        return
    }
    defer file.Close()
    
    env := evaluator.NewEnvironment()
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := scanner.Text()
        if line != "" {
            evaluator.Eval(line, env)
        }
    }
}

func runREPL() {
    fmt.Println("=================================")
    fmt.Println("  Selamet Rabu e  SakeraLang v0.1.0")
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
