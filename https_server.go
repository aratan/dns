package main

import (
    "crypto/tls"
    "github.com/miekg/dns"
    "log"
    "net/http"
    "io/ioutil"
    "time"
)

type dnsHTTPHandler struct {
    dnsServer string
    timeout   time.Duration
}

func (h *dnsHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
    }

    msg := new(dns.Msg)
    if err := msg.Unpack(body); err != nil {
        http.Error(w, "Failed to unpack DNS message", http.StatusBadRequest)
        return
    }

    response := new(dns.Msg)
    response.SetReply(msg)
    response.Authoritative = true

    for _, question := range msg.Question {
        answers := resolver(question.Name, question.Qtype, h.dnsServer, h.timeout)
        response.Answer = append(response.Answer, answers...)
    }

    respBody, err := response.Pack()
    if err != nil {
        http.Error(w, "Failed to pack DNS response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/dns-message")
    w.Write(respBody)
}

func StartHTTPSServer(config *Config) {
    handler := &dnsHTTPHandler{
        dnsServer: config.DNSServer.DNSClient.Server,
        timeout:   time.Duration(config.DNSServer.DNSClient.Timeout) * time.Second,
    }

    server := &http.Server{
        Addr: config.DNSServer.HTTPS.Address,
        Handler: handler,
        TLSConfig: &tls.Config{
            MinVersion: tls.VersionTLS12,
        },
    }

    log.Println("Starting HTTPS DNS server on port 443")

    err := server.ListenAndServeTLS(config.DNSServer.HTTPS.CertFile, config.DNSServer.HTTPS.KeyFile)
    if err != nil {
        log.Fatalf("[ERROR] Failed to start HTTPS server: %s\n", err.Error())
    }
}
