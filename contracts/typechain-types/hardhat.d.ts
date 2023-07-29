/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { ethers } from "ethers";
import {
  FactoryOptions,
  HardhatEthersHelpers as HardhatEthersHelpersBase,
} from "@nomiclabs/hardhat-ethers/types";

import * as Contracts from ".";

declare module "hardhat/types/runtime" {
  interface HardhatEthersHelpers extends HardhatEthersHelpersBase {
    getContractFactory(
      name: "Ownable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Ownable__factory>;
    getContractFactory(
      name: "ERC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20__factory>;
    getContractFactory(
      name: "IERC20Metadata",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC20Metadata__factory>;
    getContractFactory(
      name: "IERC20Permit",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC20Permit__factory>;
    getContractFactory(
      name: "IERC20",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.IERC20__factory>;
    getContractFactory(
      name: "Bridge",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Bridge__factory>;
    getContractFactory(
      name: "ERC20MintableBurnable",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20MintableBurnable__factory>;
    getContractFactory(
      name: "ERC20EvilMock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20EvilMock__factory>;
    getContractFactory(
      name: "ERC20Mock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20Mock__factory>;
    getContractFactory(
      name: "ERC20NoReturnMock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20NoReturnMock__factory>;
    getContractFactory(
      name: "ERC20ReturnFalseMock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20ReturnFalseMock__factory>;
    getContractFactory(
      name: "ERC20ReturnTrueMock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.ERC20ReturnTrueMock__factory>;
    getContractFactory(
      name: "SequenceMock",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.SequenceMock__factory>;
    getContractFactory(
      name: "Withdrawer",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Withdrawer__factory>;
    getContractFactory(
      name: "Multicall",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Multicall__factory>;
    getContractFactory(
      name: "Multicall2",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.Multicall2__factory>;
    getContractFactory(
      name: "WETH9",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.WETH9__factory>;
    getContractFactory(
      name: "WFURY",
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<Contracts.WFURY__factory>;

    getContractAt(
      name: "Ownable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Ownable>;
    getContractAt(
      name: "ERC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20>;
    getContractAt(
      name: "IERC20Metadata",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC20Metadata>;
    getContractAt(
      name: "IERC20Permit",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC20Permit>;
    getContractAt(
      name: "IERC20",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.IERC20>;
    getContractAt(
      name: "Bridge",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Bridge>;
    getContractAt(
      name: "ERC20MintableBurnable",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20MintableBurnable>;
    getContractAt(
      name: "ERC20EvilMock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20EvilMock>;
    getContractAt(
      name: "ERC20Mock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20Mock>;
    getContractAt(
      name: "ERC20NoReturnMock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20NoReturnMock>;
    getContractAt(
      name: "ERC20ReturnFalseMock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20ReturnFalseMock>;
    getContractAt(
      name: "ERC20ReturnTrueMock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.ERC20ReturnTrueMock>;
    getContractAt(
      name: "SequenceMock",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.SequenceMock>;
    getContractAt(
      name: "Withdrawer",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Withdrawer>;
    getContractAt(
      name: "Multicall",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Multicall>;
    getContractAt(
      name: "Multicall2",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.Multicall2>;
    getContractAt(
      name: "WETH9",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.WETH9>;
    getContractAt(
      name: "WFURY",
      address: string,
      signer?: ethers.Signer
    ): Promise<Contracts.WFURY>;

    // default types
    getContractFactory(
      name: string,
      signerOrOptions?: ethers.Signer | FactoryOptions
    ): Promise<ethers.ContractFactory>;
    getContractFactory(
      abi: any[],
      bytecode: ethers.utils.BytesLike,
      signer?: ethers.Signer
    ): Promise<ethers.ContractFactory>;
    getContractAt(
      nameOrAbi: string | any[],
      address: string,
      signer?: ethers.Signer
    ): Promise<ethers.Contract>;
  }
}
