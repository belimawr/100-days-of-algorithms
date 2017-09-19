Next Permutation
================

This algorithm calculates the next [permutation](https://en.wikipedia.org/wiki/Permutation) of a given sequence of integers.

The algorithm is:

1. Let ``sequence`` be the sequence to find the next permutation
1. Let ``pivot`` be the first element that is not in a non-decreasing-sub-sequence from the end to the beginning
2. Let ``sub-sequence`` be the elements after the pivot until the end
3. Swap the pivot with the first element (from the end to the beginning) bigger than it in the sub-sequence
4. Revert the ``sub-sequence`` in-place

Example:
--------
* Given the **sequence** ``[4, 2, 5, 3, 1]``
* The **sub-sequence** is ``[5, 3, 1]``
* The **pivot** is **2**
* The next permutation is ``[4, 3, 1, 2, 5]``


Build
-----

```
$ go build
```

Run
---

```
$ ./005_next_permutation
```

Input format
------------

The program reads from stdin the slice of integers to find the next permutation, the format is:

* The size of the slice
* The elements of the slice, one per line

**Example:**
To enter the slice [4, 2, 5, 3, 1], the input is
```
5
4
2
5
3
1
```

References
----------
* https://en.wikipedia.org/wiki/Permutation
* https://medium.com/100-days-of-algorithms/day-3-next-permutation-ce817f5004e3

