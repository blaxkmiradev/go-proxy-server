# go-proxy-server

> A blazing-fast, multi-protocol proxy server written in Go — HTTP, HTTPS, SOCKS4, and SOCKS5 with authentication baked in.

---

## Overview

`go-proxy-server` is a modular, production-ready proxy server supporting multiple protocols from a single binary. Built for VPS deployment as a personal or team proxy with zero external infrastructure dependencies.

```
HTTP/HTTPS  ──►  :8080   (CONNECT tunneling + Basic Auth)
SOCKS4/5    ──►  :1080   (Username/password auth)
```

---

## Features

| Category       | Details |
|----------------|---------|
| **HTTP/HTTPS** | GET, POST, and all standard methods; CONNECT tunneling for HTTPS |
| **SOCKS**      | SOCKS4 and SOCKS5 with credential auth |
| **Auth**       | Per-user username/password, configurable in `config.go` |
| **Logging**    | Request timestamps, client IP, destination URLs |
| **Security**   | Optional IP restriction, modular for throttling & filtering |

---

## Project Structure

```
go-proxy-server/
├── main.go           # Entry point — starts HTTP and SOCKS listeners
├── config.go         # Ports, users, and protocol settings
├── auth.go           # Authentication logic
├── http_proxy.go     # HTTP/HTTPS proxy handler
├── socks_proxy.go    # SOCKS4/5 proxy handler
├── utils.go          # Logging and error handling helpers
├── go.mod            # Go module definition
└── README.md
```

---

## Installation

**1. Install Go (v1.21+ recommended)**

```bash
sudo apt update && sudo apt install golang-go
```

**2. Clone the repository**

```bash
git clone https://github.com/blaxkmiradev/go-proxy-server.git
cd go-proxy-server
```

**3. Install dependencies**

```bash
go mod tidy
```

---

## Running

```bash
go run main.go
```

The proxy starts two listeners:
- **HTTP/HTTPS** → `0.0.0.0:8080`
- **SOCKS4/5** → `0.0.0.0:1080`

---

## Configuration

Edit `config.go` to customize ports and credentials:

```go
var AppConfig = Config{
    HTTPPort:  "8080",
    SOCKSPort: "1080",
    Users: map[string]string{
        "admin": "secret123",
    },
}
```

| Field       | Description                          |
|-------------|--------------------------------------|
| `HTTPPort`  | Port for HTTP/HTTPS proxy            |
| `SOCKSPort` | Port for SOCKS4/5 proxy              |
| `Users`     | Map of `username → password` pairs   |

---

## Usage Examples

**HTTP proxy**
```bash
curl -x http://admin:secret123@YOUR_VPS_IP:8080 http://example.com
```

**HTTPS proxy (CONNECT tunnel)**
```bash
curl -x http://admin:secret123@YOUR_VPS_IP:8080 https://example.com
```

**SOCKS5 proxy**
```bash
curl --socks5 admin:secret123@YOUR_VPS_IP:1080 http://example.com
```

---

## Security Recommendations

- **Change default credentials** — never ship `admin:secret123` to production
- **Firewall the ports** — use `ufw` or `iptables` to whitelist trusted IPs only
- **Don't expose publicly** without authentication enabled
- **Rotate credentials** regularly for shared team deployments

```bash
# Example: allow only a specific IP to reach port 8080
ufw allow from 203.0.113.42 to any port 8080
ufw deny 8080
```

---

## Extending the Proxy

The modular structure makes it straightforward to add:

- **Rate limiting** — cap requests per user or per IP
- **Bandwidth throttling** — limit transfer speeds
- **Domain filtering** — whitelist/blacklist URLs or TLDs
- **Enhanced logging** — structured JSON logs, log rotation
- **Single-port multiplexing** — auto-detect protocol on one port

---

## Dependencies

| Package | Purpose |
|---------|---------|
| [`github.com/armon/go-socks5`](https://github.com/armon/go-socks5) | SOCKS4/5 server implementation |

```bash
go get github.com/armon/go-socks5
```

Standard library packages used: `net/http`, `net`, `log`, `bufio`, `io`.

---

## References

- [Go `net/http` package](https://pkg.go.dev/net/http)
- [Go `net` package](https://pkg.go.dev/net)
- [armon/go-socks5](https://github.com/armon/go-socks5)

---

## License

MIT — free to use, modify, and distribute.
