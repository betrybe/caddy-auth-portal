FROM caddy:builder AS builder

RUN xcaddy build \
    --with github.com/betrybe/caddy-authorize \
    --with github.com/betrybe/caddy-auth-portal=./

FROM caddy:latest

COPY --from=builder /usr/bin/caddy /usr/bin/caddy
