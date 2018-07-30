#!/bin/bash
set -u
set -e

BASE_PATH=$(pwd)
./constellation-start.sh 1

echo "[*] Starting Ethereum node"
set -v
ARGS="--raft --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints"
IPC=$BASE_PATH/qdata/dd1/geth.ipc PRIVATE_CONFIG=qdata/c1/tm.ipc nohup ./bin/geth --datadir qdata/dd1 $ARGS --permissioned --raftport 50401 --rpcport 22000 --port 21000 --debug --unlock 0 --password passwords.txt 2>>qdata/logs/1.log &
set +v

echo
echo "Great! This means that the first node in the DES network has been successfully setup!"
echo "If you want to see the logs, you can navigate to 'qdata/logs/1.log', or you can run 'geth attach qdata/dd1/geth.ipc' to attach to interact with the Geth console."
echo "To deploy the DES network contract that will keep track of regulators and whitelisted nodes, run './runscript.sh deploy-and-init.js'."
echo "Note that you will not be able to make any private transactions or add nodes to the whitelist until this is done."
