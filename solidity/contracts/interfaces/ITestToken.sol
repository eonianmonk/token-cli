// SPDX-License-Identifier: TODO
pragma solidity ^0.8.20;


interface ITestToken {
    event Transfer(address indexed sender, address indexed receiver, uint256 value);
    event Mint(address indexed to, uint256 amount);

    function mint(address receiver, uint256 amount) external;
    function balanceOf(address owner) external view returns (uint256 balance);
    function transfer(address receiver, uint256 amount) external;
}
