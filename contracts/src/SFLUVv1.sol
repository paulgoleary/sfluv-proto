pragma solidity ^0.8.0;

import "../lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Wrapper.sol";
import "../lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol";

contract SFLUVv1 is ERC20Wrapper, ERC20Permit {

    constructor(IERC20 underlyingToken) ERC20Wrapper(underlyingToken) ERC20Permit("SFLUV V1") ERC20("SFLUV V1", "SFLUV") {}

    function decimals() public view override(ERC20, ERC20Wrapper) returns (uint8) {
        return 18;
    }

}
