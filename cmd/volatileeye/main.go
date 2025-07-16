// cmd/volatileeye/main.go
package main

import (
"flag"
"log"
"os"

"volatileeye/internal/volatileeye"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := volatileeye.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
