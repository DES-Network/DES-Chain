# Test DES Network

These scripts help configure a test DES network and allow you to evaluate the functionality with easy to use, verbose examples.

Given below is a description of the scripts, what they do and the flow of how you should use these scripts, in what order, to get a test network up and running.

  1. `1-init.sh`: Cleanup previous directories (delete any previous chain data) and (re)initialize accounts and keystores for configuring a new network and its nodes. This will set up the environment required to launch the network.
  2. `2-start-des-network.sh`: Launch `constellation` and `geth` node for the network owner. This will launch the DES network with a single node, which is the owner of the network.
  3. `deploy-and-init.js`: Deploy the DES network contract that will help add/remove and keep track of the regulators in the network and the nodes in the network. Since only permissioned nodes are allowed in the DES network, the owner will have to add any new nodes to this contract to ensure their participation in the network. 

  Note that this is written in javascript and requires attachment to a geth node for execution. We have created a script that will make this easy, so you can run this as `./run-as-owner.sh deploy-and-init.js`.

  - `test-des-contract.js`: Once the aforementioned contract has been deployed, you can run this test to confirm that the contract has been successfully deployed and to confirm if `Node 1`, also known as the Owner, has been added as a Regulator and a Permissioned Node.
  - `wrong-private-contract.js`: You can now try to deploy a private contract. However, this won't work, since the `privateFor` array only contains `"ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="`, which is not a regulator's public key.
  - `private-contract.js`: This is the correct private contract deployment script. The contract creation will be allowed as the regulator's key `BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=` is now party to the transaction.
  - `stop.sh`: Stop all `constellation` and `geth` nodes.

All logs and temporary data will be written to the `qdata` folder.

## Testing Privacy
You can run the 7node example to test the privacy features of Quorum. As described in the Quick Start section of the [README](https://github.com/jpmorganchase/quorum), the final step of the 7node `start.sh` script was the sending of a private transaction to generate a (private) smart contract (SimpleStorage) sent from node 1 "for" node 7 (denoted by the public key passed via privateFor: ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="] in the sendTransaction call). We'll begin by demonstrating only nodes 1 and 7 are able to view the initial state of the contract. Next we have node 1 update the state of this contract and again verify only nodes 1 and 7 are able to see the updated state of the contract after the block containing the update transaction is validated by the network.

For this test it is recommended to use separate terminal windows running geth JavaScript console attached to node 1, node 7, and any node 2-6 (in our example we'll choose node 4).

In each terminal, ensure you are in the 7nodes dir before running the below.

* If you aren't already running the 7nodes example, in terminal 1 run ``$ ./init.sh `` followed by ``$ ./start.sh ``
* In terminal 1 run ``$ geth attach ipc:qdata/dd1/geth.ipc``
* In terminal 2 run ``$ geth attach ipc:qdata/dd4/geth.ipc``
* In terminal 3 run ``$ geth attach ipc:qdata/dd7/geth.ipc``

For each of the 3 nodes we'll use the built-in JavaScript by setting the variable ```address``` assigned to the simpleStorage contract address created by the node 1. The address can be found in the node 1's log file 7nodes/qdata/logs/1.log, or alternatively by reading the `contractAddress` param after calling `eth.getTransactionReceipt(txHash)` ([Ethereum API documentation](https://github.com/ethereum/wiki/wiki/JavaScript-API#web3ethgettransactionreceipt)), passing in the transaction hash that was logged in the console that you ran the 7nodes example in. Replace the address below with the address value found in the 1.log file:
```
> var address = "0x1932c48b2bf8102ba33b4a6b545c32236e342f34";
```
Next we'll use ```eth.contract``` to define a contract class with the simpleStorage ABI definition as follows:
```
> var abi = [{"constant":true,"inputs":[],"name":"storedData","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"uint256"}],"name":"set","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"initVal","type":"uint256"}],"type":"constructor"}];
> var private = eth.contract(abi).at(address)
```
The function calls are now available on the contract instance and you can call those methods on the contract. Let's start by examining the initial value of the contract to make sure that only nodes 1 and 7 can see the initialized value.
* In terminal window 1 (node 1)
```
> private.get()
42
```
* In terminal window 2 (node 4)
```
> private.get()
0
```
* In terminal window 3 (node 7)
```
> private.get()
42
```

So we now see nodes 1 and 7 are able to read the state of the private contract and it's initial value is 42. Node 4 is unable to read the state. Next we'll have node 1 set the state to the value `4` and verify only nodes 1 and 7 are able to view the new state.

In terminal window 1 (node 1)
```
> private.set(4,{from:eth.coinbase,privateFor:["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]});
"0xacf293b491cccd1b99d0cfb08464a68791cc7b5bc14a9b6e4ff44b46889a8f70"
```
You can check the log files in ~7nodes/qdata/logs directory to see each node validating the block with this new private transaction. Once the block containing the transaction has been validated we can once again check the state from each node 1, 4, and 7.
* In terminal window 1 (node 1)
```
> private.get()
4
```
* In terminal window 2 (node 4)
```
> private.get()
0
```
* In terminal window 3 (node 7)
```
> private.get()
4
```
And there you have it. All 7 nodes are validating the same blockchain of transactions, the private transactions carrying nothing other than a 512 bit hash, and only parties to private transactions are able to view and update the state of private contracts.


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

You can also inspect the log files under `qdata/logs/*.log` for further diagnostics messages around incoming / outgoing connection requests. Grep for `ALLOWED-BY` or `DENIED-BY`. Please be sure to enable verobsity for p2p module.

### Permissioning configuration

Permissioning is granted based on the remote key of the geth node. The remote keys are specified in the permissioned-nodes.json and is placed under individual nodes <datadir>.

The below sample permissioned-nodes.json provides a list of nodes permissioned to join the network ( node ids truncated for clarity)

```
[
   "enode://8475a01f22a1f48116dc1f0d22ecaaaf77e@127.0.0.1:30301",
   "enode://b5660501f496e60e59ded734a889c97b7da@127.0.0.1:30302",
   "enode://54bd7ff4bd971fb80493cf4706455395917@127.0.0.1:30303"
]
```

### Enabling/Disabling permissions

An individual node can enable/disable permissioning by passing the `-permissioned` command line flag. If enabled, then only the nodes that are in the `<datadir>/permissioned-nodes.json` can connect to it. Further, these are the only nodes that this node can make outbound connections to as well.

```
MISCELLANEOUS OPTIONS:
--permissioned          If enabled, the node will allow only a defined list of nodes to connect
```


