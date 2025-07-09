// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/TriangularArbitrageCalculator.sol";

contract TriangularArbitrageCalculatorTest is Test {
    function testFindOptimalTriangular() public {
        TriangularArbitrageCalculator calc = new TriangularArbitrageCalculator();
        (uint256 amt, uint256 profit) = calc.findOptimal(1000, 1000, 1000, 1000, 800, 1200, 500, 1);
        assertGt(amt, 0);
        assertGt(profit, 0);
    }
}
