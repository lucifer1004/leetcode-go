# 1079: Letter Tile Possibilities

## Solution

First, a tile is converted to a list `f` of letter frequency, e.g., "AAB" becomes [2,1], "AAABBC" becomes [3, 2, 1].

Then, we define "+" on letter frequency. This is similar to normal addition, while the upper limit for the i-th digit is `f[i]`.

We will start from [0, 0, ..., 0], and stop at `f`. For example, in the case of `tile` = "AAB" and `f` = [2, 1], we will have

[0, 0] => [0, 1] => [1, 0] => [1, 1] => [2, 0] => [2, 1]

Each node in this path represents a selection of letters, e.g., [1, 1] means we choose 1 "A" and 1 "B".

Then we count the number of permutations for each node, and the sum of all nodes except [0, 0] is the answer to this question.

For each node, the number of permutations can be easily calculated via:

n = sum(a_i)!/((a_1)!(a_2)!...(a_n)!)

For faster calculation, all needed factorial numbers are precalculated.