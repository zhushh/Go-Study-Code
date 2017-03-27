package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
            os.Exit(1)
        }
        if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("Status code: %s\n", resp.Status)
        resp.Body.Close()
    }
}

