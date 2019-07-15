# 1125: Smallest Sufficient Team

Set cover is NPC, but since there are at most 16 unique skills, this problem is solvable.

## Solution

- Skills are represented bitwise. (A | B means A cooperates with B)
- People with unique skills are picked in advance.
- People whose skills are overlapped by others are excluded in advance.
- Brute force search among the rest people.
  - Keep record of the current minimum, and skip a condition if it uses more people than the current minimum.

The answer is the people who are picked in advance, and the minimal combination of the rest people.
