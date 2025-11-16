package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createAccountPoolCmd = &cobra.Command{
	Use:   "create-account-pool",
	Short: "Creates an account pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createAccountPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createAccountPoolCmd).Standalone()

		datazone_createAccountPoolCmd.Flags().String("account-source", "", "The source of accounts for the account pool.")
		datazone_createAccountPoolCmd.Flags().String("description", "", "The description of the account pool.")
		datazone_createAccountPoolCmd.Flags().String("domain-identifier", "", "The ID of the domain where the account pool is created.")
		datazone_createAccountPoolCmd.Flags().String("name", "", "The name of the account pool.")
		datazone_createAccountPoolCmd.Flags().String("resolution-strategy", "", "The mechanism used to resolve the account selection from the account pool.")
		datazone_createAccountPoolCmd.MarkFlagRequired("account-source")
		datazone_createAccountPoolCmd.MarkFlagRequired("domain-identifier")
		datazone_createAccountPoolCmd.MarkFlagRequired("name")
		datazone_createAccountPoolCmd.MarkFlagRequired("resolution-strategy")
	})
	datazoneCmd.AddCommand(datazone_createAccountPoolCmd)
}
