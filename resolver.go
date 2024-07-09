package main

import (
    "github.com/miekg/dns"
    "log"
    "time"
)

func resolver(domain string, qtype uint16, dnsServer string, timeout time.Duration) []dns.RR {
    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(domain), qtype)
    m.RecursionDesired = true

    c := &dns.Client{Timeout: timeout}

    response, _, err := c.Exchange(m, dnsServer)
    if err != nil {
        log.Printf("[ERROR] DNS Exchange failed: %v\n", err)
        return nil
    }

    if response == nil {
        log.Println("[ERROR] No response from DNS server")
        return nil
    }

    for _, answer := range response.Answer {
        log.Printf("[INFO] %s\n", answer.String())
    }

    return response.Answer
}
