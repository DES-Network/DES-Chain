#!/bin/bash
set -u
set -e

echo "[*] Cleaning up temporary data directories"
rm -rf qdata
mkdir -p qdata/logs

echo "[*] Configuring node 1"
mkdir -p qdata/dd1/{keystore,geth}
cp permissioned-nodes.json qdata/dd1/static-nodes.json
cp permissioned-nodes.json qdata/dd1/permissioned-nodes.json
cp keys/key1 qdata/dd1/keystore
cp raft/nodekey1 qdata/dd1/geth/nodekey
geth --datadir qdata/dd1 init genesis.json

echo "[*] Configuring node 2"
mkdir -p qdata/dd2/{keystore,geth}
cp permissioned-nodes.json qdata/dd2/static-nodes.json
cp permissioned-nodes.json qdata/dd2/permissioned-nodes.json
cp keys/key2 qdata/dd2/keystore
cp raft/nodekey2 qdata/dd2/geth/nodekey
geth --datadir qdata/dd2 init genesis.json

echo "[*] Configuring node 3"
mkdir -p qdata/dd3/{keystore,geth}
cp permissioned-nodes.json qdata/dd3/static-nodes.json
cp permissioned-nodes.json qdata/dd3/permissioned-nodes.json
cp keys/key3 qdata/dd3/keystore
cp raft/nodekey3 qdata/dd3/geth/nodekey
geth --datadir qdata/dd3 init genesis.json

echo "[*] Configuring node 4"
mkdir -p qdata/dd4/{keystore,geth}
cp permissioned-nodes.json qdata/dd4/static-nodes.json
cp permissioned-nodes.json qdata/dd4/permissioned-nodes.json
cp keys/key4 qdata/dd4/keystore
cp raft/nodekey4 qdata/dd4/geth/nodekey
geth --datadir qdata/dd4 init genesis.json

echo "[*] Configuring node 5"
mkdir -p qdata/dd5/{keystore,geth}
cp permissioned-nodes.json qdata/dd5/static-nodes.json
cp permissioned-nodes.json qdata/dd5/permissioned-nodes.json
cp keys/key5 qdata/dd5/keystore
cp raft/nodekey5 qdata/dd5/geth/nodekey
geth --datadir qdata/dd5 init genesis.json

echo "Cleaned up previous data, which includes configuration files and blockchain data, and created new directories and configuration files required for geth and constellation nodes."
echo "Now, when you start the network, it will start one from scratch. Use the './start-des-network.sh' script for this purpose"