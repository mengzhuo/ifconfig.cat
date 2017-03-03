package main

import (
	"flag"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	acceptCode = "en"
)

var (
	addr         = flag.String("addr", ":8080", "listen addr")
	tlsAddr      = flag.String("tlsaddr", "", "tls listen addr")
	geoDBPath    = flag.String("geo", "", "geo DB")
	certFile     = flag.String("cert", "", "TLS certfile")
	keyFile      = flag.String("key", "", "TLS keyfile")
	versionBool  = flag.Bool("v", false, "version")
	templatePath = flag.String("tmpl", "templates/*.tpl", "template path glob")
	icoPath      = flag.String("favicon", "favicon.ico", "favicon path")
	pro          = flag.String("prometheus", "", "prometheus listener")

	Version = "dev"
)

func main() {
	flag.Parse()
	if *versionBool {
		log.Printf("Ifconfig.cat service version:%s", Version)
		flag.PrintDefaults()
		return
	}

	handler, err := NewHandler(*geoDBPath, *pro)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.LoadHTMLGlob(*templatePath)
	router.GET("/", handler.Root)
	router.GET("/json", handler.RootJson)

	if *icoPath != "" {
		router.StaticFile("favicon.ico", *icoPath)
	}

	// for multi-listen
	if *tlsAddr != "" {
		tlsAddrGroup := strings.Split(*tlsAddr, ",")
		for _, a := range tlsAddrGroup {
			go func(a string) {
				log.Print(router.RunTLS(a, *certFile, *keyFile))
			}(a)
		}
	}

	addrGroup := strings.Split(*addr, ",")
	for _, a := range addrGroup {
		go func(a string) {
			log.Print(router.Run(a))
		}(a)
	}
	wait := make(chan struct{})
	<-wait
}
