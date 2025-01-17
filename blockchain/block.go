package blockchain

imports (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks [] *Block
}

type Block struct {
	Hash		[]byte
	Data		[]byte
	PrevHash	[]byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][] byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, PrevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), PrevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	preBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, preBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[] *Block{Genesis()}}

}