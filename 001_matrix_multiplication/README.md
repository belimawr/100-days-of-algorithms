Matrix multiplication
=====================

This is a simple implementation of the iterative algorithm
for [matrix multiplication](https://en.wikipedia.org/wiki/Matrix_multiplication_algorithm#Iterative_algorithm)

Build
-----

```
$ go build
```

Run
---

```
$ ./001_matrix_multiplication < input
```

Input format
------------

The program reads from stdin the two matrices, the format is:

* The size of the first matrix, separated by comma
* Each line of the first matrix, elements separated by spaces
* The size of the second matrix, separated by comma
* Each line of the second matrix, elements separated by spaces

**Example:**
To multiply the matrix:
```
1 2
3 4
```
by
```
5 6
7 8
```

the input is:
```
2, 2
1 2
3 4
2, 2
2 0
1 2

```

