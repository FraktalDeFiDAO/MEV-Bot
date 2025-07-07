// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/Registry.sol";

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

    function testAddExchangeAndPool() public {
        uint256 id = registry.addExchange("UniV2", address(11));
        registry.addPool(address(99), address(1), address(2), id);
        assertTrue(registry.isPoolEnabled(address(99)));
        registry.setPoolEnabled(address(99), false);
        assertTrue(!registry.isPoolEnabled(address(99)));
    }
}
