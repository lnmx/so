package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {

    flag.Usage = usage

    flag.Parse()

    args := flag.Args()

    if 0 == len(args) {
        usage()
    }

    cmd := args[0]

    switch cmd {
    case "help":
        usage()

    case "serve":
        serve()
    }

    os.Exit(0)
}

func usage() {
    fmt.Println("Usage: so command [arguments]")
    fmt.Println()
    fmt.Println("serve     run a webserver")

    os.Exit(2)
}

func serve() {

    ifc := "127.0.0.1:8080"
    dir, err := os.Getwd()

    if nil != err {
        die(err)
    }

    fmt.Printf("serve %s on %s\n", dir, ifc)

    http.Handle("/", http.FileServer(http.Dir(dir)))

    err = http.ListenAndServe(ifc, nil)

    die(err)
}

func die(err error) {
    log.Fatal(err)
}
