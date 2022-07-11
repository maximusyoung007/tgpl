package chapter2

import (
    "fmt"
    "os"
    "strconv"
    "tgpl/tempconv"
)

func Ex2_2() {
    for _, arg := range os.Args[1:] {
        _, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
            os.Exit(1)
        }
        m := tempconv.Meter(1000)
        k := tempconv.KiloMeter(1)
        fmt.Printf("%s = %s, %s = %s\n", m, tempconv.MToK(m), k, tempconv.KToM(k))
    }
}
