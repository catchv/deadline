# deadlin v1.0.0


## connect-go에서 지원되지 않는 read, write timeout을 지원하는 middleware

---


* https://github.com/connectrpc/connect-go/issues/879

* http.Server{ WriteTimeout: } is set as the timeout for all handlers.

* https://connectrpc.com/docs/go/deployment#timeouts-and-connection-pools

* net/http: no way of manipulating timeouts in Handler golang/go#16100

* http.NewResponseController is supported starting with Go 1.20.

```
go get -u github.com/cathcv/deadline 
```

```
import github.com/cathcv/deadline

...

	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter,
		handlerOptions...,
	)
	mux.Handle(path, handler)

	mux1 := deadline.TimeoutMiddleware(mux, 20*time.Second)

...

	srv := &http.Server{
		Addr:              "localhost:8080",
		Handler:           h2c.NewHandler(mux1, &http2.Server{}), // Use h2c so we can serve HTTP/2 without TLS.
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      3 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
```