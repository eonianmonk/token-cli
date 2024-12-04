const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("TestToken", async () => {
  let token, owner, addr1, addr2;

  beforeEach(async () => {
    const TokenContract = await ethers.getContractFactory("TestToken");
    token = await TokenContract.deploy();
    [owner, addr1, addr2] = await ethers.getSigners();
  });

  // Mint Tests
  it("Should mint tokens to an address", async () => {
    await token.mint(addr1.address, 1000);
    expect(await token.balanceOf(addr1.address)).to.equal(1000);
  });

  it("Should prevent non-owner from minting", async () => {
    await expect(
      token.connect(addr1).mint(addr2.address, 1000)
    ).to.be.revertedWithCustomError(token, "OwnableUnauthorizedAccount").withArgs(addr1);
  });

  it("Shouldn't mint to 0 address", async () => {
    await expect(
      token.mint(ethers.ZeroAddress, 1)
    ).to.be.revertedWith("minting to zero address is not allowed")
  })

  it("Shouldn't mint to frozen account", async () => {
    await token.freeze(addr1.address);
    await expect(
      token.mint(addr1.address, 10)
    ).to.be.revertedWith("account is frozen");
  })

  // Transfer Tests
  it("Should transfer tokens between accounts", async () => {
    await token.mint(addr1.address, 1000);
    await token.connect(addr1).transfer(addr2.address, 500);
    // check new balances
    expect(await token.balanceOf(addr1.address)).to.equal(500);
    expect(await token.balanceOf(addr2.address)).to.equal(500);
  });

  it("Should prevent from transfering more than available", async () => {
    await token.mint(addr1.address, 1000);
    await expect(
      token.connect(addr1).transfer(addr2.address, 1500)
    ).to.be.revertedWith("insufficient balance for transfer amount")
  })

  it("Should be able to transfer to 0 address (burn)", async () => {
    await token.mint(addr1.address, 1000);
    await token.connect(addr1).transfer(ethers.ZeroAddress, 1000);
    expect( await token.balanceOf(addr1.address)).to.equal(0); 
  })

  // Balance Tests
  it("Shouldn't be able to check 0-address balance", async ()=> {
    await expect(
      token.balanceOf(ethers.ZeroAddress)
    ).to.be.revertedWith("invalid account address")
  })

  // Freeze Tests
  it("Shouldn't transfer from frozen accounts", async ()=> {
    await token.mint(addr1.address, 1000);
    await token.freeze(addr1.address);
    await expect(
      token.connect(addr1).transfer(addr2.address, 500)
    ).to.be.revertedWith("account is frozen");
  });

  it("Should allow transfers after unfreezing", async () => {
    await token.mint(addr1.address, 1000);
    await token.freeze(addr1.address);
    await expect(
      token.connect(addr1).transfer(addr2.address, 500)
    ).to.be.revertedWith("account is frozen");
    await token.unfreeze(addr1.address);
    await token.connect(addr1).transfer(addr2.address, 500);
    expect(await token.balanceOf(addr2.address)).to.equal(500);
  });
});