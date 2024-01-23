// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Account} from "../lib/smartcontracts/contracts/accounts/Account.sol";
import {ITokenEntryPoint} from "../lib/smartcontracts/contracts/accounts/interfaces/ITokenEntryPoint.sol";
import {TokenEntryPoint} from "../lib/smartcontracts/contracts/accounts/TokenEntryPoint.sol"; // to build the ABI

import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

contract SFLUVAccount is Account {
    constructor(IEntryPoint anEntryPoint, ITokenEntryPoint aTokenEntryPoint)
    Account(anEntryPoint, aTokenEntryPoint) {}
}
