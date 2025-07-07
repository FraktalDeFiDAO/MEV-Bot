// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

library TokenLib {
    bytes32 constant TOKEN_STORAGE_POSITION = keccak256("mev.bot.token.storage");

    struct TokenInfo {
        uint8 decimals;
        bool enabled;
    }

    struct TokenStorage {
        mapping(address => TokenInfo) info;
        address[] tokens;
    }

    function tokenStorage() internal pure returns (TokenStorage storage ts) {
        bytes32 position = TOKEN_STORAGE_POSITION;
        assembly {
            ts.slot := position
        }
    }

    function addToken(address token, uint8 decimals) internal {
        TokenStorage storage ts = tokenStorage();
        require(token != address(0), "zero token");
        if (ts.info[token].decimals == 0) {
            ts.tokens.push(token);
        }
        ts.info[token] = TokenInfo({decimals: decimals, enabled: true});
    }

    function addTokens(address[] memory tokens, uint8[] memory decimals) internal {
        require(tokens.length == decimals.length, "length mismatch");
        for (uint256 i = 0; i < tokens.length; i++) {
            addToken(tokens[i], decimals[i]);
        }
    }

    function setTokenEnabled(address token, bool enabled) internal {
        TokenStorage storage ts = tokenStorage();
        require(ts.info[token].decimals != 0, "token missing");
        ts.info[token].enabled = enabled;
    }

    function isTokenEnabled(address token) internal view returns (bool) {
        return tokenStorage().info[token].enabled;
    }
}
