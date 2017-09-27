Calc Stats
==========

This is a simple solution to the "Calc Stats" problem of [CyberDojo](http://cyber-dojo.org/).

The task is:

>Your task is to process a sequence of integer numbers
>to determine the following statistics:
>
>    * minimum value
>    * maximum value
>    * number of elements in the sequence
>    * average value
>
>For example: [6, 9, 15, -2, 92, 11]
>
>    * minimum value = -2
>    * maximum value = 92
>    * number of elements in the sequence = 6
>    * average value = 21.833333

It is really simple and can be done in less than 30 minutes. I decided to solvethis problem today because:

1. It was (maybe is?) used in a company I worked for as coding test to hire developers of any level. It is unbelievable
the number of people that could not solve that :/

2. I wanted something really quick and interesting to implement today.

Build
-----

```
$ go build
```

Run
---

```
$ ./010_calc_stats
```

Input format
------------

The program reads from stdin the size of the slice and the elements, the format is:

* The size of the slice
* The elements of the slice, one per line

**Example:**
To input the slice ``[6, 9, 15, -2, 92, 11]``

```
6
6
9
15
-2
92
11
```
