// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title TriangularArbitrageCalculator
/// @notice Library and helper contract to compute optimal trade size for a three pool arbitrage.
library TriangularArbitrageLib {
    uint256 constant FEE_NUMERATOR = 997;
    uint256 constant FEE_DENOMINATOR = 1000;

    function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) internal pure returns (uint256) {
        uint256 amountInWithFee = amountIn * FEE_NUMERATOR;
        return (amountInWithFee * reserveOut) / (reserveIn * FEE_DENOMINATOR + amountInWithFee);
    }

    function getProfit(
        uint256 amountIn,
        uint256 rAB0,
        uint256 rAB1,
        uint256 rBC0,
        uint256 rBC1,
        uint256 rCA0,
        uint256 rCA1
    ) internal pure returns (uint256) {
        uint256 outAB = getAmountOut(amountIn, rAB0, rAB1);
        uint256 outBC = getAmountOut(outAB, rBC0, rBC1);
        uint256 outCA = getAmountOut(outBC, rCA0, rCA1);
        return outCA > amountIn ? outCA - amountIn : 0;
    }

    function findBestInput(
        uint256 rAB0,
        uint256 rAB1,
        uint256 rBC0,
        uint256 rBC1,
        uint256 rCA0,
        uint256 rCA1,
        uint256 maxIn,
        uint256 step
    ) internal pure returns (uint256 bestIn, uint256 bestProfit) {
        for (uint256 i = step; i <= maxIn; i += step) {
            uint256 p = getProfit(i, rAB0, rAB1, rBC0, rBC1, rCA0, rCA1);
            if (p > bestProfit) {
                bestProfit = p;
                bestIn = i;
            }
        }
    }
}

contract TriangularArbitrageCalculator {
    function findOptimal(
        uint256 rAB0,
        uint256 rAB1,
        uint256 rBC0,
        uint256 rBC1,
        uint256 rCA0,
        uint256 rCA1,
        uint256 maxIn,
        uint256 step
    ) external pure returns (uint256 bestIn, uint256 bestProfit) {
        return TriangularArbitrageLib.findBestInput(rAB0, rAB1, rBC0, rBC1, rCA0, rCA1, maxIn, step);
    }
}
