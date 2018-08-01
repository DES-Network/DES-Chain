# Test DES Network

This folder contains scripts to help configure a test DES network and allow you to evaluate the functionality with easy to use, verbose examples. This test network contains 3 nodes, namely `Node 1`, `Node 2` and `Node 3`. `Node 1` is the a Regulator. 

Initially, the network permissioning works like Quorum, where every node that is listed in the `permissioned-nodes.js` file, gets to participate in the network. However, after the deployment of the DES contract, which keeps track of permissioned nodes and the regulators in the network, the permissioning will be handled by the DES contract.

## Requirements and Initial Setup

Since the binaries required for running these scripts are inside the bin folder, you won't need to build anything from scratch. 

However, this has only been tested on a Linux system (Ubuntu 16.04), so you might want to make sure you try these scripts out on it. Otherwise, you will have to provide `geth` and `constellation-node` binaries for your platform and place them into the `bin` folder to effectively use these scripts. 

In order to get these scripts, you will have to use `git`, so if you don’t have it installed already, please use `apt-get install git`. You shouldn’t need to install anything else, since all the required binaries are inside the `bin` directory. **(Needs confirmation)**

  1. Clone this repository in your workspace using this command:
  `git clone https://github.com/DES-Network/quorum.git`
  This will create a `quorum` directory in your workspace, which contains the source code, as well as the test scripts.
  2. If you want to setup or use the test network, please navigate to the `test-network` subdirectory using this command: `cd quorum/test-network`. Note that all the commands shown in this document assume that you are in this directory.

## Network Setup

Given below is a description of the scripts, what they do and the flow of how you should use these scripts, in what order, to get a new test network up and running.

  1. `init.sh`: Cleanup previous directories (delete any previous chain data) and (re)initialize accounts and keystores for configuring a new network and its nodes. This will set up the environment required to launch the network from scratch. 

  Please note that the keys and configuration used here are all pre-generated. If you want to use your own, please see the `raft-setup.sh` script found (here)[https://github.com/Szkered/quorum-raft-cluster].

  *If you don't want to setup a new network, and just want to restart from a previous point, you can skip this step.*

  2. `start-des-network.sh`: Launch `constellation` and `geth` node for the network owner. This will launch the DES network with a single node, which is the owner of the network. You can stop all the nodes on the network using the `stop.sh` script.

  3. `deploy-and-init.js`: Deploy the DES network contract that will help add/remove and keep track of the regulators in the network and the nodes in the network. Since only permissioned nodes are allowed in the DES network, the owner will have to add any new nodes to this contract to ensure their participation in the network. 

Note that this is written in javascript and requires attachment to a geth node for execution. We have created a script that will make this easy, so you can run this as `./run-as-owner.sh deploy-and-init.js`.

  This means that your DES network is up and running, and since the contract is also up, this ensures that DES permissioning and transaction restrictions are activated. 

  **NOTE: All logs and temporary data will be written to the `qdata` folder.**

## Play around with DES
  The following are some scripts you can use to play around with the DES network.

  - `test-des-contract.js`: Once the aforementioned contract has been deployed, you can run this test to confirm that the contract has been successfully deployed and to confirm if `Node 1`, also known as the Owner, has been added as a Regulator and a Permissioned Node.
  - `wrong-private-contract.js`: You can now try to deploy a private contract. However, this won't work, since the `privateFor` array only contains `"ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="`, which is not a regulator's public key.
  - `private-contract.js`: This is the correct private contract deployment script. The contract creation will be allowed as the regulator's key `BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=` is now party to the transaction.
  - `whitelist.js`: This helps you whitelist or add a given enode address, so that it has permission to participate in the network.
  - `blacklist.js`: This helps you blacklist or remove a given enode address, so that it no longer has permission to participate in the network.
  - `stop.sh`: Stop all `constellation` and `geth` nodes.


You may also use separate terminal windows to run the geth consoles for any node in the network.

In each terminal, ensure you are in the test-network dir before running the below.

* If you aren't already running the test-network example, in terminal 1 run 
``$ ./init.sh `` followed by ``$ ./start-des-network.sh ``
* In terminal 1 run ``$ geth attach ipc:qdata/dd1/geth.ipc``
* In terminal 2 run ``$ geth attach ipc:qdata/dd2/geth.ipc``
* In terminal 3 run ``$ geth attach ipc:qdata/dd3/geth.ipc``

## Permissions

Node Permissioning is a feature in Quorum that allows only a pre-defined set of nodes (as identified by their remotekey/enodes) to connect to the permissioned network.

This example demonstrates the following:
* Sets up a network with a combination of permissioned and non-permissioned nodes in the cluster.
* Details of permissioned-nodes.json file.
* Demonstrate that only the nodes that are specified in permissioned-nodes.json can connect to the network.

### Verify only permissioned nodes are connected to the network.

Attach to the individual nodes via
  `geth attach path/to/geth.ipc` and use `admin.peers` to check the connected nodes.

```
❯ geth attach qdata/dd1/geth.ipc
Welcome to the Geth JavaScript console!

instance: Geth/v1.7.2-stable/darwin-amd64/go1.9.2
coinbase: 0xed9d02e382b34818e88b88a309c7fe71e65f419d
at block: 1 (Mon, 29 Oct 47909665359 22:09:51 EST)
 datadir: /Users/joel/jpm/quorum-examples/examples/7nodes/qdata/dd1
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

Once the DES contract is deployed, enabling and disabling of permissions is handled by the DES contract. See the  `whitelist.js` script to see how to interact with the contract. You will have to run it as: `./run-as-owner.sh whitelist.js` to include the mentioned enable the node's permissions. To remove a node from the whitelist, use `./run-as-owner.sh blacklist.js`.

### Testing Permissions

To test whether this permissioning mechanism is working, start the network from scratch and get the geth console for `Node 1` using `geth attach qdata/dd1/geth.ipc`.

Once you have the console, check for your self how many nodes it has initially, via `admin.peers`

Now exit this console using `exit`.

Once you are back in the `test-network` directory, deploy the DES contract via `./run-as-owner.sh deploy-and-init.js`, which deploys the DES contract and initializes it so that the permissioned nodes are only `Node 1` and `Node 2`. Give the contract time to be deployed and around a minute or so for the permission update check, then reattach to `Node 1`'s geth and try out the `admin.peers` again. It should have removed `Node 3` as a peer, which had the enode id `579f786d4e2830bbcc02815a27e8a9bacccc9605df4dc6f20bcc1a6eb391e7225fff7cb83e5b4ecd1f3a94d8b733803f2f66b7e871961e7b029e22c155c3a778`. 

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
}]


```

And if you were to attach to geth console for `Node 3` using `geth attach qdata/dd3/geth.ipc`, you would find that it no longer has any peers, since it would now be unable to connect to either `Node 1` or `Node 2`:

```
> admin.peers
[]
```

