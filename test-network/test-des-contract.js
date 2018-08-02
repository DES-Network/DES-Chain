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