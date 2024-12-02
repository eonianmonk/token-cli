// SPDX-License-Identifier: TODO
pragma solidity ^0.8.20;

abstract contract Freezable {
    mapping(address=>bool) private frozenAccounts;

    event Freeze(address indexed account);
    event Unfreeze(address indexed account);

    function freeze(address account) public virtual;
    function unfreeze(address account) public virtual;

    function _freezeAccount(address account) internal {
        frozenAccounts[account] = true;
        emit Freeze(account);
    }

    function _unfreezeAccount(address account) internal {
        frozenAccounts[account] = false;
        emit Unfreeze(account);
    }
    
    modifier notFrozen(address account) {
        require(!frozenAccounts[account],"account is frozen");
        _;
    }
}