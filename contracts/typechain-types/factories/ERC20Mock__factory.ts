/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Signer,
  utils,
  Contract,
  ContractFactory,
  PayableOverrides,
  BigNumberish,
} from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type { ERC20Mock, ERC20MockInterface } from "../ERC20Mock";

const _abi = [
  {
    inputs: [
      {
        internalType: "string",
        name: "name",
        type: "string",
      },
      {
        internalType: "string",
        name: "symbol",
        type: "string",
      },
      {
        internalType: "address",
        name: "initialAccount",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "initialBalance",
        type: "uint256",
      },
    ],
    stateMutability: "payable",
    type: "constructor",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "spender",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "value",
        type: "uint256",
      },
    ],
    name: "Approval",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "value",
        type: "uint256",
      },
    ],
    name: "Transfer",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        internalType: "address",
        name: "spender",
        type: "address",
      },
    ],
    name: "allowance",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "spender",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "approve",
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
        name: "owner",
        type: "address",
      },
      {
        internalType: "address",
        name: "spender",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "value",
        type: "uint256",
      },
    ],
    name: "approveInternal",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "account",
        type: "address",
      },
    ],
    name: "balanceOf",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "account",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "burn",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "decimals",
    outputs: [
      {
        internalType: "uint8",
        name: "",
        type: "uint8",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "spender",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "subtractedValue",
        type: "uint256",
      },
    ],
    name: "decreaseAllowance",
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
        name: "spender",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "addedValue",
        type: "uint256",
      },
    ],
    name: "increaseAllowance",
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
        name: "account",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "mint",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "name",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "symbol",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "totalSupply",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
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
        name: "from",
        type: "address",
      },
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
  {
    inputs: [
      {
        internalType: "address",
        name: "from",
        type: "address",
      },
      {
        internalType: "address",
        name: "to",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "value",
        type: "uint256",
      },
    ],
    name: "transferInternal",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
];

const _bytecode =
  "0x608060405260405162001f4e38038062001f4e8339818101604052810190620000299190620004e0565b8383816003908051906020019062000043929190620001f3565b5080600490805190602001906200005c929190620001f3565b5050506200007182826200007b60201b60201c565b5050505062000732565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620000ee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000e590620005f1565b60405180910390fd5b6200010260008383620001e960201b60201c565b806002600082825462000116919062000642565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620001c99190620006b0565b60405180910390a3620001e560008383620001ee60201b60201c565b5050565b505050565b505050565b8280546200020190620006fc565b90600052602060002090601f01602090048101928262000225576000855562000271565b82601f106200024057805160ff191683800117855562000271565b8280016001018555821562000271579182015b828111156200027057825182559160200191906001019062000253565b5b50905062000280919062000284565b5090565b5b808211156200029f57600081600090555060010162000285565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200030c82620002c1565b810181811067ffffffffffffffff821117156200032e576200032d620002d2565b5b80604052505050565b600062000343620002a3565b905062000351828262000301565b919050565b600067ffffffffffffffff821115620003745762000373620002d2565b5b6200037f82620002c1565b9050602081019050919050565b60005b83811015620003ac5780820151818401526020810190506200038f565b83811115620003bc576000848401525b50505050565b6000620003d9620003d38462000356565b62000337565b905082815260208101848484011115620003f857620003f7620002bc565b5b620004058482856200038c565b509392505050565b600082601f830112620004255762000424620002b7565b5b815162000437848260208601620003c2565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200046d8262000440565b9050919050565b6200047f8162000460565b81146200048b57600080fd5b50565b6000815190506200049f8162000474565b92915050565b6000819050919050565b620004ba81620004a5565b8114620004c657600080fd5b50565b600081519050620004da81620004af565b92915050565b60008060008060808587031215620004fd57620004fc620002ad565b5b600085015167ffffffffffffffff8111156200051e576200051d620002b2565b5b6200052c878288016200040d565b945050602085015167ffffffffffffffff81111562000550576200054f620002b2565b5b6200055e878288016200040d565b935050604062000571878288016200048e565b92505060606200058487828801620004c9565b91505092959194509250565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000620005d9601f8362000590565b9150620005e682620005a1565b602082019050919050565b600060208201905081810360008301526200060c81620005ca565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200064f82620004a5565b91506200065c83620004a5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562000694576200069362000613565b5b828201905092915050565b620006aa81620004a5565b82525050565b6000602082019050620006c760008301846200069f565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200071557607f821691505b602082108114156200072c576200072b620006cd565b5b50919050565b61180c80620007426000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806340c10f19116100975780639dc29fac116100665780639dc29fac14610286578063a457c2d7146102a2578063a9059cbb146102d2578063dd62ed3e14610302576100f5565b806340c10f191461020057806356189cb41461021c57806370a082311461023857806395d89b4114610268576100f5565b8063222f5be0116100d3578063222f5be01461016657806323b872dd14610182578063313ce567146101b257806339509351146101d0576100f5565b806306fdde03146100fa578063095ea7b31461011857806318160ddd14610148575b600080fd5b610102610332565b60405161010f9190610f36565b60405180910390f35b610132600480360381019061012d9190610ff1565b6103c4565b60405161013f919061104c565b60405180910390f35b6101506103e7565b60405161015d9190611076565b60405180910390f35b610180600480360381019061017b9190611091565b6103f1565b005b61019c60048036038101906101979190611091565b610401565b6040516101a9919061104c565b60405180910390f35b6101ba610430565b6040516101c79190611100565b60405180910390f35b6101ea60048036038101906101e59190610ff1565b610439565b6040516101f7919061104c565b60405180910390f35b61021a60048036038101906102159190610ff1565b610470565b005b61023660048036038101906102319190611091565b61047e565b005b610252600480360381019061024d919061111b565b61048e565b60405161025f9190611076565b60405180910390f35b6102706104d6565b60405161027d9190610f36565b60405180910390f35b6102a0600480360381019061029b9190610ff1565b610568565b005b6102bc60048036038101906102b79190610ff1565b610576565b6040516102c9919061104c565b60405180910390f35b6102ec60048036038101906102e79190610ff1565b6105ed565b6040516102f9919061104c565b60405180910390f35b61031c60048036038101906103179190611148565b610610565b6040516103299190611076565b60405180910390f35b606060038054610341906111b7565b80601f016020809104026020016040519081016040528092919081815260200182805461036d906111b7565b80156103ba5780601f1061038f576101008083540402835291602001916103ba565b820191906000526020600020905b81548152906001019060200180831161039d57829003601f168201915b5050505050905090565b6000806103cf610697565b90506103dc81858561069f565b600191505092915050565b6000600254905090565b6103fc83838361086a565b505050565b60008061040c610697565b9050610419858285610ae2565b61042485858561086a565b60019150509392505050565b60006012905090565b600080610444610697565b90506104658185856104568589610610565b6104609190611218565b61069f565b600191505092915050565b61047a8282610b6e565b5050565b61048983838361069f565b505050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546104e5906111b7565b80601f0160208091040260200160405190810160405280929190818152602001828054610511906111b7565b801561055e5780601f106105335761010080835404028352916020019161055e565b820191906000526020600020905b81548152906001019060200180831161054157829003601f168201915b5050505050905090565b6105728282610cc5565b5050565b600080610581610697565b9050600061058f8286610610565b9050838110156105d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105cb906112e0565b60405180910390fd5b6105e1828686840361069f565b60019250505092915050565b6000806105f8610697565b905061060581858561086a565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561070f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070690611372565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561077f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077690611404565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161085d9190611076565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156108da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d190611496565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561094a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094190611528565b60405180910390fd5b610955838383610e93565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156109db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d2906115ba565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ac99190611076565b60405180910390a3610adc848484610e98565b50505050565b6000610aee8484610610565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610b685781811015610b5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5190611626565b60405180910390fd5b610b67848484840361069f565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610bde576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bd590611692565b60405180910390fd5b610bea60008383610e93565b8060026000828254610bfc9190611218565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610cad9190611076565b60405180910390a3610cc160008383610e98565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610d35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2c90611724565b60405180910390fd5b610d4182600083610e93565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610dc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dbe906117b6565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e7a9190611076565b60405180910390a3610e8e83600084610e98565b505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610ed7578082015181840152602081019050610ebc565b83811115610ee6576000848401525b50505050565b6000601f19601f8301169050919050565b6000610f0882610e9d565b610f128185610ea8565b9350610f22818560208601610eb9565b610f2b81610eec565b840191505092915050565b60006020820190508181036000830152610f508184610efd565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f8882610f5d565b9050919050565b610f9881610f7d565b8114610fa357600080fd5b50565b600081359050610fb581610f8f565b92915050565b6000819050919050565b610fce81610fbb565b8114610fd957600080fd5b50565b600081359050610feb81610fc5565b92915050565b6000806040838503121561100857611007610f58565b5b600061101685828601610fa6565b925050602061102785828601610fdc565b9150509250929050565b60008115159050919050565b61104681611031565b82525050565b6000602082019050611061600083018461103d565b92915050565b61107081610fbb565b82525050565b600060208201905061108b6000830184611067565b92915050565b6000806000606084860312156110aa576110a9610f58565b5b60006110b886828701610fa6565b93505060206110c986828701610fa6565b92505060406110da86828701610fdc565b9150509250925092565b600060ff82169050919050565b6110fa816110e4565b82525050565b600060208201905061111560008301846110f1565b92915050565b60006020828403121561113157611130610f58565b5b600061113f84828501610fa6565b91505092915050565b6000806040838503121561115f5761115e610f58565b5b600061116d85828601610fa6565b925050602061117e85828601610fa6565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806111cf57607f821691505b602082108114156111e3576111e2611188565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061122382610fbb565b915061122e83610fbb565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611263576112626111e9565b5b828201905092915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006112ca602583610ea8565b91506112d58261126e565b604082019050919050565b600060208201905081810360008301526112f9816112bd565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061135c602483610ea8565b915061136782611300565b604082019050919050565b6000602082019050818103600083015261138b8161134f565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006113ee602283610ea8565b91506113f982611392565b604082019050919050565b6000602082019050818103600083015261141d816113e1565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611480602583610ea8565b915061148b82611424565b604082019050919050565b600060208201905081810360008301526114af81611473565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611512602383610ea8565b915061151d826114b6565b604082019050919050565b6000602082019050818103600083015261154181611505565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006115a4602683610ea8565b91506115af82611548565b604082019050919050565b600060208201905081810360008301526115d381611597565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611610601d83610ea8565b915061161b826115da565b602082019050919050565b6000602082019050818103600083015261163f81611603565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b600061167c601f83610ea8565b915061168782611646565b602082019050919050565b600060208201905081810360008301526116ab8161166f565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b600061170e602183610ea8565b9150611719826116b2565b604082019050919050565b6000602082019050818103600083015261173d81611701565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b60006117a0602283610ea8565b91506117ab82611744565b604082019050919050565b600060208201905081810360008301526117cf81611793565b905091905056fea2646970667358221220f568b25606cb08a9aa716723ce604792431c8e5639a0738bb308e0447b94bca264736f6c63430008090033";

type ERC20MockConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ERC20MockConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ERC20Mock__factory extends ContractFactory {
  constructor(...args: ERC20MockConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
    this.contractName = "ERC20Mock";
  }

  deploy(
    name: string,
    symbol: string,
    initialAccount: string,
    initialBalance: BigNumberish,
    overrides?: PayableOverrides & { from?: string | Promise<string> }
  ): Promise<ERC20Mock> {
    return super.deploy(
      name,
      symbol,
      initialAccount,
      initialBalance,
      overrides || {}
    ) as Promise<ERC20Mock>;
  }
  getDeployTransaction(
    name: string,
    symbol: string,
    initialAccount: string,
    initialBalance: BigNumberish,
    overrides?: PayableOverrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      name,
      symbol,
      initialAccount,
      initialBalance,
      overrides || {}
    );
  }
  attach(address: string): ERC20Mock {
    return super.attach(address) as ERC20Mock;
  }
  connect(signer: Signer): ERC20Mock__factory {
    return super.connect(signer) as ERC20Mock__factory;
  }
  static readonly contractName: "ERC20Mock";
  public readonly contractName: "ERC20Mock";
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ERC20MockInterface {
    return new utils.Interface(_abi) as ERC20MockInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ERC20Mock {
    return new Contract(address, _abi, signerOrProvider) as ERC20Mock;
  }
}
