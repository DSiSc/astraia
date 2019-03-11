package api

import (
	"flag"
	"fmt"
	"github.com/DSiSc/craft/types"
	"github.com/DSiSc/lightClient/config"
	"github.com/DSiSc/wallet/common"
	"github.com/DSiSc/web3go/provider"
	"github.com/DSiSc/web3go/rpc"
	"github.com/DSiSc/web3go/web3"
	web3cmn "github.com/DSiSc/web3go/common"
	local "github.com/DSiSc/wallet/core/types"
	"strconv"
)

//Send a signed transaction
func SendTransaction(tx *types.Transaction) (common.Hash, error) {
	//format 0x string
	from := fmt.Sprintf("0x%x", *(tx.Data.From))
	to := from
	gas := "0x" + strconv.FormatInt(int64(tx.Data.GasLimit),16)
	gasprice := "0x" + tx.Data.Price.String()
	value := "0x" + tx.Data.Amount.String()
	data := ""

	if tx.Data.Payload != nil {
		data = "0x" + string(tx.Data.Payload)
	} else {
		data = ""
	}

	configHostName := config.GetApiGatewayHostName()
	hostname := flag.String("hostname", configHostName, "The ethereum client RPC host")
	configPort := config.GetApiGatewayPort()
	port := flag.String("port", configPort, "The ethereum client RPC port")
	verbose := flag.Bool("verbose", true, "Print verbose messages")

	if *verbose {
		fmt.Printf("Connect to %s:%s\n", *hostname, *port)
	}

	provider := provider.NewHTTPProvider(*hostname+":"+*port, rpc.GetDefaultMethod())
	web3 := web3.NewWeb3(provider)

	req := &web3cmn.TransactionRequest{
		From:     from,
		To:       to,
		Gas:      gas,
		GasPrice: gasprice,
		Value:    value,
		Data:     data,
	}

	hash, err := web3.Eth.SendTransaction(req)
	return common.Hash(hash), err
}

func SendRawTransaction(tx *types.Transaction) (common.Hash, error) {

	configHostName := config.GetApiGatewayHostName()
	hostname := flag.String("hostname", configHostName, "The ethereum client RPC host")
	configPort := config.GetApiGatewayPort()
	port := flag.String("port", configPort, "The ethereum client RPC port")
	verbose := flag.Bool("verbose", false, "Print verbose messages")

	if *verbose {
		fmt.Printf("Connect to %s:%s\n", *hostname, *port)
	}

	provider := provider.NewHTTPProvider(*hostname+":"+*port, rpc.GetDefaultMethod())
	web3 := web3.NewWeb3(provider)

	txBytes, _ := local.EncodeToRLP(tx)
	hash, err := web3.Eth.SendRawTransaction(txBytes)

	return common.Hash(hash), err
}

