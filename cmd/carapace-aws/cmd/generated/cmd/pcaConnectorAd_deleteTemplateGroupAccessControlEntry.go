package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_deleteTemplateGroupAccessControlEntryCmd = &cobra.Command{
	Use:   "delete-template-group-access-control-entry",
	Short: "Deletes a group access control entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_deleteTemplateGroupAccessControlEntryCmd).Standalone()

	pcaConnectorAd_deleteTemplateGroupAccessControlEntryCmd.Flags().String("group-security-identifier", "", "Security identifier (SID) of the group object from Active Directory.")
	pcaConnectorAd_deleteTemplateGroupAccessControlEntryCmd.Flags().String("template-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html).")
	pcaConnectorAd_deleteTemplateGroupAccessControlEntryCmd.MarkFlagRequired("group-security-identifier")
	pcaConnectorAd_deleteTemplateGroupAccessControlEntryCmd.MarkFlagRequired("template-arn")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_deleteTemplateGroupAccessControlEntryCmd)
}
