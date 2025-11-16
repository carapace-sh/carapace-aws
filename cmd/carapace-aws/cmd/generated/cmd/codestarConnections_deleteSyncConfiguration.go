package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_deleteSyncConfigurationCmd = &cobra.Command{
	Use:   "delete-sync-configuration",
	Short: "Deletes the sync configuration for a specified repository and connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_deleteSyncConfigurationCmd).Standalone()

	codestarConnections_deleteSyncConfigurationCmd.Flags().String("resource-name", "", "The name of the Amazon Web Services resource associated with the sync configuration to be deleted.")
	codestarConnections_deleteSyncConfigurationCmd.Flags().String("sync-type", "", "The type of sync configuration to be deleted.")
	codestarConnections_deleteSyncConfigurationCmd.MarkFlagRequired("resource-name")
	codestarConnections_deleteSyncConfigurationCmd.MarkFlagRequired("sync-type")
	codestarConnectionsCmd.AddCommand(codestarConnections_deleteSyncConfigurationCmd)
}
