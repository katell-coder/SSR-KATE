package main

import (
	"fmt"
	"minibc/blockchain"
	"minibc/util"
)

func main() {
	//新建一条区块链，里面隐含创建一个创世区块（初始区块）
	bc := blockchain.NewBlockChain()

	//添加3个区块
	bc.AddBlock("Mini block 01")
	bc.AddBlock("Mini Block 02")
	bc.AddBlock("Mini Block 03")

	//从区块链中应该有4个区块 1创世 3新添加
	for _, block := range bc.Blocks {
		fmt.Println("前一区块哈希值：", util.BytesToHex(block.HashPrevBlock))
		fmt.Println("当前区块内容为：", string(block.Data))
		fmt.Println("当前区块哈希值：", util.BytesToHex(block.GetHash()))
		fmt.Println("=============================================")
	}

}
