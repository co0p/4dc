package main

import (
    "flag"
    "log"

    "pomodoro/internal/icon"
)

func main() {
    size := flag.Int("size", 1024, "square icon size in pixels")
    out := flag.String("out", "icon.png", "output PNG path")
    flag.Parse()
    if err := icon.Generate(*size, *out); err != nil {
        log.Fatal(err)
    }
}
