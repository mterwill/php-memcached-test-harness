# php-memcached-test-harness

A test harness to play around with PHP's Memcached client and persistent
connections in Apache/mod_php.

`script/setup` starts up an Apache server with 5 workers and a fake Memcached
server that only responds to the command `get foo`. Any other requests hangs
indefinitely so you can explore PHP's Memcached client timeouts.

The Memcached server logs commands as they're received to paint a picture of
what the PHP client is doing. For example, the following request:

```
$ curl http://localhost:8080/00-no-pool.php
string(3) "bar"
```

yields the logs:

```
fakememcached_1  | 2019/11/18 22:00:22 172.21.0.3:53570: Got message "get foo\r\n"
fakememcached_1  | 2019/11/18 22:00:22 172.21.0.3:53570: Got message "quit\r\n"
fakememcached_1  | 2019/11/18 22:00:22 172.21.0.3:53570: Closing connection
```

where `172.21.0.3:53570` is the source IP and port of the connection.

Notice instead, if you request `01-pool.php`, you will see that the connections
are left open and that the requests round-robin through the 5 Apache workers.

`02-bad-no-pool.php` consistently takes 5 seconds:

```
$ time curl http://localhost:8080/02-bad-no-pool.php
string(51) "SERVER HAS FAILED AND IS DISABLED UNTIL TIMED RETRY"
curl http://localhost:8080/02-bad-no-pool.php  0.00s user 0.00s system 0% cpu 5.021 total
```

`03-bad-pool.php` takes 5 seconds for the first 5 requests (one to warm up each
Apache worker) then will return immediately for the next 100 seconds,
`OPT_RETRY_TIMEOUT`.

```
$ time curl http://localhost:8080/03-bad-pool.php
string(51) "SERVER HAS FAILED AND IS DISABLED UNTIL TIMED RETRY"
curl http://localhost:8080/03-bad-pool.php  0.00s user 0.00s system 49% cpu 0.016 total
```
