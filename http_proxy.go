package main

import (
    "io"
    "log"
    "net"
    "net/http"
    "time"
)

func HandleHTTP(w http.ResponseWriter, r *http.Request) {
    if !IsAuthorized(r) {
        RequireAuth(w)
        return
    }

    if r.Method == http.MethodConnect {
        handleHTTPS(w, r)
    } else {
        handleHTTP(w, r)
    }
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
    log.Printf("[HTTP] %s %s", r.Method, r.URL)
    client := &http.Client{Timeout: 30 * time.Second}
    req, _ := http.NewRequest(r.Method, r.URL.String(), r.Body)
    req.Header = r.Header

    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }
    defer resp.Body.Close()

    for k, v := range resp.Header {
        for _, vv := range v {
            w.Header().Add(k, vv)
        }
    }

    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}

func handleHTTPS(w http.ResponseWriter, r *http.Request) {
    log.Printf("[HTTPS] CONNECT %s", r.Host)
    destConn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }

    hijacker, ok := w.(http.Hijacker)
    if !ok {
        http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
        return
    }

    clientConn, _, err := hijacker.Hijack()
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }

    clientConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

    go transfer(destConn, clientConn)
    go transfer(clientConn, destConn)
}

func transfer(dst io.WriteCloser, src io.ReadCloser) {
    defer dst.Close()
    defer src.Close()
    io.Copy(dst, src)
}
