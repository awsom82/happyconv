# XML & JSON 2-Way Converter [![GitHub Actions](https://github.com/awsom82/happyconv/workflows/Go/badge.svg)](https://github.com/awsom82/happyconv/actions?workflow=Go)
This work is done just in a test case. So, it have some pitfails and flaws.

## Installing
Just execute in terminal `git clone github.com/awsom82/happyconv && cd happyconv` and run command `go build ./cmd/webconv && ./webconv`.
This will run conversion http service at port 8080.

## Using
After your run this app, you should able to send any JSON or XML file to `http://localhost:8080/`.

Notice, there no specific path for JSON or XML. The application will detect an input type of file by a mime-type header, or if it lacks that info. It will try to detect that by file signature [MIME Sniffing](https://mimesniff.spec.whatwg.org)

### Configuration
You can just type `./webconv --help` to get help message.
```
Usage of ./webconv:

  -hostname string
  	Bind server address (default "localhost")
  -keep-alive
  	HTTP Keep-Alive
  -port int
  	Port number (default 8080)
  -rate float
  	Rate limiter (default 200000)
  -read-timeout duration
  	HTTP Read timeout (default 5s)
  -ttl duration
  	Rate limiter TTL (default 5s)
  -write-timeout duration
  	HTTP Write timeout (default 10s)
```

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
BenchmarkParallelTestServeHTTP-8   	   43688	     45218 ns/op	   20037 B/op	     154 allocs/op
BenchmarkJson2Xml-8                	   41541	     33855 ns/op	   41864 B/op	     211 allocs/op
BenchmarkXml2Json-8                	   31566	     38070 ns/op	   19704 B/op	     406 allocs/op
PASS
ok  	github.com/awsom82/happyconv	5.550s
```

### wrk
Good start point for testing is run app with next keys: `./webconv -rate 3e5 -read-timeout 0.5s -write-timeout 1.3s -keep-alive 1`.
Use `wrk -t12 -c400 -d30s -swrk-post.lua  http://localhost:8080/` for simply load test

#### Beat next results
JSON  *104.00k rps* is best and XML 14.05k rps is best result 

#### json-small
```
./wrk -t1 -c9000 -d30s -s./../happyconv/json.lua http://localhost:8080/
Running 30s test @ http://localhost:8080/
  1 threads and 9000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    33.93ms   70.66ms   1.07s    95.39%
    Req/Sec    37.26k    26.31k  104.00k    64.34%
  925155 requests in 30.06s, 292.92MB read
  Socket errors: connect 7980, read 5, write 0, timeout 1
Requests/sec:  30774.50
Transfer/sec:      9.74MB
```

#### books.xml
```
./wrk -t1 -c9400 -d30s -s./../happyconv/wrk-post.lua http://localhost:8080/
Running 30s test @ http://localhost:8080/
  1 threads and 9400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    80.75ms   95.77ms   1.93s    85.83%
    Req/Sec    12.70k     0.87k   14.05k    97.98%
  376629 requests in 30.07s, 1.24GB read
  Socket errors: connect 8380, read 0, write 0, timeout 4
  Non-2xx or 3xx responses: 379
Requests/sec:  12524.26
Transfer/sec:     42.39MB

```
