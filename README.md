godaemon
========

Daemonize Go applications deviously.

You can't daemonize in Go. Daemonizing is a Unix concept that requires
some specific things you can't do in Go. But you can simulate it pretty
accurately, if you don't mind that your program will start copies of itself
several times.

An earlier version of this concept with a slightly different interface was
developed internally at VividCortex. If you use this package, omit the parameter
to `Daemonize()`.
