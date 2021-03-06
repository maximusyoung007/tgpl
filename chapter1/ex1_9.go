package chapter1

import (
    "fmt"
    "net/http"
    "os"
    "strings"
)

func Ex1_9() {
    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url, "http://") {
            url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        //_, err = io.Copy(io.Closer() ,resp.Body)
        fmt.Printf("status code: %s", resp.Status)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
    }
}
