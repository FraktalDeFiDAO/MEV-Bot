// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./libraries/TokenLib.sol";
import "./libraries/ExchangeLib.sol";
import "./libraries/PoolLib.sol";

/// @title Registry contract holding tokens, exchanges and pools
/// @notice Uses library-based diamond storage for modularity
contract Registry {
    using TokenLib for *;
    using ExchangeLib for *;
    using PoolLib for *;

    // ---- token management ----
    function addToken(address token, uint8 decimals) external {
        TokenLib.addToken(token, decimals);
    }

    function addTokens(address[] calldata tokens, uint8[] calldata decimals) external {
        TokenLib.addTokens(tokens, decimals);
    }

    function setTokenEnabled(address token, bool enabled) external {
        TokenLib.setTokenEnabled(token, enabled);
    }

    function isTokenEnabled(address token) external view returns (bool) {
        return TokenLib.isTokenEnabled(token);
    }

    function getToken(address token) external view returns (TokenLib.TokenInfo memory) {
        return TokenLib.getToken(token);
    }

    function getTokens() external view returns (address[] memory) {
        return TokenLib.getTokens();
    }

    function getTokenCount() external view returns (uint256) {
        return TokenLib.getTokenCount();
    }

    // ---- exchange management ----
    function addExchange(string calldata name, address router) external returns (uint256 id) {
        return ExchangeLib.addExchange(name, router);
    }

    function setExchangeEnabled(uint256 id, bool enabled) external {
        ExchangeLib.setExchangeEnabled(id, enabled);
    }

    function getExchange(uint256 id) external view returns (ExchangeLib.ExchangeInfo memory) {
        return ExchangeLib.getExchange(id);
    }

    function getExchangeCount() external view returns (uint256) {
        return ExchangeLib.getExchangeCount();
    }

    function getExchanges() external view returns (ExchangeLib.ExchangeInfo[] memory) {
        return ExchangeLib.getExchanges();
    }

    // ---- pool management ----
    function addPool(address pool, address token0, address token1, uint256 exchangeId) external {
        PoolLib.addPool(pool, token0, token1, exchangeId);
    }

    function setPoolEnabled(address pool, bool enabled) external {
        PoolLib.setPoolEnabled(pool, enabled);
    }

    function isPoolEnabled(address pool) external view returns (bool) {
        return PoolLib.isPoolEnabled(pool);
    }

    function getPool(address pool) external view returns (PoolLib.PoolInfo memory) {
        return PoolLib.getPool(pool);
    }

    function getPools() external view returns (address[] memory) {
        return PoolLib.getPools();
    }

    function getPoolCount() external view returns (uint256) {
        return PoolLib.getPoolCount();
    }
}
