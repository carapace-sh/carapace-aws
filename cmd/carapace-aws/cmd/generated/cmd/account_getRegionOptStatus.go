package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_getRegionOptStatusCmd = &cobra.Command{
	Use:   "get-region-opt-status",
	Short: "Retrieves the opt-in status of a particular Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_getRegionOptStatusCmd).Standalone()

	account_getRegionOptStatusCmd.Flags().String("account-id", "", "Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
	account_getRegionOptStatusCmd.Flags().String("region-name", "", "Specifies the Region-code for a given Region name (for example, `af-south-1`).")
	account_getRegionOptStatusCmd.MarkFlagRequired("region-name")
	accountCmd.AddCommand(account_getRegionOptStatusCmd)
}
