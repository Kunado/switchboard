package main

import (
	"switchboard/dns"
	"switchboard/web"
)

func main() {
	go dns.DnsServer()
	web.HttpServer()
}
