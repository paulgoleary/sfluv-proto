pragma solidity ^0.8.0;

import "../lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

contract MockCoin is ERC20 {
    constructor() ERC20("Mock USDC", "MXUSDC") {}
}
