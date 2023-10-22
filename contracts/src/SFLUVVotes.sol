// SPDX-License-Identifier: MIT
pragma solidity <=0.8.19;

import "../lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Votes.sol";
import "../lib/openzeppelin-contracts/contracts/access/AccessControlDefaultAdminRules.sol";

contract SFLUVVotesV1 is ERC721Votes, AccessControlDefaultAdminRules {
    uint48 constant private initialDelay = 60 * 60 * 24 * 7; // 7 days?
    constructor()
    ERC721("SFLUVVotes V1.0", "SFLUVVotes")
    EIP712("SFLUVVotes", "0.0.1")
    AccessControlDefaultAdminRules(initialDelay, msg.sender) {}

    bytes32 public constant MINTER_ROLE = keccak256("MINTER");
    bytes32 public constant BURNER_ROLE = keccak256("BURNER");

    uint256 private _tokenId;

    /**
    * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721, AccessControlDefaultAdminRules) returns (bool) {
        return ERC721.supportsInterface(interfaceId) || AccessControlDefaultAdminRules.supportsInterface(interfaceId);
    }

    // function _safeMint(address to, uint256 tokenId) internal virtual
    function mint(address to) public {
        require(hasRole(MINTER_ROLE, _msgSender()));
        _safeMint(to, _tokenId++);
    }

    function burn(uint256 tokenId) public {
        require(hasRole(BURNER_ROLE, _msgSender()));
        _burn(tokenId);
    }
}
