# DNS

## version basica sin cifrado 

```
./dns.exe
```

```bash
Starting DNS server on port 53
example.com.    2189    IN      A       93.184.215.14
example.com.    2960    IN      AAAA    2606:2800:21f:cb07:6820:80da:af6b:8b2c
youtube.es.     300     IN      A       216.58.215.174
youtube.es.     300     IN      AAAA    2a00:1450:4003:806::200e
```

## nslookup youtube.es 127.0.0.1
## nslookup -query=TXT dataarchitectai.link 127.0.0.1

Servidor:  UnKnown
Address:  127.0.0.1

Nombre:  youtube.es
Addresses:  2a00:1450:4003:806::200e
          216.58.215.174

##  go get github.com/miekg/dns

# La ultima actualizacion va con cifrado:

### Instala openssl en windows:  https://slproweb.com/download/Win64OpenSSL_Light-3_3_1.msi

```bash
openssl genpkey -algorithm RSA -out key.pem
openssl req -new -key key.pem -out cert.csr
openssl x509 -req -days 365 -in cert.csr -signkey key.pem -out cert.pem
```

``` bash
netstat -an | findstr :53
netstat -an | findstr :443
nslookup -port=53 www.example.com 127.0.0.1
curl -X POST -H "Content-Type: application/dns-message" --data-binary @dns_query.bin https://dns.google/dns-query
```

## Respuesta por curl:

```bash
<!DOCTYPE html>
<html lang=en>
  <meta charset=utf-8>
  <meta name=viewport content="initial-scale=1, minimum-scale=1, width=device-width">
  <title>Error 400 (Bad Request)!!1</title>
  <style>
    *{margin:0;padding:0}html,code{font:15px/22px arial,sans-serif}html{background:#fff;color:#222;padding:15px}body{margin:7% auto 0;max-width:390px;min-height:180px;padding:30px 0 15px}* > body{background:url(//www.google.com/images/errors/robot.png) 100% 5px no-repeat;padding-right:205px}p{margin:11px 0 22px;overflow:hidden}ins{color:#777;text-decoration:none}a img{border:0}@media screen and (max-width:772px){body{background:none;margin-top:0;max-width:none;padding-right:0}}#logo{background:url(//www.google.com/images/branding/googlelogo/1x/googlelogo_color_150x54dp.png) no-repeat;margin-left:-5px}@media only screen and (min-resolution:192dpi){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat 0% 0%/100% 100%;-moz-border-image:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) 0}}@media only screen and (-webkit-min-device-pixel-ratio:2){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat;-webkit-background-size:100% 100%}}#logo{display:inline-block;height:54px;width:150px}
  </style>
  <a href=//www.google.com/><span id=logo aria-label=Google></span></a>
  <p><b>400.</b> <ins>That’s an error.</ins>
  <p>Your client has issued a malformed or illegal request. DNS message less than 12 bytes long. <ins>That’s all we know.</ins>
```

