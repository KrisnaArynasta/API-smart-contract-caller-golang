package main

import (
	"fmt"
    "math/big"
    "net/http"
    "strings"
    // "strconv"

    "github.com/ethereum/go-ethereum/common"
    "github.com/labstack/echo/v4"
    SmartContractCaller "smartcontractcallermodule"
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

type DataResAllowance struct {
    Owner             string  `json:"owner_address"`
    Spender           string  `json:"spender_address"`
    Allowance         *big.Int `json:"allowance"`
}

type DataResTransaction struct {
    TransactionHash     string  `json:"transaction_hash"`
}

func main(){
    
    //"PrivateKey-tvYwRHzhBX9p9D2oYsYQxDxozLqAhvPBR5hWhXmmHDmadMgiz" //Krisna Walet Avalanche
    //"PrivateKey-uVPEjPMvVthUgbQy6g9eg5qkedWtXjFwtApeViGPPMvmCmjwx" //Krisna Walet Metamask
    //"PrivateKey-v9CgXqT77DgJjWNLLpJP5td7v2DjPrdMTEsK6sCDjBGoRcisz" //Kepeng Wallet

    route := echo.New()
	route.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome!")
	})

    route.GET("/getgassprice", getGassPrice)
    route.GET("/getcallerpubaddress", getCallerPubAddress)
    route.GET("/gettokeninfo", getTokenInfo)
    route.GET("/getaddressbalance", getAddressBalance)
    route.GET("/getallowance", getAllowance)
    route.POST("/transfer", transfer)
    route.POST("/approve", approve)

	route.Logger.Fatal(route.Start(":3725"))
}

func rpcURL() string {
    return "https://api.avax-test.network/ext/bc/C/rpc" //"https://api.avax.network/ext/bc/C/rpc"
}

func smartContractAddress() string {
    //"0x836F91Ab1aE172958E494e8407B0F88d05166A03" //KEPENG TOKEN (MAIN NET)
    //0xeccd68E23CA0D0DD2184F20Db728BA08339FE602 //KEPENG TOKEN (TESTNET NET)
    //"0xCDE8A483758CA28a78267fc13832aB31b88F78C1" //KRISNA TOKEN
    return "0xeccd68E23CA0D0DD2184F20Db728BA08339FE602"
}

func decimalBalance(smartContractAddress string, rpcURL string) *big.Int{
    callerPrivateKeyString := "PrivateKey-tvYwRHzhBX9p9D2oYsYQxDxozLqAhvPBR5hWhXmmHDmadMgiz"
	traDecimals, err := SmartContractCaller.GetTokenDecimal(smartContractAddress, rpcURL, callerPrivateKeyString)
    if(err != ""){
        fmt.Println("error \t \t:",err)
        return big.NewInt(0)
    }
    decimalBalance := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(int(traDecimals))), nil)
    return decimalBalance
}

func getGassPrice(c echo.Context) error {
    rpcURL := rpcURL();
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
    }else{
        data.PubAddress = callerPubAddress
        res.Data = data
        res.Message = "success"
        res.Status = http.StatusOK
        
        fmt.Println("Caller Pub Address \t:", callerPubAddress)
    }
    

    return c.JSON(http.StatusOK, res)
}

func getTokenInfo(c echo.Context) error {
    var res Response
    var data DataResTokenInfo

    smartContractAddress := smartContractAddress();
    rpcURL := rpcURL();
    //callerPrivateKeyString := "PrivateKey-tvYwRHzhBX9p9D2oYsYQxDxozLqAhvPBR5hWhXmmHDmadMgiz"
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

func getAddressBalance(c echo.Context) error {
    var res Response
    var data DataResAddressBalance

    smartContractAddress := smartContractAddress();
    rpcURL := rpcURL();
    decimalBalance := decimalBalance(smartContractAddress, rpcURL)

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
    data.Balance = new(big.Int).Div(traAddrBlnc,decimalBalance)
    res.Data = data
	res.Message = "success"
    res.Status = http.StatusOK

    fmt.Printf("Saldo \t \t \t: %d (%s)\n",new(big.Int).Div(traAddrBlnc,decimalBalance), common.HexToAddress(holderAddress))
    return c.JSON(http.StatusOK, res)
}

func getAllowance(c echo.Context) error {
    var res Response
    var data DataResAllowance

    smartContractAddress := smartContractAddress();
    rpcURL := rpcURL();
    decimalBalance := decimalBalance(smartContractAddress, rpcURL)

    callerPrivateKeyString := c.QueryParam("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
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

func transfer(c echo.Context) error {
    var res Response
    var data DataResTransaction

    smartContractAddress := smartContractAddress();
    rpcURL := rpcURL();
    decimalBalance := decimalBalance(smartContractAddress, rpcURL)

    //context.Request().Header
    callerPrivateKeyString := c.FormValue("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
	recipient := c.FormValue("Recipient")
    if(recipient == "" || len(recipient) == 0){
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

	traTransfer, err := SmartContractCaller.Transfer(smartContractAddress, rpcURL, callerPrivateKeyString, recipient, amount)
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

func approve(c echo.Context) error {
    var res Response
    var data DataResTransaction

    smartContractAddress := smartContractAddress();
    rpcURL := rpcURL();
    decimalBalance := decimalBalance(smartContractAddress, rpcURL)

    //context.Request().Header
    callerPrivateKeyString := c.FormValue("PrivateKey")
    if(callerPrivateKeyString == "" || len(callerPrivateKeyString) == 0){
        res.Message = "Caller Caller credential is not valid"
        res.Status = http.StatusUnauthorized
        return echo.NewHTTPError(http.StatusUnauthorized, res)
    }
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
//     callerPrivateKeyString := "PrivateKey-tvYwRHzhBX9p9D2oYsYQxDxozLqAhvPBR5hWhXmmHDmadMgiz"
// 	sender := "0x3f75899e56b639106205c78a4e64c6a3c1e68b5e"
// 	recipient := "0xACd458287AE86e7B35050e398E361FAAc6b949AC"
//     amount := big.NewInt(12);
//     amount.Mul(amount,decimalBalance)
// 	traTransferFrom := SmartContractCaller.TransferFrom(smartContractAddress, rpcURL, callerPrivateKeyString, sender, recipient, amount)
//     fmt.Println()
//     fmt.Println("Transaction \t \t:",traTransferFrom)
// }
