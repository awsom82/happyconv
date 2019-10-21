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
BenchmarkParallelTestServeHTTP-4   	    7641	    339199 ns/op	   19842 B/op	     154 allocs/op
BenchmarkJson2Xml-4                	   43954	     26005 ns/op	   41864 B/op	     211 allocs/op
BenchmarkXml2Json-4                	   25051	     47962 ns/op	   19701 B/op	     406 allocs/op
PASS
ok  	github.com/awsom82/happyconv	6.689s
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

### hey

#### example.xml
`/hey_linux_amd64  -n 1200 -c 400 -z 5s -m POST -t 3 -D ./../happyconv/example.xml -T text/xml http://localhost:8080/`

```
Summary:
  Total:	5.5673 secs
  Slowest:	1.5055 secs
  Fastest:	0.0003 secs
  Average:	0.0237 secs
  Requests/sec:	14928.8483
  
  Total data:	43530776 bytes
  Size/request:	524 bytes

Response time histogram:
  0.000 [1]	|
  0.151 [81854]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.301 [8]	|
  0.452 [0]	|
  0.602 [7]	|
  0.753 [26]	|
  0.903 [45]	|
  1.054 [1007]	|
  1.204 [6]	|
  1.355 [97]	|
  1.506 [23]	|


Latency distribution:
  10% in 0.0021 secs
  25% in 0.0043 secs
  50% in 0.0082 secs
  75% in 0.0117 secs
  90% in 0.0159 secs
  95% in 0.0203 secs
  99% in 1.0163 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0134 secs, 0.0003 secs, 1.5055 secs
  DNS-lookup:	0.0004 secs, 0.0000 secs, 0.0638 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0124 secs
  resp wait:	0.0078 secs, 0.0001 secs, 0.4161 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0124 secs

Status code distribution:
  [200]	83074 responses

Error distribution:
  [39]	Post http://localhost:8080/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
```

#### books.xml
`./hey_linux_amd64  -n 1200 -c 400 -z 30s -m POST -t 3 -D ./../happyconv/books.xml -T text/xml http://localhost:8080/`
```
root@Whirlpool:/var/local/hey# 

Summary:
  Total:	30.9133 secs
  Slowest:	2.6862 secs
  Fastest:	0.0005 secs
  Average:	0.0630 secs
  Requests/sec:	6105.4221
  

Response time histogram:
  0.000 [1]	|
  0.269 [180788]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.538 [129]	|
  0.806 [1509]	|
  1.075 [5679]	|■
  1.343 [530]	|
  1.612 [0]	|
  1.881 [0]	|
  2.149 [0]	|
  2.418 [0]	|
  2.686 [14]	|


Latency distribution:
  10% in 0.0025 secs
  25% in 0.0151 secs
  50% in 0.0239 secs
  75% in 0.0359 secs
  90% in 0.0550 secs
  95% in 0.0834 secs
  99% in 1.0402 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0085 secs, 0.0005 secs, 2.6862 secs
  DNS-lookup:	0.0001 secs, 0.0000 secs, 0.0082 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0411 secs
  resp wait:	0.0256 secs, 0.0003 secs, 0.3068 secs
  resp read:	0.0000 secs, 0.0000 secs, 0.0214 secs

Status code distribution:
  [200]	188650 responses

Error distribution:
  [68]	Post http://localhost:8080/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49206->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49214->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49222->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49228->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49232->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49236->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49242->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49246->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49252->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49260->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49266->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49272->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49320->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49334->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49340->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49346->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49372->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49376->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49380->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49438->127.0.0.1:8080: read: connection reset by peer
  [1]	Post http://localhost:8080/: read tcp 127.0.0.1:49452->127.0.0.1:8080: read: connection reset by peer

```
