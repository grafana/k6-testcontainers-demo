# k6 Testcontainers demo

This repository contains a demo of using the [Testcontainers k6 module](https://golang.testcontainers.org/modules/k6/) for running k6 test scripts.

The demo also uses the Testcontainers [K3s module](https://golang.testcontainers.org/modules/k3s/) for running a Kubernetes cluster that is used for deploying the [QuickPizza](https://github.com/grafana/quickpizza) application.

> Note: Most of the boilerplate code for deploying the application will be simplified by an [upcoming feature in the K3s module](https://github.com/testcontainers/testcontainers-go/issues/1915).


To run the demo use the command below. Notice the `TC_BUILD_CACHE` environment variable that enables the [build cache of the k6x tool](https://golang.testcontainers.org/modules/k6/#withcache) used by the k6 module.


```sh
TC_K6_BUILD_CACHE=k6x-cache  go test demo_test.go -v
```
```
=== RUN   Test_Demo
2024/01/10 17:31:50 github.com/testcontainers/testcontainers-go - Connected to docker: 
  Server Version: 24.0.7 (via Testcontainers Desktop 1.5.5)
  API Version: 1.43
  Operating System: Ubuntu 22.04.3 LTS
  Total Memory: 63966 MB
  Resolved Docker Host: tcp://127.0.0.1:34601
  Resolved Docker Socket Path: /var/run/docker.sock
  Test SessionID: e3530bdce668381fad15d90d280e0d663a5268717b97fd10a7a5f73905864553
  Test ProcessID: 989ac33a-121c-48e8-bd67-d029c3c070a2
2024/01/10 17:31:50 🐳 Creating container for image docker.io/testcontainers/ryuk:0.5.1
2024/01/10 17:31:50 ✅ Container created: c92e1e27e221
2024/01/10 17:31:50 🐳 Starting container: c92e1e27e221
2024/01/10 17:31:51 ✅ Container started: c92e1e27e221
2024/01/10 17:31:51 🚧 Waiting for container id c92e1e27e221 image: docker.io/testcontainers/ryuk:0.5.1. Waiting for: &{Port:8080/tcp timeout:<nil> PollInterval:100ms}
2024/01/10 17:31:51 🐳 Creating container for image docker.io/rancher/k3s:v1.27.1-k3s1
2024/01/10 17:31:51 ✅ Container created: 0dc05d6fe3c7
2024/01/10 17:31:51 🐳 Starting container: 0dc05d6fe3c7
2024/01/10 17:31:51 ✅ Container started: 0dc05d6fe3c7
2024/01/10 17:31:51 🚧 Waiting for container id 0dc05d6fe3c7 image: docker.io/rancher/k3s:v1.27.1-k3s1. Waiting for: &{timeout:<nil> Log:.*Node controller sync successful.* IsRegexp:true Occurrence:1 PollInterval:100ms}
2024/01/10 17:32:19 🐳 Creating container for image szkiba/k6x:v0.3.1
2024/01/10 17:32:19 ✅ Container created: 10bb3d6db341
2024/01/10 17:32:19 🐳 Starting container: 10bb3d6db341
2024/01/10 17:32:19 ✅ Container started: 10bb3d6db341
2024/01/10 17:32:19 🚧 Waiting for container id 10bb3d6db341 image: szkiba/k6x:v0.3.1. Waiting for: &{timeout:<nil> PollInterval:100ms}
2024/01/10 17:32:30 🐳 Terminating container: 0dc05d6fe3c7
2024/01/10 17:32:31 🚫 Container terminated: 0dc05d6fe3c7
2024/01/10 17:32:31 🐳 Terminating container: 10bb3d6db341
2024/01/10 17:32:31 🚫 Container terminated: 10bb3d6db341
--- PASS: Test_Demo (41.00s)
PASS
ok      command-line-arguments  41.024s
```

> Note: the execution time the first time would be higher due to the k6x build process


## TODO

- [] Run multiple test scripts using a table driven approach
- [] Show how to reuse the `k3s` container across multiple tests
- [] Incorporate enhancements from [#1915](https://github.com/testcontainers/testcontainers-go/issues/1915)