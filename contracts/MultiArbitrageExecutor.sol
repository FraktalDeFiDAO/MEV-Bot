// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./MockERC20.sol";

interface IPair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function getReserves() external view returns (uint112, uint112, uint32);
    function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes calldata data) external;
}

library MultiArbitrageLib {
    function getAmountOut(
        uint256 amountIn,
        uint256 reserveIn,
        uint256 reserveOut,
        uint256 feeNumerator,
        uint256 feeDenominator
    ) internal pure returns (uint256) {
        uint256 amountInWithFee = amountIn * feeNumerator;
        return (amountInWithFee * reserveOut) / (reserveIn * feeDenominator + amountInWithFee);
    }

    function getProfit(
        uint256 amountIn,
        uint256[] memory r0,
        uint256[] memory r1,
        uint256[] memory feeN,
        uint256[] memory feeD
    ) internal pure returns (uint256) {
        uint256 out = amountIn;
        for (uint256 i = 0; i < r0.length; i++) {
            out = getAmountOut(out, r0[i], r1[i], feeN[i], feeD[i]);
        }
        return out > amountIn ? out - amountIn : 0;
    }

    function findBestInput(
        uint256[] memory r0,
        uint256[] memory r1,
        uint256[] memory feeN,
        uint256[] memory feeD,
        uint256 maxIn,
        uint256 step
    ) internal pure returns (uint256 bestIn, uint256 bestProfit) {
        for (uint256 i = step; i <= maxIn; i += step) {
            uint256 p = getProfit(i, r0, r1, feeN, feeD);
            if (p > bestProfit) {
                bestProfit = p;
                bestIn = i;
            }
        }
    }
}

contract MultiArbitrageExecutor {
    using MultiArbitrageLib for uint256[];

    event TradeExecuted(uint256 amountIn, uint256 profit);

    function execute(
        address[] calldata pairs,
        uint256[] calldata feeNumerators,
        uint256[] calldata feeDenominators,
        uint256 maxIn,
        uint256 step
    ) external {
        require(pairs.length >= 2, "need >=2 pairs");
        require(pairs.length == feeNumerators.length && pairs.length == feeDenominators.length, "fee length");

        uint256[] memory r0 = new uint256[](pairs.length);
        uint256[] memory r1 = new uint256[](pairs.length);
        address inputToken = IPair(pairs[0]).token0();
        for (uint256 i = 0; i < pairs.length; i++) {
            (uint112 a, uint112 b,) = IPair(pairs[i]).getReserves();
            r0[i] = a;
            r1[i] = b;
            if (i < pairs.length - 1) {
                require(IPair(pairs[i]).token1() == IPair(pairs[i+1]).token0(), "pair mismatch");
            } else {
                require(IPair(pairs[i]).token1() == inputToken, "cycle mismatch");
            }
        }

        (uint256 bestIn, uint256 profit) = MultiArbitrageLib.findBestInput(r0, r1, feeNumerators, feeDenominators, maxIn, step);
        require(profit > 0, "no profit");

        address currentToken = inputToken;
        MockERC20(currentToken).transferFrom(msg.sender, pairs[0], bestIn);
        uint256 amountIn = bestIn;
        for (uint256 i = 0; i < pairs.length; i++) {
            address tokenOut = IPair(pairs[i]).token1();
            uint256 amountOut = MultiArbitrageLib.getAmountOut(amountIn, r0[i], r1[i], feeNumerators[i], feeDenominators[i]);
            IPair(pairs[i]).swap(0, amountOut, i == pairs.length - 1 ? msg.sender : address(this), new bytes(0));
            if (i < pairs.length - 1) {
                MockERC20(tokenOut).transfer(pairs[i+1], amountOut);
            }
            amountIn = amountOut;
        }

        emit TradeExecuted(bestIn, profit);
    }
}

