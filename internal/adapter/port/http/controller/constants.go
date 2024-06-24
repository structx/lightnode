package controller

const (
	base    = 10
	bitSize = 64

	v1 = "/1.0"

	blockPath     = v1 + "/blocks"
	blockHashPath = blockPath + "/{blockHash}"

	transactionPath     = blockHashPath + "/tx"
	transactionHashPath = transactionPath + "/{txHash}"

	health  = "/health"
	metrics = "/metrics"
)
