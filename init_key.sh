#!/bin/sh
KEYRING="file"
KEYALGO="eth_secp256k1"

KEY=$1
PASSWORD=$2

# check arguments
if [ "$#" -ne 2 ]; then
    echo "Error. Wrong number of arguments"
    echo "Using: $0 [KEY] [PASSWORD]"
    exit 1
fi

# if $KEY exists it should be deleted
yes $PASSWORD | entangled keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO 


