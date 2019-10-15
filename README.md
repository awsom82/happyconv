# XML & JSON 2-Way Converter [![GitHub Actions](https://github.com/awsom82/happyconv/workflows/Go/badge.svg)](https://github.com/awsom82/happyconv/actions?workflow=Go) [![Go Report Card](https://goreportcard.com/badge/github.com/awsom82/happyconv)](https://goreportcard.com/report/github.com/awsom82/happyconv) [![Coverage](http://gocover.io/_badge/github.com/awsom82/happyconv)](http://gocover.io/github.com/awsom82/happyconv) [![license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/awsom82/happyconv/master/LICENSE)
This work is done just in a test case

## Installing
Just execute in terminal `git clone github.com/awsom82/happyconv && cd happyconv` and run command `go build && ./happyconv`.
This will run conversion http service at port 8080.

## Using
After your run this app, you should able to send any JSON or XML file to `http://localhost:8080/`.

Notice, there no specific path for JSON or XML. The application will detect an input type of file by a mime-type header, or if it lacks that info. it will try to detect that by file signature [MIME Sniffing](https://mimesniff.spec.whatwg.org)

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
