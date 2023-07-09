pragma solidity ^0.8.0;

import "../lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Wrapper.sol";
import "../lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol";
import "../lib/openzeppelin-contracts/contracts/access/AccessControlDefaultAdminRules.sol";

contract SFLUVv1 is ERC20Wrapper, ERC20Permit, AccessControlDefaultAdminRules {

    uint48 constant private initialDelay = 60 * 60 * 24 * 7; // 7 days?

    constructor(IERC20 underlyingToken)
        ERC20Wrapper(underlyingToken)
        ERC20Permit("SFLUV V1")
        ERC20("SFLUV V1", "SFLUV")
        AccessControlDefaultAdminRules(initialDelay, msg.sender) {}

    function decimals() public pure override(ERC20, ERC20Wrapper) returns (uint8) {
        return 18;
    }

    // this role allows the holder to mint (wrap) underlying USDC into SFLUV
    bytes32 public constant MINTER_ROLE = keccak256("MINTER");

    function depositFor(address account, uint256 amount) public override returns (bool) {
        require(hasRole(MINTER_ROLE, _msgSender()));
        return ERC20Wrapper.depositFor(account, amount);
    }

}
