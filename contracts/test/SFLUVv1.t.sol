// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/MockCoin.sol";
import "../src/SFLUVv1.sol";

contract SFLUVv1Test is Test {
    MockCoin public mockCoin;
    SFLUVv1 public testCoin;

    StdCheats cheats;

    address internal payer;
    address internal payee;

    uint public oneEther = 1 ether; // TODO: elsewhere?

    function setUp() public {
        mockCoin = new MockCoin();
        testCoin = new SFLUVv1(mockCoin);

        cheats = StdCheats(HEVM_ADDRESS);

        payer = makeAddr("payer");
        payee = makeAddr("payee");
    }

    function testBasic() public {
        mockCoin.mint(payer, 100 * oneEther);

        uint checkBalance = mockCoin.balanceOf(payer);
        assertEq(checkBalance, 100 * oneEther);

        assertTrue(testCoin.owner() == address(this));
        assertFalse(testCoin.hasRole(testCoin.MINTER_ROLE(), testCoin.owner()));
        testCoin.grantRole(testCoin.MINTER_ROLE(), testCoin.owner());
        assertTrue(testCoin.hasRole(testCoin.MINTER_ROLE(), testCoin.owner()));

        vm.expectRevert("ERC20: insufficient allowance");
        testCoin.depositFor(payer, 100 * oneEther);

        // allow the payer to wrap (mint) SFLUV
        testCoin.grantRole(testCoin.MINTER_ROLE(), payer);

        vm.startPrank(payer);
        mockCoin.approve(address(testCoin), 100 * oneEther);
        testCoin.depositFor(payer, 100 * oneEther);
        vm.stopPrank();

        checkBalance = testCoin.balanceOf(payer); // payer now has luv
        assertEq(checkBalance, 100 * oneEther);

        checkBalance = mockCoin.balanceOf(address(testCoin)); // wrapper contract is holding base coin
        assertEq(checkBalance, 100 * oneEther);

        vm.prank(payer);
        testCoin.transfer(payee, 50 * oneEther); // spread some luv

        checkBalance = testCoin.balanceOf(payer);
        assertEq(checkBalance, 50 * oneEther, "payer gave some luv");

        checkBalance = testCoin.balanceOf(payee); // wrapper contract is holding base coin
        assertEq(checkBalance, 50 * oneEther, "payee has some luv");
    }
}
