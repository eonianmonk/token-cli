const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const TestTokenModule = buildModule("TestTokenModule", (m) => {
    const token = m.contract("TestToken");
    return {token};
})

module.exports = TestTokenModule;