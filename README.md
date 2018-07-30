# DES-Network

The DES Network is a fork of Quorum that adds some restrictions to its protocol, by limiting permissioned nodes to a curated whitelist controlled by a set of Regulators. It also ensures that all private transactions are privy to at least one Regulator.

These restrictions are controlled via a smart contract that is deployed after the network is launched by the network owner.

Key enhancements over Quorum:
*__Regulators__ - DES requires Regulators in the network, which are considered central authorities responsible for providing mediation, if required. 
*__Permissioned__ - DES enhances Quorum's permissioning to make permissioned nodes mandatory and adds a mechanism whereby only Regulators may whitelist a node that wishes to participate in the network and only nodes that are whitelisted would be part of the network.

## Quorum

Quorum is a fork of [go-ethereum](https://github.com/ethereum/go-ethereum) and is updated in line with go-ethereum releases.

Key enhancements over go-ethereum:

  * __Privacy__ - Quorum supports private transactions and private contracts through public/private state separation and utilising [Constellation](https://github.com/jpmorganchase/constellation), a peer-to-peer encrypted message exchange for directed transfer of private data to network participants
  * __Alternative Consensus Mechanisms__ - with no need for POW/POS in a permissioned network, Quorum instead offers multiple consensus mechanisms that are more appropriate for consortium chains:
    * __Raft-based Consensus__ - a consensus model for faster blocktimes, transaction finality, and on-demand block creation
    * __Istanbul BFT__ - a PBFT-inspired consensus algorithm with transaction finality, by AMIS.
  * __Peer Permissioning__ - node/peer permissioning using smart contracts, ensuring only known parties can join the network
  * __Higher Performance__ - Quorum offers significantly higher performance than public geth

Note: The QuorumChain consensus algorithm is not yet supported by this release.

## Architecture

<a href="https://github.com/jpmorganchase/quorum/wiki/Transaction-Processing#private-transaction-process-flow">![Quorum privacy architecture](https://github.com/jpmorganchase/quorum-docs/raw/master/images/QuorumTransactionProcessing.JPG)</a>

The above diagram is a high-level overview of the privacy architecture used by Quorum. For more in-depth discussion of the components, refer to the [wiki](https://github.com/jpmorganchase/quorum/wiki/) pages.

## Quickstart

The quickest way to get started with DES is to clone this repo and navigate to the test-network folder:

```sh
git clone https://github.com/jpmorganchase/quorum-examples
cd test-network
```

There you will find detailed instructions on how to set up a test network and test scripts for analyzing different scenarios.

## Further Reading

Further documentation can be found in the [docs](docs/) folder and on the [wiki](https://github.com/jpmorganchase/quorum/wiki/).

## See also

* [Quorum](https://github.com/jpmorganchase/quorum): this repository
* [Constellation](https://github.com/jpmorganchase/constellation): peer-to-peer encrypted message exchange for transaction privacy
* [Raft Consensus Documentation](raft/doc.md)
* [Istanbul BFT Consensus Documentation](https://github.com/ethereum/EIPs/issues/650): [RPC API](https://github.com/getamis/go-ethereum/wiki/RPC-API) and [technical article](https://medium.com/getamis/istanbul-bft-ibft-c2758b7fe6ff)
* [ZSL](https://github.com/jpmorganchase/quorum/wiki/ZSL) wiki page and [documentation](https://github.com/jpmorganchase/zsl-q/blob/master/README.md)
* [quorum-examples](https://github.com/jpmorganchase/quorum-examples): example quorum clusters
* [quorum-tools](https://github.com/jpmorganchase/quorum-tools): local cluster orchestration, and integration testing tool
* [Quorum Wiki](https://github.com/jpmorganchase/quorum/wiki)

## Third Party Tools/Libraries

The following Quorum-related libraries/applications have been created by Third Parties and as such are not specifically endorsed by J.P. Morgan.  A big thanks to the developers for improving the tooling around Quorum!

* [Quorum Blockchain Explorer](https://github.com/blk-io/blk-explorer-free) - a Blockchain Explorer for Quorum which supports viewing private transactions
* [Quorum-Genesis](https://github.com/davebryson/quorum-genesis) - A simple CL utility for Quorum to help populate the genesis file with voters and makers
* [Quorum Maker](https://github.com/synechron-finlabs/quorum-maker/) - a utility to create Quorum nodes
* [QuorumNetworkManager](https://github.com/ConsenSys/QuorumNetworkManager) - makes creating & managing Quorum networks easy
* [ERC20 REST service](https://github.com/blk-io/erc20-rest-service) - a Quorum-supported RESTful service for creating and managing ERC-20 tokens
* [Nethereum Quorum](https://github.com/Nethereum/Nethereum/tree/master/src/Nethereum.Quorum) - a .NET Quorum adapter
* [web3j-quorum](https://github.com/web3j/quorum) - an extension to the web3j Java library providing support for the Quorum API

## Contributing

Thank you for your interest in contributing to Quorum!

Quorum is built on open source and we invite you to contribute enhancements. Upon review you will be required to complete a Contributor License Agreement (CLA) before we are able to merge. If you have any questions about the contribution process, please feel free to send an email to [quorum_info@jpmorgan.com](mailto:quorum_info@jpmorgan.com).

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
