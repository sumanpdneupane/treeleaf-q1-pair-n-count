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