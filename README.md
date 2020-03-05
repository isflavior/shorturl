# shorturl

*shorturl* is a HTTP-based RESTful API for managing short URLs and redirecting clients, similar to bit.ly or goo.gl.

### API

- GET /key
  - 302 redirect to the original long url
- POST /
  - body: long url
  
### Quick Start

Starts all services and setup nginx as a reverse proxy for the URL shortener service.

```
./scripts/start.sh
```
  
### ./shorturl

```
Usage: ./shorturl

  -counter string
        Endpoint to the counter API
  -db string
        Path to leveldb
  -host string
        Host of the url shortener (default "localhost")
  -port int
        Port for the api server (default 8080)
```

### ./counter

```
Usage: ./counter

  -path string
        Path to the counter data file
  -port int
        Port for the counter server (default 8181)
```

### Examples

```
# generating a short url from a long url
curl -v X POST -d "http://google.pt/" localhost

# redirecting a short url to a long url
curl -v localhost/A61hA
```

### Performance

```
ab -n 1000 -c 100 http://localhost/A61hA
```
