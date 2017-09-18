Matrix Chain Multiplication
===========================

This is an implementation of the optimised algorithm to solve the [Matrix Chain Multiplication](https://en.wikipedia.org/wiki/Matrix_chain_multiplication) problem.

Build
-----

```
$ go build
```

Run
---

```
$ ./004_matrix_chain_multiplication
```

Input format
------------

The program reads from stdin the slice containing the matrices sizes.
Matrix A<sub>i</sub> has dimension ``dims[i-1] x dims[i]`` for ``i = 1 ... n``. To represent the matrices:
* A:``40x20``
* B:``20x30``
* C:``30x10`` 
* D:``10x30`` 

The input slice is: ``[40, 20, 30, 10, 30]``

The input format of the program is:
* The size of the slice
* The elements of the slice, one per line

**Example:**
To enter the slice [40, 20, 30, 10, 30], the input is
```
5
40
20
30
10
30
```

Resources
---------
* https://en.wikipedia.org/wiki/Matrix_chain_multiplication
* http://www.geeksforgeeks.org/dynamic-programming-set-8-matrix-chain-multiplication/
* http://www.geeksforgeeks.org/printing-brackets-matrix-chain-multiplication-problem/
* https://medium.com/100-days-of-algorithms/day-2-matrix-chain-multiplication-3ae6349c34ab


