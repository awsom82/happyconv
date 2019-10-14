# XML & JSON 2-Way Converter [![Go Report Card](https://goreportcard.com/badge/github.com/awsom82/happyconv)](https://goreportcard.com/report/github.com/awsom82/happyconv)
This work done just in test case

## Installing

Just run `git clone github.com/awsom82/happyconv && cd happyconv` and run by `go build && ./happyconv`.
This will run conversion http service at port 8080.

## Using
After you run app, you shoud able send any JSON or XML file to `http://localhost:8080/`.

Notice, there no specifict path for JSON or XML. The application will detect input type of file by mime-type header, ot, if it lack that info. it will try to detect that by file signature [MIME Sniffing](https://mimesniff.spec.whatwg.org)

### Examples
```
$ http :8080 Content-type:application/xml < example.xml
$ http :8080 Content-type:application/json < example.json
```


## Testing & Benchmarking

Test with `go test -v .`

and run benchmarks 
`go test -bench . -benchmem -parallel 24 -cpu 4`

```
BenchmarkParallelTestServeHTTP-4   	    7641	    339199 ns/op	   19842 B/op	     154 allocs/op
BenchmarkJson2Xml-4                	   43954	     26005 ns/op	   41864 B/op	     211 allocs/op
BenchmarkXml2Json-4                	   25051	     47962 ns/op	   19701 B/op	     406 allocs/op
PASS
ok  	github.com/awsom82/happyconv	6.689s
```
