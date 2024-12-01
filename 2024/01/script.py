import sys
from collections import Counter

INPUT_FILE = f"inputs/{sys.argv[-1]}.txt"
left, right = [], []

with open(INPUT_FILE, "r") as input_file:
    for line in input_file:
        a, b = [int(val) for val in line.split()]
        left.append(a)
        right.append(b)

# Part 1
paired = zip(sorted(left), sorted(right))
total_dist = sum(abs(l - r) for l, r in paired)
print(f"The sum of distances is: {total_dist}.")

# Part 2
counter = Counter(right)
total_similarity = sum([value * counter[value] for value in left])
print(f"The total similarity is: {total_similarity}.")
