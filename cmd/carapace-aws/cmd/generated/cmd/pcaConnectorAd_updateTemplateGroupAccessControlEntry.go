package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd = &cobra.Command{
	Use:   "update-template-group-access-control-entry",
	Short: "Update a group access control entry you created using [CreateTemplateGroupAccessControlEntry](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplateGroupAccessControlEntry.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd).Standalone()

		pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd.Flags().String("access-rights", "", "Allow or deny permissions for an Active Directory group to enroll or autoenroll certificates for a template.")
		pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd.Flags().String("group-display-name", "", "Name of the Active Directory group.")
		pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd.Flags().String("group-security-identifier", "", "Security identifier (SID) of the group object from Active Directory.")
		pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd.Flags().String("template-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html).")
		pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd.MarkFlagRequired("group-security-identifier")
		pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd.MarkFlagRequired("template-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_updateTemplateGroupAccessControlEntryCmd)
}
