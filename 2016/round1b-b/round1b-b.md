# Problem B. Close Match

Problem

You are attending the most important game in sports history. The Oceania Coders are playing the Eurasia Jammers in the Centrifugal Bumble-Puppy world finals. Unfortunately, you were sleep deprived from all the anticipation, so you fell asleep during the game!

The scoreboard is currently displaying both scores, perhaps with one or more leading zeroes (because the scoreboard displays a fixed number of digits). While you were asleep, some of the lights on the scoreboard were damaged by strong ball hits, so one or more of the digits in one or both scores are not being displayed.

You think close games are more exciting, and you would like to imagine that the scores are as close as possible. Can you fill in all of the missing digits in a way that minimizes the absolute difference between the scores? If there is more than one way to attain the minimum absolute difference, choose the way that minimizes the Coders' score. If there is more than one way to attain the minimum absolute difference while also minimizing the Coders' score, choose the way that minimizes the Jammers' score.

Input

The first line of the input gives the number of test cases, T. T cases follow. Each case consists of one line with two non-empty strings C and J of the same length, composed only of decimal digits and question marks, representing the score as you see it for the Coders and the Jammers, respectively. There will be at least one question mark in each test case.

Output

For each test case, output one line containing Case #x: c j, where x is the test case number (starting from 1), c is C with the question marks replaced by digits, and j is J with the question marks replaced by digits, such that the absolute difference between the integers represented by c and j is minimized. If there are multiple solutions with the same absolute difference, use the one in which c is minimized; if there are multiple solutions with the same absolute difference and the same value of c, use the one in which j is minimized.

Limits

1 ≤ T ≤ 200.
C and J have the same length.
Small dataset

1 ≤ the length of C and J ≤ 3.

Large dataset

1 ≤ the length of C and J ≤ 18.

Sample

```
Input 
    
4
1? 2?
?2? ??3
? ?
?5 ?0

Output 

Case #1: 19 20
Case #2: 023 023
Case #3: 0 0
Case #4: 05 00
```

In sample case #4, note that the answer cannot be 15 10; that minimizes the absolute difference, but does not minimize the Coders' score. Nor can the answer be 05 10; that minimizes the absolute difference and the Coders' score, but does not minimize the Jammers' score.