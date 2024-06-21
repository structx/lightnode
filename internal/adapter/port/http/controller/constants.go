package controller

const (
	base    = 10
	bitSize = 64

	v1 = "/api/v1"

	blockPath     = v1 + "/blocks"
	blockHashPath = v1 + blockPath + "/{blockHash}"

	transactionPath     = blockHashPath + "/tx"
	transactionHashPath = transactionPath + "/{txHash}"

	health = "/health"
)
