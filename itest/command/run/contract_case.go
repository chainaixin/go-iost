package run

import (
	"github.com/iost-official/go-iost/itest"
	"github.com/urfave/cli"
)

// ContractCaseCommand is the command of contract test case
var ContractCaseCommand = cli.Command{
	Name:      "contract_case",
	ShortName: "c_case",
	Usage:     "run contract test case",
	Flags:     ContractCaseFlags,
	Action:    ContractCaseAction,
}

// ContractCaseFlags ...
var ContractCaseFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "account, a",
		Value: 10,
		Usage: "number of account",
	},
	cli.IntFlag{
		Name:  "transaction, t",
		Value: 1000,
		Usage: "number of transaction",
	},
}

// ContractCaseAction is the action of contract test case
var ContractCaseAction = func(c *cli.Context) error {
	anum := c.Int("account")
	tnum := c.Int("transaction")
	keysfile := c.GlobalString("keys")
	configfile := c.GlobalString("config")
	codefile := c.GlobalString("code")
	abifile := c.GlobalString("abi")

	it, err := itest.Load(keysfile, configfile)
	if err != nil {
		return err
	}

	contract, err := itest.LoadContract(codefile, abifile)
	if err != nil {
		return err
	}

	accounts, err := it.CreateAccountN(anum)
	if err != nil {
		return err
	}

	cid, err := it.SetContract(contract)
	if err != nil {
		return err
	}

	if err := it.ContractTransferN(cid, tnum, accounts); err != nil {
		return err
	}

	if err := it.CheckAccounts(accounts); err != nil {
		return err
	}

	return nil
}