solidity_dir:=./solidity
cli_dir:=./token-cli
npm_prefix:=--prefix $(solidity_dir)
testnet:=holesky

install:
	npm install $(npm_prefix)

# build artifacts, abi for go
build:
	npm run $(npm_prefix) build 
	jq -c '.abi' $(solidity_dir)/artifacts/contracts/TestToken.sol/TestToken.json | \
		abigen --abi=- --pkg=token_cli --type TestToken --out=$(cli_dir)/test_token.go

build-docker:
	docker build -t token-cli:latest $(cli_dir)

test:
	npm run $(npm_prefix) test

deploy-token:
	npm run $(npm_prefix) deploy-token -- --network $(testnet)