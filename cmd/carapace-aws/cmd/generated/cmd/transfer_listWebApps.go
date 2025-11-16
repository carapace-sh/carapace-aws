package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listWebAppsCmd = &cobra.Command{
	Use:   "list-web-apps",
	Short: "Lists all web apps associated with your Amazon Web Services account for your current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listWebAppsCmd).Standalone()

	transfer_listWebAppsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	transfer_listWebAppsCmd.Flags().String("next-token", "", "Returns the `NextToken` parameter in the output.")
	transferCmd.AddCommand(transfer_listWebAppsCmd)
}
