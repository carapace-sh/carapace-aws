package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_createTemplateGroupAccessControlEntryCmd = &cobra.Command{
	Use:   "create-template-group-access-control-entry",
	Short: "Create a group access control entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_createTemplateGroupAccessControlEntryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_createTemplateGroupAccessControlEntryCmd).Standalone()

		pcaConnectorAd_createTemplateGroupAccessControlEntryCmd.Flags().String("access-rights", "", "Allow or deny permissions for an Active Directory group to enroll or autoenroll certificates for a template.")
		pcaConnectorAd_createTemplateGroupAccessControlEntryCmd.Flags().String("client-token", "", "Idempotency token.")
		pcaConnectorAd_createTemplateGroupAccessControlEntryCmd.Flags().String("group-display-name", "", "Name of the Active Directory group.")
		pcaConnectorAd_createTemplateGroupAccessControlEntryCmd.Flags().String("group-security-identifier", "", "Security identifier (SID) of the group object from Active Directory.")
		pcaConnectorAd_createTemplateGroupAccessControlEntryCmd.Flags().String("template-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html).")
		pcaConnectorAd_createTemplateGroupAccessControlEntryCmd.MarkFlagRequired("access-rights")
		pcaConnectorAd_createTemplateGroupAccessControlEntryCmd.MarkFlagRequired("group-display-name")
		pcaConnectorAd_createTemplateGroupAccessControlEntryCmd.MarkFlagRequired("group-security-identifier")
		pcaConnectorAd_createTemplateGroupAccessControlEntryCmd.MarkFlagRequired("template-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_createTemplateGroupAccessControlEntryCmd)
}
