package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/awsom82/happyconv"
)

var (
	gitHash string = "NOBUILD"
	gitTime string = "0"
)

var webconvUsage = func() {
	var useText string = `You should able to send any JSON or XML file localhost:8080.
Notice, there no specific path for JSON or XML.

The application will detect an input type of file by a mime-type header.
If it lacks that info, it will try to detect that by file signature.

Examples:
> http :8080 Content-type:application/xml < example.xml
> http :8080 Content-type:application/json < example.json`

	appVersion := fmt.Sprintf("Version:\n  Build %s at %s\n\n", strings.ToUpper(gitHash[:7]), gitTime)

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n"+useText+"\n\n"+appVersion, os.Args[0])
	flag.PrintDefaults()
}

func main() {

	flag.Usage = webconvUsage

	conf := webconv.NewConfig()

	flag.StringVar(&conf.Hostname, "hostname", conf.Hostname, "Bind server address")
	flag.UintVar(&conf.Port, "port", conf.Port, "Port number")
	flag.Float64Var(&conf.RateLimit, "rate", conf.RateLimit, "Rate limiter")
	flag.DurationVar(&conf.RateLimitTTL, "ttl", conf.RateLimitTTL, "Rate limiter TTL")
	flag.BoolVar(&conf.KeepAlive, "keep-alive", conf.KeepAlive, "HTTP Keep-Alive")
	flag.DurationVar(&conf.ReadTimeout, "read-timeout", conf.ReadTimeout, "HTTP Read timeout")
	flag.DurationVar(&conf.WriteTimeout, "write-timeout", conf.WriteTimeout, "HTTP Write timeout")

	flag.Parse()

	srv := webconv.NewServer(conf)
	log.Fatal(srv.ListenAndServe())

}
