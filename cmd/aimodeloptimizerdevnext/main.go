// cmd/aimodeloptimizerdevnext/main.go
package main

import (
"flag"
"log"
"os"

"aimodeloptimizerdevnext/internal/aimodeloptimizerdevnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := aimodeloptimizerdevnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
