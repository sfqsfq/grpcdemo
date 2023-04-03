openssl genrsa -out server.key 2048
openssl req -new -x509 -days 3650 -subj "/CN=sfq.me" -addext "subjectAltName = DNS:sfq.me"  -key server.key -out server.crt