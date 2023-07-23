package testutil

import (
	"fmt"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

func init() {
	// read the .env file, if present
	gotenv.Load()
}

type SuiteConfig struct {
	// A funded account used to fnd all other accounts.
	FundedAccountMnemonic string
	// The fury.configTemplate flag to be passed to futool, usually "master".
	// This allows one to change the base genesis used to start the chain.
	FuryConfigTemplate string
	// Whether or not to start an IBC chain. Use `suite.SkipIfIbcDisabled()` in IBC tests in IBC tests.
	IncludeIbcTests bool

	// Whether or not to run a chain upgrade & run post-upgrade tests. Use `suite.SkipIfUpgradeDisabled()` in post-upgrade tests.
	IncludeAutomatedUpgrade bool
	// Name of the upgrade, if upgrade is enabled.
	FuryUpgradeName string
	// Height upgrade will be applied to the test chain, if upgrade is enabled.
	FuryUpgradeHeight int64
	// Tag of fury docker image that will be upgraded to the current image before tests are run, if upgrade is enabled.
	FuryUpgradeBaseImageTag string

	// The contract address of a deployed ERC-20 token
	FuryErc20Address string

	// When true, the chains will remain running after tests complete (pass or fail)
	SkipShutdown bool
}

func ParseSuiteConfig() SuiteConfig {
	config := SuiteConfig{
		// this mnemonic is expected to be a funded account that can seed the funds for all
		// new accounts created during tests. it will be available under Accounts["whale"]
		FundedAccountMnemonic:   nonemptyStringEnv("E2E_FURY_FUNDED_ACCOUNT_MNEMONIC"),
		FuryConfigTemplate:      nonemptyStringEnv("E2E_FUTOOL_FURY_CONFIG_TEMPLATE"),
		FuryErc20Address:        nonemptyStringEnv("E2E_FURY_ERC20_ADDRESS"),
		IncludeIbcTests:         mustParseBool("E2E_INCLUDE_IBC_TESTS"),
		IncludeAutomatedUpgrade: mustParseBool("E2E_INCLUDE_AUTOMATED_UPGRADE"),
	}

	skipShutdownEnv := os.Getenv("E2E_SKIP_SHUTDOWN")
	if skipShutdownEnv != "" {
		config.SkipShutdown = mustParseBool("E2E_SKIP_SHUTDOWN")
	}

	if config.IncludeAutomatedUpgrade {
		config.FuryUpgradeName = nonemptyStringEnv("E2E_FURY_UPGRADE_NAME")
		config.FuryUpgradeBaseImageTag = nonemptyStringEnv("E2E_FURY_UPGRADE_BASE_IMAGE_TAG")
		upgradeHeight, err := strconv.ParseInt(nonemptyStringEnv("E2E_FURY_UPGRADE_HEIGHT"), 10, 64)
		if err != nil {
			panic(fmt.Sprintf("E2E_FURY_UPGRADE_HEIGHT must be a number: %s", err))
		}
		config.FuryUpgradeHeight = upgradeHeight
	}

	return config
}

func mustParseBool(name string) bool {
	envValue := os.Getenv(name)
	if envValue == "" {
		panic(fmt.Sprintf("%s is unset but expected a bool", name))
	}
	value, err := strconv.ParseBool(envValue)
	if err != nil {
		panic(fmt.Sprintf("%s (%s) cannot be parsed to a bool: %s", name, envValue, err))
	}
	return value
}

func nonemptyStringEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("no %s env variable provided", name))
	}
	return value
}
