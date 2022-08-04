package SmartContractApiERC721Baliola

import (
	"fmt"
    "math/big"
    "net/http"
    "strings"
    "strconv"

    //"github.com/ethereum/go-ethereum/common"
    "github.com/labstack/echo/v4"
    SmartContractCaller "SmartContractCallerERC721Baliola"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DataResGassPrice struct {
    GasPrice          string `json:"gas_price"`
}

type DataResCallerPubAddress struct {
    PubAddress        string `json:"pub_address"`
}

type DataResTokenInfo struct {
    TokenName          string   `json:"toke_name"`
    TokenSymbol        string   `json:"token_symbol"`
    TokenDecimals      string   `json:"token_decimals"`
    TotalSupply        *big.Int `json:"total_supply"`
}

type DataResAddressBalance struct {
    Address             string  `json:"pub_address"`
    Balance             *big.Int `json:"balance"`
}

type DataResTransaction struct {
    TransactionHash     string  `json:"transaction_hash"`
}

type DataResTokenURI struct {
    Kind                    string  `json:"kind"`
    SmartContractAddress    string  `json:"smart_contract_address"`
    Id                      *big.Int  `json:"token_id"`
    TokenURI                string    `json:"token_uri"`
    Balance                 *big.Int   `json:"balance"`
}

func rpcURL(isMainNet bool) string {
    if(isMainNet == true){
        return "https://api.avax.network/ext/bc/C/rpc"
    }else{
        return "https://api.avax-test.network/ext/bc/C/rpc"
    }
     
}

func smartContractAddress(isMainNet bool) string {   
    if(isMainNet == true){
        return "0x9d9c79b5060cdF2A8954EfC9FD307030e879f4ee"
    }else{
        return "0x0259C1E0b1Bf123754056d243CF5AD5Fb1427de0"
    }
}

func GetGassPrice(c echo.Context) error {
    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    rpcURL := rpcURL(isMainNet);
    gasPrice := SmartContractCaller.GetGassPrice(rpcURL)
    gasPriceStr := gasPrice.String()

    var res Response
    var data DataResGassPrice

    data.GasPrice = gasPriceStr

    res.Data = data
	res.Message = "success"
    res.Status = http.StatusOK

    fmt.Println("Gas Price \t \t:",gasPrice)
    return c.JSON(http.StatusOK, res)
}

func GetCallerPubAddress(c echo.Context) error {
    var res Response
    var data DataResCallerPubAddress

    callerPrivateKeyString := c.QueryParam("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

    callerPubAddress := SmartContractCaller.GetCallerPubAddress(callerPrivateKeyString)
    if(strings.Contains(callerPubAddress, "invalid input checksum") || strings.Contains(callerPubAddress, "Invalid base58 digit")){
        res.Message = "Caller credential is not valid"
        res.Status = http.StatusUnauthorized

        fmt.Println("error \t \t \t:",callerPubAddress)
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }else{
        data.PubAddress = callerPubAddress
        res.Data = data
        res.Message = "success"
        res.Status = http.StatusOK
        
        fmt.Println("Caller Pub Address \t:", callerPubAddress)
        return c.JSON(http.StatusOK, res)
    }
}

func GetTokenInfo(c echo.Context) error {
    var res Response
    var data DataResTokenInfo

    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    smartContractAddress := smartContractAddress(isMainNet);
    rpcURL := rpcURL(isMainNet);
    callerPrivateKeyString := c.QueryParam("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

	traName,traSym,traTtlSup, err := SmartContractCaller.GetTokenInfo(smartContractAddress, rpcURL, callerPrivateKeyString)
    if(err != ""){
        res.Message = err
        if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit")){
            res.Message = "Caller credential is not valid"
        }
        res.Status = http.StatusUnauthorized

        fmt.Println("error \t \t \t:",traName)
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

    fmt.Println("Token Name \t \t:",traName)
    fmt.Println("Token Symbol \t \t:",traSym)
    fmt.Println("Total Supply \t \t:",traTtlSup)

    data.TokenName = traName
    data.TokenSymbol = traSym
    data.TotalSupply = traTtlSup
    res.Data = data
	res.Message = "success"
    res.Status = http.StatusOK

    return c.JSON(http.StatusOK, res)
}

func GetOwerByNftId(c echo.Context) error {
    var res Response
    var data DataResCallerPubAddress

    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    smartContractAddress := smartContractAddress(isMainNet);
    rpcURL := rpcURL(isMainNet);
    callerPrivateKeyString := c.QueryParam("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

    tokenId := new(big.Int)
    tokenId, _ = tokenId.SetString(c.QueryParam("tokenId"),10)

	traOwnerAddress, err := SmartContractCaller.GetOwerByNftId(smartContractAddress, rpcURL, callerPrivateKeyString, tokenId)
    if(err != ""){
        res.Message = err
        if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit")){
            res.Message = "Caller credential is not valid"
        }
        res.Status = http.StatusUnauthorized

        fmt.Println("error \t \t \t:",traOwnerAddress)
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

    fmt.Println("Address Owner \t \t:",traOwnerAddress)

    data.PubAddress = traOwnerAddress.String()
	res.Message = "success"
    res.Status = http.StatusOK
    res.Data = data

    return c.JSON(http.StatusOK, res)
}

func GetListOfNftByOwner(c echo.Context) error {
    var res Response
    var data []DataResTokenURI = make([]DataResTokenURI,0)

    pubAddress := c.QueryParam("PubAddress")
    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    smartContractAddress := smartContractAddress(isMainNet);
    rpcURL := rpcURL(isMainNet);
    callerPrivateKeyString := c.QueryParam("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

    errString,_,traTtlSup, err := SmartContractCaller.GetTokenInfo(smartContractAddress, rpcURL, callerPrivateKeyString)
    if(err != ""){
        res.Message = err
        if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit")){
            res.Message = "Caller credential is not valid"
        }
        res.Status = http.StatusUnauthorized

        fmt.Println("error \t \t \t:",errString)
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

    n, _ := strconv.Atoi(traTtlSup.Text(10))
    
    for i := 1; i < n; i++ {
        tokenId := big.NewInt(int64(i))
        traTokenURI, err := SmartContractCaller.GetListOfNftByOwner(smartContractAddress, rpcURL, callerPrivateKeyString, pubAddress, tokenId)
        if(err != ""){
            res.Message = err
            if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit")){
                res.Message = "Caller credential is not valid"
            }
            res.Status = http.StatusUnauthorized
    
            fmt.Println("error \t \t \t:",errString)
            return echo.NewHTTPError(http.StatusUnauthorized, res)
        }
        
        if(traTokenURI != ""){
            var result DataResTokenURI
            result.Kind = "single" // "multiples"
            result.SmartContractAddress = smartContractAddress
            result.Id = tokenId
            result.TokenURI = traTokenURI
            result.Balance = big.NewInt(int64(1))

            data = append(data,result)

            
            fmt.Println("Smart Contract \t \t:",result.Kind)
            fmt.Println("Asset Kind \t \t:",result.SmartContractAddress)
            fmt.Println("Token ID \t \t:",tokenId)
            fmt.Println("Token URI \t \t:",traTokenURI)
            fmt.Println("Balance \t \t:",result.Balance)
            fmt.Println("")
        }
    }

	res.Message = "success"
    res.Status = http.StatusOK
    res.Data = data

    return c.JSON(http.StatusOK, res)
}

func Transfer(c echo.Context) error {
    var res Response
    var data DataResTransaction

    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    smartContractAddress := smartContractAddress(isMainNet);
    rpcURL := rpcURL(isMainNet);

    callerPrivateKeyString := c.FormValue("privateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

    sender := c.FormValue("sender")
    if(sender == "" || len(sender) == 0){
        res.Message = "Holder Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

	recipient := c.FormValue("recipient")
    if(recipient == "" || len(recipient) == 0){
        res.Message = "Recipient Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
    
    tokenId := new(big.Int)
    tokenId.SetString(c.FormValue("tokenId"), 10)
    
	traTransfer, err := SmartContractCaller.TransferFrom(smartContractAddress, rpcURL, callerPrivateKeyString, sender, recipient, tokenId)
    if(err != ""){
        res.Message = err
        if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit") || strings.Contains(err, "invalid hex character")){
            res.Message = "Caller credential is not valid"
        }

        res.Status = http.StatusUnauthorized

        fmt.Println("error \t \t \t:",err)
        return c.JSON(http.StatusUnauthorized, res)
    }

    data.TransactionHash = traTransfer
    res.Data = data
	res.Message = "success"
    res.Status = http.StatusOK

    fmt.Println("Transaction Hash \t:", traTransfer)
    return c.JSON(http.StatusOK, res)

}
