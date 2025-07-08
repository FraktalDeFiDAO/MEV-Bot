// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./ArbitrageCalculator.sol";
import "./MockERC20.sol";

interface IPair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function getReserves() external view returns (uint112, uint112, uint32);
    function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes calldata data) external;
}

contract ArbitrageExecutor {
    uint256 constant FEE_NUMERATOR = 997;
    uint256 constant FEE_DENOMINATOR = 1000;

    function _getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) internal pure returns (uint256) {
        uint256 amountInWithFee = amountIn * FEE_NUMERATOR;
        return (amountInWithFee * reserveOut) / (reserveIn * FEE_DENOMINATOR + amountInWithFee);
    }

    /// @notice execute arbitrage across two constant product pairs
    /// @param pairA token0-token1 pair
    /// @param pairB token1-token0 pair
    /// @param maxIn maximum input amount
    /// @param step step size for search
    function execute(address pairA, address pairB, uint256 maxIn, uint256 step) external {
        (uint112 a0, uint112 a1,) = IPair(pairA).getReserves();
        (uint112 b0, uint112 b1,) = IPair(pairB).getReserves();
        require(IPair(pairA).token0() == IPair(pairB).token1(), "pair mismatch");
        require(IPair(pairA).token1() == IPair(pairB).token0(), "pair mismatch");

        (uint256 bestIn, uint256 profit) = _findBestInput(a0, a1, b0, b1, maxIn, step);
        require(profit > 0, "no profit");

        address tokenA = IPair(pairA).token0();
        address tokenB = IPair(pairA).token1();

        MockERC20(tokenA).transferFrom(msg.sender, pairA, bestIn);
        uint256 out1 = _getAmountOut(bestIn, a0, a1);
        IPair(pairA).swap(0, out1, address(this), new bytes(0));

        MockERC20(tokenB).transfer(pairB, out1);
        uint256 out2 = _getAmountOut(out1, b0, b1);
        IPair(pairB).swap(0, out2, msg.sender, new bytes(0));
    }

    function _findBestInput(uint256 rA0, uint256 rA1, uint256 rB0, uint256 rB1, uint256 maxIn, uint256 step)
        internal
        pure
        returns (uint256 bestIn, uint256 bestProfit)
    {
        for (uint256 i = step; i <= maxIn; i += step) {
            uint256 out1 = _getAmountOut(i, rA0, rA1);
            uint256 out2 = _getAmountOut(out1, rB0, rB1);
            if (out2 > i) {
                uint256 profit = out2 - i;
                if (profit > bestProfit) {
                    bestProfit = profit;
                    bestIn = i;
                }
            }
        }
    }
}
