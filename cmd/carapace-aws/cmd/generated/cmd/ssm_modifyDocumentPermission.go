package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_modifyDocumentPermissionCmd = &cobra.Command{
	Use:   "modify-document-permission",
	Short: "Shares a Amazon Web Services Systems Manager document (SSM document)publicly or privately.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_modifyDocumentPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_modifyDocumentPermissionCmd).Standalone()

		ssm_modifyDocumentPermissionCmd.Flags().String("account-ids-to-add", "", "The Amazon Web Services users that should have access to the document.")
		ssm_modifyDocumentPermissionCmd.Flags().String("account-ids-to-remove", "", "The Amazon Web Services users that should no longer have access to the document.")
		ssm_modifyDocumentPermissionCmd.Flags().String("name", "", "The name of the document that you want to share.")
		ssm_modifyDocumentPermissionCmd.Flags().String("permission-type", "", "The permission type for the document.")
		ssm_modifyDocumentPermissionCmd.Flags().String("shared-document-version", "", "(Optional) The version of the document to share.")
		ssm_modifyDocumentPermissionCmd.MarkFlagRequired("name")
		ssm_modifyDocumentPermissionCmd.MarkFlagRequired("permission-type")
	})
	ssmCmd.AddCommand(ssm_modifyDocumentPermissionCmd)
}
