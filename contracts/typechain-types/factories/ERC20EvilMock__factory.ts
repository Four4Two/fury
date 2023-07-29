/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type { ERC20EvilMock, ERC20EvilMockInterface } from "../ERC20EvilMock";

const _abi = [
  {
    inputs: [],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "target",
        type: "address",
      },
      {
        internalType: "address",
        name: "toAddr",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "attackLock",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "transfer",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
      {
        internalType: "address",
        name: "",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "transferFrom",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
];

const _bytecode =
  "0x608060405234801561001057600080fd5b50600080819055506105d1806100276000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806323b872dd14610046578063a9059cbb14610076578063a9532ab5146100a6575b600080fd5b610060600480360381019061005b919061036e565b6100c2565b60405161006d91906103dc565b60405180910390f35b610090600480360381019061008b91906103f7565b610181565b60405161009d91906103dc565b60405180910390f35b6100c060048036038101906100bb919061036e565b610220565b005b6000600360005410156101765760016000546100de9190610466565b6000819055503373ffffffffffffffffffffffffffffffffffffffff16637750c9f030600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518463ffffffff1660e01b8152600401610143939291906104da565b600060405180830381600087803b15801561015d57600080fd5b505af1158015610171573d6000803e3d6000fd5b505050505b600190509392505050565b60006003600054101561021657600160005461019d9190610466565b6000819055503373ffffffffffffffffffffffffffffffffffffffff16634202e90730858560016040518563ffffffff1660e01b81526004016101e39493929190610556565b600060405180830381600087803b1580156101fd57600080fd5b505af1158015610211573d6000803e3d6000fd5b505050505b6001905092915050565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff16637750c9f03084846040518463ffffffff1660e01b815260040161029e939291906104da565b600060405180830381600087803b1580156102b857600080fd5b505af11580156102cc573d6000803e3d6000fd5b50505050505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610305826102da565b9050919050565b610315816102fa565b811461032057600080fd5b50565b6000813590506103328161030c565b92915050565b6000819050919050565b61034b81610338565b811461035657600080fd5b50565b60008135905061036881610342565b92915050565b600080600060608486031215610387576103866102d5565b5b600061039586828701610323565b93505060206103a686828701610323565b92505060406103b786828701610359565b9150509250925092565b60008115159050919050565b6103d6816103c1565b82525050565b60006020820190506103f160008301846103cd565b92915050565b6000806040838503121561040e5761040d6102d5565b5b600061041c85828601610323565b925050602061042d85828601610359565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061047182610338565b915061047c83610338565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156104b1576104b0610437565b5b828201905092915050565b6104c5816102fa565b82525050565b6104d481610338565b82525050565b60006060820190506104ef60008301866104bc565b6104fc60208301856104bc565b61050960408301846104cb565b949350505050565b6000819050919050565b6000819050919050565b600061054061053b61053684610511565b61051b565b610338565b9050919050565b61055081610525565b82525050565b600060808201905061056b60008301876104bc565b61057860208301866104bc565b61058560408301856104cb565b6105926060830184610547565b9594505050505056fea26469706673582212201a25fd7a50f13b4c28cb0e3b516ed73bb9ca4ebe593e1344b834f411cc0c27d064736f6c63430008090033";

type ERC20EvilMockConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ERC20EvilMockConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ERC20EvilMock__factory extends ContractFactory {
  constructor(...args: ERC20EvilMockConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
    this.contractName = "ERC20EvilMock";
  }

  deploy(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ERC20EvilMock> {
    return super.deploy(overrides || {}) as Promise<ERC20EvilMock>;
  }
  getDeployTransaction(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  attach(address: string): ERC20EvilMock {
    return super.attach(address) as ERC20EvilMock;
  }
  connect(signer: Signer): ERC20EvilMock__factory {
    return super.connect(signer) as ERC20EvilMock__factory;
  }
  static readonly contractName: "ERC20EvilMock";
  public readonly contractName: "ERC20EvilMock";
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ERC20EvilMockInterface {
    return new utils.Interface(_abi) as ERC20EvilMockInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ERC20EvilMock {
    return new Contract(address, _abi, signerOrProvider) as ERC20EvilMock;
  }
}