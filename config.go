package main

// Configuration
type Config struct {
    HTTPPort  string
    SOCKSPort string
    Users     map[string]string // username:password
}

var AppConfig = Config{
    HTTPPort:  "8080",
    SOCKSPort: "1080",
    Users: map[string]string{
        "admin": "secret123",
    },
}
