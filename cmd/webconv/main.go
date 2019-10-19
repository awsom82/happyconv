package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/awsom82/happyconv"
)

var webconvUsage = func() {
	var useText string = `You should able to send any JSON or XML file localhost:8080.
Notice, there no specific path for JSON or XML.

The application will detect an input type of file by a mime-type header,
or if it lacks that info. it will try to detect that by file signature MIME Sniffing

Examples:
> http :8080 Content-type:application/xml < example.xml
> http :8080 Content-type:application/json < example.json`

	//var CommandLine = NewFlagSet(os.Args[0], flag.ExitOnError)

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n"+useText+"\n\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	flag.Usage = webconvUsage

	conf := webconv.NewConfig()

	flag.StringVar(&conf.Hostname, "hostname", "localhost", "bind server hostname")
	flag.UintVar(&conf.Port, "port", 8080, "port number")
	flag.Float64Var(&conf.RateLimit, "rate", 2e5, "Rate limiter")
	flag.DurationVar(&conf.RateLimitTTL, "ttl", conf.RateLimitTTL, "Rate limiter TTL")
	flag.Parse()

	srv := webconv.NewServer(conf)
	log.Fatal(srv.ListenAndServe())

}
