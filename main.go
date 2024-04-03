package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type UTXO struct {
	TransactionID string
	Index         int
	Amount        float64
}

type Transaction struct {
	TransactionID string
	Inputs        []UTXO
	Outputs       []UTXO
}

type Block struct {
	BlockID      int
	Transactions []Transaction
}

type Blockchain struct {
	Blocks  []Block
	UTXOSet map[string]UTXO
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks:  []Block{},
		UTXOSet: make(map[string]UTXO),
	}
}

func (bc *Blockchain) UpdateUTXOSet(block Block) {
	for _, tx := range block.Transactions {
		for _, inputUTXO := range tx.Inputs { //mark the input UTXO as spent.
			delete(bc.UTXOSet, inputUTXO.TransactionID)
		}
		for _, output := range tx.Outputs { //mark UTXO as unspent.
			//generate a unique key based on the transaction ID and index
			key := fmt.Sprintf("%s-%d", output.TransactionID, output.Index)
			//add the output UTXO to the UTXO set with the generated key
			bc.UTXOSet[key] = output
		}
	}
}

func (bc *Blockchain) FindLongestChain() []Block {
	var longestChain []Block
	var currentChain []Block

	for _, block := range bc.Blocks {
		// Start a new chain for each block
		currentChain = []Block{block}

		//ensure that any changes made to the UTXO set during the processing of transactions in each block are reflected accurately
		//when determining the length of the chain.
		bc.UpdateUTXOSet(block)

		if len(currentChain) > len(longestChain) {
			longestChain = append([]Block(nil), currentChain...)
		}
	}
	return longestChain
}

func (bc *Blockchain) FilterTransactions() []Transaction {
	var filteredTransactions []Transaction

	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if len(tx.Inputs) == 1 && len(tx.Outputs) == 2 {
				filteredTransactions = append(filteredTransactions, tx)
			}
		}
	}
	return filteredTransactions
}

func getRandomNum() int {
	return rand.Intn(1000)
}

func main() {
	blockchain := NewBlockchain()

	// Generate unique block and transaction IDs for the first chain
	blockID1 := getRandomNum()
	txID1 := "tx" + strconv.Itoa(getRandomNum())
	txID2 := "tx" + strconv.Itoa(getRandomNum())
	amount1 := float64(getRandomNum())
	amount2 := float64(getRandomNum())
	amount3 := float64(getRandomNum())
	amount4 := float64(getRandomNum())

	// Generate unique block and transaction IDs for the second chain
	blockID2 := getRandomNum()
	txID4 := "tx" + strconv.Itoa(getRandomNum())
	txID5 := "tx" + strconv.Itoa(getRandomNum())
	txID6 := "tx" + strconv.Itoa(getRandomNum())
	amount5 := float64(getRandomNum())
	amount6 := float64(getRandomNum())
	amount7 := float64(getRandomNum())

	// Create the first block and transactions for the first chain
	block1 := Block{
		BlockID: blockID1,
		Transactions: []Transaction{
			{TransactionID: txID1, Inputs: []UTXO{{TransactionID: "input1", Index: 0, Amount: amount1}}, Outputs: []UTXO{{TransactionID: txID1, Index: 0, Amount: amount2}, {TransactionID: txID1, Index: 1, Amount: amount3}}},
			{TransactionID: txID5, Inputs: []UTXO{{TransactionID: "input4", Index: 1, Amount: amount7}, {TransactionID: "input1", Index: 4, Amount: amount3}}, Outputs: []UTXO{{TransactionID: txID5, Index: 2, Amount: amount4}, {TransactionID: txID5, Index: 4, Amount: amount1}, {TransactionID: txID6, Index: 2, Amount: amount2}}},
		},
	}

	// Create the second block and transactions for the second chain
	block2 := Block{
		BlockID: blockID2,
		Transactions: []Transaction{
			{TransactionID: txID4, Inputs: []UTXO{{TransactionID: "input3", Index: 2, Amount: amount5}}, Outputs: []UTXO{{TransactionID: txID4, Index: 5, Amount: amount6}}},
			{TransactionID: txID2, Inputs: []UTXO{{TransactionID: "input2", Index: 0, Amount: amount4}}, Outputs: []UTXO{{TransactionID: txID2, Index: 0, Amount: amount1}, {TransactionID: txID2, Index: 1, Amount: amount2}}},
		}}

	// Add the blocks to the blockchain
	blockchain.Blocks = append(blockchain.Blocks, block1)
	blockchain.Blocks = append(blockchain.Blocks, block2)

	// Find the longest chain in the blockchain
	longestChain := blockchain.FindLongestChain()

	// Filter transactions based on the specified criteria
	filteredTransactions := blockchain.FilterTransactions()

	fmt.Println("Longest Chain:")
	for _, block := range longestChain {
		fmt.Printf("BlockID: %d\n", block.BlockID)
		for _, tx := range block.Transactions {
			fmt.Printf("\tTransactionID: %s\n", tx.TransactionID)
			fmt.Println("\tInputs:")
			for _, input := range tx.Inputs {
				fmt.Printf("\t\tTransactionID: %s, Index: %d\n", input.TransactionID, input.Index)
			}
			fmt.Println("\tOutputs:")
			for _, output := range tx.Outputs {
				fmt.Printf("\t\tTransactionID: %s, Index: %d, Amount: %.2f\n", output.TransactionID, output.Index, output.Amount)
			}
		}
	}

	fmt.Println("Filtered Transactions:")
	for _, tx := range filteredTransactions {
		fmt.Printf("TransactionID: %s\n", tx.TransactionID)
		fmt.Println("\tInputs:")
		for _, input := range tx.Inputs {
			fmt.Printf("\t\tTransactionID: %s, Index: %d\n", input.TransactionID, input.Index)
		}
		fmt.Println("\tOutputs:")
		for _, output := range tx.Outputs {
			fmt.Printf("\t\tTransactionID: %s, Index: %d, Amount: %.2f\n", output.TransactionID, output.Index, output.Amount)
		}
	}

}
