package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeDocumentPermissionCmd = &cobra.Command{
	Use:   "describe-document-permission",
	Short: "Describes the permissions for a Amazon Web Services Systems Manager document (SSM document).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeDocumentPermissionCmd).Standalone()

	ssm_describeDocumentPermissionCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeDocumentPermissionCmd.Flags().String("name", "", "The name of the document for which you are the owner.")
	ssm_describeDocumentPermissionCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_describeDocumentPermissionCmd.Flags().String("permission-type", "", "The permission type for the document.")
	ssm_describeDocumentPermissionCmd.MarkFlagRequired("name")
	ssm_describeDocumentPermissionCmd.MarkFlagRequired("permission-type")
	ssmCmd.AddCommand(ssm_describeDocumentPermissionCmd)
}
