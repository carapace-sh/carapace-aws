package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteResourceDataSyncCmd = &cobra.Command{
	Use:   "delete-resource-data-sync",
	Short: "Deletes a resource data sync configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteResourceDataSyncCmd).Standalone()

	ssm_deleteResourceDataSyncCmd.Flags().String("sync-name", "", "The name of the configuration to delete.")
	ssm_deleteResourceDataSyncCmd.Flags().String("sync-type", "", "Specify the type of resource data sync to delete.")
	ssm_deleteResourceDataSyncCmd.MarkFlagRequired("sync-name")
	ssmCmd.AddCommand(ssm_deleteResourceDataSyncCmd)
}
