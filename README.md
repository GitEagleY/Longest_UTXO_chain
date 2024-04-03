# Blockchain Example

## How to Run

1.  Make sure you have Go installed on your system. If not, you can download and install it from the [official Go website](https://go.dev/doc/install).
2.  Clone this repository to your local machine.
3.  Navigate to the project directory.
4.  Run the following command to compile and execute the program:

         go run main.go

## Algorithm Description and Efficiency

The greedy algorithm for finding the longest chain in a blockchain involves iteratively examining each block in the blockchain and determining whether it satisfies certain criteria to be considered as part of the longest chain or not. Here's how it typically works:

1. Iterate through each block in the blockchain.
2. For each block, check if it meets certain conditions to be considered valid for inclusion in the longest chain. These conditions might involve properties of the transactions within the block, such as the number of inputs and outputs.
3. If the block meets the conditions, add it to the current chain.
4. Compare the length of the current chain with the length of the longest chain found so far. If the current chain is longer, update the longest chain.
5. Repeat steps 1-4 for all blocks in the blockchain.

The greedy algorithm is called so because it makes decisions locally at each step, choosing the block that appears to be the best option at that moment without considering the overall global picture.

It's worth noting that in the context of blockchain, finding the longest chain is typically associated with finding the chain with the most accumulated proof-of-work. However, in this specific algorithm, the focus is on the structure of transactions within blocks rather than proof-of-work.

The time complexity of the greedy algorithm for finding the longest chain in a blockchain depends on the size of the blockchain, denoted as n, where n is the number of blocks.

In the worst case, the greedy algorithm needs to iterate through each block in the blockchain once to determine whether it satisfies the conditions to be considered part of the longest chain. Therefore, the time complexity of the greedy algorithm can be expressed as O(n), where n is the number of blocks in the blockchain.

## Time

~5 hours spent on this project
