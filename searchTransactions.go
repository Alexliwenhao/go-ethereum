package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/aA9qPRm8NbL7i7xunQlnKByn7G3Nf28m")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainID:", chainID)

	blockNumber := big.NewInt(5671744)

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		continue
		fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())
		fmt.Printf("交易值: %s\n", tx.Value().String())
		fmt.Printf("Gas价格: %s\n", tx.GasPrice().String())
		fmt.Printf("Gas限制: %d\n", tx.Gas())
		fmt.Printf("Nonce: %d\n", tx.Nonce())
		fmt.Println("Data:", tx.Data())

		if tx.To() != nil {
			fmt.Printf("接收地址: %s\n", tx.To().Hex())
		}
		fmt.Printf("ChainID: %d\n", tx.ChainId())
		fmt.Println("------------------------")

		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Printf("发送地址: %s\n", sender.Hex())
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("状态码:", receipt.Status)
		fmt.Println("日志:", receipt.Logs)
		fmt.Println("------------------------")
	}

	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("第几个事务： %d, 交易哈希: %s\n", idx, tx.Hash().Hex())
	}
}
