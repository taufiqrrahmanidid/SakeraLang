package main

import (
    "bufio"
    "fmt"
    "os"
    
    "github.com/taufiqrrahmanidid/SakeraLang/pkg/evaluator"
)

func main() {

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
