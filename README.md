# Go vs Java reactor load test


| Service                   | Technologies                   |
| :------------------------ | :----------------------------- |
| go-service                | Go 1.13 + Chi                  |
| reactive-netty-service    | Spring Boot 2 web-flux + Netty |
| blocking-undertow-service | Spring Boot 2 web + Undertow   |

## Vegeta load test
* Download `vegeta` from https://github.com/tsenart/vegeta/releases
* Start services: `docker-compose up`
* Test: `echo "GET {target}/sleep?time=1s" | vegeta attack -rate=1000/s -duration=60s | vegeta report`, where `{target}` is:
  * `http://dockerhost:8081` for _go-service_
  * `http://dockerhost:8082` for _reactive-netty-service_
  * `http://dockerhost:8083` for _blocking-undertow-service_