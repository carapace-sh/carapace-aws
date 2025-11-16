package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_updateSyncConfigurationCmd = &cobra.Command{
	Use:   "update-sync-configuration",
	Short: "Updates the sync configuration for your connection and a specified external Git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_updateSyncConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_updateSyncConfigurationCmd).Standalone()

		codestarConnections_updateSyncConfigurationCmd.Flags().String("branch", "", "The branch for the sync configuration to be updated.")
		codestarConnections_updateSyncConfigurationCmd.Flags().String("config-file", "", "The configuration file for the sync configuration to be updated.")
		codestarConnections_updateSyncConfigurationCmd.Flags().String("publish-deployment-status", "", "Whether to enable or disable publishing of deployment status to source providers.")
		codestarConnections_updateSyncConfigurationCmd.Flags().String("repository-link-id", "", "The ID of the repository link for the sync configuration to be updated.")
		codestarConnections_updateSyncConfigurationCmd.Flags().String("resource-name", "", "The name of the Amazon Web Services resource for the sync configuration to be updated.")
		codestarConnections_updateSyncConfigurationCmd.Flags().String("role-arn", "", "The ARN of the IAM role for the sync configuration to be updated.")
		codestarConnections_updateSyncConfigurationCmd.Flags().String("sync-type", "", "The sync type for the sync configuration to be updated.")
		codestarConnections_updateSyncConfigurationCmd.Flags().String("trigger-resource-update-on", "", "When to trigger Git sync to begin the stack update.")
		codestarConnections_updateSyncConfigurationCmd.MarkFlagRequired("resource-name")
		codestarConnections_updateSyncConfigurationCmd.MarkFlagRequired("sync-type")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_updateSyncConfigurationCmd)
}
