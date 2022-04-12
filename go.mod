module github.com/greenpau/caddy-auth-portal

go 1.16

replace (
	github.com/greenpau/caddy-auth-portal => ./
)

require (
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/caddyserver/caddy/v2 v2.4.6
	github.com/crewjam/saml v0.4.5
	github.com/go-ldap/ldap/v3 v3.4.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v4 v4.1.0
	github.com/google/go-cmp v0.5.6
	github.com/greenpau/caddy-authorize v1.3.24
	github.com/greenpau/caddy-trace v1.1.8
	github.com/greenpau/go-identity v1.1.6
	github.com/iancoleman/strcase v0.2.0
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattrobenolt/go-memcached v0.0.0-20130819063329-ed41d12c2de1
	github.com/satori/go.uuid v1.2.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/stretchr/testify v1.7.0
	github.com/yuin/goldmark v1.4.11 // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412071739-889880a91fd5 // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
)
