// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./MockERC20.sol";

contract MockPair {
    address public token0;
    address public token1;
    uint112 private reserve0;
    uint112 private reserve1;

    constructor(address _token0, address _token1) {
        token0 = _token0;
        token1 = _token1;
    }

    function setReserves(uint112 r0, uint112 r1) external {
        if (MockERC20(token0).balanceOf(address(this)) < r0) {
            MockERC20(token0).mint(address(this), r0 - MockERC20(token0).balanceOf(address(this)));
        }
        if (MockERC20(token1).balanceOf(address(this)) < r1) {
            MockERC20(token1).mint(address(this), r1 - MockERC20(token1).balanceOf(address(this)));
        }
        reserve0 = r0;
        reserve1 = r1;
    }

    function getReserves() external view returns (uint112, uint112, uint32) {
        return (reserve0, reserve1, 0);
    }

    function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes calldata) external {
        if (amount0Out > 0) {
            require(amount0Out <= reserve0, "insufficient0");
            reserve0 -= uint112(amount0Out);
            MockERC20(token0).transfer(to, amount0Out);
        }
        if (amount1Out > 0) {
            require(amount1Out <= reserve1, "insufficient1");
            reserve1 -= uint112(amount1Out);
            MockERC20(token1).transfer(to, amount1Out);
        }
        reserve0 = uint112(MockERC20(token0).balanceOf(address(this)));
        reserve1 = uint112(MockERC20(token1).balanceOf(address(this)));
    }
}
