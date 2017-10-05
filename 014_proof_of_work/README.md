Proof of Work
=============

This is a very simple implementation of the "proof of work" described on
Satoshi's paper [Bitcoin: A Peer-to-Peer Electronic Cash System](https://bitcoin.org/bitcoin.pdf)

My basic implementation is:

* Get an string as data
* Add an integer to the end of this string
* Calculate the SHA256 of the resulting string
* Stop when the number of desired leading zeros is found

Build
-----

```
$ go build
```

Run
---

```
$ ./014_proof_of_work -d "Hello, world\!" -z 4
```

WARNING
-------

Depending on the chosen input it can take a very long time to find a hash with
the desired number of leading zeros


References
----------

* https://en.bitcoin.it/wiki/Proof_of_work
