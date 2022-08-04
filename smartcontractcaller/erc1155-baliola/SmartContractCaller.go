package SmartContractCaller1155Baliola

import (
	"fmt"
    "context"
    // "log"
    "math/big"
    //"strings"
    "strconv"
    // "time"

    // "github.com/ava-labs/avalanchego/utils/constants"
    //"github.com/ava-labs/avalanchego/utils/formatting"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"//"github.com/ava-labs/coreth/accounts/abi/bind"
    // "github.com/ava-labs/coreth/core/types"
    // "github.com/ava-labs/coreth/params"
    // "github.com/ava-labs/coreth/rpc"
    //"github.com/decred/dcrd/dcrec/secp256k1/v3"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"// "github.com/ava-labs/coreth/ethclient"
    BaliolaMultiplesCollections "BaliolaMultiplesCollections"
)

func configSC(smartContractAddress string, rpcURL string, callerPrivateKeyString string) (*bind.TransactOpts, *BaliolaMultiplesCollections.BaliolaMultiplesCollections, error) {
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
    callerPrivateKeyECDSA, err := crypto.HexToECDSA(callerPrivateKeyString)
    if err != nil {
        fmt.Println(err)
        return nil, nil, err
    }
    
    cAddress := crypto.PubkeyToAddress(callerPrivateKeyECDSA.PublicKey)

    // derive 'c' address
    nonce, err := client.PendingNonceAt(ctx, cAddress)

    scAddress := common.HexToAddress(smartContractAddress)
    contract, err := BaliolaMultiplesCollections.NewBaliolaMultiplesCollections(scAddress, client)
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

    callerPrivateKeyECDSA, err := crypto.HexToECDSA(callerPrivateKeyString)
    if err != nil {
        return err.Error()
    }
    // derive 'c' address
    cAddress := crypto.PubkeyToAddress(callerPrivateKeyECDSA.PublicKey)
    if err != nil {
        //fmt.Println(err)
        return err.Error()
    }

    return cAddress.String()
}

func GetTokenInfo(smartContractAddress string, rpcURL string, callerPrivateKeyString string) (string, string) {

    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return "",  err.Error()
    }

    //get name token
    traName, err := contract.Name(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err.Error())
        return "",  err.Error()
    }
    
    return traName, ""
}

func GetAddressBalance(smartContractAddress string, rpcURL string, callerPrivateKeyString string, holderAddress string, tokenId *big.Int) (*big.Int, string){
    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return nil, err.Error()
    }
    
    traBalanceOf, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(holderAddress), tokenId)
    if err != nil {
        fmt.Println(err)
        return nil, err.Error()
    }
    
    return traBalanceOf, ""
}

func SafeTransferFrom(smartContractAddress string, rpcURL string, callerPrivateKeyString string, sender string, recipient string, tokenId *big.Int, amount *big.Int) (string, string){
    tx, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return "", err.Error()
    }
    
    traExists, err := contract.Exists(&bind.CallOpts{}, tokenId)
    if err != nil {
        fmt.Println(err)
        return "", err.Error()
    }
    if(traExists == false){
        return "", "Token is not exists"
    }


    traBalanceOf, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(sender), tokenId)
    if err != nil {
        fmt.Println(err)
        return "", err.Error()
    }

    if(amount.Cmp(traBalanceOf) > 0){
        return "", "Amount must be less than Balance"
    }
    trasafeTransferFrom, err := contract.SafeTransferFrom(
        tx, 
        common.HexToAddress(sender), 
        common.HexToAddress(recipient), 
        tokenId,
        amount,
        []byte{0x00, 0x00},
    )
    if err != nil {
        return "", err.Error()
    }

    return trasafeTransferFrom.Hash().String(), ""
}

func GetListOfNftByOwner(smartContractAddress string, rpcURL string, callerPrivateKeyString string, pubAddress string) (int, []*big.Int, string){
    var accounts []common.Address
    var ids []*big.Int

    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return 0, nil, err.Error()
    }

    traCurrentTokenId, err := contract.GetCurrentTokenId(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err)
        return 0, nil, err.Error()
    }

    n, _ := strconv.Atoi(traCurrentTokenId.Text(10))
    for i := 0; i < n; i++ {
        tokenId := big.NewInt(int64(i))
        accounts = append(accounts,common.HexToAddress(pubAddress))
        ids = append(ids,tokenId)
    }

    traBalanceOfNFT, err := contract.BalanceOfBatch(&bind.CallOpts{}, accounts, ids)
    if err != nil {
        fmt.Println(err)
        return 0, nil, err.Error()
    }

    return n, traBalanceOfNFT, ""
}