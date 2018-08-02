pragma solidity ^0.4.24;

/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
  address public owner;


  event OwnershipRenounced(address indexed previousOwner);
  event OwnershipTransferred(
    address indexed previousOwner,
    address indexed newOwner
  );


  /**
   * @dev The Ownable constructor sets the original `owner` of the contract to the sender
   * account.
   */
  constructor() public {
    owner = msg.sender;
  }

  /**
   * @dev Throws if called by any account other than the owner.
   */
  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  /**
   * @dev Allows the current owner to relinquish control of the contract.
   * @notice Renouncing to ownership will leave the contract without an owner.
   * It will not be possible to call the functions with the `onlyOwner`
   * modifier anymore.
   */
  function renounceOwnership() public onlyOwner {
    emit OwnershipRenounced(owner);
    owner = address(0);
  }

  /**
   * @dev Allows the current owner to transfer control of the contract to a newOwner.
   * @param _newOwner The address to transfer ownership to.
   */
  function transferOwnership(address _newOwner) public onlyOwner {
    _transferOwnership(_newOwner);
  }

  /**
   * @dev Transfers control of the contract to a newOwner.
   * @param _newOwner The address to transfer ownership to.
   */
  function _transferOwnership(address _newOwner) internal {
    require(_newOwner != address(0));
    emit OwnershipTransferred(owner, _newOwner);
    owner = _newOwner;
  }
}

/**
 * @title DESContract
 * @dev DESContract is Ownable and allows the owner to keep track of regulator 
 * addresses and permissioned nodes in the DES network.
 */
contract DESContract is Ownable {
  
  // enode addresses
  mapping(string => bool) permissionedNodes;
  
  // regulator public key
  mapping(string => bool) regulators;
  
  event RegulatorRegistered(string indexed _address);
  event RegulatorRemoved(string indexed _address);
  event NodeAdded(string indexed enode);
  event NodeDeleted(string indexed enode);
  
  function DESContract (string reg, string enode, uint64 raftid) {
      register(reg);
        addNode(enode, raftid);
  }

  /**
   * @dev Allows anyone to check if the given public key belongs to a 
   * registered regulator
   * @param _address The public key to check for regulator privileges
   */
  function exists (string _address) public view returns (bool) {
    return regulators[_address];
  }


  /**
   * @dev Allows only the owner to add the given public key as a  
   * regulator
   * @param _address The address to register
   */ 
  function register (string _address) onlyOwner public {
    require(!exists(_address));
    regulators[_address] = true;
    emit RegulatorRegistered(_address);
  }

  /**
   * @dev Allows only the owner to remove the given public key from 
   * the regulator registry
   * @param _address The address to remove from registry
   */ 
  function remove (string _address) onlyOwner public {
    require(exists(_address));
    regulators[_address] = false;
    emit RegulatorRemoved(_address);
  }

  /**
   * @dev Allows anyone to check if the given enode address is a
   * registered node
   * @param enode The enode address to check
   */
  function nodeExists (string enode) public view returns (bool) {
    return permissionedNodes[enode];

  }

  /**
   * @dev Add an enode to the permissioned list (only owner can add)
   * @param enode The enode address to add
   */ 
  function addNode (string enode, uint64 raftid) onlyOwner public {
    require(!nodeExists(enode));
    permissionedNodes[enode] = true;
    emit NodeAdded(enode);
  }

  /**
   * @dev Remove enode from the permissioned list (only owner can remove)
   * @param enode The enode address to check
   */
  function deleteNode (string enode) onlyOwner public {
    require(nodeExists(enode));
    permissionedNodes[enode] = false;
    emit NodeDeleted(enode);
  }
  
}


