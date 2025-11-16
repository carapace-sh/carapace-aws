package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_createSyncConfigurationCmd = &cobra.Command{
	Use:   "create-sync-configuration",
	Short: "Creates a sync configuration which allows Amazon Web Services to sync content from a Git repository to update a specified Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_createSyncConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_createSyncConfigurationCmd).Standalone()

		codestarConnections_createSyncConfigurationCmd.Flags().String("branch", "", "The branch in the repository from which changes will be synced.")
		codestarConnections_createSyncConfigurationCmd.Flags().String("config-file", "", "The file name of the configuration file that manages syncing between the connection and the repository.")
		codestarConnections_createSyncConfigurationCmd.Flags().String("publish-deployment-status", "", "Whether to enable or disable publishing of deployment status to source providers.")
		codestarConnections_createSyncConfigurationCmd.Flags().String("repository-link-id", "", "The ID of the repository link created for the connection.")
		codestarConnections_createSyncConfigurationCmd.Flags().String("resource-name", "", "The name of the Amazon Web Services resource (for example, a CloudFormation stack in the case of CFN\\_STACK\\_SYNC) that will be synchronized from the linked repository.")
		codestarConnections_createSyncConfigurationCmd.Flags().String("role-arn", "", "The ARN of the IAM role that grants permission for Amazon Web Services to use Git sync to update a given Amazon Web Services resource on your behalf.")
		codestarConnections_createSyncConfigurationCmd.Flags().String("sync-type", "", "The type of sync configuration.")
		codestarConnections_createSyncConfigurationCmd.Flags().String("trigger-resource-update-on", "", "When to trigger Git sync to begin the stack update.")
		codestarConnections_createSyncConfigurationCmd.MarkFlagRequired("branch")
		codestarConnections_createSyncConfigurationCmd.MarkFlagRequired("config-file")
		codestarConnections_createSyncConfigurationCmd.MarkFlagRequired("repository-link-id")
		codestarConnections_createSyncConfigurationCmd.MarkFlagRequired("resource-name")
		codestarConnections_createSyncConfigurationCmd.MarkFlagRequired("role-arn")
		codestarConnections_createSyncConfigurationCmd.MarkFlagRequired("sync-type")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_createSyncConfigurationCmd)
}
