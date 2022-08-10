package SmartContractApiERC20Kepeng

import (
	"fmt"
    "math/big"
    "net/http"
    "strings"
    "strconv"

    "github.com/ethereum/go-ethereum/common"
    "github.com/labstack/echo/v4"
    SmartContractCaller "SmartContractCallerERC20Kepeng"
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
    Balance             *big.Float `json:"balance"`
}

type DataResAllowance struct {
    Owner             string  `json:"owner_address"`
    Spender           string  `json:"spender_address"`
    Allowance         *big.Int `json:"allowance"`
}

type DataResTransaction struct {
    TransactionHash     string  `json:"transaction_hash"`
}

func RpcURL(isMainNet bool) string {
    if(isMainNet == true){
        return "https://api.avax.network/ext/bc/C/rpc"
        //return "https://api.avax.network/ext/bc/X"
    }else{
        return "https://api.avax-test.network/ext/bc/C/rpc"
    }
     
}

func SmartContractAddress(isMainNet bool) string {
    //"0x836F91Ab1aE172958E494e8407B0F88d05166A03" //KEPENG TOKEN (MAIN NET)
    //0xeccd68E23CA0D0DD2184F20Db728BA08339FE602 //KEPENG TOKEN (TEST NET)
    //"0xCDE8A483758CA28a78267fc13832aB31b88F78C1" //KRISNA TOKEN
   
    if(isMainNet == true){
        return "0x836F91Ab1aE172958E494e8407B0F88d05166A03"
        //return "0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7" //wavax
        //return "2oYMBNV4eNHyqk2fjjV5nVQLDbtmNJzq5s3qs3Lo6ftnC6FByM"
    }else{
        return "0xeccd68E23CA0D0DD2184F20Db728BA08339FE602"
    }
}

func DecimalBalance(smartContractAddress string, rpcURL string, callerPrivateKeyString string) *big.Int{
	traDecimals, err := SmartContractCaller.GetTokenDecimal(smartContractAddress, rpcURL, callerPrivateKeyString)
    if(err != ""){
        fmt.Println("error \t \t:",err)
        return big.NewInt(0)
    }
    decimalBalance := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(int(traDecimals))), nil)
    return decimalBalance
}

func GetGassPrice(c echo.Context) error {
    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    rpcURL := RpcURL(isMainNet);
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
    smartContractAddress := SmartContractAddress(isMainNet);
    rpcURL := RpcURL(isMainNet);
    callerPrivateKeyString := c.QueryParam("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

	traName,traDecimals,traSym,traTtlSup, err := SmartContractCaller.GetTokenInfo(smartContractAddress, rpcURL, callerPrivateKeyString)
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
    fmt.Println("Token Decimals \t \t:",traDecimals)
    fmt.Println("Total Supply \t \t:",traTtlSup)

    data.TokenName = traName
    data.TokenSymbol = traSym
    data.TokenDecimals = traDecimals
    data.TotalSupply = traTtlSup
    res.Data = data
	res.Message = "success"
    res.Status = http.StatusOK

    return c.JSON(http.StatusOK, res)
}

func GetAddressBalance(c echo.Context) error {
    var res Response
    var data DataResAddressBalance

    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    smartContractAddress := SmartContractAddress(isMainNet);
    rpcURL := RpcURL(isMainNet);

    callerPrivateKeyString := c.QueryParam("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
	holderAddress := c.QueryParam("PubAddress")
    if(holderAddress == "" || len(holderAddress) == 0){
        res.Message = "Holder Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

    decimalBalance := DecimalBalance(smartContractAddress, rpcURL, callerPrivateKeyString)
	traAddrBlnc, err := SmartContractCaller.GetAddressBalance(smartContractAddress, rpcURL, callerPrivateKeyString, holderAddress)
    if(err != ""){
        res.Message = err
        if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit")){
            res.Message = "Caller credential is not valid"
        }
        res.Status = http.StatusUnauthorized

        fmt.Println("error \t \t \t:",err)
        return c.JSON(http.StatusUnauthorized, res)
    }

    data.Address = holderAddress
    data.Balance = new(big.Float).Quo(new(big.Float).SetInt(traAddrBlnc),new(big.Float).SetInt(decimalBalance))
    res.Data = data
	res.Message = "success"
    res.Status = http.StatusOK

    fmt.Printf("Saldo \t \t \t: %d (%s)\n",traAddrBlnc, common.HexToAddress(holderAddress))
    return c.JSON(http.StatusOK, res)
}

func GetAllowance(c echo.Context) error {
    var res Response
    var data DataResAllowance

    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    smartContractAddress := SmartContractAddress(isMainNet);
    rpcURL := RpcURL(isMainNet);

    callerPrivateKeyString := c.QueryParam("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
    decimalBalance := DecimalBalance(smartContractAddress, rpcURL, callerPrivateKeyString)

	ownerAddress := c.QueryParam("Owner")
    if(ownerAddress == "" || len(ownerAddress) == 0){
        res.Message = "Holder Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

	spenderAddress := c.QueryParam("Spender")
    if(spenderAddress == "" || len(spenderAddress) == 0){
        res.Message = "Holder Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

	traAllowence, err := SmartContractCaller.GetAllowance(smartContractAddress, rpcURL, callerPrivateKeyString, ownerAddress, spenderAddress)
    if(err != ""){
        res.Message = err
        if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit")){
            res.Message = "Caller credential is not valid"
        }
        res.Status = http.StatusUnauthorized

        fmt.Println("error \t \t \t:",err)
        return c.JSON(http.StatusUnauthorized, res)
    }

    data.Owner = ownerAddress
    data.Spender = spenderAddress
    data.Allowance = new(big.Int).Div(traAllowence,decimalBalance)
    res.Data = data
	res.Message = "success"
    res.Status = http.StatusOK

    fmt.Println("Allowance \t \t \t:",new(big.Int).Div(traAllowence,decimalBalance))
    return c.JSON(http.StatusOK, res)
}

func Transfer(c echo.Context) error {
    var res Response
    var data DataResTransaction

    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    smartContractAddress := SmartContractAddress(isMainNet);
    rpcURL := RpcURL(isMainNet);

    //context.Request().Header
    callerPrivateKeyString := c.FormValue("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
    decimalBalance := DecimalBalance(smartContractAddress, rpcURL, callerPrivateKeyString)

	recipient := c.FormValue("Recipient")
    if(recipient == "" || len(recipient) == 0){
        res.Message = "Holder Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
    
    //jadiin big int param dari post
    // amount := new(big.Int)
    // amount.SetString(c.FormValue("amount"), 10)
    amount, _ := strconv.ParseFloat(c.FormValue("amount"), 64)
    
    //jadiin float decimal token dan amount, trus d kali antara param amount dan decimal
    amountFloat := new(big.Float)
    amountFloat.Mul(new(big.Float).SetFloat64(amount), new(big.Float).SetInt(decimalBalance))

    //hasil perkalian dijadiin big int
    amountBigInt := new(big.Int)
    amountFloat.Int(amountBigInt)
    fmt.Println(amountBigInt)
    if(len(amountBigInt.Bits()) == 0){
        res.Message = "amount must bigger than 0"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

	traTransfer, err := SmartContractCaller.Transfer(smartContractAddress, rpcURL, callerPrivateKeyString, recipient, amountBigInt)
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

func Approve(c echo.Context) error {
    var res Response
    var data DataResTransaction

    isMainNet, _ := strconv.ParseBool(c.QueryParam("isMainNet"))
    smartContractAddress := SmartContractAddress(isMainNet);
    rpcURL := RpcURL(isMainNet);

    //context.Request().Header
    callerPrivateKeyString := c.FormValue("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
    decimalBalance := DecimalBalance(smartContractAddress, rpcURL, callerPrivateKeyString)

    spender := c.FormValue("Spender")
    if(spender == "" || len(spender) == 0){
        res.Message = "Holder Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }

    amount := new(big.Int)
    amount.SetString(c.FormValue("amount"), 10)
    if(len(amount.Bits()) == 0){
        res.Message = "amount must bigger than 0"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
    amount.Mul(amount,decimalBalance)

	traApprove, err := SmartContractCaller.Approve(smartContractAddress, rpcURL, callerPrivateKeyString, spender, amount)
    if(err != ""){
        res.Message = err
        if(strings.Contains(err, "invalid input checksum") || strings.Contains(err, "Invalid base58 digit")){
            res.Message = "Caller credential is not valid"
        }

        res.Status = http.StatusUnauthorized

        fmt.Println("error \t \t \t:",err)
        return c.JSON(http.StatusUnauthorized, res)
    }

    data.TransactionHash = traApprove
    res.Data = data
	res.Message = "success"
    res.Status = http.StatusOK

    fmt.Println("Transaction Hash \t:", traApprove)
    return c.JSON(http.StatusOK, res)
}

// func transferFrom(smartContractAddress string, rpcURL string, decimalBalance *big.Int){
// 	sender := "0x3f75899e56b639106205c78a4e64c6a3c1e68b5e"
// 	recipient := "0xACd458287AE86e7B35050e398E361FAAc6b949AC"
//     amount := big.NewInt(12);
//     amount.Mul(amount,decimalBalance)
// 	traTransferFrom := SmartContractCaller.TransferFrom(smartContractAddress, rpcURL, callerPrivateKeyString, sender, recipient, amount)
//     fmt.Println()
//     fmt.Println("Transaction \t \t:",traTransferFrom)
// }
