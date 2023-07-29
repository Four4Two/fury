#! /bin/bash
set -e

DATADIR=./contrib/devnet/geth/data
GENESIS=./contrib/devnet/geth/genesis.json

# Delete state
rm -rf $DATADIR/geth

# To create the account, if more are needed
# geth account new --datadir $DATADIR --password ./contrib/devnet/eth-password

geth init --datadir $DATADIR $GENESIS

geth --datadir $DATADIR \
     --unlock 9B277763809483FA65F7177BC8B3AF1CBB8AEE46 \
     --password ./contrib/devnet/eth-password \
     --mine \
     --allow-insecure-unlock \
     --http \
     --http.addr 0.0.0.0 \
     --http.corsdomain '*' \
     --http.port 8555 \
     --http.api eth,net,rpc,web3,debug \
     --ws.addr 0.0.0.0 \
     --ws.port 8556 \
     --ws.origins '*' &

# Deploy contracts after geth started
sleep 5 \
    && cd ./contract \
    && npx hardhat run scripts/init_dev_env.ts --network localhost &

wait
