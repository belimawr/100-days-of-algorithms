Shuffle
======

A small algorithm to shuffle an array with uniform distribution


Build
-----

```
$ go build
```

Run
---

```
$ ./017_shuffle
```

Input format
------------

The program reads from stdin:
* the seed for the random number generator
* the slice to be sorted, the format is:
 * The size of the slice
 * The elements of the slice, one per line

**Example:**
To enter the seed 42 and the slice [5, 4, 3, 2, 1], the input is
```
42
5
5
4
3
2
1
```


References
----------
* https://medium.com/100-days-of-algorithms/day-43-shuffle-b5abe4644c23
