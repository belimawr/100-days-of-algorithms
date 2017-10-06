Vector Clocks
=============

On 28 September of 2017
[Women Who Go Berlin](https://www.meetup.com/Women-Who-Go-Berlin/events/242953531/)
organised a meetup open for everyone where the challenge was to implement some 
functions of a vector clock. Those functions are:
* ``Timestamp.HappensBefore``
* ``Timestamp.Merge``
* ``MajorityResolver.Resolve``

Original repository, with the base code, is [wwgberlin/timelord](https://github.com/wwgberlin/timelord) and my implementations are on [vc/implemented.go](vc/implemented.go). I also implemented a few more tests that are on []vc/vc_test.go](vc/vc_test.go)

Build
-----

```
$ go build
```

Run
---

```
$ ./011_vector_clocks
```

TODO
----
* Double check if the code compiles in a "clean environment"

References
----------
* https://en.wikipedia.org/wiki/Vector_clock
* https://github.com/wwgberlin/timelord
* http://basho.com/posts/technical/why-vector-clocks-are-easy/
