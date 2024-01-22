// var postMethods = map[string]bool{
// 	// custom RPCs
// 	"relay_sendTransaction":      true,
// 	"relay_getTransactionStatus": true,
// 	"relay_getBalance":           true,
// 	// standard RPCs
// 	"web3_clientVersion":  true,
// 	"web3_sha3":           true,
// 	"net_version":         true,
// 	"net_listening":       true,
// 	"net_peerCount":       true,
// 	"eth_chainId":         true,
// 	"eth_protocolVersion": true,
// 	"eth_syncing":         true,
// 	// "#eth_coinbase": true,
// 	"eth_mining":                           true,
// 	"eth_hashrate":                         true,
// 	"eth_gasPrice":                         true,
// 	"eth_maxPriorityFeePerGas":             true,
// 	"eth_feeHistory":                       true,
// 	"eth_accounts":                         true,
// 	"eth_blockNumber":                      true,
// 	"eth_getBalance":                       true,
// 	"eth_getStorageAt":                     true,
// 	"eth_getTransactionCount":              true,
// 	"eth_getBlockTransactionCountByHash":   true,
// 	"eth_getBlockTransactionCountByNumber": true,
// 	"eth_getUncleCountByBlockHash":         true,
// 	"eth_getUncleCountByBlockNumber":       true,
// 	"eth_getCode":                          true,
// 	"eth_getProof":                         true,
// 	// "#eth_sign": true,
// 	// "#eth_sendTransaction": true,
// 	"eth_sendRawTransaction":                  true,
// 	"eth_call":                                true,
// 	"eth_estimateGas":                         true,
// 	"eth_getBlockByHash":                      true,
// 	"eth_getBlockByNumber":                    true,
// 	"eth_getTransactionByHash":                true,
// 	"eth_getTransactionByBlockHashAndIndex":   true,
// 	"eth_getTransactionByBlockNumberAndIndex": true,
// 	"eth_getTransactionReceipt":               true,
// 	"eth_getBlockReceipts":                    true,
// 	"eth_getUncleByBlockHashAndIndex":         true,
// 	"eth_getUncleByBlockNumberAndIndex":       true,
// 	"eth_getCompilers":                        true,
// 	"eth_compileSolidity":                     true,
// 	"eth_compileLLL":                          true,
// 	"eth_compileSerpent":                      true,
// 	"eth_newFilter":                           true,
// 	"eth_newBlockFilter":                      true,
// 	// "#eth_newPendingTransactionFilter": true,
// 	"eth_uninstallFilter":  true,
// 	"eth_getFilterChanges": true,
// 	"eth_getFilterLogs":    true,
// 	"eth_getLogs":          true,
// 	"eth_getWork":          true,
// 	"eth_submitWork":       true,
// 	"eth_submitHashrate":   true,
// 	"eth_createAccessList": true,
// 	// tx-sentinel
// 	"eth_sendProtectedTransaction":           true,
// 	"eth_cancelProtectedTransaction":         true,
// 	"eth_getBlockSupplyChainByNumber":        true,
// 	"eth_getProtectedTransactionSubmissions": true,
// 	"eth_getProtectedTransactionCount":       true,
// 	"infura_simulateTransactions":            true,
// 	// Debug namespace RPCs
// 	"debug_traceCall":          true,
// 	"debug_traceBlockByHash":   true,
// 	"debug_traceBlockByNumber": true,
// 	"debug_traceTransaction":   true,
// 	// "#db_putString": true,
// 	// "#db_getString": true,
// 	// "#db_putHex": true,
// 	// "#db_getHex": true,
// 	"shh_version":          true,
// 	"shh_post":             true,
// 	"shh_newIdentity":      true,
// 	"shh_hasIdentity":      true,
// 	"shh_newGroup":         true,
// 	"shh_addToGroup":       true,
// 	"shh_newFilter":        true,
// 	"shh_uninstallFilter":  true,
// 	"shh_getFilterChanges": true,
// 	"shh_getMessages":      true,
// 	"parity_nextNonce":     true,
// 	// Trace endpoints
// 	"trace_block":       true,
// 	"trace_transaction": true,
// 	"trace_filter":      true,
// 	"trace_call":        true,
// 	"trace_callMany":    true,
// }

package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strconv",
    "os",
    "log"
)

type j struct {
    Cl []string `json:"cl"`
    Gr []string `json:"gr"`
    Cr []string `json:"cr"`
}

func main() {
    // create an instance of j as a slice
    var data []j
    // using a for loop for create dummy data fast
    for i := 0; i < 5; i++ {
        v := strconv.Itoa(i)
        data = append(data, j{
            Cl: []string{"foo " + v},
            Gr: []string{"bar " + v},
            Cr: []string{"foobar " + v},
        })
    }

    // printing out json neatly to demonstrate
    b, _ := json.MarshalIndent(data, "", " ")
    fmt.Println(string(b))

    // writing json to file

    _ = ioutil.WriteFile("file.json", b, 0644)

    // to append to a file
    // create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
    f, err := os.OpenFile("/var/log/debug-web.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
    if err != nil {
        log.Fatal(err)
    }
    // write to file, f.Write()
    f.Write(b)

    // if you are doing alot of I/O work you may not want to write out to file often instead load up a bytes.Buffer and write to file when you are done... assuming you don't run out of memory when loading to bytes.Buffer


}