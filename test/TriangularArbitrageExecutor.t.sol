// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/TriangularArbitrageExecutor.sol";
import "../contracts/MockERC20.sol";
import "../contracts/MockPair.sol";

contract TriangularArbitrageExecutorTest is Test {
    TriangularArbitrageExecutor exec;
    MockERC20 tokenA;
    MockERC20 tokenB;
    MockERC20 tokenC;
    MockPair pairAB;
    MockPair pairBC;
    MockPair pairCA;

    function setUp() public {
        tokenA = new MockERC20("A", "A", 18);
        tokenB = new MockERC20("B", "B", 18);
        tokenC = new MockERC20("C", "C", 18);
        pairAB = new MockPair(address(tokenA), address(tokenB));
        pairBC = new MockPair(address(tokenB), address(tokenC));
        pairCA = new MockPair(address(tokenC), address(tokenA));
        pairAB.setReserves(1000, 1000);
        pairBC.setReserves(1000, 1000);
        pairCA.setReserves(800, 1200);
        exec = new TriangularArbitrageExecutor();
        tokenA.mint(address(this), 1000);
        tokenA.approve(address(exec), type(uint256).max);
        tokenB.approve(address(pairBC), type(uint256).max);
        tokenC.approve(address(pairCA), type(uint256).max);
    }

    function testExecuteTriangular() public {
        uint256 beforeBal = tokenA.balanceOf(address(this));
        exec.execute(address(pairAB), address(pairBC), address(pairCA), 500, 1);
        uint256 afterBal = tokenA.balanceOf(address(this));
        assertGt(afterBal, beforeBal);
    }
}
