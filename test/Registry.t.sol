// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/Registry.sol";
import "../contracts/libraries/TokenLib.sol";
import "../contracts/libraries/ExchangeLib.sol";
import "../contracts/libraries/PoolLib.sol";

contract RegistryTest is Test {
    Registry registry;

    function setUp() public {
        registry = new Registry();
    }

    function testAddToken() public {
        registry.addToken(address(1), 18);
        assertTrue(registry.isTokenEnabled(address(1)));
    }

    function testDisableToken() public {
        registry.addToken(address(2), 6);
        registry.setTokenEnabled(address(2), false);
        assertTrue(!registry.isTokenEnabled(address(2)));
    }

    function testGetTokens() public {
        address[] memory addrs = new address[](2);
        addrs[0] = address(3);
        addrs[1] = address(4);
        uint8[] memory dec = new uint8[](2);
        dec[0] = 18;
        dec[1] = 6;
        registry.addTokens(addrs, dec);
        address[] memory tokens = registry.getTokens();
        assertEq(tokens.length, 2);
        assertEq(tokens[0], address(3));
        TokenLib.TokenInfo memory info = registry.getToken(address(3));
        assertEq(info.decimals, 18);
        assertTrue(info.enabled);
    }

    function testAddExchangeAndPool() public {
        uint256 id = registry.addExchange("UniV2", address(11));
        registry.addPool(address(99), address(1), address(2), id);
        assertTrue(registry.isPoolEnabled(address(99)));
        registry.setPoolEnabled(address(99), false);
        assertTrue(!registry.isPoolEnabled(address(99)));
        assertEq(registry.getExchangeCount(), 1);
        ExchangeLib.ExchangeInfo[] memory exs = registry.getExchanges();
        assertEq(exs.length, 1);
        assertEq(exs[0].router, address(11));
        address[] memory pools = registry.getPools();
        assertEq(pools[0], address(99));
        PoolLib.PoolInfo memory pinfo = registry.getPool(address(99));
        assertEq(pinfo.token0, address(1));
    }
}
