import { ethers } from "hardhat";

async function main() {
  const tokenName = "Jinx";
  const tokenSymbol = "JINX";
  const tokenDecimals = 18;

  const ERC20FuryWrappedCosmosCoin = await ethers.getContractFactory("Jinx");
  const token = await ERC20FuryWrappedCosmosCoin.deploy(
    tokenName,
    tokenSymbol,
    tokenDecimals
  );

  await token.deployed();

  console.log(
    `Token "${tokenName}" (${tokenSymbol}) with ${tokenDecimals} decimals is deployed to ${token.address}!`
  );
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
