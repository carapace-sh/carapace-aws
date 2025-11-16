package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateDocumentMetadataCmd = &cobra.Command{
	Use:   "update-document-metadata",
	Short: "Amazon Web Services Systems Manager Change Manager will no longer be open to new customers starting November 7, 2025. If you would like to use Change Manager, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateDocumentMetadataCmd).Standalone()

	ssm_updateDocumentMetadataCmd.Flags().String("document-reviews", "", "The change template review details to update.")
	ssm_updateDocumentMetadataCmd.Flags().String("document-version", "", "The version of a change template in which to update approval metadata.")
	ssm_updateDocumentMetadataCmd.Flags().String("name", "", "The name of the change template for which a version's metadata is to be updated.")
	ssm_updateDocumentMetadataCmd.MarkFlagRequired("document-reviews")
	ssm_updateDocumentMetadataCmd.MarkFlagRequired("name")
	ssmCmd.AddCommand(ssm_updateDocumentMetadataCmd)
}
