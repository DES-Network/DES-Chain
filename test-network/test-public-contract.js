a = eth.accounts[0]
web3.eth.defaultAccount = a;

var abi = [{"constant":true,"inputs":[],"name":"storedData","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"uint256"}],"name":"set","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"initVal","type":"uint256"}],"payable":false,"type":"constructor"}];

var address = "0x1349F3e1B8D71eFfb47B840594Ff27dA7E603d17";

var regContract = web3.eth.contract(abi);
var contractInstance = regContract.at(address);
var result = contractInstance.storedData();
console.log(result);

console.log("Change the value and check again");

contractInstance.set(55, {gas: 0x47b760, from:a});
// var getData = contractInstance.set.getData(550);
// //finally paas this data parameter to send Transaction
// web3.eth.sendTransaction({to:address, from:a, data: getData, gas: 0x47b760}, function (err, res) {
// 	if (!err) {
// 		console.log(res)
// 	}
// });
// console.log(getData);
// setTimeout(function(){
// 	result = contractInstance.storedData();
// 	console.log(result);
// }, 5000);


