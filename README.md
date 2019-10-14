# XML & JSON 2-Way Converter
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
`go test -bench . -benchmem`

```
goos: darwin
goarch: amd64
BenchmarkJson2Xml-4   	   42492	     26760 ns/op	   43400 B/op	     214 allocs/op
BenchmarkXml2Json-4   	   23481	     50462 ns/op	   21045 B/op	     408 allocs/op
PASS
ok  	_/Users/awsom82/Sources/happyconv	3.151s
```
