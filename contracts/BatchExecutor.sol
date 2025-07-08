// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title BatchExecutor
/// @notice Executes a sequence of arbitrary calls. Used by the bot to dispatch
/// multiple swaps or actions in a single transaction.
contract BatchExecutor {
    /// @notice Execute a batch of calls
    /// @param targets The target addresses for each call
    /// @param data Calldata for each call
    function execute(address[] calldata targets, bytes[] calldata data) external {
        require(targets.length == data.length, "length mismatch");
        for (uint256 i = 0; i < targets.length; i++) {
            (bool success,) = targets[i].call(data[i]);
            require(success, "call failed");
        }
    }
}
