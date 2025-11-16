package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_createSyncConfigurationCmd = &cobra.Command{
	Use:   "create-sync-configuration",
	Short: "Creates a sync configuration which allows Amazon Web Services to sync content from a Git repository to update a specified Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_createSyncConfigurationCmd).Standalone()

	codeconnections_createSyncConfigurationCmd.Flags().String("branch", "", "The branch in the repository from which changes will be synced.")
	codeconnections_createSyncConfigurationCmd.Flags().String("config-file", "", "The file name of the configuration file that manages syncing between the connection and the repository.")
	codeconnections_createSyncConfigurationCmd.Flags().String("publish-deployment-status", "", "Whether to enable or disable publishing of deployment status to source providers.")
	codeconnections_createSyncConfigurationCmd.Flags().String("pull-request-comment", "", "A toggle that specifies whether to enable or disable pull request comments for the sync configuration to be created.")
	codeconnections_createSyncConfigurationCmd.Flags().String("repository-link-id", "", "The ID of the repository link created for the connection.")
	codeconnections_createSyncConfigurationCmd.Flags().String("resource-name", "", "The name of the Amazon Web Services resource (for example, a CloudFormation stack in the case of CFN\\_STACK\\_SYNC) that will be synchronized from the linked repository.")
	codeconnections_createSyncConfigurationCmd.Flags().String("role-arn", "", "The ARN of the IAM role that grants permission for Amazon Web Services to use Git sync to update a given Amazon Web Services resource on your behalf.")
	codeconnections_createSyncConfigurationCmd.Flags().String("sync-type", "", "The type of sync configuration.")
	codeconnections_createSyncConfigurationCmd.Flags().String("trigger-resource-update-on", "", "When to trigger Git sync to begin the stack update.")
	codeconnections_createSyncConfigurationCmd.MarkFlagRequired("branch")
	codeconnections_createSyncConfigurationCmd.MarkFlagRequired("config-file")
	codeconnections_createSyncConfigurationCmd.MarkFlagRequired("repository-link-id")
	codeconnections_createSyncConfigurationCmd.MarkFlagRequired("resource-name")
	codeconnections_createSyncConfigurationCmd.MarkFlagRequired("role-arn")
	codeconnections_createSyncConfigurationCmd.MarkFlagRequired("sync-type")
	codeconnectionsCmd.AddCommand(codeconnections_createSyncConfigurationCmd)
}
