package main

import (
    "fmt"
    "github.com/miekg/dns"
    "log"
    "os"
    "time"
)

// resolver realiza la resolución DNS para un dominio y tipo de consulta dados.
func resolver(domain string, qtype uint16) []dns.RR {
    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(domain), qtype)
    m.RecursionDesired = true

    c := &dns.Client{Timeout: 5 * time.Second}

    response, _, err := c.Exchange(m, "8.8.8.8:53")
    if err != nil {
        log.Printf("[ERROR] DNS Exchange failed: %v\n", err)
        return nil
    }

    if response == nil {
        log.Println("[ERROR] No response from DNS server")
        return nil
    }

    for _, answer := range response.Answer {
        fmt.Printf("%s\n", answer.String())
    }

    return response.Answer
}

type dnsHandler struct{}

// ServeDNS maneja las consultas DNS entrantes y envía respuestas.
func (h *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
    msg := new(dns.Msg)
    msg.SetReply(r)
    msg.Authoritative = true

    for _, question := range r.Question {
        answers := resolver(question.Name, question.Qtype)
        msg.Answer = append(msg.Answer, answers...)
    }

    err := w.WriteMsg(msg)
    if err != nil {
        log.Printf("[ERROR] Failed to write DNS response: %v\n", err)
    }
}

// StartDNSServer inicia el servidor DNS.
func StartDNSServer() {
    handler := new(dnsHandler)
    server := &dns.Server{
        Addr:      ":53",
        Net:       "udp",
        Handler:   handler,
        UDPSize:   65535,
        ReusePort: true,
    }

    fmt.Println("Starting DNS server on port 53")

    err := server.ListenAndServe()
    if err != nil {
        log.Fatalf("[ERROR] Failed to start server: %s\n", err.Error())
    }
}

func main() {
    logFile, err := os.OpenFile("dns_server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("[ERROR] Failed to open log file: %v\n", err)
    }
    defer logFile.Close()

    log.SetOutput(logFile)
    log.Println("DNS server is starting...")

    StartDNSServer()
}
