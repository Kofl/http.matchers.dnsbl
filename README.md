caddy-matcher-dnsbl
===================

All Cudos goes to https://git.madhouse-project.org/caddy/http.matchers.dnsbl

A [Caddy](https://caddyserver.com/) module that adds a `dnsbl` HTTP matcher. The matcher looks up the request's `remote_addr` in the configured DNS blocklist providers, and returns a match if it is found in any of them.

## Installation

Build Caddy using [xcaddy](https://github.com/caddyserver/xcaddy):

``` shellsession
xcaddy build --with git.madhouse-project.org/caddy/http.matchers.dnsbl
```

## Syntax

Cudos to https://404.tw/archives/2025/07/31/148/use-dnsbl-to-block-ai-crawlers-in-caddy/

``` caddyfile
example.com {
  @dnsbl dnsbl {
    providers "dnsbl.dronebl.org."
  }
  respond @dnsbl 403
}

wiki.gslin.org {
        @dnsbl dnsbl {
                providers "b.barracudacentral.org." "spam.spamrats.com."
        }
        respond @dnsbl 403
}
```
