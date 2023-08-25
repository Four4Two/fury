import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.18",
    settings: {
      // istanbul upgrade occurred before the london jinxfork, so is compatible with fury's evm
      evmVersion: "istanbul",
      // optimize build for deployment to mainnet!
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    },
  },
  networks: {
    // futool's local network
    fury: {
      url: "http://127.0.0.1:8545",
      accounts: [
        // fury keys unsafe-export-eth-key whale2
        "2E6CBA62B30E51E805D8F4D5999E06B65B4A3BF1520D660782E775073F518E1B",
      ],
    },
    protonet: {
      url: "https://evm.app.protonet.us-east.production.fury.io:443",
      accounts: [
        "bfb5222ff6f3a553d3d74c160d744311315f56728977828e00755521a0bb9768",
      ],
    },
    internal_testnet: {
      url: "https://evm.data.internal.testnet.us-east.production.fury.io:443",
      accounts: [
        "bfb5222ff6f3a553d3d74c160d744311315f56728977828e00755521a0bb9768",
      ],
    },
  },
};

export default config;
