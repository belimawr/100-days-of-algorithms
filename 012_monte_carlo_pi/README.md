Monte Carlo - &pi;
==================

Monte Carlo is a method of solving problems using statistics. The key idea is that
given the probability `P` of an event occur meeting certain conditions, a computer
can be used to generate the events and `P` is equal to the number of events meeting
the conditions divided by the number of events. So:

* Let `N` be the number of generated events
* Let `n` be the number of events meeting some constraints
* Then `P = n/N`

If we get a square of length `2R` and draw a circle inside it with radius `R` we
can calculate &pi; using the ratio between the area of the circle and the area of
the square:

* The area <code>A<sub>c</sub></code> of a circle of radius `R` 
is <code>A<sub>c</sub> = &pi;R<sup>2</sup></code>
* The area of a square <code>A<sub>s</sub></code> with side length
`2R` is<code>A<sub>s</sub> = (2R)<sup>2</sup></code>
* The ration between <code>A<sub>c</sub></code> and <code>A<sub>s</sub></code> is
<code>&pi;/4</code>, here is the maths:
  * <code>k = <code>A<sub>c</sub></code>/<code>A<sub>s</sub></code></code>
  * <code>k = &pi;R<sup>2</sup>/(2R)<sup>2</sup></code>
  * <code>k = &pi;R<sup>2</sup>/4R<sup>2</sup></code>
  * <code>k = &pi;/4</code>

This means that if we pick a random point inside the square, the probability `P` of it
also be inside the circle is equals to the ratio between the areas: <code>P = &pi;/4</code>.

We also know that if we pick `N` random points inside the square, and `n` points are
inside the circle, we can calculate the provability of a point being inside the circle:
`P = n/N`.

Then we have that <code>n/N = &pi;/4</code>, therefore <code>&pi; = 4n/N</code>


Build
-----

```
$ go build
```

Run
---

```
$ ./012_monte_carlo_pi -n 10000000
```

References
----------
* http://mathfaculty.fullerton.edu/mathews/n2003/montecarlopimod.html
* http://www.eveandersson.com/pi/monte-carlo-circle
* https://medium.com/100-days-of-algorithms/day-9-monte-carlo-%CF%80-7ae010743bde
