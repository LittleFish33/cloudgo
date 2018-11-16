# 简单Web应用开发 Cloudgo

[项目地址](https://github.com/LittleFish33/Cloudgo)

## 1. 项目展示

该项目实现了一个简单的Web应用，具体效果如下：

1. 输入地址，进入欢迎界面，时间为当前时间

![](https://littlefish33.cn/image/temp/193.png)

2. 输入"/static/" + 资源名字可以访问静态资源

![](https://littlefish33.cn/image/temp/194.png)

3. hello/{id}

![](https://littlefish33.cn/image/temp/195.png)

## 2. curl 测试和ab测试

**curl测试**

```
F:\Users\HP\GOPATH>curl -v http://localhost:8080/hello/1233
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /hello/1233 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Fri, 16 Nov 2018 12:43:54 GMT
< Content-Length: 27
<
{
  "Test": "Hello 1233"
}
* Connection #0 to host localhost left intact

F:\Users\HP\GOPATH>curl http://localhost:8080/hello/1233
{
  "Test": "Hello 1233"
}
```

**ab测试**

部分参数：

* -n 在测试会话中所执行的请求个数（总数）
* -c 一次产生的请求个数（单次请求次数）
* -t 测试所进行的最大秒数。其内部隐含值是-n 50000，它可以使对服务器的测试限制在一个固定的总时间以内。默认时，没有时间限制。
* -p 包含了需要POST的数据的文件。



```
C:\Users\HP>ab -n 1000 -c 100 http://localhost:8080/hello/1233
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /hello/1233
Document Length:        27 bytes

Concurrency Level:      100
Time taken for tests:   0.981 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      150000 bytes
HTML transferred:       27000 bytes
Requests per second:    1019.76 [#/sec] (mean)
Time per request:       98.062 [ms] (mean)
Time per request:       0.981 [ms] (mean, across all concurrent requests)
Transfer rate:          149.38 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.9      1       8
Processing:     9   94  40.3     75     188
Waiting:        0   55  38.7     46     172
Total:         10   94  40.6     75     189

Percentage of the requests served within a certain time (ms)
  50%     75
  66%     96
  75%    132
  80%    139
  90%    164
  95%    177
  98%    184
  99%    187
 100%    189 (longest request)
```