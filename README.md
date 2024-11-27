## Question
```
Q1. Given an array, find the total number of inversions of it. If (i < j) and (A[i] >
		A[j]), then pair (i, j) is called an inversion of an array A. Write a program to identify
		the number of inversion in the array and count all such pairs in the array.
		Sample input: Input: A[] = [1, 9, 6, 4, 5].
```

## Answer
### Explanation of the Problem and Code
```
Problem Overview
An inversion in an array is a condition where:
    1. 𝑖<𝑗, and
    2. A[i]>A[j]
For a given array 𝐴 [ ] A[], we need to:
    1. Count the total number of inversions in the array.
    2. Identify all such pairs ( 𝑖 , 𝑗 ) (i,j) where the inversion occurs.

Sample Input
For the input array: A[]=[1,9,6,4,5]
    * Total inversions: 5
    * Inversion pairs: [(1,2),(1,3),(1,4),(2,3),(2,4)]
```

### Code Walkthrough

#### Input
The input array is passed as array in the main function:
``` go
array := []int{1, 9, 6, 4, 5}
```

#### Core Function
The function CountAndGetPairInversionFromArray is called, which does the following:
1. Initialize Counters:
    * count: Keeps track of the total number of inversions.
    * pairs: A slice of pairs to store the inversion indices (i,j).
2. Iterate Through the Array:
    * Use a nested loop to compare every pair of indices (i,j) such that i<j.
    * Check if A[i]>A[j]. If true:
        * Increment count.
        * Append the pair (i,j) to pairs.
3. Return the Results:
    * The function returns the total count of inversions and the list of inversion pairs.

#### Main Function
The main function calls CountAndGetPairInversionFromArray:
``` go
count, pairs := CountAndGetPairInversionFromArray(array)
```
It receives:
    * count: Total number of inversions.
    * pairs: Slice containing all inversion pairs.
These results are printed to the console.


### Output
``` cmd
Total Inversions: 5
Inversion Pairs: [[1 2] [1 3] [1 4] [2 3] [2 4]]
```

## Example Walkthrough
```
Array:[1, 9, 6, 4, 5]
1. Compare A[0] =  withA[1], A[2], A[3], A[4]: No inversions.
2. Compare A[1] = 9:
    * A[1] > A[2] Pair (1,2)(1, 2)(1,2).
    * A[1] > A[3] Pair (1,3)(1, 3)(1,3).
    * A[1] > A[4] Pair (1,4)(1, 4)(1,4).
3. Compare A[2] = 6:
    * A[2] > A[3]: Pair (2,3)(2, 3)(2,3).
    * A[2] > A[4]: Pair (2,4)(2, 4)(2,4).
4. Compare A[3] = 4: No more inversions.
```

## Time Complexity
```
The code uses a nested loop:
    * Outer loop runs 𝑛 times.
    * Inner loop runs 𝑛 − 𝑖 − 1 times.
Thus, the time complexity is: O(n^2)
```