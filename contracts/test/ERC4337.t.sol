// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../lib/account-abstraction/contracts/samples/SimpleAccountFactory.sol";
import "../lib/account-abstraction/contracts/interfaces/IEntryPoint.sol";

contract ERC4337Test is Test {

    SimpleAccountFactory public factory;

    function setUp() public {
        factory = new SimpleAccountFactory(IEntryPoint(0x0576a174D229E3cFA37253523E645A78A0C91B57));
    }
}
