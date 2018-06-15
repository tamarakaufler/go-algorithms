# Quicksort

Quicksort is a divide-and-conquer algorithm, which solve a recursive problem, that requires defining
a base base, which defines the end of the processing and the recursive case, swhich
calls the same function on a subset of the original problem/data, until the base case is reached.
Base case avoids creating an infinite processing loop. It can also handle irregular cases (like duplicatation etc).

Data structure:  array/list

Quicksort uses a pivot, an element from the data set, to create two partitions/subsets, that will be then recursively
processed until reaching the base case. As the dats distribution is unknown, randomizing the pilot selection optimizes
the solution time.

Steps:
  - decide on the base case(s)
  - choose a pivot
  - solution: apply the quick sort on all partitions (ie quicksort(left partition) + quicksort(right partition))