package core

import "encoding/hex"

// UTXOSet represents UTXO set
type UTXOSet struct {
	Blockchain *Blockchain
}

func (chain *Blockchain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	//unspentTxs := chain.FindUnspentTransactions(address)
	accumulated := 0

	return accumulated, unspentOutputs
}

// FindUnspentTransactions finds all unspent transaction outputs for a given address.
func (chain *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction        // List to store unspent transactions
	spentTXOs := make(map[string][]int) // Map to track spent transaction outputs
	iterator := chain.Iterator()        // Get an iterator for the blockchain

	// Iterate through each block in the blockchain
	// TODO complex Iterate will cost much time, need to refactor
	for {
		block := iterator.Next() // Get the next block

		// Iterate through each transaction in the block
		for _, tx := range block.Transactions {
			txId := hex.EncodeToString(tx.ID) // Convert transaction ID to string

		Outputs:
			// Iterate through each output in the transaction
			for outIdx, out := range tx.VOut {
				// Check if the output was spent
				if spentTXOs[txId] != nil {
					for _, spentOut := range spentTXOs[txId] {
						if spentOut == outIdx {
							continue Outputs // Skip if the output was spent
						}
					}
				}

				// Check if the output can be unlocked with the given address
				if out.CanBeUnlockedWith(address) {
					unspentTXs = append(unspentTXs, *tx) // Add the transaction to the list of unspent transactions
				}
			}

			// If the transaction is not a coinbase transaction, track spent inputs
			if !tx.IsCoinbase() {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}

		// Break the loop if the genesis block is reached
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return unspentTXs // Return the list of unspent transactions
}

// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

func (chain *Blockchain) FindUTXO(address string) []TxOutput {
	var UTXOs []TxOutput
	unspentTxs := chain.FindUnspentTransactions(address)

	for _, tx := range unspentTxs {
		for _, out := range tx.VOut {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}
	return UTXOs
}
