package controller

const (
	base    = 10
	bitSize = 64

	blockPath     = "/blocks"
	blockHashPath = "/blocks/{blockHash}"

	transactionPath     = blockHashPath + "/tx"
	transactionHashPath = transactionPath + "/{txHash}"
)
