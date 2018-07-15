package main

// A sequence of Blocks
type Blockchain struct {
	blocks []*Block
}

// Save provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// Create a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}


