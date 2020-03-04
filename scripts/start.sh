#!/bin/bash

kill $(pgrep -f nginx)

CONF=$(mktemp)
echo "
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    server {
        listen       80;
        server_name  localhost;

        location / {
          mirror /stats;
          proxy_pass http://localhost:8080;
        }

        location /stats {
          internal;
          proxy_pass http://localhost:8282;
        }
    }
}
" > $CONF

nginx -c $CONF

./counter -port 8181 -path /usr/local/counter | ./shorturl -port 8080 -host localhost -db /usr/local/indexdb/ -counter http://localhost:8181/