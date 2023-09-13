// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../src/SFLUVGovernorV1.sol";

contract SFLUVv1Test is Test {
    SFLUVVotesV1 public testVotes;
    SFLUVGovernorV1 public testLUVGov;
    TimelockController public testTimelock;

    StdCheats cheats;

    address internal citizen;

//    uint public oneEther = 1 ether; // TODO: elsewhere?

    function setUp() public {
        testVotes = new SFLUVVotesV1();
        address[] memory emptyAddrs;
        testTimelock = new TimelockController(43200, emptyAddrs, emptyAddrs, address(this));
        testLUVGov = new SFLUVGovernorV1(testVotes, testTimelock);

        cheats = StdCheats(HEVM_ADDRESS);

        citizen = makeAddr("citizen");
    }

    function testBasic() public {
        // referring to https://docs.openzeppelin.com/contracts/4.x/governance:
        // The Proposer role is in charge of queueing operations: this is the role the Governor instance should be granted,
        //  and it should likely be the only proposer in the system.
        assertFalse(testTimelock.hasRole(testTimelock.PROPOSER_ROLE(), address(testLUVGov)));
        testTimelock.grantRole(testTimelock.PROPOSER_ROLE(), address(testLUVGov));

        address[] memory proposalTargets = new address[](1);
        proposalTargets[0] = address(this); // doesn't matter

        uint256[] memory proposalValues = new uint256[](1);
        proposalValues[0] = 0;

        bytes[] memory proposalCalldata = new bytes[](1);
        proposalCalldata[0] = "0xDEADBEEF";

        string memory proposalDesc = "This is a test proposal.";

        // mint dao token to citizen
        testVotes.mint(citizen);

        // citizen needs to self delegate
        vm.startPrank(citizen);
        testVotes.delegate(citizen);
        vm.stopPrank();

        uint256 currentWeight = testLUVGov.getVotes(citizen, testLUVGov.clock() - 1);
        assertEq(0, currentWeight);

        vm.startPrank(citizen);
        // vm.expectEmit();
        // function propose(address[] memory targets, uint256[] memory values, bytes[] memory calldatas, string memory description)
        testLUVGov.propose(proposalTargets, proposalValues, proposalCalldata, proposalDesc);
        vm.stopPrank();

        uint256 proposalId = testLUVGov.hashProposal(proposalTargets, proposalValues, proposalCalldata, keccak256(bytes(proposalDesc)));
        address checkAddr = testLUVGov.proposalProposer(proposalId);
        console.log(checkAddr);
        assertEq(checkAddr, citizen, "citizen should have made a proposal");

    }
}
