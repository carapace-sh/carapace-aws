package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listResourceDataSyncCmd = &cobra.Command{
	Use:   "list-resource-data-sync",
	Short: "Lists your resource data sync configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listResourceDataSyncCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listResourceDataSyncCmd).Standalone()

		ssm_listResourceDataSyncCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listResourceDataSyncCmd.Flags().String("next-token", "", "A token to start the list.")
		ssm_listResourceDataSyncCmd.Flags().String("sync-type", "", "View a list of resource data syncs according to the sync type.")
	})
	ssmCmd.AddCommand(ssm_listResourceDataSyncCmd)
}
