// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title Simple Counter contract
contract Counter {
    uint256 public number;

    /// @notice increment the counter
    function increment() external {
        number += 1;
    }

    /// @notice set the counter to a specific value
    function setNumber(uint256 newNumber) external {
        number = newNumber;
    }
}
