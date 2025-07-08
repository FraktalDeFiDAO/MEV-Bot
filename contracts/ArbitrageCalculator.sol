// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title ArbitrageCalculator
/// @notice Simple library and contract to compute best trade size for two constant product pools.

library ArbitrageLib {
    uint256 constant FEE_NUMERATOR = 997;
    uint256 constant FEE_DENOMINATOR = 1000;

    function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) internal pure returns (uint256) {
        uint256 amountInWithFee = amountIn * FEE_NUMERATOR;
        return (amountInWithFee * reserveOut) / (reserveIn * FEE_DENOMINATOR + amountInWithFee);
    }

    function getProfit(uint256 amountIn, uint256 rA0, uint256 rB0, uint256 rA1, uint256 rB1) internal pure returns (uint256) {
        uint256 out1 = getAmountOut(amountIn, rA0, rB0);
        uint256 out2 = getAmountOut(out1, rB1, rA1);
        return out2 > amountIn ? out2 - amountIn : 0;
    }

    function findBestInput(
        uint256 rA0,
        uint256 rB0,
        uint256 rA1,
        uint256 rB1,
        uint256 maxIn,
        uint256 step
    ) internal pure returns (uint256 bestIn, uint256 bestProfit) {
        for (uint256 i = step; i <= maxIn; i += step) {
            uint256 p = getProfit(i, rA0, rB0, rA1, rB1);
            if (p > bestProfit) {
                bestProfit = p;
                bestIn = i;
            }
        }
    }
}

contract ArbitrageCalculator {
    function findOptimal(
        uint256 rA0,
        uint256 rB0,
        uint256 rA1,
        uint256 rB1,
        uint256 maxIn,
        uint256 step
    ) external pure returns (uint256 bestIn, uint256 bestProfit) {
        return ArbitrageLib.findBestInput(rA0, rB0, rA1, rB1, maxIn, step);
    }
}

