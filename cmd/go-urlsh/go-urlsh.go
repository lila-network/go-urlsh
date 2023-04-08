package main

import (
	"codeberg.org/lauralani/go-urlsh/internal/api"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var bindaddr string
	var trustedproxies []string
	if os.Getenv("PORT") != "" {
		bindaddr = ":" + os.Getenv("PORT")
	} else {
		bindaddr = ":2345"
	}

	if os.Getenv("BINDADDR") != "" {
		bindaddr = os.Getenv("BINDADDR") + bindaddr
	} else {
		bindaddr = "127.0.0.1" + bindaddr
	}

	if os.Getenv("TRUSTEDPROXIES") != "" {
		trustedproxies = strings.Split(os.Getenv("TRUSTEDPROXIES"), ",")
	} else {
		trustedproxies = []string{"127.0.0.1", "::1"}
	}

	r := api.SetupRouter()

	proxyerr := r.SetTrustedProxies(trustedproxies)
	if proxyerr != nil {
		panic(fmt.Errorf("value for TRUSTEDPROXIES env variable is not valid: %q", proxyerr))
	}

	err := r.Run(bindaddr)
	if err != nil {
		log.Fatal(err)
	}
}
