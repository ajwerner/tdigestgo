# tdigest
[![Documentation](https://godoc.org/github.com/ajwerner/tdigestgo?status.svg)](http://godoc.org/github.com/ajwerner/tdigestgo)

This is a concurrent go implementation of the t-digest data structure
for streaming quantile estimation. The code implements the zero-allocation,
merging algorithm from Ted Dunning's paper ([here][histo.pdf]).

It utilizes the later iteration of the scale functions, currently just exposing
scale function `k_2` from the paper.

The implementation strives to make concurrent writes cheap. In the common case
a write needs only increment an atomic and write two floats to a buffer.
Occasionlly, when the buffer fills, a caller will have to perform the merge
operation.

[histo.pdf]: https://github.com/tdunning/t-digest/raw/d7427ee41be6a6fd271206f26a0cad42f74f30bf/docs/t-digest-paper/histo.pdf
