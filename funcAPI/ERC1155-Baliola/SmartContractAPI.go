package SmartContractApiERC1155Baliola

import (
	"fmt"
    "math/big"
    "net/http"
    "strings"
    "strconv"

    //"github.com/ethereum/go-ethereum/common"
    "github.com/labstack/echo/v4"
    SmartContractCaller "SmartContractCallerERC1155Baliola"
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
        return "0xd70c7F73fA7366c4aAF8F4A5D2912822C788C19E" //1155
    }else{
        return "0xd70c7F73fA7366c4aAF8F4A5D2912822C788C19E"
    }
}

func baseURIMetaData(isMainNet bool) string{
    if(isMainNet == true){
        return "https://api.baliola.com/metadata?id="
    }else{
        return "https://api.baliola.com/metadata?id="
    }
}

func getGassPrice(c echo.Context) error {
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

func getCallerPubAddress(c echo.Context) error {
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

	traName, err := SmartContractCaller.GetTokenInfo(smartContractAddress, rpcURL, callerPrivateKeyString)
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

    data.TokenName = traName
    res.Data = data
	res.Message = "success"
    res.Status = http.StatusOK

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

    totalMintedNFT, traBalanceOfNFT, err := SmartContractCaller.GetListOfNftByOwner(smartContractAddress, rpcURL, callerPrivateKeyString, pubAddress)
    if(err != ""){
        res.Message = err
        if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit")){
            res.Message = "Caller credential is not valid"
        }
        res.Status = http.StatusUnauthorized

        fmt.Println("error \t \t \t:",err)
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
    
    for i := 0; i < totalMintedNFT; i++ {
        if(traBalanceOfNFT[i].Cmp(big.NewInt(int64(0))) > 0){
            var result DataResTokenURI
            
            result.Kind = "multiples"             
            result.SmartContractAddress = smartContractAddress
            result.Id = big.NewInt(int64(i))
            result.TokenURI = baseURIMetaData(isMainNet) + strconv.FormatInt(int64(i),10) +".json"
            result.Balance = traBalanceOfNFT[i]

            data = append(data,result)
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

    amount := new(big.Int)
    amount.SetString(c.FormValue("amount"), 10)
    if(len(amount.Bits()) == 0){
        res.Message = "amount must bigger than 0"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
    
	traTransfer, err := SmartContractCaller.SafeTransferFrom(smartContractAddress, rpcURL, callerPrivateKeyString, sender, recipient, tokenId, amount)
    if(err != ""){
        res.Message = err
        if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit")){
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
