a = eth.accounts[0]
web3.eth.defaultAccount = a;

var abi = [{"constant":true,"inputs":[{"name":"enode","type":"string"}],"name":"nodeExists","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_address","type":"string"}],"name":"exists","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"renounceOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_address","type":"string"}],"name":"remove","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"enode","type":"string"}],"name":"addNode","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"enode","type":"string"}],"name":"deleteNode","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_address","type":"string"}],"name":"register","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_address","type":"string"}],"name":"RegulatorRegistered","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_address","type":"string"}],"name":"RegulatorRemoved","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"enode","type":"string"}],"name":"NodeAdded","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"enode","type":"string"}],"name":"NodeDeleted","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"previousOwner","type":"address"}],"name":"OwnershipRenounced","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"previousOwner","type":"address"},{"indexed":true,"name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"}];

var address = "0x1932c48b2bF8102Ba33B4A6B545C32236e342f34";

var regContract = web3.eth.contract(abi);
var contractInstance = regContract.at(address);
var result = contractInstance.exists("BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=");
if (result) {
	console.log("Node 1 is a regulator")
} else {
	console.log("Node 1 is not a regulator");
}

result = contractInstance.nodeExists('ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef');
if (result) {
	console.log("Node 1 has been added as a permissioned node")
} else {
	console.log("Node 1 has not been added as a permissioned node");
}

result = contractInstance.nodeExists('0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416');
if (result) {
	console.log("Node 2 has been added as a permissioned node")
} else {
	console.log("Node 2 has not been added as a permissioned node");
}

result = contractInstance.nodeExists('579f786d4e2830bbcc02815a27e8a9bacccc9605df4dc6f20bcc1a6eb391e7225fff7cb83e5b4ecd1f3a94d8b733803f2f66b7e871961e7b029e22c155c3a778');
if (result) {
	console.log("Node 3 has been added as a permissioned node")
} else {
	console.log("Node 3 has not been added as a permissioned node");
}

