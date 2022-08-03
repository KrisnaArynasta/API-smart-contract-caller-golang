package main

import (
	"net/http"

    "github.com/labstack/echo/v4"
    ApiERC20Kepeng "smartContractApiERC20Kepeng"
    ApiERC721Baliola "smartContractApiERC721Baliola"
    ApiERC1155Baliola "smartContractApiERC1155Baliola"
)

func main(){

    route := echo.New()
	route.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to Betamax API. To see our documentation please contact our developer!")
	})

    route.GET("/getgassprice", ApiERC20Kepeng.GetGassPrice)
    route.GET("/getcallerpubaddress", ApiERC20Kepeng.GetCallerPubAddress)
    route.GET("/gettokeninfo", ApiERC20Kepeng.GetTokenInfo)
    route.GET("/getaddressbalance", ApiERC20Kepeng.GetAddressBalance)
    route.GET("/getallowance", ApiERC20Kepeng.GetAllowance)
    route.POST("/transfer", ApiERC20Kepeng.Transfer)
    route.POST("/approve", ApiERC20Kepeng.Approve)

    route.GET("/721/gettokeninfo", ApiERC721Baliola.GetTokenInfo)
    route.GET("/721/getOwerByNftId", ApiERC721Baliola.GetOwerByNftId)
    route.GET("/721/getListOfNftByOwner", ApiERC721Baliola.GetListOfNftByOwner)
    route.POST("/721/transfer", ApiERC721Baliola.Transfer)

    route.GET("/1155/gettokeninfo", ApiERC1155Baliola.GetTokenInfo)
    route.GET("/1155/getListOfNftByOwner", ApiERC1155Baliola.GetListOfNftByOwner)
    route.POST("/1155/transfer", ApiERC1155Baliola.Transfer)

	route.Logger.Fatal(route.Start(":3725"))
}