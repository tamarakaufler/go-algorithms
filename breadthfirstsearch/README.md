# Breadth first search

Search algorithm, that works on graphs/trees with each node having its
own set of child nodes. Child nodes can be duplicates. The aim is to find
the shortest path to the node that satisfies given criteria.

## Input

Data structure representing the problem is a map of lists:

                     node1
                       |
           -------------------------------------------
           |                   |                     |
         node2               node3                 node4                                  LEVEL1
           |                   |                     |
       ----------------        ---------            ---------------------------
       |       |      |        |       |            |        |        |       |
     node21  node22 node23   node31  node32       node41   node42   node43  node44        LEVEL2
       |       |
    ------   -----------
    |    |   |    |    |
                                                                                          LEVEL3
 etc    


## Solution

The implementation searches each level first and descents into the next one only
if the solution is not found on the current leve. The problem uses a queue to arrive
at a solution.

User supplies his own criteria function that chooses the first acceptable node. The provided implementation works with string nodes.
