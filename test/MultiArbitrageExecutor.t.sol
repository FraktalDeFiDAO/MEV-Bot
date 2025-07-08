// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/MultiArbitrageExecutor.sol";
import "../contracts/MockERC20.sol";
import "../contracts/MockPair.sol";

contract MultiArbitrageExecutorTest is Test {
    MultiArbitrageExecutor exec;
    MockERC20 tokenA;
    MockERC20 tokenB;
    MockERC20 tokenC;
    MockPair pairAB;
    MockPair pairBC;
    MockPair pairCA;

    function setUp() public {
        tokenA = new MockERC20("A","A",18);
        tokenB = new MockERC20("B","B",18);
        tokenC = new MockERC20("C","C",18);
        pairAB = new MockPair(address(tokenA), address(tokenB));
        pairBC = new MockPair(address(tokenB), address(tokenC));
        pairCA = new MockPair(address(tokenC), address(tokenA));
        pairAB.setReserves(1000, 1000);
        pairBC.setReserves(1000, 1000);
        pairCA.setReserves(800, 1200);
        exec = new MultiArbitrageExecutor();
        tokenA.mint(address(this), 1000);
        tokenA.approve(address(exec), type(uint256).max);
        tokenB.approve(address(pairBC), type(uint256).max);
        tokenC.approve(address(pairCA), type(uint256).max);
    }

    function testExecute() public {
        address[] memory pairs = new address[](3);
        pairs[0] = address(pairAB);
        pairs[1] = address(pairBC);
        pairs[2] = address(pairCA);
        uint256[] memory fN = new uint256[](3);
        uint256[] memory fD = new uint256[](3);
        for (uint256 i = 0; i < 3; i++) {
            fN[i] = 997;
            fD[i] = 1000;
        }
        uint256 beforeBal = tokenA.balanceOf(address(this));
        exec.execute(pairs, fN, fD, 500, 1);
        uint256 afterBal = tokenA.balanceOf(address(this));
        assertGt(afterBal, beforeBal);
    }
}
