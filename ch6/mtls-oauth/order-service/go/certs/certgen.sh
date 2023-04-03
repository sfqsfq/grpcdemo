#!/bin/bash

domains="server.sfq.me client.sfq.me"
ips=""

mkdir -p ca

cat >> ca/ca.cnf << EOF
[req]
prompt = no
distinguished_name = dn
req_extensions = ext

[dn]
CN = root.com
emailAddress = sfq@qq.com
O = ABC
L = Pudong
ST = Shanghai
C = CN

[ext]
subjectAltName = @alt_names

[alt_names]
DNS.1 = *.root.com
DNS.2 = root.com
EOF


openssl req -x509 -nodes -days 3650 -newkey rsa:2048 \
 -keyout ca/ca.key -out ca/ca.crt -config ca/ca.cnf -sha256

for domain in $domains
do
  mkdir -p $domain
  openssl genrsa -out $domain/$domain.key 2048
cat >> $domain/$domain.cnf <<EOF
[req]
prompt = no
distinguished_name = dn
req_extensions = ext

[dn]
CN = $domain
emailAddress = sfq@qq.com
O = ABC
L = Pudong
ST = Shanghai
C = CN

[ext]
subjectAltName = @alt_names

[alt_names]
DNS.1 = *.$domain
DNS.2 = $domain
EOF

cat >> $domain/$domain.ext << EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names

[alt_names]
DNS.1 = $domain
EOF


openssl req -new -key $domain/$domain.key -out $domain/$domain.csr -config $domain/$domain.cnf
openssl x509 -req -in $domain/$domain.csr -CA ca/ca.crt -CAkey ca/ca.key -CAcreateserial -out $domain/$domain.crt -days 3650 -sha256 -extfile $domain/$domain.ext
done





for ip in $ips
do
  mkdir -p $ip
  openssl genrsa -out $ip/$ip.key 2048
cat >> $ip/$ip.cnf <<EOF
[req]
prompt = no
distinguished_name = dn
req_extensions = ext

[dn]
CN = $ip
emailAddress = sfq@qq.com
O = ABC
L = Pudong
ST = Shanghai
C = CN

[ext]
subjectAltName = @alt_names

[alt_names]
IP.1 = $ip
EOF

cat >> $ip/$ip.ext << EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names

[alt_names]
IP.1 = $ip
EOF


openssl req -new -key $ip/$ip.key -out $ip/$ip.csr -config $ip/$ip.cnf
openssl x509 -req -in $ip/$ip.csr -CA ca/ca.crt -CAkey ca/ca.key -CAcreateserial -out $ip/$ip.crt -days 3650 -sha256 -extfile $ip/$ip.ext
done
