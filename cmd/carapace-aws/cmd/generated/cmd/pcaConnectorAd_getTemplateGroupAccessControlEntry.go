package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_getTemplateGroupAccessControlEntryCmd = &cobra.Command{
	Use:   "get-template-group-access-control-entry",
	Short: "Retrieves the group access control entries for a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_getTemplateGroupAccessControlEntryCmd).Standalone()

	pcaConnectorAd_getTemplateGroupAccessControlEntryCmd.Flags().String("group-security-identifier", "", "Security identifier (SID) of the group object from Active Directory.")
	pcaConnectorAd_getTemplateGroupAccessControlEntryCmd.Flags().String("template-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html).")
	pcaConnectorAd_getTemplateGroupAccessControlEntryCmd.MarkFlagRequired("group-security-identifier")
	pcaConnectorAd_getTemplateGroupAccessControlEntryCmd.MarkFlagRequired("template-arn")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_getTemplateGroupAccessControlEntryCmd)
}
