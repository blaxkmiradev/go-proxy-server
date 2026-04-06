package main

import (
    "encoding/base64"
    "net/http"
    "strings"
)

func IsAuthorized(r *http.Request) bool {
    auth := r.Header.Get("Proxy-Authorization")
    if auth == "" {
        return false
    }

    const prefix = "Basic "
    if !strings.HasPrefix(auth, prefix) {
        return false
    }

    payload, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
    if err != nil {
        return false
    }

    pair := strings.SplitN(string(payload), ":", 2)
    if len(pair) != 2 {
        return false
    }

    pwd, ok := AppConfig.Users[pair[0]]
    return ok && pwd == pair[1]
}

func RequireAuth(w http.ResponseWriter) {
    w.Header().Set("Proxy-Authenticate", `Basic realm="Proxy"`)
    http.Error(w, "Proxy Authentication Required", http.StatusProxyAuthRequired)
}
