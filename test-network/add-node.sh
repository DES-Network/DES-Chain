#!/bin/bash
set -u
set -e

echo "[*]  To add a node to the network, an existing node in the network has to add the peer"
echo "PRIVATE_CONFIG=qdata/c1/tm.ipc ./bin/geth --exec 'raft.addPeer('enode://0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416@127.0.0.1:21001?discport=0&raftport=50402')' attach ipc:qdata/dd1/geth.ipc"

# RAFT_ID=$(PRIVATE_CONFIG=qdata/c1/tm.ipc ./bin/geth --exec "raft.addPeer('enode://0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416@127.0.0.1:21001?discport=0&raftport=50402')" attach ipc:qdata/dd1/geth.ipc)
# echo $RAFT_ID

RAFT_ID=2
BASE_PATH=$(pwd)
./constellation-start.sh $RAFT_ID

echo "[*] Starting Ethereum node"
set -v
ARGS="--raft --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum --emitcheckpoints"
IPC=$BASE_PATH/qdata/dd$RAFT_ID/geth.ipc PRIVATE_CONFIG=qdata/c$RAFT_ID/tm.ipc nohup geth --datadir qdata/dd2 $ARGS --rpccorsdomain "*" --permissioned --debug --vmodule eth/*=5,p2p/*=5  --raftport 50402 --rpcport 22001 --port 21001 2>>qdata/logs/$RAFT_ID.log &
set +v

echo
echo "A second node has been successfully setup!"
echo "But wait... even though Node 2 is running now, it can't communicate with the owner's node, because it hasn't been added to the whitelist."
echo

echo "This will have to be done using the DES contract. Run './run-as-owner.sh whitelist.js'"