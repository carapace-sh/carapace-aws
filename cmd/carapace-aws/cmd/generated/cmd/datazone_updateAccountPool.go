package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateAccountPoolCmd = &cobra.Command{
	Use:   "update-account-pool",
	Short: "Updates the account pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateAccountPoolCmd).Standalone()

	datazone_updateAccountPoolCmd.Flags().String("account-source", "", "The source of accounts for the account pool.")
	datazone_updateAccountPoolCmd.Flags().String("description", "", "The description of the account pool that is to be udpated.")
	datazone_updateAccountPoolCmd.Flags().String("domain-identifier", "", "The domain ID where the account pool that is to be updated lives.")
	datazone_updateAccountPoolCmd.Flags().String("identifier", "", "The ID of the account pool that is to be updated.")
	datazone_updateAccountPoolCmd.Flags().String("name", "", "The name of the account pool that is to be updated.")
	datazone_updateAccountPoolCmd.Flags().String("resolution-strategy", "", "The mechanism used to resolve the account selection from the account pool.")
	datazone_updateAccountPoolCmd.MarkFlagRequired("domain-identifier")
	datazone_updateAccountPoolCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_updateAccountPoolCmd)
}
