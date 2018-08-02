# Changelog
All notable changes to this project will be documented in this file.

## [1.0.0-RC] - 2018-08-02
### Added
- contracts/des/des.sol: This is the DES contract file, which contains logic for keeping track of, adding and deleting regulators and permissioned nodes.
- contracts/des/des.go: The file was added, which is code generated against the DES contract's ABI for native go binding. It allows interaction with the DES contract from within the geth code. 
- private/regulators.go: This file was added to provide an interface to interact with the contract with a layer of abstraction for creating IPC client, handling contract pre-deployment, etc.
- test-network: This directory was added to provide easy access to testing the DES network via scripts. 

### Changed
- p2p/server.go: Removed permissioning as an option - DES is always permissioned.
- p2p/permissions.go: This is Quorum's addition, to which we made changes so that we could allow the DES contract to override once the contract has been deployed.
- raft/handler.go: This file was also Quorum's addition, but we changed it to handle updates in permissioning, so that it would actively add or remove raft peers based on the current permissions.
- internal/ethapi/api.go: This is where we check for the regulator's presence in a private transaction. Will reject transaction where regulator doesn't exist.
