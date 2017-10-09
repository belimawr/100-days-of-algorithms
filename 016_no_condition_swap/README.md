No Condition Swap
=================

This one is a quite interesting piece of code, it's a mix of bit operations
and maths manipulation to swap two number if the first is greater then the
second without doing any comparison (using an if statement)

If statements are not good for the speed of a code when it's being run by
the CPU because those decision branches break the pipeline of
prefetching the next operations done by the CPU.

Tricks like these can be used by compiles to optimise code

Build
-----

```
$ go build
```

Run
---

```
$ ./016_no_condition_swap -x 42 -y 1
```

References
----------

* https://medium.com/100-days-of-algorithms/day-16-no-condition-swap-99ff930de0b4
