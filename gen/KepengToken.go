// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KepengToken

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KepengTokenMetaData contains all meta data concerning the KepengToken contract.
var KepengTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600681526020017f4b4550454e4700000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4b5047000000000000000000000000000000000000000000000000000000000081525081600390816200008f91906200067f565b508060049081620000a191906200067f565b5050506000600560006101000a81548160ff021916908315150217905550620000df620000d36200012560201b60201c565b6200012d60201b60201c565b6200011f33620000f4620001f360201b60201c565b600a620001029190620008f6565b633b9aca0062000113919062000947565b620001fc60201b60201c565b62000b28565b600033905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60006004905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036200026e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002659062000a09565b60405180910390fd5b62000282600083836200037460201b60201c565b806002600082825462000296919062000a2b565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620002ed919062000a2b565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405162000354919062000a99565b60405180910390a36200037060008383620003e460201b60201c565b5050565b62000384620003e960201b60201c565b15620003c7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003be9062000b06565b60405180910390fd5b620003df8383836200040060201b62000bc81760201c565b505050565b505050565b6000600560009054906101000a900460ff16905090565b505050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200048757607f821691505b6020821081036200049d576200049c6200043f565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005077fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620004c8565b620005138683620004c8565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620005606200055a62000554846200052b565b62000535565b6200052b565b9050919050565b6000819050919050565b6200057c836200053f565b620005946200058b8262000567565b848454620004d5565b825550505050565b600090565b620005ab6200059c565b620005b881848462000571565b505050565b5b81811015620005e057620005d4600082620005a1565b600181019050620005be565b5050565b601f8211156200062f57620005f981620004a3565b6200060484620004b8565b8101602085101562000614578190505b6200062c6200062385620004b8565b830182620005bd565b50505b505050565b600082821c905092915050565b6000620006546000198460080262000634565b1980831691505092915050565b60006200066f838362000641565b9150826002028217905092915050565b6200068a8262000405565b67ffffffffffffffff811115620006a657620006a562000410565b5b620006b282546200046e565b620006bf828285620005e4565b600060209050601f831160018114620006f75760008415620006e2578287015190505b620006ee858262000661565b8655506200075e565b601f1984166200070786620004a3565b60005b8281101562000731578489015182556001820191506020850194506020810190506200070a565b868310156200075157848901516200074d601f89168262000641565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b6001851115620007f457808604811115620007cc57620007cb62000766565b5b6001851615620007dc5780820291505b8081029050620007ec8562000795565b9450620007ac565b94509492505050565b6000826200080f5760019050620008e2565b816200081f5760009050620008e2565b8160018114620008385760028114620008435762000879565b6001915050620008e2565b60ff84111562000858576200085762000766565b5b8360020a91508482111562000872576200087162000766565b5b50620008e2565b5060208310610133831016604e8410600b8410161715620008b35782820a905083811115620008ad57620008ac62000766565b5b620008e2565b620008c28484846001620007a2565b92509050818404811115620008dc57620008db62000766565b5b81810290505b9392505050565b600060ff82169050919050565b600062000903826200052b565b91506200091083620008e9565b92506200093f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484620007fd565b905092915050565b600062000954826200052b565b915062000961836200052b565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156200099d576200099c62000766565b5b828202905092915050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000620009f1601f83620009a8565b9150620009fe82620009b9565b602082019050919050565b6000602082019050818103600083015262000a2481620009e2565b9050919050565b600062000a38826200052b565b915062000a45836200052b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562000a7d5762000a7c62000766565b5b828201905092915050565b62000a93816200052b565b82525050565b600060208201905062000ab0600083018462000a88565b92915050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b600062000aee601083620009a8565b915062000afb8262000ab6565b602082019050919050565b6000602082019050818103600083015262000b218162000adf565b9050919050565b6120768062000b386000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806370a08231116100ad57806395d89b411161007157806395d89b41146102d2578063a457c2d7146102f0578063a9059cbb14610320578063dd62ed3e14610350578063f2fde38b1461038057610121565b806370a0823114610254578063715018a61461028457806379cc67901461028e5780638456cb59146102aa5780638da5cb5b146102b457610121565b8063313ce567116100f4578063313ce567146101c257806339509351146101e05780633f4ba83a1461021057806342966c681461021a5780635c975abb1461023657610121565b806306fdde0314610126578063095ea7b31461014457806318160ddd1461017457806323b872dd14610192575b600080fd5b61012e61039c565b60405161013b91906114f4565b60405180910390f35b61015e600480360381019061015991906115af565b61042e565b60405161016b919061160a565b60405180910390f35b61017c61044c565b6040516101899190611634565b60405180910390f35b6101ac60048036038101906101a7919061164f565b610456565b6040516101b9919061160a565b60405180910390f35b6101ca61054e565b6040516101d791906116be565b60405180910390f35b6101fa60048036038101906101f591906115af565b610557565b604051610207919061160a565b60405180910390f35b610218610603565b005b610234600480360381019061022f91906116d9565b610689565b005b61023e61069d565b60405161024b919061160a565b60405180910390f35b61026e60048036038101906102699190611706565b6106b4565b60405161027b9190611634565b60405180910390f35b61028c6106fc565b005b6102a860048036038101906102a391906115af565b610784565b005b6102b26107ff565b005b6102bc610885565b6040516102c99190611742565b60405180910390f35b6102da6108af565b6040516102e791906114f4565b60405180910390f35b61030a600480360381019061030591906115af565b610941565b604051610317919061160a565b60405180910390f35b61033a600480360381019061033591906115af565b610a2c565b604051610347919061160a565b60405180910390f35b61036a6004803603810190610365919061175d565b610a4a565b6040516103779190611634565b60405180910390f35b61039a60048036038101906103959190611706565b610ad1565b005b6060600380546103ab906117cc565b80601f01602080910402602001604051908101604052809291908181526020018280546103d7906117cc565b80156104245780601f106103f957610100808354040283529160200191610424565b820191906000526020600020905b81548152906001019060200180831161040757829003601f168201915b5050505050905090565b600061044261043b610bcd565b8484610bd5565b6001905092915050565b6000600254905090565b6000610463848484610d9e565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006104ae610bcd565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561052e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105259061186f565b60405180910390fd5b6105428561053a610bcd565b858403610bd5565b60019150509392505050565b60006004905090565b60006105f9610564610bcd565b848460016000610572610bcd565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105f491906118be565b610bd5565b6001905092915050565b61060b610bcd565b73ffffffffffffffffffffffffffffffffffffffff16610629610885565b73ffffffffffffffffffffffffffffffffffffffff161461067f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067690611960565b60405180910390fd5b61068761101d565b565b61069a610694610bcd565b826110bf565b50565b6000600560009054906101000a900460ff16905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610704610bcd565b73ffffffffffffffffffffffffffffffffffffffff16610722610885565b73ffffffffffffffffffffffffffffffffffffffff1614610778576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076f90611960565b60405180910390fd5b6107826000611295565b565b600061079783610792610bcd565b610a4a565b9050818110156107dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d3906119f2565b60405180910390fd5b6107f0836107e8610bcd565b848403610bd5565b6107fa83836110bf565b505050565b610807610bcd565b73ffffffffffffffffffffffffffffffffffffffff16610825610885565b73ffffffffffffffffffffffffffffffffffffffff161461087b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087290611960565b60405180910390fd5b61088361135b565b565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546108be906117cc565b80601f01602080910402602001604051908101604052809291908181526020018280546108ea906117cc565b80156109375780601f1061090c57610100808354040283529160200191610937565b820191906000526020600020905b81548152906001019060200180831161091a57829003601f168201915b5050505050905090565b60008060016000610950610bcd565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610a0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0490611a84565b60405180910390fd5b610a21610a18610bcd565b85858403610bd5565b600191505092915050565b6000610a40610a39610bcd565b8484610d9e565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610ad9610bcd565b73ffffffffffffffffffffffffffffffffffffffff16610af7610885565b73ffffffffffffffffffffffffffffffffffffffff1614610b4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4490611960565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bbc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb390611b16565b60405180910390fd5b610bc581611295565b50565b505050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3b90611ba8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610cb3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610caa90611c3a565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610d919190611634565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610e0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e0490611ccc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7390611d5e565b60405180910390fd5b610e878383836113fe565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610f0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f0490611df0565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610fa091906118be565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516110049190611634565b60405180910390a3611017848484611456565b50505050565b61102561069d565b611064576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105b90611e5c565b60405180910390fd5b6000600560006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6110a8610bcd565b6040516110b59190611742565b60405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361112e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112590611eee565b60405180910390fd5b61113a826000836113fe565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156111c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b790611f80565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546112179190611fa0565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161127c9190611634565b60405180910390a361129083600084611456565b505050565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61136361069d565b156113a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161139a90612020565b60405180910390fd5b6001600560006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586113e7610bcd565b6040516113f49190611742565b60405180910390a1565b61140661069d565b15611446576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161143d90612020565b60405180910390fd5b611451838383610bc8565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561149557808201518184015260208101905061147a565b838111156114a4576000848401525b50505050565b6000601f19601f8301169050919050565b60006114c68261145b565b6114d08185611466565b93506114e0818560208601611477565b6114e9816114aa565b840191505092915050565b6000602082019050818103600083015261150e81846114bb565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006115468261151b565b9050919050565b6115568161153b565b811461156157600080fd5b50565b6000813590506115738161154d565b92915050565b6000819050919050565b61158c81611579565b811461159757600080fd5b50565b6000813590506115a981611583565b92915050565b600080604083850312156115c6576115c5611516565b5b60006115d485828601611564565b92505060206115e58582860161159a565b9150509250929050565b60008115159050919050565b611604816115ef565b82525050565b600060208201905061161f60008301846115fb565b92915050565b61162e81611579565b82525050565b60006020820190506116496000830184611625565b92915050565b60008060006060848603121561166857611667611516565b5b600061167686828701611564565b935050602061168786828701611564565b92505060406116988682870161159a565b9150509250925092565b600060ff82169050919050565b6116b8816116a2565b82525050565b60006020820190506116d360008301846116af565b92915050565b6000602082840312156116ef576116ee611516565b5b60006116fd8482850161159a565b91505092915050565b60006020828403121561171c5761171b611516565b5b600061172a84828501611564565b91505092915050565b61173c8161153b565b82525050565b60006020820190506117576000830184611733565b92915050565b6000806040838503121561177457611773611516565b5b600061178285828601611564565b925050602061179385828601611564565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806117e457607f821691505b6020821081036117f7576117f661179d565b5b50919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b6000611859602883611466565b9150611864826117fd565b604082019050919050565b600060208201905081810360008301526118888161184c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118c982611579565b91506118d483611579565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156119095761190861188f565b5b828201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061194a602083611466565b915061195582611914565b602082019050919050565b600060208201905081810360008301526119798161193d565b9050919050565b7f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f7760008201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b60006119dc602483611466565b91506119e782611980565b604082019050919050565b60006020820190508181036000830152611a0b816119cf565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000611a6e602583611466565b9150611a7982611a12565b604082019050919050565b60006020820190508181036000830152611a9d81611a61565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611b00602683611466565b9150611b0b82611aa4565b604082019050919050565b60006020820190508181036000830152611b2f81611af3565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611b92602483611466565b9150611b9d82611b36565b604082019050919050565b60006020820190508181036000830152611bc181611b85565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000611c24602283611466565b9150611c2f82611bc8565b604082019050919050565b60006020820190508181036000830152611c5381611c17565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611cb6602583611466565b9150611cc182611c5a565b604082019050919050565b60006020820190508181036000830152611ce581611ca9565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611d48602383611466565b9150611d5382611cec565b604082019050919050565b60006020820190508181036000830152611d7781611d3b565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611dda602683611466565b9150611de582611d7e565b604082019050919050565b60006020820190508181036000830152611e0981611dcd565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b6000611e46601483611466565b9150611e5182611e10565b602082019050919050565b60006020820190508181036000830152611e7581611e39565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611ed8602183611466565b9150611ee382611e7c565b604082019050919050565b60006020820190508181036000830152611f0781611ecb565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000611f6a602283611466565b9150611f7582611f0e565b604082019050919050565b60006020820190508181036000830152611f9981611f5d565b9050919050565b6000611fab82611579565b9150611fb683611579565b925082821015611fc957611fc861188f565b5b828203905092915050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b600061200a601083611466565b915061201582611fd4565b602082019050919050565b6000602082019050818103600083015261203981611ffd565b905091905056fea2646970667358221220d1dae3a5fe0f4f2f1435d09abb142db43c4368513e4389a7b51fd7d107beb33964736f6c634300080f0033",
}

// KepengTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use KepengTokenMetaData.ABI instead.
var KepengTokenABI = KepengTokenMetaData.ABI

// KepengTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KepengTokenMetaData.Bin instead.
var KepengTokenBin = KepengTokenMetaData.Bin

// DeployKepengToken deploys a new Ethereum contract, binding an instance of KepengToken to it.
func DeployKepengToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KepengToken, error) {
	parsed, err := KepengTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KepengTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KepengToken{KepengTokenCaller: KepengTokenCaller{contract: contract}, KepengTokenTransactor: KepengTokenTransactor{contract: contract}, KepengTokenFilterer: KepengTokenFilterer{contract: contract}}, nil
}

// KepengToken is an auto generated Go binding around an Ethereum contract.
type KepengToken struct {
	KepengTokenCaller     // Read-only binding to the contract
	KepengTokenTransactor // Write-only binding to the contract
	KepengTokenFilterer   // Log filterer for contract events
}

// KepengTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type KepengTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KepengTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KepengTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KepengTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KepengTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KepengTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KepengTokenSession struct {
	Contract     *KepengToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KepengTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KepengTokenCallerSession struct {
	Contract *KepengTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// KepengTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KepengTokenTransactorSession struct {
	Contract     *KepengTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// KepengTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type KepengTokenRaw struct {
	Contract *KepengToken // Generic contract binding to access the raw methods on
}

// KepengTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KepengTokenCallerRaw struct {
	Contract *KepengTokenCaller // Generic read-only contract binding to access the raw methods on
}

// KepengTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KepengTokenTransactorRaw struct {
	Contract *KepengTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKepengToken creates a new instance of KepengToken, bound to a specific deployed contract.
func NewKepengToken(address common.Address, backend bind.ContractBackend) (*KepengToken, error) {
	contract, err := bindKepengToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KepengToken{KepengTokenCaller: KepengTokenCaller{contract: contract}, KepengTokenTransactor: KepengTokenTransactor{contract: contract}, KepengTokenFilterer: KepengTokenFilterer{contract: contract}}, nil
}

// NewKepengTokenCaller creates a new read-only instance of KepengToken, bound to a specific deployed contract.
func NewKepengTokenCaller(address common.Address, caller bind.ContractCaller) (*KepengTokenCaller, error) {
	contract, err := bindKepengToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KepengTokenCaller{contract: contract}, nil
}

// NewKepengTokenTransactor creates a new write-only instance of KepengToken, bound to a specific deployed contract.
func NewKepengTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*KepengTokenTransactor, error) {
	contract, err := bindKepengToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KepengTokenTransactor{contract: contract}, nil
}

// NewKepengTokenFilterer creates a new log filterer instance of KepengToken, bound to a specific deployed contract.
func NewKepengTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*KepengTokenFilterer, error) {
	contract, err := bindKepengToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KepengTokenFilterer{contract: contract}, nil
}

// bindKepengToken binds a generic wrapper to an already deployed contract.
func bindKepengToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KepengTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KepengToken *KepengTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KepengToken.Contract.KepengTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KepengToken *KepengTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KepengToken.Contract.KepengTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KepengToken *KepengTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KepengToken.Contract.KepengTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KepengToken *KepengTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KepengToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KepengToken *KepengTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KepengToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KepengToken *KepengTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KepengToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KepengToken *KepengTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KepengToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KepengToken *KepengTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KepengToken.Contract.Allowance(&_KepengToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KepengToken *KepengTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KepengToken.Contract.Allowance(&_KepengToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KepengToken *KepengTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KepengToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KepengToken *KepengTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KepengToken.Contract.BalanceOf(&_KepengToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KepengToken *KepengTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KepengToken.Contract.BalanceOf(&_KepengToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KepengToken *KepengTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KepengToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KepengToken *KepengTokenSession) Decimals() (uint8, error) {
	return _KepengToken.Contract.Decimals(&_KepengToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KepengToken *KepengTokenCallerSession) Decimals() (uint8, error) {
	return _KepengToken.Contract.Decimals(&_KepengToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KepengToken *KepengTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KepengToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KepengToken *KepengTokenSession) Name() (string, error) {
	return _KepengToken.Contract.Name(&_KepengToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KepengToken *KepengTokenCallerSession) Name() (string, error) {
	return _KepengToken.Contract.Name(&_KepengToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KepengToken *KepengTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KepengToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KepengToken *KepengTokenSession) Owner() (common.Address, error) {
	return _KepengToken.Contract.Owner(&_KepengToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KepengToken *KepengTokenCallerSession) Owner() (common.Address, error) {
	return _KepengToken.Contract.Owner(&_KepengToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KepengToken *KepengTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KepengToken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KepengToken *KepengTokenSession) Paused() (bool, error) {
	return _KepengToken.Contract.Paused(&_KepengToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KepengToken *KepengTokenCallerSession) Paused() (bool, error) {
	return _KepengToken.Contract.Paused(&_KepengToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KepengToken *KepengTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KepengToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KepengToken *KepengTokenSession) Symbol() (string, error) {
	return _KepengToken.Contract.Symbol(&_KepengToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KepengToken *KepengTokenCallerSession) Symbol() (string, error) {
	return _KepengToken.Contract.Symbol(&_KepengToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KepengToken *KepengTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KepengToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KepengToken *KepengTokenSession) TotalSupply() (*big.Int, error) {
	return _KepengToken.Contract.TotalSupply(&_KepengToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KepengToken *KepengTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _KepengToken.Contract.TotalSupply(&_KepengToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KepengToken *KepengTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KepengToken *KepengTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.Approve(&_KepengToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KepengToken *KepengTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.Approve(&_KepengToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_KepengToken *KepengTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_KepengToken *KepengTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.Burn(&_KepengToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_KepengToken *KepengTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.Burn(&_KepengToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_KepengToken *KepengTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_KepengToken *KepengTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.BurnFrom(&_KepengToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_KepengToken *KepengTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.BurnFrom(&_KepengToken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KepengToken *KepengTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KepengToken *KepengTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.DecreaseAllowance(&_KepengToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KepengToken *KepengTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.DecreaseAllowance(&_KepengToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KepengToken *KepengTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KepengToken *KepengTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.IncreaseAllowance(&_KepengToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KepengToken *KepengTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.IncreaseAllowance(&_KepengToken.TransactOpts, spender, addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KepengToken *KepengTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KepengToken *KepengTokenSession) Pause() (*types.Transaction, error) {
	return _KepengToken.Contract.Pause(&_KepengToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KepengToken *KepengTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _KepengToken.Contract.Pause(&_KepengToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KepengToken *KepengTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KepengToken *KepengTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _KepengToken.Contract.RenounceOwnership(&_KepengToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KepengToken *KepengTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KepengToken.Contract.RenounceOwnership(&_KepengToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KepengToken *KepengTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KepengToken *KepengTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.Transfer(&_KepengToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KepengToken *KepengTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.Transfer(&_KepengToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KepengToken *KepengTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KepengToken *KepengTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.TransferFrom(&_KepengToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KepengToken *KepengTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KepengToken.Contract.TransferFrom(&_KepengToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KepengToken *KepengTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KepengToken *KepengTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KepengToken.Contract.TransferOwnership(&_KepengToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KepengToken *KepengTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KepengToken.Contract.TransferOwnership(&_KepengToken.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KepengToken *KepengTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KepengToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KepengToken *KepengTokenSession) Unpause() (*types.Transaction, error) {
	return _KepengToken.Contract.Unpause(&_KepengToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KepengToken *KepengTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _KepengToken.Contract.Unpause(&_KepengToken.TransactOpts)
}

// KepengTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KepengToken contract.
type KepengTokenApprovalIterator struct {
	Event *KepengTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KepengTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KepengTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KepengTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KepengTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KepengTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KepengTokenApproval represents a Approval event raised by the KepengToken contract.
type KepengTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KepengToken *KepengTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KepengTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KepengToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KepengTokenApprovalIterator{contract: _KepengToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KepengToken *KepengTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KepengTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KepengToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KepengTokenApproval)
				if err := _KepengToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KepengToken *KepengTokenFilterer) ParseApproval(log types.Log) (*KepengTokenApproval, error) {
	event := new(KepengTokenApproval)
	if err := _KepengToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KepengTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KepengToken contract.
type KepengTokenOwnershipTransferredIterator struct {
	Event *KepengTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KepengTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KepengTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KepengTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KepengTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KepengTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KepengTokenOwnershipTransferred represents a OwnershipTransferred event raised by the KepengToken contract.
type KepengTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KepengToken *KepengTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KepengTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KepengToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KepengTokenOwnershipTransferredIterator{contract: _KepengToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KepengToken *KepengTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KepengTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KepengToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KepengTokenOwnershipTransferred)
				if err := _KepengToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KepengToken *KepengTokenFilterer) ParseOwnershipTransferred(log types.Log) (*KepengTokenOwnershipTransferred, error) {
	event := new(KepengTokenOwnershipTransferred)
	if err := _KepengToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KepengTokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KepengToken contract.
type KepengTokenPausedIterator struct {
	Event *KepengTokenPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KepengTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KepengTokenPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KepengTokenPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KepengTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KepengTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KepengTokenPaused represents a Paused event raised by the KepengToken contract.
type KepengTokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KepengToken *KepengTokenFilterer) FilterPaused(opts *bind.FilterOpts) (*KepengTokenPausedIterator, error) {

	logs, sub, err := _KepengToken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KepengTokenPausedIterator{contract: _KepengToken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KepengToken *KepengTokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KepengTokenPaused) (event.Subscription, error) {

	logs, sub, err := _KepengToken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KepengTokenPaused)
				if err := _KepengToken.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KepengToken *KepengTokenFilterer) ParsePaused(log types.Log) (*KepengTokenPaused, error) {
	event := new(KepengTokenPaused)
	if err := _KepengToken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KepengTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KepengToken contract.
type KepengTokenTransferIterator struct {
	Event *KepengTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KepengTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KepengTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KepengTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KepengTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KepengTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KepengTokenTransfer represents a Transfer event raised by the KepengToken contract.
type KepengTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KepengToken *KepengTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KepengTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KepengToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KepengTokenTransferIterator{contract: _KepengToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KepengToken *KepengTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KepengTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KepengToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KepengTokenTransfer)
				if err := _KepengToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KepengToken *KepengTokenFilterer) ParseTransfer(log types.Log) (*KepengTokenTransfer, error) {
	event := new(KepengTokenTransfer)
	if err := _KepengToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KepengTokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KepengToken contract.
type KepengTokenUnpausedIterator struct {
	Event *KepengTokenUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KepengTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KepengTokenUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KepengTokenUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KepengTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KepengTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KepengTokenUnpaused represents a Unpaused event raised by the KepengToken contract.
type KepengTokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KepengToken *KepengTokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KepengTokenUnpausedIterator, error) {

	logs, sub, err := _KepengToken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KepengTokenUnpausedIterator{contract: _KepengToken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KepengToken *KepengTokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KepengTokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _KepengToken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KepengTokenUnpaused)
				if err := _KepengToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KepengToken *KepengTokenFilterer) ParseUnpaused(log types.Log) (*KepengTokenUnpaused, error) {
	event := new(KepengTokenUnpaused)
	if err := _KepengToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
