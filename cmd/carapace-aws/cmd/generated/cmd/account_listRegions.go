package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_listRegionsCmd = &cobra.Command{
	Use:   "list-regions",
	Short: "Lists all the Regions for a given account and their respective opt-in statuses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_listRegionsCmd).Standalone()

	account_listRegionsCmd.Flags().String("account-id", "", "Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
	account_listRegionsCmd.Flags().String("max-results", "", "The total number of items to return in the command’s output.")
	account_listRegionsCmd.Flags().String("next-token", "", "A token used to specify where to start paginating.")
	account_listRegionsCmd.Flags().String("region-opt-status-contains", "", "A list of Region statuses (Enabling, Enabled, Disabling, Disabled, Enabled\\_by\\_default) to use to filter the list of Regions for a given account.")
	accountCmd.AddCommand(account_listRegionsCmd)
}
