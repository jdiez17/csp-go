csp-go
======

Constraint satisfying problem solver in Go. Uses backtracking and concurrency.

FAQ
===

Is it any good?
----------------

It does solve all constraint satisfying problems with finite variables and a finite domain. It does so by checking all the possible solutions and using concurrency to test the constraints concurrently.
But it does not do constraint propagation, so it's not as efficient as it should be.

Why should I use it?
---------------------

You probably shouldn't it, I only wrote it for my own amusement.
