#!/bin/sh
CHAINID="entangle_33133-1"
MONIKER="validator"
LOGLEVEL="info"

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# Set moniker and chain-id for Entangle (Moniker can be anything, chain-id must be an integer)
entangled init $MONIKER --chain-id $CHAINID

cp -f config/genesis.json $HOME/.entangled/config/
cp -f config/config.toml $HOME/.entangled/config/ 

entangled start --pruning=default --log_level $LOGLEVEL --minimum-gas-prices=0.0001aNGL --json-rpc.api eth,net,web3,miner
