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


    function testAddTokenDuplicate() public {
        registry.addToken(address(1), 18);
        registry.addToken(address(1), 18);
        address[] memory tokens = registry.getTokens();
        assertEq(tokens.length, 1);
    }

    function testExchangeIdsIncrement() public {
        uint256 id0 = registry.addExchange("A", address(10));
        uint256 id1 = registry.addExchange("B", address(11));
        uint256 id2 = registry.addExchange("C", address(12));
        assertEq(id0, 0);
        assertEq(id1, 1);
        assertEq(id2, 2);
        assertEq(registry.getExchangeCount(), 3);
        ExchangeLib.ExchangeInfo[] memory exs = registry.getExchanges();
        assertEq(exs.length, 3);
    }

    function testAddPoolDuplicate() public {
        uint256 id = registry.addExchange("UniV2", address(11));
        registry.addPool(address(1), address(2), address(3), id);
        registry.addPool(address(1), address(2), address(3), id);
        address[] memory pools = registry.getPools();
        assertEq(pools.length, 1);
    }

    function testSetExchangeEnabled() public {
        uint256 id = registry.addExchange("UniV2", address(11));
        registry.setExchangeEnabled(id, false);
        ExchangeLib.ExchangeInfo memory info = registry.getExchange(id);
        assertTrue(!info.enabled);
        registry.setExchangeEnabled(id, true);
        assertTrue(registry.getExchange(id).enabled);
    }

    function testGetExchange() public {
        uint256 id = registry.addExchange("UniV2", address(11));
        ExchangeLib.ExchangeInfo memory info = registry.getExchange(id);
        assertEq(info.name, "UniV2");
        assertEq(info.router, address(11));
        assertTrue(info.enabled);
    }

    function testSetExchangeEnabledMissing() public {
        vm.expectRevert("exchange missing");
        registry.setExchangeEnabled(1, true);
    }

    function testAddTokensLengthMismatch() public {
        address[] memory addrs = new address[](2);
        addrs[0] = address(5);
        addrs[1] = address(6);
        uint8[] memory dec = new uint8[](1);
        dec[0] = 18;
        vm.expectRevert("length mismatch");
        registry.addTokens(addrs, dec);
    }

    function testAddTokenZeroAddress() public {
        vm.expectRevert("zero token");
        registry.addToken(address(0), 18);
    }

    function testSetTokenEnabledMissing() public {
        vm.expectRevert("token missing");
        registry.setTokenEnabled(address(77), true);
    }

    function testAddExchangeZeroRouter() public {
        vm.expectRevert("zero router");
        registry.addExchange("bad", address(0));
    }

    function testAddPoolZeroPool() public {
        uint256 id = registry.addExchange("UniV2", address(11));
        vm.expectRevert("zero pool");
        registry.addPool(address(0), address(1), address(2), id);
    }

    function testSetPoolEnabledMissing() public {
        vm.expectRevert("pool missing");
        registry.setPoolEnabled(address(123), true);
    }
}
