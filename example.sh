#!/bin/bash
set -o verbose

curl localhost:8080/storage -XPOST -d '{"method":"SET", "key":"key_1", "value":"value_1"}'

echo '{"method":"GET", "key":"key_1"}' | nc localhost 8081

curl localhost:8080/storage -XPOST -d '{"method":"EXISTS", "key":"key_1"}'

curl localhost:8080/storage -XPOST -d '{"method":"DELETE", "key":"key_1"}'

echo '{"method":"EXISTS", "key":"key_1"}' | nc localhost 8081

echo '{"method":"GET", "key":"key_55555555555555555555555555555555555555555555555555551"}' | nc localhost 8081