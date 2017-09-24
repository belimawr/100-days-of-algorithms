Binary Search
=============


This is a simple implementation of the [binary search algorithm](https://en.wikipedia.org/wiki/Binary_search_algorithm).


Build
-----

```
$ go build
```

Run
---

```
$ ./009_binary_search
```

Input format
------------

The program reads from stdin the number to search for and the slice to search in, the format is:

* The number to search
* The size of the slice
* The elements of the slice, one per line

**Example:**
To search for ``42`` in the slice ``[-100, 42, 100]``, the input is
```
42
3
-100
42
100
```

References
----------

* https://en.wikipedia.org/wiki/Binary_search_algorithm
* https://medium.com/100-days-of-algorithms/day-8-binary-search-842f3b700555