// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/BatchExecutor.sol";
import "../contracts/MockTarget.sol";

contract BatchExecutorTest is Test {
    BatchExecutor exec;
    MockTarget t1;
    MockTarget t2;

    function setUp() public {
        exec = new BatchExecutor();
        t1 = new MockTarget();
        t2 = new MockTarget();
    }

    function testExecute() public {
        address[] memory targets = new address[](2);
        bytes[] memory data = new bytes[](2);
        targets[0] = address(t1);
        targets[1] = address(t2);
        data[0] = abi.encodeWithSelector(MockTarget.setValue.selector, 1);
        data[1] = abi.encodeWithSelector(MockTarget.setValue.selector, 2);
        exec.execute(targets, data);
        assertEq(t1.value(), 1);
        assertEq(t2.value(), 2);
    }

    function testLengthMismatch() public {
        address[] memory targets = new address[](1);
        bytes[] memory data = new bytes[](0);
        vm.expectRevert("length mismatch");
        exec.execute(targets, data);
    }

    function testCallFailure() public {
        address[] memory targets = new address[](1);
        bytes[] memory data = new bytes[](1);
        targets[0] = address(t1);
        data[0] = hex"deadbeef";
        vm.expectRevert("call failed");
        exec.execute(targets, data);
    }
}
