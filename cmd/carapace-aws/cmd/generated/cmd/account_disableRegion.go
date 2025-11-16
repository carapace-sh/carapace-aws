package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_disableRegionCmd = &cobra.Command{
	Use:   "disable-region",
	Short: "Disables (opts-out) a particular Region for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_disableRegionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(account_disableRegionCmd).Standalone()

		account_disableRegionCmd.Flags().String("account-id", "", "Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
		account_disableRegionCmd.Flags().String("region-name", "", "Specifies the Region-code for a given Region name (for example, `af-south-1`).")
		account_disableRegionCmd.MarkFlagRequired("region-name")
	})
	accountCmd.AddCommand(account_disableRegionCmd)
}
