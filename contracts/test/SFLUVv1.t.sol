// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/MockCoin.sol";
import "../src/SFLUVv1.sol";

contract SFLUVv1Test is Test {
    MockCoin public mockCoin;
    SFLUVv1 public testLUVCoin;

    StdCheats cheats;

    address internal payer;
    address internal payee;

    uint public oneEther = 1 ether; // TODO: elsewhere?

    function setUp() public {
        mockCoin = new MockCoin();
        testLUVCoin = new SFLUVv1(mockCoin);

        cheats = StdCheats(HEVM_ADDRESS);

        payer = makeAddr("payer");
        payee = makeAddr("payee");
    }

    function testBasic() public {
        mockCoin.mint(payer, 100 * oneEther);

        uint checkBalance = mockCoin.balanceOf(payer);
        assertEq(checkBalance, 100 * oneEther);

        // contract owner has DEFAULT_ADMIN_ROLE by default
        assertTrue(testLUVCoin.owner() == address(this));
        assertTrue(testLUVCoin.hasRole(testLUVCoin.DEFAULT_ADMIN_ROLE(), testLUVCoin.owner()));

        // contract owner does *not* have minter role by default
        assertFalse(testLUVCoin.hasRole(testLUVCoin.MINTER_ROLE(), testLUVCoin.owner()));
        testLUVCoin.grantRole(testLUVCoin.MINTER_ROLE(), testLUVCoin.owner());
        assertTrue(testLUVCoin.hasRole(testLUVCoin.MINTER_ROLE(), testLUVCoin.owner()));

        // unfunded (default) account cannot wrap underlying coin w/o balance
        vm.expectRevert("ERC20: insufficient allowance");
        testLUVCoin.depositFor(payer, 100 * oneEther);

        // allow the payer to wrap (mint) SFLUV
        testLUVCoin.grantRole(testLUVCoin.MINTER_ROLE(), payer);

        vm.startPrank(payer);
        mockCoin.approve(address(testLUVCoin), 100 * oneEther);
        testLUVCoin.depositFor(payer, 100 * oneEther);
        vm.stopPrank();

        checkBalance = testLUVCoin.balanceOf(payer); // payer now has luv
        assertEq(checkBalance, 100 * oneEther);

        checkBalance = mockCoin.balanceOf(address(testLUVCoin)); // wrapper contract is holding base coin
        assertEq(checkBalance, 100 * oneEther);

        vm.prank(payer);
        testLUVCoin.transfer(payee, 50 * oneEther); // spread some luv

        checkBalance = testLUVCoin.balanceOf(payer);
        assertEq(checkBalance, 50 * oneEther, "payer gave some luv");

        checkBalance = testLUVCoin.balanceOf(payee); // wrapper contract is holding base coin
        assertEq(checkBalance, 50 * oneEther, "payee has some luv");
    }
}
