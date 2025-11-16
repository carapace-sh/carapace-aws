package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_createResourceDataSyncCmd = &cobra.Command{
	Use:   "create-resource-data-sync",
	Short: "A resource data sync helps you view data from multiple sources in a single location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_createResourceDataSyncCmd).Standalone()

	ssm_createResourceDataSyncCmd.Flags().String("s3-destination", "", "Amazon S3 configuration details for the sync.")
	ssm_createResourceDataSyncCmd.Flags().String("sync-name", "", "A name for the configuration.")
	ssm_createResourceDataSyncCmd.Flags().String("sync-source", "", "Specify information about the data sources to synchronize.")
	ssm_createResourceDataSyncCmd.Flags().String("sync-type", "", "Specify `SyncToDestination` to create a resource data sync that synchronizes data to an S3 bucket for Inventory.")
	ssm_createResourceDataSyncCmd.MarkFlagRequired("sync-name")
	ssmCmd.AddCommand(ssm_createResourceDataSyncCmd)
}
