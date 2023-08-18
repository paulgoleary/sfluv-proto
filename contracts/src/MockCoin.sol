pragma solidity <=0.8.19;

import "../lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

contract MockCoin is ERC20 {
    constructor() ERC20("Mock USDC", "MXUSDC") {}

    function mint(address account, uint256 amount) external {
        _mint(account, amount);
    }

    function burn(address account, uint256 amount) external {
        _burn(account, amount);
    }
}
