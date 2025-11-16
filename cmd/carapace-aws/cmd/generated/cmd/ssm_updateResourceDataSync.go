package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateResourceDataSyncCmd = &cobra.Command{
	Use:   "update-resource-data-sync",
	Short: "Update a resource data sync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateResourceDataSyncCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_updateResourceDataSyncCmd).Standalone()

		ssm_updateResourceDataSyncCmd.Flags().String("sync-name", "", "The name of the resource data sync you want to update.")
		ssm_updateResourceDataSyncCmd.Flags().String("sync-source", "", "Specify information about the data sources to synchronize.")
		ssm_updateResourceDataSyncCmd.Flags().String("sync-type", "", "The type of resource data sync.")
		ssm_updateResourceDataSyncCmd.MarkFlagRequired("sync-name")
		ssm_updateResourceDataSyncCmd.MarkFlagRequired("sync-source")
		ssm_updateResourceDataSyncCmd.MarkFlagRequired("sync-type")
	})
	ssmCmd.AddCommand(ssm_updateResourceDataSyncCmd)
}
