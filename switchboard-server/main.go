package main

import (
	"switchboard-server/dns"
	"switchboard-server/web"
)

func main() {
	go dns.DnsServer()
	web.HttpServer()
}
