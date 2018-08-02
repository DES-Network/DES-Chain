# Test DES Network

This folder contains scripts to help configure a test DES network and allow you to evaluate the functionality with easy to use, verbose examples. This test network contains 5 nodes, out of which `Node 1` is the Regulator. 

Initially, the network permissioning works like Quorum, where every node that is listed in the `permissioned-nodes.js` file, gets to participate in the network. However, after the deployment of the DES contract, which keeps track of permissioned nodes and the regulators in the network, the permissioning will be handled by the DES contract.

## Requirements and Initial Setup

Since the binaries required for running these scripts are inside the bin folder, you won't need to build anything from scratch. 

However, this has only been tested on a Linux system (Ubuntu 16.04), so you might want to make sure you try these scripts out on it. Otherwise, you will have to provide `geth` and `constellation-node` binaries for your platform and place them into the `bin` folder to effectively use these scripts. 

In order to get these scripts, you will have to use `git`, so if you don’t have it installed already, please use `apt-get install git`. You shouldn’t need to install anything else, since all the required binaries are inside the `bin` directory. 

  1. Clone this repository in your workspace using this command:
  `git clone https://github.com/DES-Network/des-quorum.git`
  This will create a `des-quorum` directory in your workspace, which contains the source code, as well as the test scripts.
  2. If you want to setup or use the test network, please navigate to the `test-network` subdirectory using this command: `cd des-quorum/test-network`. Note that all the commands shown in this document assume that you are in this directory.

## Network Setup

Given below is a description of the scripts, what they do and the flow of how you should use these scripts, in what order, to get a new test network up and running.

  1. `init.sh`: Cleanup previous directories (delete any previous chain data) and (re)initialize accounts and keystores for configuring a new network and its nodes. This will set up the environment required to launch the network from scratch. 

  Please note that the keys and configuration used here are all pre-generated. If you want to use your own, please see the `raft-setup.sh` script found (here)[https://github.com/Szkered/quorum-raft-cluster].

  *If you don't want to setup a new network, and just want to restart from a previous point, you can skip this step.*

  2. `start-des-network.sh`: Launch `constellation` and `geth` node for the network owner. This will launch the DES network with a single node, which is the owner of the network. You can stop all the nodes on the network using the `stop.sh` script.

  3. `deploy-and-init.js`: Deploy the DES network contract that will help add/remove and keep track of the regulators in the network and the nodes in the network. Since only permissioned nodes are allowed in the DES network, the owner will have to add any new nodes to this contract to ensure their participation in the network. 

  Note that this is written in javascript and requires attachment to a geth node for execution. We have created a script that will make this easy, so you can run this as `./runscript.sh deploy-and-init.js`.

This means that your DES network is up and running, and since the contract is also up, this ensures that DES permissioning and transaction restrictions are activated. 

**NOTE: All logs and temporary data will be written to the `qdata` folder.**

## Play around with DES
  The following are some scripts you can use to play around with the DES network.

  - `test-des-contract.js`: Once the aforementioned contract has been deployed, you can run this test to confirm that the contract has been successfully deployed and to confirm if `Node 1`, also known as the Owner, has been added as a Regulator and a Permissioned Node.
  - `wrong-private-contract.js`: You can now try to deploy a private contract. However, this won't work, since the `privateFor` array only contains `"ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="`, which is not a regulator's public key.
  - `private-contract.js`: This is the correct private contract deployment script. The contract creation will be allowed as the regulator's key `BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=` is now party to the transaction.
  - `whitelist.js`: This helps you whitelist or add a given enode address, so that it has permission to participate in the network.
  - `blacklist.js`: This helps you blacklist or remove a given enode address, so that it no longer has permission to participate in the network.
  - `runscript.sh`: This will allow you to execute a JS script with the owner's address.
  - `stop.sh`: Stop all `constellation` and `geth` nodes.


You may also use separate terminal windows to run the geth consoles for any node in the network.

In each terminal, ensure you are in the test-network dir before running the below.

* If you aren't already running the test-network example, in terminal 1 run 
``$ ./init.sh `` followed by ``$ ./start-des-network.sh ``
* In a new terminal navigate to the `des-quorum/test-net` directory and run the following: ``$ ./bin/geth attach qdata/dd1/geth.ipc``, where `qdata/dd1/geth.ipc` can be replaced by the respective data directory of the node you want to attach to, like `qdata/dd2/geth.ipc`, for connecting to Node 2, etc.


## Private Transactions

There are two private contract scripts available. The `wrong-private-contract.js` is the one that doesn't contain the Regulator key in either the `privateFrom` or the `privateTo` fields, and is, therefore, rejected. You can try to run it as `./runscript.sh wrong-private-contract.js`, which will yield the following result:

```
err creating contract Error: invalid transaction because no regulator provided

```

When you use `private-contract.js`, though, it works out without a hitch. Once it has been sent to mine. Note that this transaction is between `Node 1` and `Node 2`. The rest of the nodes shouldn't be privy to it. 

Please run `./runscript.sh private-contract.js` and copy the transaction hash it prints out. Then attach to geth console for `Node 1`, `Node 2` and `Node 3` in different terminals using `./bin/geth attach path/to/geth.ipc`.

First, we need to get the contract's address, so we'll use the following command using the transaction hash for the contract submission:

```
> eth.getTransactionReceipt("0x6f6...")

```
Note that `0x6f6...` is where you paste the transaction hash. This will return something similar to the following result:

```
{
  blockHash: "0xd11326c5501af22a9228742928b3c5ce71108432128d6b014fc398b74e0ef94d",
  blockNumber: 8,
  contractAddress: "0x4d3bfd7821e237ffe84209d8e638f9f309865b87",
  cumulativeGasUsed: 0,
  from: "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
  gasUsed: 0,
  logs: [],
  logsBloom: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  status: "0x1",
  to: null,
  transactionHash: "0x6f69adef793d32249ab0ae0f8202d50f87c5e0c526999f84ced4863aae2998e4",
  transactionIndex: 0
}

```

This can be done on either terminal. Now, you want to copy the `contractAddress`, so we can interact with the contract. You will name the `contractAddress` as `var address` and execute the following in all terminals:

```
var address = "0x4d3bfd7821e237ffe84209d8e638f9f309865b87";
var abi = [{"constant":true,"inputs":[],"name":"storedData","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"uint256"}],"name":"set","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"initVal","type":"uint256"}],"payable":false,"type":"constructor"}];
var private = eth.contract(abi).at(address)
```

* In terminal window 1 (node 1)
```
> private.get()
42
```
* In terminal window 2 (node 2)
```
> private.get()
42
``` 
* In terminal window 3 (node 3) - not privy to the contract
```
> private.get()
0
```


## Permissions

Node Permissioning is a feature in Quorum that allows only a pre-defined set of nodes (as identified by their remotekey/enodes) to connect to the permissioned network.

This example demonstrates the following:
* Sets up a network with a combination of permissioned and non-permissioned nodes in the cluster.
* Details of permissioned-nodes.json file.
* Demonstrate that only the nodes that are specified in permissioned-nodes.json can connect to the network.

### Verify only permissioned nodes are connected to the network.

Attach to the individual nodes via
  `./bin/geth attach path/to/geth.ipc` and use `admin.peers` to check the connected nodes.

```
❯ ./geth attach qdata/dd1/geth.ipc
Welcome to the Geth JavaScript console!

instance: Geth/v1.7.2-stable/darwin-amd64/go1.9.2
coinbase: 0xed9d02e382b34818e88b88a309c7fe71e65f419d
at block: 1 (Mon, 29 Oct 47909665359 22:09:51 EST)
 datadir: /home/xyz/workspace/des-quorum/test-network/qdata/dd1
 modules: admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 raft:1.0 rpc:1.0 txpool:1.0 web3:1.0

> admin.peers
[{
    caps: ["eth/63"],
    id: "0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416",
    name: "Geth/v1.7.2-stable/darwin-amd64/go1.9.2",
    network: {
      localAddress: "127.0.0.1:65188",
      remoteAddress: "127.0.0.1:21001"
    },
    protocols: {
      eth: {
        difficulty: 0,
        head: "0xc23b4ebccc79e2636d66939924d46e618269ca1beac5cf1ec83cc862b88b1b71",
        version: 63
      }
    }
},
...
]
```

You can also inspect the log files under `qdata/logs/*.log` for further diagnostics messages around incoming / outgoing connection requests. Grep for `ALLOWED-BY` or `DENIED-BY`. Please be sure to enable verobsity for p2p module, which is done by adding the option `--vmodule p2p=5` when you run geth, as you can see in the `start-des-network` script.

### Permissioning

Permissioning is granted based on the remote key (enode address) of the geth node. The pre-generated remote keys are specified in the permissioned-nodes.json and are placed under individual nodes <datadir>.

Once the DES contract is deployed, enabling and disabling of permissions is handled by the DES contract. See the  `whitelist.js` script to see how to interact with the contract. You will have to run it as: `./runscript.sh whitelist.js` to include the mentioned enable the node's permissions. To remove a node from the whitelist, use `./runscript.sh blacklist.js`.

### Testing Permissions

To test whether this permissioning mechanism is working, start the network from scratch and get the geth console for `Node 1` using `geth attach qdata/dd1/geth.ipc`.

Once you have the console, check for your self how many nodes it has initially, via `admin.peers`

Now exit this console using `exit`.

Once you are back in the `test-network` directory, deploy the DES contract via `./runscript.sh deploy-and-init.js`, which deploys the DES contract and initializes it so that the permissioned nodes are only the first four nodes. Give the contract time to be deployed and around a minute or so for the permission update check, then reattach to `Node 1`'s geth and try out the `admin.peers` command again. It should have removed `Node 5` as a peer, so there will be only three peers remaining for the permissioned nodes, aka `Node 1`, `Node 2`, `Node 3` and `Node 4`. 

```
> admin.peers
[{
    caps: ["eth/62", "eth/63"],
    id: "0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416",
    name: "Geth/v1.7.2-stable-153cb1ea/linux-amd64/go1.10.1",
    network: {
      localAddress: "127.0.0.1:46146",
      remoteAddress: "127.0.0.1:21001"
    },
    protocols: {
      eth: {
        difficulty: 262144,
        head: "0x1caf2a909961a378d6a824f93e8af0716d4979a0e1729f8b3a0af9f1ccfa0d33",
        version: 63
      }
    }
}
...
]


```

And if you were to attach to geth console for `Node 5` using `geth attach qdata/dd5/geth.ipc`, you would find that it no longer has any peers, since it would now be unable to connect to any node in the cluster:

```
> admin.peers
[]
```
Exit this console and run `./runscript.sh whitelist.sh`. This will make `Node 5` permissioned. Wait a minute or so and then check again for `admin.peers` on `Node 5`. You will find that it now has peers again.