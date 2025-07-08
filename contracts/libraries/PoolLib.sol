// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

library PoolLib {
    bytes32 constant POOL_STORAGE_POSITION = keccak256("mev.bot.pool.storage");

    struct PoolInfo {
        address token0;
        address token1;
        uint256 exchangeId;
        bool enabled;
    }

    struct PoolStorage {
        mapping(address => PoolInfo) info;
        address[] pools;
    }

    function poolStorage() internal pure returns (PoolStorage storage ps) {
        bytes32 position = POOL_STORAGE_POSITION;
        assembly {
            ps.slot := position
        }
    }

    function addPool(address pool, address token0, address token1, uint256 exchangeId) internal {
        PoolStorage storage ps = poolStorage();
        require(pool != address(0), "zero pool");
        if (ps.info[pool].token0 == address(0)) {
            ps.pools.push(pool);
        }
        ps.info[pool] = PoolInfo({token0: token0, token1: token1, exchangeId: exchangeId, enabled: true});
    }

    function setPoolEnabled(address pool, bool enabled) internal {
        PoolStorage storage ps = poolStorage();
        require(ps.info[pool].token0 != address(0), "pool missing");
        ps.info[pool].enabled = enabled;
    }

    function isPoolEnabled(address pool) internal view returns (bool) {
        return poolStorage().info[pool].enabled;
    }

    function getPool(address pool) internal view returns (PoolInfo memory) {
        return poolStorage().info[pool];
    }

    function getPools() internal view returns (address[] memory pools) {
        PoolStorage storage ps = poolStorage();
        pools = new address[](ps.pools.length);
        for (uint256 i = 0; i < ps.pools.length; i++) {
            pools[i] = ps.pools[i];
        }
    }

    function getPoolCount() internal view returns (uint256) {
        return poolStorage().pools.length;
    }
}
