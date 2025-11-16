package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_deleteSyncConfigurationCmd = &cobra.Command{
	Use:   "delete-sync-configuration",
	Short: "Deletes the sync configuration for a specified repository and connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_deleteSyncConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_deleteSyncConfigurationCmd).Standalone()

		codeconnections_deleteSyncConfigurationCmd.Flags().String("resource-name", "", "The name of the Amazon Web Services resource associated with the sync configuration to be deleted.")
		codeconnections_deleteSyncConfigurationCmd.Flags().String("sync-type", "", "The type of sync configuration to be deleted.")
		codeconnections_deleteSyncConfigurationCmd.MarkFlagRequired("resource-name")
		codeconnections_deleteSyncConfigurationCmd.MarkFlagRequired("sync-type")
	})
	codeconnectionsCmd.AddCommand(codeconnections_deleteSyncConfigurationCmd)
}
