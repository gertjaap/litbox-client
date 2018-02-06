package main

import (
	"github.com/gertjaap/litbox-client/pairing"
)

func main() {
	pairing.Pair("http://ya2qhylozur3rcux.onion:49280", "abc123", "socks5://172.17.0.1:9050/")
}