package main

import (
    "github.com/miekg/dns"
    "log"
    "time"
)

type dnsHandler struct {
    dnsServer string
    timeout   time.Duration
}

func (h *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
    msg := new(dns.Msg)
    msg.SetReply(r)
    msg.Authoritative = true

    for _, question := range r.Question {
        answers := resolver(question.Name, question.Qtype, h.dnsServer, h.timeout)
        msg.Answer = append(msg.Answer, answers...)
    }

    err := w.WriteMsg(msg)
    if err != nil {
        log.Printf("[ERROR] Failed to write DNS response: %v\n", err)
    }
}

func StartDNSServer(config *Config) {
    handler := &dnsHandler{
        dnsServer: config.DNSServer.DNSClient.Server,
        timeout:   time.Duration(config.DNSServer.DNSClient.Timeout) * time.Second,
    }
    server := &dns.Server{
        Addr:      config.DNSServer.Address,
        Net:       "udp",
        Handler:   handler,
        UDPSize:   65535,
        ReusePort: true,
    }

    log.Println("Starting DNS server on port 53")

    err := server.ListenAndServe()
    if err != nil {
        log.Fatalf("[ERROR] Failed to start server: %s\n", err.Error())
    }
}
