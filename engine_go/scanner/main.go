package main

import (
    "fmt"
    "os/exec"
    "strings"
)

func pingHost(ip string) {
    cmd := exec.Command("ping", "-n", "1", "-w", "1000", ip) // Windows args
    out, err := cmd.Output()
    if err == nil && strings.Contains(string(out), "TTL=") {
        fmt.Println("Host up:", ip)
    } else {
        fmt.Println("No reply:", ip)
    }
}

func main() {
    hosts := []string{
        "8.8.8.8",   // Google DNS
        "8.8.4.4",   // Google DNS
        "1.1.1.1",   // Cloudflare DNS
        "9.9.9.9",   // Quad9 DNS
    }

    for _, ip := range hosts {
        pingHost(ip)
    }
}
