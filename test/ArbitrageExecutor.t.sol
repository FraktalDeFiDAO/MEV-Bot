// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/ArbitrageExecutor.sol";
import "../contracts/MockERC20.sol";
import "../contracts/MockPair.sol";

contract ArbitrageExecutorTest is Test {
    ArbitrageExecutor exec;
    MockERC20 tokenA;
    MockERC20 tokenB;
    MockPair pairA;
    MockPair pairB;

    function setUp() public {
        tokenA = new MockERC20("A", "A", 18);
        tokenB = new MockERC20("B", "B", 18);
        pairA = new MockPair(address(tokenA), address(tokenB));
        pairB = new MockPair(address(tokenB), address(tokenA));
        pairA.setReserves(1000, 1000);
        pairB.setReserves(800, 1200);
        exec = new ArbitrageExecutor();
        tokenA.mint(address(this), 1000);
        tokenA.approve(address(exec), type(uint256).max);
        tokenB.approve(address(pairB), type(uint256).max);
    }

    function testExecute() public {
        vm.recordLogs();
        uint256 beforeBal = tokenA.balanceOf(address(this));
        exec.execute(address(pairA), address(pairB), 500, 1);
        uint256 afterBal = tokenA.balanceOf(address(this));
        assertGt(afterBal, beforeBal);
        Vm.Log[] memory logs = vm.getRecordedLogs();
        bytes32 sig = keccak256("TradeExecuted(uint256,uint256)");
        bool found;
        for (uint256 i = 0; i < logs.length; i++) {
            if (logs[i].topics[0] == sig) {
                found = true;
                break;
            }
        }
        assertTrue(found);
    }
}
