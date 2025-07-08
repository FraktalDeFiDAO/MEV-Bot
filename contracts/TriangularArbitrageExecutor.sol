// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./TriangularArbitrageCalculator.sol";
import "./MockERC20.sol";

interface IPair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function getReserves() external view returns (uint112, uint112, uint32);
    function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes calldata data) external;
}

contract TriangularArbitrageExecutor {
    uint256 constant FEE_NUMERATOR = 997;
    uint256 constant FEE_DENOMINATOR = 1000;

    function _getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) internal pure returns (uint256) {
        uint256 amountInWithFee = amountIn * FEE_NUMERATOR;
        return (amountInWithFee * reserveOut) / (reserveIn * FEE_DENOMINATOR + amountInWithFee);
    }

    function _findBestInput(
        uint256 ab0,
        uint256 ab1,
        uint256 bc0,
        uint256 bc1,
        uint256 ca0,
        uint256 ca1,
        uint256 maxIn,
        uint256 step
    ) internal pure returns (uint256 bestIn, uint256 bestProfit) {
        for (uint256 i = step; i <= maxIn; i += step) {
            uint256 p = TriangularArbitrageLib.getProfit(i, ab0, ab1, bc0, bc1, ca0, ca1);
            if (p > bestProfit) {
                bestProfit = p;
                bestIn = i;
            }
        }
    }

    /// @notice Execute a triangular arbitrage across three constant product pairs
    /// @param pairAB tokenA-tokenB pair
    /// @param pairBC tokenB-tokenC pair
    /// @param pairCA tokenC-tokenA pair
    /// @param maxIn maximum input amount
    /// @param step step size for search
    function execute(address pairAB, address pairBC, address pairCA, uint256 maxIn, uint256 step) external {
        (uint112 ab0, uint112 ab1,) = IPair(pairAB).getReserves();
        (uint112 bc0, uint112 bc1,) = IPair(pairBC).getReserves();
        (uint112 ca0, uint112 ca1,) = IPair(pairCA).getReserves();

        require(IPair(pairAB).token0() == IPair(pairCA).token1(), "pair mismatch");
        require(IPair(pairAB).token1() == IPair(pairBC).token0(), "pair mismatch");
        require(IPair(pairBC).token1() == IPair(pairCA).token0(), "pair mismatch");

        (uint256 bestIn, uint256 profit) = _findBestInput(ab0, ab1, bc0, bc1, ca0, ca1, maxIn, step);
        require(profit > 0, "no profit");

        address tokenA = IPair(pairAB).token0();
        MockERC20(tokenA).transferFrom(msg.sender, pairAB, bestIn);
        uint256 outAB = _getAmountOut(bestIn, ab0, ab1);
        IPair(pairAB).swap(0, outAB, address(this), new bytes(0));

        MockERC20(IPair(pairAB).token1()).transfer(pairBC, outAB);
        uint256 outBC = _getAmountOut(outAB, bc0, bc1);
        IPair(pairBC).swap(0, outBC, address(this), new bytes(0));

        MockERC20(IPair(pairBC).token1()).transfer(pairCA, outBC);
        uint256 outCA = _getAmountOut(outBC, ca0, ca1);
        IPair(pairCA).swap(0, outCA, msg.sender, new bytes(0));
    }
}
