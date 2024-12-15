// SPDX-License-Identifier: TODO (MIT I guess)
pragma solidity ^0.8.20;

// Ownable implementation is straightforward: 
// store owner in constructor and check when onlyOwner modifier is applied 
// https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/access/Ownable.sol
import "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./Freezable.sol";
import "./interfaces/ITestToken.sol";

contract TestToken is Ownable,Freezable,ITestToken {
    mapping(address => uint256) private balances;

    constructor() Ownable(msg.sender) {}

    function freeze(address account) public override onlyOwner {
        _freezeAccount(account);
    }

    function unfreeze(address account) public override onlyOwner {
        _unfreezeAccount(account);
    }

    function mint(address receiver, uint256 amount) public onlyOwner notFrozen(receiver) {
        // why mint to immediately burn?
        require(receiver != address(0), "minting to zero address is not allowed");
        // if we add totalSupply, we need to add to it here
        balances[receiver] += amount;
        emit Mint(receiver, amount);
    }

    function balanceOf(address account) public view returns (uint256) {
        // can't check  burnt supply
        require(account != address(0), "invalid account address");
        return balances[account];
    }

    // if receiver is 0, we are burning tokens, so it can be a valid address 
    function transfer(address receiver, uint256 amount) public notFrozen(msg.sender) notFrozen(receiver) {
        require(amount > 0, "amount must be positive");
        require(balances[msg.sender] >= amount, "insufficient balance for transfer amount");
        // TODO: forbid transfer to yourself?

        balances[msg.sender] -= amount;

        // no need to track burnt tokens
        // if we operate with totalSupply, need to subtract from total supply when burning
        if (receiver != address(0)) {
            balances[receiver] += amount;
        }

        emit Transfer(msg.sender, receiver, amount);
    }
}