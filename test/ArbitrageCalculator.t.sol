// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/ArbitrageCalculator.sol";

contract ArbitrageCalculatorTest is Test {
    function testFindOptimal() public {
        ArbitrageCalculator calc = new ArbitrageCalculator();
        (uint256 amountIn, uint256 profit) = calc.findOptimal(1000, 1000, 1200, 800, 500, 1);
        assertGt(amountIn, 0);
        assertGt(profit, 0);
    }
}
