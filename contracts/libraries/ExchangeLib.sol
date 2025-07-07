// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

library ExchangeLib {
    bytes32 constant EXCHANGE_STORAGE_POSITION = keccak256("mev.bot.exchange.storage");

    struct ExchangeInfo {
        string name;
        address router;
        bool enabled;
    }

    struct ExchangeStorage {
        mapping(uint256 => ExchangeInfo) info;
        uint256 nextId;
    }

    function exchangeStorage() internal pure returns (ExchangeStorage storage es) {
        bytes32 position = EXCHANGE_STORAGE_POSITION;
        assembly {
            es.slot := position
        }
    }

    function addExchange(string memory name, address router) internal returns (uint256 id) {
        require(router != address(0), "zero router");
        ExchangeStorage storage es = exchangeStorage();
        id = es.nextId++;
        es.info[id] = ExchangeInfo({name: name, router: router, enabled: true});
    }

    function setExchangeEnabled(uint256 id, bool enabled) internal {
        ExchangeStorage storage es = exchangeStorage();
        require(bytes(es.info[id].name).length != 0, "exchange missing");
        es.info[id].enabled = enabled;
    }

    function getExchange(uint256 id) internal view returns (ExchangeInfo memory) {
        return exchangeStorage().info[id];
    }

    function getExchangeCount() internal view returns (uint256) {
        return exchangeStorage().nextId;
    }

    function getExchanges() internal view returns (ExchangeInfo[] memory exchanges) {
        ExchangeStorage storage es = exchangeStorage();
        exchanges = new ExchangeInfo[](es.nextId);
        for (uint256 i = 0; i < es.nextId; i++) {
            exchanges[i] = es.info[i];
        }
    }

}
