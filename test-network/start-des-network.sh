#!/bin/bash
set -u
set -e

BASE_PATH=$(pwd)
./constellation-start.sh

echo "[*] Starting Ethereum node"
set -v
ARGS="--raft --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints"
IPC=$BASE_PATH/qdata/dd1/geth.ipc PRIVATE_CONFIG=qdata/c1/tm.ipc nohup ./bin/geth --vmodule p2p=5,private=5 --datadir qdata/dd1 $ARGS --permissioned --raftport 50401 --rpcport 22000 --port 21000 --debug --unlock 0 --password passwords.txt 2>>qdata/logs/1.log &
IPC=$BASE_PATH/qdata/dd2/geth.ipc PRIVATE_CONFIG=qdata/c2/tm.ipc nohup ./bin/geth --datadir qdata/dd2 $ARGS --permissioned --raftport 50402 --rpcport 22001 --port 21001 2>>qdata/logs/2.log &
IPC=$BASE_PATH/qdata/dd3/geth.ipc PRIVATE_CONFIG=qdata/c3/tm.ipc nohup ./bin/geth --vmodule p2p=5,private=5 --datadir qdata/dd3 $ARGS --permissioned --raftport 50403 --rpcport 22002 --port 21002 2>>qdata/logs/3.log &
set +v

echo
echo "Great! This means that the a test DES network has been successfully setup, with three nodes, where Node 1 is the only Regulator!"
echo "If you want to see the logs, you can navigate to 'qdata/logs', or you can run e.g. 'geth attach qdata/dd1/geth.ipc' to attach to interact with a Geth console."
echo "To deploy the DES network contract that will keep track of regulators and whitelisted nodes, run './run-as-owner.sh deploy-and-init.js'."
echo "Note that you will not be able to make any private transactions or add nodes to the whitelist until this is done."
