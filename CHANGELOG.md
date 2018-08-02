
# Changelog
Des-Quorum is a fork of JPMorganChase's Quorum. All notable changes to this fork will be documented in this file.

## [1.0.0-beta] - 2018-08-02
### Added
- `contracts/des/des.sol`: This is the DES contract file, which contains logic for keeping track of, adding and deleting regulators and permissioned nodes.
- `contracts/des/des.go`: The file was added, which is code generated against the DES contract's ABI for native go binding. It allows interaction with the DES contract from within the geth code. 
- `private/regulators.go`: This file was added to provide an interface to interact with the contract with a layer of abstraction for creating IPC client, handling contract pre-deployment, etc.
- `cmd/des-genesis/main.go`: DES genesis generation tool.
- `test-network`: This directory was added to provide easy access to testing the DES network via scripts. More information [here](https://github.com/DES-Network/des-quorum/tree/master/test-network/README.md).
	- `init.sh`: Cleanup previous directories (delete any previous chain data) and (re)initialize accounts and keystores for configuring a new network and its nodes. 
	- `start-des-network.sh`: Launch `constellation` and `geth` node for the network owner. 
	- `stop.sh`: Shuts down the network. 
	- `deploy-and-init.js`: Deploy the DES network contract that will help add/remove and keep track of the regulators in the network and the nodes in the network. 
	- `runscript.sh`:  Used to run javascript code in geth console.
	- `test-des-contract.js`: Script to test DES contract.
  - `wrong-private-contract.js`: Bad example of a private contract.
  - `private-contract.js`: Good example of private contract.
  -  `public-contract.js`: Good example of private contract.
  - `whitelist.js`: Whitelist the given node on the des network.
  - `blacklist.js`: Blacklist the given node on the des network.
  -  `test-public-contract`: Write to/Read from public contract.
  - `simplestorage.sol`: Simple storage contract being used in all examples.
  -  Others: Files and folders meant for configuration or holding binaries.

### Changed
- `p2p/server.go`: Removed permissioning as an option - DES is always permissioned.
- `p2p/permissions.go`: This is Quorum's addition, to which we made changes so that we could allow the DES contract to override once the contract has been deployed.
- `raft/handler.go`: This file was also Quorum's addition, but we changed it to handle updates in permissioning, so that it would actively add or remove raft peers based on the current permissions.
- `internal/ethapi/api.go`: This is where we check for the regulator's presence in a private transaction. Will reject transaction where regulator doesn't exist.
- `Makefile`: Updated to add make target for des-genesis.
- `cmd/geth/chaincmd.go`: Write out genesis block to a file, for setup purposes.
- `core/genesis.go`: Hardcoded DES genesis.
- `core/genesis_alloc.go`: Hardcoded DES genesis alloc.
-  `params/config.go`: Hardcoded DES chain config.
- `eth/config.go`: Remove permission enabling flag.
- `node/config.go`: Remove permission enabling flag.
- `node/node.go`: Remove permission enabling flag.