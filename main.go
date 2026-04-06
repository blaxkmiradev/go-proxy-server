package main

import (
    "log"
    "net/http"
)

func main() {
    // Start HTTP/HTTPS proxy
    go func() {
        log.Println("HTTP/HTTPS proxy running on port", AppConfig.HTTPPort)
        http.ListenAndServe(":"+AppConfig.HTTPPort, http.HandlerFunc(HandleHTTP))
    }()

    // Start SOCKS proxy
    go StartSOCKS()

    // Keep main alive
    select {}
}
