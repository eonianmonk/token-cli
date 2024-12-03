solidity_dir:=./solidity
npm_prefix:=--prefix $(solidity_dir)
testnet:=holesky

# .PHONY: test
test:
	npm run $(npm_prefix) test

deploy-token:
	npm run $(npm_prefix) deploy-token -- --network $(testnet)

build:
	npm install $(npm_prefix)
	npm run $(npm_prefix) build 