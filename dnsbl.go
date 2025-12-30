// SPDX-FileCopyrightText: 2024 Gergely Nagy
// SPDX-FileContributor: Gergely Nagy
//
// SPDX-License-Identifier: EUPL-1.2

package dnsbl

import (
	"fmt"
	"net"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(DNSBL{})
}

type DNSBL struct {
	Providers []string `json:"providers"`

	logger *zap.Logger
}

func (DNSBL) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.matchers.dnsbl",
		New: func() caddy.Module { return new(DNSBL) },
	}
}

func (m *DNSBL) Provision(ctx caddy.Context) error {
	m.logger = ctx.Logger()
	return nil
}

func (m *DNSBL) Validate() error {
	if m.Providers == nil || len(m.Providers) == 0 {
		return fmt.Errorf("Specifying at least one Provider is required")
	}

	return nil
}

func (m *DNSBL) Match(req *http.Request) bool {
	m.logger.Debug("matching against", zap.String("remote_addr", req.RemoteAddr))

	remote_ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		m.logger.Error("Error parsing remote addr into IP & port", zap.String("remote_addr", req.RemoteAddr), zap.Error(err))
		return false
	}

	reverse_addr, err := reverseaddr(remote_ip)
	if err != nil {
		m.logger.Error("Error reversing remote addr", zap.String("remote_addr", remote_ip), zap.Error(err))
		return false
	}
	for _, provider := range m.Providers {
		addr := reverse_addr + "." + provider
		m.logger.Debug("Looking up in DNSBL", zap.String("addr", addr), zap.String("provider", provider))
		result, err := net.LookupHost(addr)
		if err == nil || len(result) > 0 {
			m.logger.Info("Found in DNSBL", zap.String("addr", addr), zap.String("provider", provider))
			return true
		}
	}

	m.logger.Debug("Not found in DNSBL", zap.String("remote_addr", remote_ip))
	return false
}

func (m *DNSBL) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	d.Next()

	for nesting := d.Nesting(); d.NextBlock(nesting); {
		switch d.Val() {
		case "providers":
			for d.NextArg() {
				m.Providers = append(m.Providers, d.Val())
			}
		default:
			return d.SyntaxErr(fmt.Sprintf("unexpected directive: %s", d.Val()))
		}
	}

	return nil
}

var (
	_ caddy.Validator          = (*DNSBL)(nil)
	_ caddy.Provisioner        = (*DNSBL)(nil)
	_ caddyhttp.RequestMatcher = (*DNSBL)(nil)
	_ caddyfile.Unmarshaler    = (*DNSBL)(nil)
)
