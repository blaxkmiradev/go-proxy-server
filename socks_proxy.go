package main

import (
    "log"
    "net"

    "github.com/armon/go-socks5"
)

func StartSOCKS() {
    conf := &socks5.Config{
        AuthMethods: []socks5.Authenticator{
            socks5.UserPassAuthenticator{
                Credentials: AppConfig.Users,
            },
        },
        Logger: log.Default(),
    }

    server, err := socks5.New(conf)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("SOCKS5 server running on port", AppConfig.SOCKSPort)
    if err := server.ListenAndServe("tcp", ":"+AppConfig.SOCKSPort); err != nil {
        log.Fatal(err)
    }
}
