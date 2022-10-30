# Experiments in async BFT

## Communication Model
[Good explanation](https://www.youtube.com/watch?v=m6nSN4_E0Dc)  

Sync - Known upper bound on latency (i.e. all messages delivered within 1 min)
Partial Async - There is an upper bound on latency, it's not known
Async - There is no upper bound on latency (eventually delivered)

With partially async model we use rounds and timeouts to achieve 
- Validity: every honest recipient terminated and outputs m
- Consistency: every honest recipient terminates with output m
  (or none of them do)

With Async communication model a replica can't know if it didn't receive a message because it wasn't sent or because it wasn't received.
Adversary controls the delay time of messages, can corrupt up to t parties