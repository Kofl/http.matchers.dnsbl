caddy-matcher-dnsbl
===================

[![Build status][ci:badge]][ci:url]

 [ci:badge]: https://git.madhouse-project.org/caddy/http.matchers.dnsbl/actions/workflows/build.yaml/badge.svg?style=for-the-badge&label=CI
 [ci:url]: https://git.madhouse-project.org/caddy/http.matchers.dnsbl/actions/workflows/build.yaml/runs/latest

A [Caddy](https://caddyserver.com/) module that adds a `dnsbl` HTTP matcher. The matcher looks up the request's `remote_addr` in the configured DNS blocklist providers, and returns a match if it is found in any of them.

## Installation

Build Caddy using [xcaddy](https://github.com/caddyserver/xcaddy):

``` shellsession
xcaddy build --with git.madhouse-project.org/caddy/http.matchers.dnsbl
```

## Syntax

``` caddyfile
example.com {
  @dnsbl dnsbl {
    providers "dnsbl.dronebl.org."
  }
  respond @dnsbl 403
}
```
