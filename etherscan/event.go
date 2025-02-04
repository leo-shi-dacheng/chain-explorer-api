package etherscan

import (
	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/etherscan/base"
)

func (c *Client) GetLogs(fromBlock, toBlock int, address, topic string) (logs []base.Log, err error) {
	param := common.M{
		"fromBlock": fromBlock,
		"toBlock":   toBlock,
		"topic0":    topic,
		"address":   address,
	}
	err = c.call("logs", "getLogs", param, &logs)
	return
}
