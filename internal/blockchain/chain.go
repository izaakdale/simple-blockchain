package blockchain

import "slices"

func New() *chain {
	return &chain{
		[]*block{genesis},
	}
}

type chain struct {
	blocks []*block
}

func (c *chain) AddBlock(data []byte) {
	pBlock := c.blocks[len(c.blocks)-1]
	c.blocks = append(c.blocks, newBlock(data, pBlock.myBlockHash))
}

func (c *chain) Verify(idx int, data []byte) bool {
	// transaction index cannot be before 1 (first tx), or after all the transactions
	if idx > 0 && idx < len(c.blocks) {
		tx := c.blocks[idx]
		return slices.Equal(tx.allData, data)
	}

	return false
}
