#!/bin/bash
set -u
set -e

BASE_PATH=$(pwd)
./constellation-start.sh

echo "[*] Starting Ethereum nodes"
set -v
ARGS="--raft --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints"
IPC=$BASE_PATH/qdata/dd1/geth.ipc PRIVATE_CONFIG=qdata/c1/tm.ipc nohup ./bin/geth --datadir qdata/dd1 $ARGS --raftport 50401 --rpcport 22000 --port 21000 --debug --unlock 0 --password passwords.txt 2>>qdata/logs/1.log &
IPC=$BASE_PATH/qdata/dd2/geth.ipc PRIVATE_CONFIG=qdata/c2/tm.ipc nohup ./bin/geth --datadir qdata/dd2 $ARGS --raftport 50402 --rpcport 22001 --port 21001 2>>qdata/logs/2.log &
IPC=$BASE_PATH/qdata/dd3/geth.ipc PRIVATE_CONFIG=qdata/c3/tm.ipc nohup ./bin/geth --debug --verbosity 5 --datadir qdata/dd3 $ARGS --raftport 50403 --rpcport 22002 --port 21002 2>>qdata/logs/3.log &
IPC=$BASE_PATH/qdata/dd4/geth.ipc PRIVATE_CONFIG=qdata/c4/tm.ipc nohup ./bin/geth --datadir qdata/dd4 $ARGS --raftport 50404 --rpcport 22003 --port 21003 2>>qdata/logs/4.log &
IPC=$BASE_PATH/qdata/dd5/geth.ipc PRIVATE_CONFIG=qdata/c5/tm.ipc nohup ./bin/geth --datadir qdata/dd5 $ARGS --raftport 50405 --rpcport 22004 --port 21004 2>>qdata/logs/5.log &
set +v

echo
echo "Great! This means that a test DES network has been successfully setup, with five nodes, where Node 1 is the only Regulator!"
echo "If you want to see the logs, you can navigate to 'qdata/logs', or you can run e.g. 'geth attach qdata/dd1/geth.ipc' to interact with a Geth console."
echo "To deploy the DES network contract that will keep track of regulators and whitelisted nodes, run './runscript.sh deploy-and-init.js'."
echo "Note that you will not be able to make any private transactions or add nodes to the whitelist until this is done."
