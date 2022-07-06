package models

type Wallet struct {
	Wallet string `json:"wallet"`
}

type WalletArray struct {
	Wallets []string `json:"wallets"`
}

type EtherScanResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  interface{}
}

type TradeRecord struct {
	blockNumber,
	timeStamp,
	Hash,
	nonce,
	blockHash,
	transactionIndex,
	from,
	to,
	value,
	gas,
	gasPrice,
	isError,
	txreceipt_status,
	input,
	contractAddress,
	cumulativeGasUsed,
	gasUsed, confirmations string
}
