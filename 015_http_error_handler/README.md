HTTP Error Handler
==================


Every time I need to write a micro-service every HTTP handler have some validations
of the input that lead to the repetition of this piece of code:

```go
w.WriteHeader(http.StatusBadRequest) // Or any other error
w.Write([]byte("my error message"))
return
```

So to avoid this repetition I always create a wrapper to use a handler that returns an error
and this error holds the appropriate status code. So today piece of code is this implementation
and an example of usage. In a near future it might become a wee library

* The implementation is on [handler.go](handler.go)
* The example is on [example.go](example.go)

I know that I should have used packages here, but I'm avoiding it so if
you can clone this repository anywhere and just ``go build`` the code

Build
-----

```
$ go build
```

Run
---

```
$ ./015_http_error_handler -p 3000
```

