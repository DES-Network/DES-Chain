#!/bin/bash
set -u
set -e

BASE_PATH=$(pwd)
./constellation-start.sh 2

echo "[*] Starting Ethereum node"
set -v
ARGS="--raft --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints"
IPC=$BASE_PATH/qdata/dd2/geth.ipc PRIVATE_CONFIG=qdata/c2/tm.ipc nohup geth --datadir qdata/dd2 $ARGS --permissioned --raftport 50402 --rpcport 22001 --port 21001 2>>qdata/logs/2.log &
set +v

echo
echo "A second node has been successfully setup!"
echo "If you want to see the logs, you can navigate to 'qdata/logs/2.log', or you can run 'geth attach qdata/dd2/geth.ipc' to attach to interact with the Geth console."