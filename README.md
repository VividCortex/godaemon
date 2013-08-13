godaemon
========

Daemonize Go applications with `exec()` instead of `fork()`.

You can't daemonize in Go. Daemonizing is a Unix concept that requires
some specific things you can't do in Go. But you can simulate it pretty
accurately, if you don't mind that your program will start copies of itself
several times. Thus, a Go Daemon isn't truly deamonic, and so we present an angelic cat picture:

![Angelic Cat](http://f.cl.ly/items/2b0y0n3W2W1H0S1K3g0g/angelic-cat.jpg)

### History

An earlier version of this concept with a slightly different interface was
developed internally at VividCortex.
