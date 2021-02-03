Exercise 1 - Observations
-----------------------------

 ### When executing the program, I get a random number between -1000000 and 1000000
 > When we have two processes (incrementing and decrementing in this case), "lost update" will occur. This happens when the two processes read the same data and try to update the data with different values. When two operations are done on the same operand (i), one of them may update the operand the other one has written to memory. This is a result of concurrency, meaning that the the operations are not actually executed simultaneously.
