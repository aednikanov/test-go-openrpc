package http_proxy

import "github.com/INFURA/inf-reverse-proxy/pkg/systemization"

var postMethods = map[string]bool{
	// custom RPCs
	"relay_sendTransaction":      true,
	"relay_getTransactionStatus": true,
	"relay_getBalance":           true,
	// standard RPCs
	"web3_clientVersion":  true,
	"web3_sha3":           true,
	"net_version":         true,
	"net_listening":       true,
	"net_peerCount":       true,
	"eth_chainId":         true,
	"eth_protocolVersion": true,
	"eth_syncing":         true,
	// "#eth_coinbase": true,
	"eth_mining":                           true,
	"eth_hashrate":                         true,
	"eth_gasPrice":                         true,
	"eth_maxPriorityFeePerGas":             true,
	"eth_feeHistory":                       true,
	"eth_accounts":                         true,
	"eth_blockNumber":                      true,
	"eth_getBalance":                       true,
	"eth_getStorageAt":                     true,
	"eth_getTransactionCount":              true,
	"eth_getBlockTransactionCountByHash":   true,
	"eth_getBlockTransactionCountByNumber": true,
	"eth_getUncleCountByBlockHash":         true,
	"eth_getUncleCountByBlockNumber":       true,
	"eth_getCode":                          true,
	"eth_getProof":                         true,
	// "#eth_sign": true,
	// "#eth_sendTransaction": true,
	"eth_sendRawTransaction":                  true,
	"eth_call":                                true,
	"eth_estimateGas":                         true,
	"eth_getBlockByHash":                      true,
	"eth_getBlockByNumber":                    true,
	"eth_getTransactionByHash":                true,
	"eth_getTransactionByBlockHashAndIndex":   true,
	"eth_getTransactionByBlockNumberAndIndex": true,
	"eth_getTransactionReceipt":               true,
	"eth_getBlockReceipts":                    true,
	"eth_getUncleByBlockHashAndIndex":         true,
	"eth_getUncleByBlockNumberAndIndex":       true,
	"eth_getCompilers":                        true,
	"eth_compileSolidity":                     true,
	"eth_compileLLL":                          true,
	"eth_compileSerpent":                      true,
	"eth_newFilter":                           true,
	"eth_newBlockFilter":                      true,
	// "#eth_newPendingTransactionFilter": true,
	"eth_uninstallFilter":  true,
	"eth_getFilterChanges": true,
	"eth_getFilterLogs":    true,
	"eth_getLogs":          true,
	"eth_getWork":          true,
	"eth_submitWork":       true,
	"eth_submitHashrate":   true,
	"eth_createAccessList": true,
	// tx-sentinel
	"eth_sendProtectedTransaction":           true,
	"eth_cancelProtectedTransaction":         true,
	"eth_getBlockSupplyChainByNumber":        true,
	"eth_getProtectedTransactionSubmissions": true,
	"eth_getProtectedTransactionCount":       true,
	"infura_simulateTransactions":            true,
	// Debug namespace RPCs
	"debug_traceCall":          true,
	"debug_traceBlockByHash":   true,
	"debug_traceBlockByNumber": true,
	"debug_traceTransaction":   true,
	// "#db_putString": true,
	// "#db_getString": true,
	// "#db_putHex": true,
	// "#db_getHex": true,
	"shh_version":          true,
	"shh_post":             true,
	"shh_newIdentity":      true,
	"shh_hasIdentity":      true,
	"shh_newGroup":         true,
	"shh_addToGroup":       true,
	"shh_newFilter":        true,
	"shh_uninstallFilter":  true,
	"shh_getFilterChanges": true,
	"shh_getMessages":      true,
	"parity_nextNonce":     true,
	// Trace endpoints
	"trace_block":       true,
	"trace_transaction": true,
	"trace_filter":      true,
	"trace_call":        true,
	"trace_callMany":    true,
}

// If and only if an RPC method is listed in BOTH `postMethods` AND in
// `selectiveAvailability` (below), an additional check will be performed
// to determine whether traffic for that method should be directed to the
// smart router, conditionally on the current network.
var selectiveAvailability = map[string]map[string]bool{
	"trace_block": {
		systemization.NETWORK_MAINNET:         true,
		systemization.NETWORK_GOERLI:          true,
		systemization.NETWORK_SEPOLIA:         true,
		systemization.NETWORK_LINEA_GOERLI:    true,
		systemization.NETWORK_LINEA_MAINNET:   true,
		systemization.NETWORK_POLYGON_MAINNET: true,
		// defaults to false for all other networks
	},
	"trace_transaction": {
		systemization.NETWORK_MAINNET:       true,
		systemization.NETWORK_GOERLI:        true,
		systemization.NETWORK_SEPOLIA:       true,
		systemization.NETWORK_LINEA_GOERLI:  true,
		systemization.NETWORK_LINEA_MAINNET: true,
		// defaults to false for all other networks
	},
	"trace_filter": {
		systemization.NETWORK_MAINNET: true,
		systemization.NETWORK_GOERLI:  true,
		systemization.NETWORK_SEPOLIA: true,
		// defaults to false for all other networks
	},
	"trace_call": {
		systemization.NETWORK_MAINNET: true,
		systemization.NETWORK_GOERLI:  true,
		systemization.NETWORK_SEPOLIA: true,
		// defaults to false for all other networks
	},
	"trace_callMany": {
		systemization.NETWORK_MAINNET: true,
		systemization.NETWORK_GOERLI:  true,
		systemization.NETWORK_SEPOLIA: true,
		// defaults to false for all other networks
	},
	"eth_createAccessList": {
		systemization.NETWORK_MAINNET:          true,
		systemization.NETWORK_GOERLI:           true,
		systemization.NETWORK_SEPOLIA:          true,
		systemization.NETWORK_ARBITRUM_MAINNET: true,
		systemization.NETWORK_ARBITRUM_GOERLI:  true,
		systemization.NETWORK_ARBITRUM_SEPOLIA: true,
		systemization.NETWORK_LINEA_MAINNET:    true,
		systemization.NETWORK_LINEA_GOERLI:     true,
		systemization.NETWORK_OPTIMISM_MAINNET: true,
		systemization.NETWORK_OPTIMISM_GOERLI:  true,
		systemization.NETWORK_OPTIMISM_SEPOLIA: true,
		systemization.NETWORK_POLYGON_MAINNET:  true,
		systemization.NETWORK_POLYGON_MUMBAI:   true,
		// defaults to false for all other networks
	},
	"eth_getBlockReceipts": {
		systemization.NETWORK_MAINNET:          true,
		systemization.NETWORK_SEPOLIA:          true,
		systemization.NETWORK_ARBITRUM_SEPOLIA: true,
		systemization.NETWORK_ARBITRUM_GOERLI:  true,
		systemization.NETWORK_ARBITRUM_MAINNET: true,
		systemization.NETWORK_LINEA_GOERLI:     true,
		systemization.NETWORK_LINEA_MAINNET:    true,
		systemization.NETWORK_POLYGON_MAINNET:  true,
		systemization.NETWORK_POLYGON_MUMBAI:   true,
		// defaults to false for all other networks
	},
	"debug_traceCall": {
		systemization.NETWORK_MAINNET:          true,
		systemization.NETWORK_GOERLI:           true,
		systemization.NETWORK_SEPOLIA:          true,
		systemization.NETWORK_LINEA_GOERLI:     true,
		systemization.NETWORK_LINEA_MAINNET:    true,
		systemization.NETWORK_ARBITRUM_MAINNET: true,
		systemization.NETWORK_POLYGON_MAINNET:  true,
		systemization.NETWORK_OPTIMISM_MAINNET: true,
		// defaults to false for all other networks
	},
	"debug_traceBlockByHash": {
		systemization.NETWORK_LINEA_GOERLI:  true,
		systemization.NETWORK_LINEA_MAINNET: true,
		// defaults to false for all other networks
	},
	"debug_traceBlockByNumber": {
		systemization.NETWORK_LINEA_GOERLI:  true,
		systemization.NETWORK_LINEA_MAINNET: true,
		// defaults to false for all other networks
	},
	"debug_traceTransaction": {
		systemization.NETWORK_LINEA_GOERLI:  true,
		systemization.NETWORK_LINEA_MAINNET: true,
		// defaults to false for all other networks
	},
	"eth_sendProtectedTransaction": {
		systemization.NETWORK_MAINNET:       true,
		systemization.NETWORK_GOERLI:        true,
		systemization.NETWORK_SEPOLIA:       true,
		systemization.NETWORK_LINEA_MAINNET: true,
		systemization.NETWORK_LINEA_GOERLI:  true,
		// defaults to false for all other networks
	},
	"eth_cancelProtectedTransaction": {
		systemization.NETWORK_MAINNET:       true,
		systemization.NETWORK_GOERLI:        true,
		systemization.NETWORK_SEPOLIA:       true,
		systemization.NETWORK_LINEA_MAINNET: true,
		systemization.NETWORK_LINEA_GOERLI:  true,
		// defaults to false for all other networks
	},
	"eth_getBlockSupplyChainByNumber": {
		systemization.NETWORK_MAINNET:       true,
		systemization.NETWORK_GOERLI:        true,
		systemization.NETWORK_SEPOLIA:       true,
		systemization.NETWORK_LINEA_MAINNET: true,
		systemization.NETWORK_LINEA_GOERLI:  true,
		// defaults to false for all other networks
	},
	"eth_getProtectedTransactionSubmissions": {
		systemization.NETWORK_MAINNET:       true,
		systemization.NETWORK_GOERLI:        true,
		systemization.NETWORK_SEPOLIA:       true,
		systemization.NETWORK_LINEA_MAINNET: true,
		systemization.NETWORK_LINEA_GOERLI:  true,
		// defaults to false for all other networks
	},
	"eth_getProtectedTransactionCount": {
		systemization.NETWORK_MAINNET:       true,
		systemization.NETWORK_GOERLI:        true,
		systemization.NETWORK_SEPOLIA:       true,
		systemization.NETWORK_LINEA_MAINNET: true,
		systemization.NETWORK_LINEA_GOERLI:  true,
		// defaults to false for all other networks
	},
	"infura_simulateTransactions": {
		systemization.NETWORK_MAINNET:       true,
		systemization.NETWORK_GOERLI:        true,
		systemization.NETWORK_SEPOLIA:       true,
		systemization.NETWORK_LINEA_MAINNET: true,
		systemization.NETWORK_LINEA_GOERLI:  true,
		// defaults to false for all other networks
	},
}
