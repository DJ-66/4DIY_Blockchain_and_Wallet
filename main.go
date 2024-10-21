
//https://www.youtube.com/watch?v=mYlHT9bB6OE&list=PLJbE2Yu2zumC5QE39TQHBLYJDB2gfFE5Q

package main
import (
	"bytes"
	"crypto/sha256"
	"fmt"
)



func main (){
	chain := InitBlockChain()

	chain.AddBlock("First block after genesis block")
	chain.AddBlock("Second block after genesis block")
	chain.AddBlock("Third block after genesis block")

	for _, block := range chain.blocks {
		//fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}