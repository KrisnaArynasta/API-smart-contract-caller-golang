package SmartContractCaller

import (
	"fmt"
    "context"
    // "log"
    "math/big"
    "strings"
    "strconv"
    // "time"

    // "github.com/ava-labs/avalanchego/utils/constants"
    "github.com/ava-labs/avalanchego/utils/formatting"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"//"github.com/ava-labs/coreth/accounts/abi/bind"
    // "github.com/ava-labs/coreth/core/types"
    // "github.com/ava-labs/coreth/params"
    // "github.com/ava-labs/coreth/rpc"
    "github.com/decred/dcrd/dcrec/secp256k1/v3"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"// "github.com/ava-labs/coreth/ethclient"
    KepengToken "kepengmodule"
)

func configSC(smartContractAddress string, rpcURL string, callerPrivateKeyString string) (*bind.TransactOpts, *KepengToken.KepengToken, error) {
    // setup client
    client, err := ethclient.Dial(rpcURL)
    // /rc, err := rpc.Dial("http://localhost:9650/ext/bc/C/rpc")
    if err != nil {
        fmt.Println(err)
        return nil, nil, err
    }
    //ec := ethclient.NewClient(rc)

    ctx := context.Background()

    gasPrice, err := client.SuggestGasPrice(ctx)
    if err != nil {
        fmt.Println(err)
        return nil, nil, err
    }

    // // fetch networkid
    // networkId, err := client.ChainID(ctx)
    // if err != nil {
    //     fmt.Println(err)
    // }

    // parse key
    callerPrivateKeyBytes, err := formatting.Decode(formatting.CB58, strings.TrimPrefix(callerPrivateKeyString, "PrivateKey-"/*constants.SecretKeyPgorefix*/))
    if err != nil {
        fmt.Println(err)
        return nil, nil, err
    }
    callerPrivateKey := secp256k1.PrivKeyFromBytes(callerPrivateKeyBytes)
    callerPrivateKeyECDSA := callerPrivateKey.ToECDSA()

    // derive 'c' address
    cAddress := crypto.PubkeyToAddress(callerPrivateKeyECDSA.PublicKey)
    nonce, err := client.PendingNonceAt(ctx, cAddress)

    scAddress := common.HexToAddress(smartContractAddress)
    contract, err := KepengToken.NewKepengToken(scAddress, client)
    if err != nil {
        fmt.Println(err)
        return nil, nil, err
    }

    chainID, err := client.NetworkID(ctx)
    if err != nil {
        fmt.Println(err)
        return nil, nil, err
    }

    tx, err := bind.NewKeyedTransactorWithChainID(callerPrivateKeyECDSA, chainID);
    if err != nil {
        fmt.Println(err)
        return nil, nil, err
    }
    tx.Nonce = big.NewInt(int64(nonce))
    tx.Value = big.NewInt(0)     // in wei
    tx.GasLimit = uint64(300000) // in units
    tx.GasPrice = gasPrice

    return tx, contract, nil
}

func GetGassPrice(rpcURL string) *big.Int{
    
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        fmt.Println(err)
    }

    ctx := context.Background()

    gasPrice, err := client.SuggestGasPrice(ctx)
    if err != nil {
        fmt.Println(err)
    }

    return gasPrice
}

func GetCallerPubAddress(callerPrivateKeyString string) string{
    
    callerPrivateKeyBytes, err := formatting.Decode(formatting.CB58, strings.TrimPrefix(callerPrivateKeyString, "PrivateKey-"/*constants.SecretKeyPgorefix*/))
    if err != nil {
        return err.Error()
    }
    callerPrivateKey := secp256k1.PrivKeyFromBytes(callerPrivateKeyBytes)
    callerPrivateKeyECDSA := callerPrivateKey.ToECDSA()

    // derive 'c' address
    cAddress := crypto.PubkeyToAddress(callerPrivateKeyECDSA.PublicKey)
    if err != nil {
        //fmt.Println(err)
        return err.Error()
    }

    return cAddress.String()
}

func GetTokenInfo(smartContractAddress string, rpcURL string, callerPrivateKeyString string) (string, string, string, *big.Int, string) {

    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return "", "", "", nil, err.Error()
    }

    //get name token
    traName, err := contract.Name(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err.Error())
        return "", "", "", nil, err.Error()
    }

    //get decimal token
    traDecimals, err := contract.Decimals(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err.Error())
        return "", "", "", nil, err.Error()
    }

    //get symbol token
    traSym, err := contract.Symbol(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err.Error())
        return "", "", "", nil, err.Error()
    }
    
    //get total supply token
    traSupply, err := contract.TotalSupply(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err.Error())
        return "", "", "", nil, err.Error()
    }
    
    return traName,strconv.Itoa(int(traDecimals)),traSym,traSupply, ""





    // // setup signer and transaction options.
    // signer := types.LatestSignerForChainID(networkId)
    // to := &bind.TransactOpts{
    //     Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
    //         return types.SignTx(transaction, signer, privateKeyECDSA)
    //     },
    //     From:     cAddress,
    //     Context:  ctx,
    //     GasLimit: params.ApricotPhase1GasLimit,
    // }

    // // deploy the contract
    // storageAddress, storageTransaction, storageContract, err := DeployStorage(to, ec)
    // if err != nil {
    //     fmt.Println(err)
    // }

    // // wait for the transaction to be accepted
    // for {
    //     r, err := ec.TransactionReceipt(ctx, storageTransaction.Hash())
    //     if err != nil {
    //         if err.Error() != "not found" {
    //             fmt.Println(err)
    //         }
    //         time.Sleep(1 * time.Second)
    //         continue
    //     }
    //     if r.Status != 0 {
    //         break
    //     }
    //     time.Sleep(1 * time.Second)
    // }

    // log.Println("storageAddress", storageAddress)
    // log.Println("storageTransaction", storageTransaction)

    // Call store on the contract
    // storeTransaction, err := KepengToken.Name(to, big.NewInt(1), common.BytesToAddress([]byte("addr1")))
    // if err != nil {
    //     fmt.Println(err)
    // }

    // // wait for the transaction
    // for {
    //     r, err := ec.TransactionReceipt(ctx, storeTransaction.Hash())
    //     if err != nil {
    //         if err.Error() != "not found" {
    //             fmt.Println(err)
    //         }
    //         time.Sleep(1 * time.Second)
    //         continue
    //     }
    //     if r.Status != 0 {
    //         break
    //     }
    //     time.Sleep(1 * time.Second)
    // }

    // log.Println("storeTransaction", storeTransaction)

    // // setup call options for storage
    // co := &bind.CallOpts{
    //     Accepted: true,
    //     Context:  ctx,
    //     From:     storageAddress,
    // }

    // // retrieve the value of the contract
    // storageValue, err := storageContract.Retrieve(co)
    // if err != nil {
    //     fmt.Println(err)
    // }

    // log.Println("storageValue", storageValue)
}

func GetTokenDecimal(smartContractAddress string, rpcURL string, callerPrivateKeyString string) (uint8, string){
    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return 0, err.Error()
    }

    //get decimal token
    traDecimals, err := contract.Decimals(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err)
        return 0, err.Error()
    }
    
    return traDecimals, ""
}

func GetAddressBalance(smartContractAddress string, rpcURL string, callerPrivateKeyString string, holderAddress string) (*big.Int, string){
    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return nil, err.Error()
    }
    
    traBalanceOf, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(holderAddress))
    if err != nil {
        fmt.Println(err)
        return nil, err.Error()
    }
    
    return traBalanceOf, ""
}

func GetAllowance(smartContractAddress string, rpcURL string, callerPrivateKeyString string, ownerAddress string, spenderAddress string) (*big.Int, string){
    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return nil, err.Error()
    }
    
    traAllowence, err := contract.Allowance(&bind.CallOpts{}, common.HexToAddress(ownerAddress), common.HexToAddress(spenderAddress))
    if err != nil {
        fmt.Println(err)
        return nil, err.Error()
    }
    
    return traAllowence, ""
}

func Approve(smartContractAddress string, rpcURL string, callerPrivateKeyString string, spender string, amount *big.Int) (string, string){
    tx, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return "", err.Error()
    }

    //send token from and to specific address
    traApprove, err := contract.Approve(
        tx, 
        common.HexToAddress(spender), 
        amount,
    )
    if err != nil {
        //fmt.Println(err)
        return "", err.Error()
    }

    return traApprove.Hash().String(), ""
}

func TransferFrom(smartContractAddress string, rpcURL string, callerPrivateKeyString string, sender string, recipient string, amount *big.Int) string{
    tx, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    
    //send token from and to specific address
    traTransferFrom, err := contract.TransferFrom(
        tx, 
        common.HexToAddress(sender), 
        common.HexToAddress(recipient), 
        amount,
    )
    if err != nil {
        //fmt.Println(err)
        return err.Error()
    }

    return traTransferFrom.Hash().String()
}

func Transfer(smartContractAddress string, rpcURL string, callerPrivateKeyString string, recipient string, amount *big.Int) (string, string){
    tx, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return "", err.Error()
    }

    //send token from SC caller
    traTransfer, err := contract.Transfer(
        tx, 
        common.HexToAddress(recipient), 
        amount,
    )
    if err != nil {
        //fmt.Println(err)
        return "", err.Error()
    }

    return traTransfer.Hash().String(), ""
}
