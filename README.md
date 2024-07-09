# DNS

## ./dns.exe
Starting DNS server on port 53
example.com.    2189    IN      A       93.184.215.14
example.com.    2960    IN      AAAA    2606:2800:21f:cb07:6820:80da:af6b:8b2c
youtube.es.     300     IN      A       216.58.215.174
youtube.es.     300     IN      AAAA    2a00:1450:4003:806::200e


## nslookup youtube.es 127.0.0.1
## nslookup -query=TXT dataarchitectai.link 127.0.0.1

Servidor:  UnKnown
Address:  127.0.0.1

Nombre:  youtube.es
Addresses:  2a00:1450:4003:806::200e
          216.58.215.174


##  go get github.com/miekg/dns


## La ultima actualizacion va con cifrado:

openssl genpkey -algorithm RSA -out key.pem
openssl req -new -key key.pem -out cert.csr
openssl x509 -req -days 365 -in cert.csr -signkey key.pem -out cert.pem



netstat -an | findstr :53
netstat -an | findstr :443

nslookup -port=53 www.example.com 127.0.0.1

curl -X POST -H "Content-Type: application/dns-message" --data-binary @dns_query.bin https://localhost:443
