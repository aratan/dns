package main

import (
    "log"
    "os"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Config struct {
    DNSServer struct {
        Address  string `yaml:"address"`
        LogFile  string `yaml:"log_file"`
        DNSClient struct {
            Timeout int    `yaml:"timeout"`
            Server  string `yaml:"server"`
        } `yaml:"dns_client"`
        HTTPS struct {
            Address  string `yaml:"address"`
            CertFile string `yaml:"cert_file"`
            KeyFile  string `yaml:"key_file"`
        } `yaml:"https"`
    } `yaml:"dns_server"`
}

func loadConfig() (*Config, error) {
    var config Config
    data, err := ioutil.ReadFile("config.yml")
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}

func main() {
    config, err := loadConfig()
    if err != nil {
        log.Fatalf("[ERROR] Failed to load config: %v\n", err)
    }

    logFile, err := os.OpenFile(config.DNSServer.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("[ERROR] Failed to open log file: %v\n", err)
    }
    defer logFile.Close()

    log.SetOutput(logFile)
    log.Println("DNS server is starting...")

    go StartDNSServer(config)
    go StartHTTPSServer(config)

    // Mantener el programa en ejecución
    select {}
}
