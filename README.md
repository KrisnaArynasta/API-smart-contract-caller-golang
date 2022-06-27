solcjs --bin --abi contract/KepengToken.sol -o build
*contract is folder source. 
*build is the folder where the abi and bin will be created.

abigen --bin=build/contract_KepengToken_sol_KEPENG.bin --abi=build/contract_KepengToken_sol_KEPENG.abi -pkg=KepengToken --out=gen/KepengToken.go
*build is folder source. 
*gen is folder where the .go file will be created

