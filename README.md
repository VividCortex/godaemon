godaemon
========

Daemonize Go applications with `exec()` instead of `fork()`.

You can't daemonize the usual way in Go. Daemonizing is a Unix concept that requires
some [specific things](http://goo.gl/vTUsVy) you can't do
easily in Go. But you can still accomplish the same goals 
if you don't mind that your program will start copies of itself
several times, as opposed to using `fork()` the way many programmers are accustomed to doing.

A Go Daemon is a good thing, and so we present an angelic cat picture:

![Angelic Cat](http://f.cl.ly/items/2b0y0n3W2W1H0S1K3g0g/angelic-cat.jpg)

### History

An earlier version of this concept with a slightly different interface was
developed internally at VividCortex.
