package SmartContractCallerERC721Baliola

import (
	"fmt"
    "context"
    // "log"
    "math/big"
    //"strings"
    //"strconv"
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
    BaliolaSeries "BaliolaSeries"
)

func configSC(smartContractAddress string, rpcURL string, callerPrivateKeyString string) (*bind.TransactOpts, *BaliolaSeries.BaliolaSeries, error) {
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
    contract, err := BaliolaSeries.NewBaliolaSeries(scAddress, client)
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

func GetTokenInfo(smartContractAddress string, rpcURL string, callerPrivateKeyString string) (string, string, *big.Int, string) {

    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return "", "", nil, err.Error()
    }

    //get name token
    traName, err := contract.Name(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err.Error())
        return "", "", nil, err.Error()
    }

    //get symbol token
    traSym, err := contract.Symbol(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err.Error())
        return "", "", nil, err.Error()
    }
    
    //get total supply token
    traSupply, err := contract.TotalSupply(&bind.CallOpts{})
    if err != nil {
        fmt.Println(err.Error())
        return "", "", nil, err.Error()
    }
    
    return traName,traSym,traSupply, ""
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

func TransferFrom(smartContractAddress string, rpcURL string, callerPrivateKeyString string, sender string, recipient string, tokenId *big.Int) (string, string){
    tx, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return "", err.Error()
    }
    
    //send token from and to specific address
    traTransferFrom, err := contract.TransferFrom(
        tx, 
        common.HexToAddress(sender), 
        common.HexToAddress(recipient), 
        tokenId,
    )
    if err != nil {
        //fmt.Println(err)
        return "", err.Error()
    }

    return traTransferFrom.Hash().String(), ""
}

func GetOwerByNftId(smartContractAddress string, rpcURL string, callerPrivateKeyString string, tokenId *big.Int) (common.Address, string){
    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return common.HexToAddress("0x00"), err.Error()
    }
    
    traOwnerAddress, err := contract.OwnerOf(&bind.CallOpts{}, tokenId)
    if err != nil {
        fmt.Println(err)
        return common.HexToAddress("0x00"), err.Error()
    }
    
    return traOwnerAddress, ""
}

func GetListOfNftByOwner(smartContractAddress string, rpcURL string, callerPrivateKeyString string, pubAddress string, tokenId *big.Int) (string, string){
    traTokenURI := ""

    _, contract, err := configSC(smartContractAddress, rpcURL, callerPrivateKeyString)
    if err != nil {
        return "", err.Error()
    }

    traOwnerAddress, err := contract.OwnerOf(&bind.CallOpts{}, tokenId)
    if err != nil {
        fmt.Println(err)
        return "", err.Error()
    }

    if(pubAddress == traOwnerAddress.String()){
        traTokenURI, err = contract.TokenURI(&bind.CallOpts{}, tokenId)
        if err != nil {
            fmt.Println(err)
            return "", err.Error()
        }
    }

    return traTokenURI, ""
}