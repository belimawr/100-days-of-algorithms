Bit counting
=====================

This is a very simple algorithm that counts the number of bits. 

The idea is that subtracting 1 from any positive integer the first ``1`` bit will become zero and all the other ``1`` bits will remain unchanged.

**Examples:**
```
   9   -    1   = 8
0b1001 - 0b0001 = 0b1000

   36    -     1    = 35
0b100100 - 0b000001 = 0b100011
```

When we do a bit-wise AND of the number and itself minus one, the result is the same as removing the least significant bit ``1`` of the number. If this is done until the number becomes zero, the number of interactions is equals to the number of bits ``1`` on the number

Build
-----

```
$ go build
```

Run
---

```
$ ./002_bit_count -n <number_to_count_bits>
```
