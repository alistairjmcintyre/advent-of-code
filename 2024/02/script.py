import sys
from itertools import pairwise, combinations

INPUT_FILE = f"inputs/{sys.argv[-1]}.txt"

inputs = []

with open(INPUT_FILE, "r") as input_file:
    for line in input_file:
        inputs.append([int(val) for val in line.split()])


def is_safe(levels):
    if sorted(levels) != levels and sorted(levels)[::-1] != levels:
        return False

    for x, y in pairwise(levels):
        value = abs(x - y)
        if not (0 < value < 4):
            return False

    return True


safe = list(filter(is_safe, inputs))
print(f"Total count of safe reports: {len(safe)}")


def is_still_safe(levels):
    if is_safe(levels):
        return True

    combs = combinations(levels, len(levels) - 1)
    return any([is_safe(list(level)) for level in combs])


safe = list(filter(is_still_safe, inputs))
print(f"Total count of safe reports after removing an entry: {len(safe)}")
